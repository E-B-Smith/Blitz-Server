

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
    "strings"
    "strconv"
    "database/sql"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/card"
    "github.com/stripe/stripe-go/charge"
    "github.com/stripe/stripe-go/customer"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func StripeCIDFromUserID(userID string) (string, error) {
    row := config.DB.QueryRow(
        `select processorID from UserTable where userID = $1;`,
        userID,
    )
    var stripeCID sql.NullString
    error := row.Scan(&stripeCID)
    if error != nil {
        Log.LogError(error)
        return "", error
    }
    if ! stripeCID.Valid {
        return "", fmt.Errorf("Not found")
    }
    return stripeCID.String, nil
}


func CreateStripeCIDFromUserIDToken(userID, token string) (string, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select processorID, name
            from UserTable where userID = $1;`,
        userID)

    var processorID, name, defaultCard sql.NullString
    error := row.Scan(&processorID, &name)
    if error != nil {
        Log.LogError(error)
        return "", error
    }

    if processorID.Valid && len(processorID.String) > 0 {
        return processorID.String, nil
    }

    customerParams := &stripe.CustomerParams {}
    customerParams.Desc = name.String
    customerParams.Meta = map[string]string {"UserID": userID}
    if len(token) > 0 {
        customerParams.SetSource(token)
    }
    newCust, error := customer.New(customerParams)
    if error != nil {
        Log.LogError(error)
        return "", error
    }
    Log.Debugf("Cust: %+v.", newCust)
    processorID.Valid  = true
    processorID.String = newCust.ID
    if newCust.DefaultSource != nil {
        defaultCard.Valid = true
        defaultCard.String = newCust.DefaultSource.ID
    }

    result, error := config.DB.Exec(
        `update UserTable set (
            processorID,
            defaultCard
        ) = ($1, $2)
            where userID = $3;`,
        processorID,
        defaultCard,
        userID,
    )
    error = pgsql.ResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return "", error
    }

    return processorID.String, nil
}


func UpdateCardForStripeCID(stripeCID string, cardInfo *BlitzMessage.CardInfo) error {
    Log.LogFunctionName()

    var error error
    var cardParams stripe.CardParams
    cardParams.Customer = stripeCID
    cardParams.Name     = StringFromStringPtr(cardInfo.CardHolderName)
    if cardInfo.ExpireMonth != nil {
        cardParams.Month = strconv.Itoa(int(*cardInfo.ExpireMonth))
    }
    if cardInfo.ExpireYear != nil {
        cardParams.Year = strconv.Itoa(int(*cardInfo.ExpireYear))
    }
    memoText := ""
    if cardInfo.MemoText != nil {
        memoText = *cardInfo.MemoText
    }
    cardParams.Meta = map[string]string { "MemoText": memoText }

    if cardInfo.Token == nil {
        return fmt.Errorf("No token")
    }
    if strings.HasPrefix(*cardInfo.Token, "tok") {

        //  Add a new card --

        var newCard *stripe.Card
        newCardParams := stripe.CardParams{
            Customer:   stripeCID,
            Token:      *cardInfo.Token,
        }
        newCard, error = card.New(&newCardParams)
        if error != nil {
            Log.LogError(error)
            cardInfo.Token = nil
        } else {
            cardInfo.Token = &newCard.ID
        }

    }

    if error == nil {
        //  Update a card --
        _, error = card.Update(*cardInfo.Token, &cardParams)
        if error != nil {
            Log.LogError(error)
        }
    }

    return error
}


func DeleteCardForStripeCID(stripeCID string, cardInfo *BlitzMessage.CardInfo) error {
    Log.LogFunctionName()
    cardParams := &stripe.CardParams {
        Customer:   stripeCID,
    }
    _, error := card.Del(*cardInfo.Token, cardParams)
    if error != nil {
        Log.LogError(error)
    }
    return nil
}


func CardsForStripeCID(stripeCID string) []*BlitzMessage.CardInfo {
    Log.LogFunctionName()

    result := make([]*BlitzMessage.CardInfo, 0)

    params := &stripe.CardListParams { Customer: stripeCID }
    iter := card.List(params)
    for iter.Next() {

        stripeCard := iter.Card()

        last4 := stripeCard.LastFour
        if len(last4) == 0 || last4 == "0000" {
            last4 = stripeCard.DynLastFour
        }
        memoText := stripeCard.Meta["MemoText"]

        card := BlitzMessage.CardInfo {
            CardStatus:         BlitzMessage.CardStatus(BlitzMessage.CardStatus_CSStandard).Enum(),
            CardHolderName:     proto.String(stripeCard.Name),
            MemoText:           proto.String(memoText),
            Brand:              proto.String(string(stripeCard.Brand)),
            Last4:              proto.String(last4),
            ExpireMonth:        proto.Int32(int32(stripeCard.Month)),
            ExpireYear:         proto.Int32(int32(stripeCard.Year)),
            Token:              proto.String(stripeCard.ID),
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

    stripeCID, error := StripeCIDFromUserID(session.UserID)
    if error != nil {
        //  Create a customer --
        stripeCID, error = CreateStripeCIDFromUserIDToken(session.UserID, "")
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
        }
    }

    for _, cardInfo := range cardUpdate.CardInfo {

        error = fmt.Errorf("Invalid card status")

        if cardInfo.CardStatus != nil {

            switch *cardInfo.CardStatus {
            case BlitzMessage.CardStatus_CSStandard:
                error = UpdateCardForStripeCID(stripeCID, cardInfo)

            case BlitzMessage.CardStatus_CSDeleted:
                error = DeleteCardForStripeCID(stripeCID, cardInfo)
            }
        }

        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
    }

    updatedInfo := BlitzMessage.UserCardInfo {
        CardInfo:    CardsForStripeCID(stripeCID),
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
        chargeReq.PayerID == nil || len(*chargeReq.PayerID) == 0 ||
        chargeReq.ChargeToken == nil || len(*chargeReq.ChargeToken) == 0 ||
        chargeReq.ChargeToken == nil ||
        amountI < 0) {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    //  Get the Stripe customerID --

    stripeCID, error := StripeCIDFromUserID(*chargeReq.PayerID)//, *chargeReq.ChargeToken)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    if strings.HasPrefix(*chargeReq.ChargeToken, "tok") {
        //  Add a new card --
        var newCard *stripe.Card
        newCardParams := stripe.CardParams{
            Customer:   stripeCID,
            Token:      *chargeReq.ChargeToken,
        }
        newCard, error = card.New(&newCardParams)
        if error != nil {
            Log.LogError(error)
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
        chargeReq.ChargeToken = &newCard.ID
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
    responseCode := BlitzMessage.ResponseCode_RCSuccess

    //  Charge stripe --

    chargeParams := &stripe.ChargeParams{
      Amount:           uint64(amountI),
      Currency:         "usd",
      Customer:         stripeCID,
    }
    if chargeReq.MemoText != nil {
        chargeParams.Desc = *chargeReq.MemoText
    }

    //  We're charging the customer rather than the card. Not Needed:
    // error = chargeParams.SetSource(stripeCID)
    // if error != nil { Log.LogError(error) }

    Log.Debugf("Charge params: %+v", *chargeParams)
    stripeCharge, stripeError := charge.New(chargeParams)
    if stripeError != nil {
        Log.LogError(stripeError)
        chargeReq.ChargeStatus = BlitzMessage.ChargeStatus(BlitzMessage.ChargeStatus_CSDeclined).Enum()
        stripeReason = stripeError.Error()
        responseCode = BlitzMessage.ResponseCode_RCPaymentError
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
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseType:       &BlitzMessage.ResponseType { ChargeResponse: chargeReq },
    }
    return response
}

