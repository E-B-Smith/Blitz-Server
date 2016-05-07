

//----------------------------------------------------------------------------------------
//
//                                                          BlitzHere-Server : Payments.go
//                                                          Stripe payment & card routines
//
//                                                                   E.B. Smith, May, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "strconv"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/charge"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func UpdateCard(userID string, card *BlitzMessage.CardInfo) error {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `insert into CardTable (
             userID
            ,cardStatus
            ,cardHolderName
            ,memoText
            ,brand
            ,last4
            ,expireMonth
            ,expireYear
            ,token
        ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        on conflict (userID, brand, last4)
        update CardTable set (
             cardStatus
            ,cardHolderName
            ,memoText
            ,expireMonth
            ,expireYear
            ,token
        ) = ($2, $3, $4, $7, $8, $9)
            where userID = $1
              and brand  = $5
              and last4  = $6;`,
            userID,
            card.CardStatus,
            card.CardHolderName,
            card.MemoText,
            card.Brand,
            card.Last4,
            card.ExpireMonth,
            card.ExpireYear,
            card.Token,
    )
    error = pgsql.ResultError(result, error)
    if error != nil { Log.LogError(error) }
    return error
}


func DeleteCard(userID string, card *BlitzMessage.CardInfo) error {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `delete from CardTable
            where userID = $1
              and brand  = $2
              and last4  = $3;`,
        userID,
        card.Brand,
        card.Last4,
    )
    error = pgsql.ResultError(result, error)
    if error != nil { Log.LogError(error) }

    return error
}


func CardsForUserID(userID string) []*BlitzMessage.CardInfo {
    Log.LogFunctionName()

    result := make([]*BlitzMessage.CardInfo, 0)

    rows, error := config.DB.Query(
        `select
             cardStatus
            ,cardHolderName
            ,memoText
            ,brand
            ,last4
            ,expireMonth
            ,expireYear
            ,token
        from CardTable
        where userID = $1;`,
        userID,
    )
    if error != nil {
        Log.LogError(error)
        return result
    }
    defer rows.Close()


    for rows.Next() {
        var (
            cardStatus          int32
            cardHolderName      string
            memoText            string
            brand               string
            last4               string
            expireMonth         int32
            expireYear          int32
            token               string
        )

        error = rows.Scan(
            &cardStatus,
            &cardHolderName,
            &memoText,
            &brand,
            &last4,
            &expireMonth,
            &expireYear,
            &token,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }

        card := BlitzMessage.CardInfo {
            CardStatus:         BlitzMessage.CardStatus(cardStatus).Enum(),
            CardHolderName:     proto.String(cardHolderName),
            MemoText:           proto.String(memoText),
            Brand:              proto.String(brand),
            Last4:              proto.String(last4),
            ExpireMonth:        proto.Int32(expireMonth),
            ExpireYear:         proto.Int32(expireYear),
            Token:              proto.String(token),
        }

        result = append(result, &card)
    }

    return result
}


//----------------------------------------------------------------------------------------
//
//                                                                             UpdateCards
//
//----------------------------------------------------------------------------------------


func UpdateCards(session *Session, cardUpdate *BlitzMessage.UserCardInfo) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var error error
    for _, cardInfo := range cardUpdate.CardInfo {

        error = fmt.Errorf("Invalid card status")

        if cardInfo.CardStatus != nil {

            switch *cardInfo.CardStatus {
            case BlitzMessage.CardStatus_CSStandard:
                error = UpdateCard(session.UserID, cardInfo)

            case BlitzMessage.CardStatus_CSDeleted:
                error = DeleteCard(session.UserID, cardInfo)
            }
        }

        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
    }

    updatedInfo := BlitzMessage.UserCardInfo {
        CardInfo:    CardsForUserID(session.UserID),
    }
    response := &BlitzMessage.ServerResponse {
        ResponseCode:        BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { UserCardInfo: &updatedInfo },
    }
    return response
}


//----------------------------------------------------------------------------------------
//
//                                                                           ChargeRequest
//
//----------------------------------------------------------------------------------------


func ChargeRequest(session *Session, chargeReq *BlitzMessage.Charge) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var amountI int = 0
    if chargeReq.Amount != nil {
        amountF, _ := strconv.ParseFloat(*chargeReq.Amount, 64)
        amountI = int(amountF * 100)
    }

    if (chargeReq.ChargeStatus == nil ||
        *chargeReq.ChargeStatus != BlitzMessage.ChargeStatus_CSChargeRequest ||
        chargeReq.ChargeToken == nil || len(*chargeReq.ChargeToken) == 0 ||
        chargeReq.ChargeToken == nil ||
        amountI < 0) {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    result, error := config.DB.Exec(
        `insert into ChargeTable (
            chargeID,
            timestamp,
            chargeStatus,
            payerID,
            payeeID,
            conversationID,
            memoText,
            amount,
            currency,
            chargeToken
        ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
        chargeReq.ChargeID,
        BlitzMessage.NullTimeFromTimestamp(chargeReq.Timestamp),
        chargeReq.ChargeStatus,
        chargeReq.PayerID,
        chargeReq.PayeeID,
        chargeReq.ConversationID,
        chargeReq.MemoText,
        chargeReq.Amount,
        chargeReq.Currency,
        chargeReq.ChargeToken,
    )
    error = pgsql.ResultError(result, error)
    if error != nil  {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    stripeReason := "Charged"
    stripeChargeID := ""
    chargeReq.ChargeStatus = BlitzMessage.ChargeStatus(BlitzMessage.ChargeStatus_CSCharged).Enum()


    //  Charge stripe --

    chargeParams := &stripe.ChargeParams{
      Amount: uint64(amountI),
      Currency: "usd",
    }
    if chargeReq.MemoText != nil {
        chargeParams.Desc = *chargeReq.MemoText
    }

    chargeParams.SetSource(chargeReq.ChargeToken)
    stripeCharge, stripeError := charge.New(chargeParams)
    if stripeError != nil {
        Log.LogError(stripeError)
        chargeReq.ChargeStatus = BlitzMessage.ChargeStatus(BlitzMessage.ChargeStatus_CSDeclined).Enum()
        stripeReason = stripeError.Error()
    } else {
        stripeChargeID = stripeCharge.ID
    }

    chargeReq.ProcessorReason = &stripeReason

    result, error = config.DB.Exec(
        `update ChargeTable set (
            chargeStatus,
            processorReason,
            processorChargeID) = ($1, $2, $3)
        where chargeID = $4;`,
        chargeReq.ChargeStatus,
        stripeReason,
        stripeChargeID,
        chargeReq.ChargeID,
    )
    error = pgsql.ResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:        BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { ChargeResponse: chargeReq },
    }
    return response
}

