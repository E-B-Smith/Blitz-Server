

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
//    "time"
//    "database/sql"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func UpdateCard(userID string, card *BlitzMessage.CardInfo) error {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `insert into CardTable
            (userID
            ,cardStatus
            ,cardHolderName
            ,memoText
            ,brand
            ,last4
            ,expireMonth
            ,expireYear
            ,token) = ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        on conflict (userID, brand, last4)
        update CardTable set
            (cardStatus
            ,cardHolderName
            ,memoText
            ,expireMonth
            ,expireYear
            ,token) values ($2, $3, $4, $7, $8, $9)
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
            case BlitzMessage.CardStatus_CSDefault, BlitzMessage.CardStatus_CSStandard:
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
//                                                                                Transact
//
//----------------------------------------------------------------------------------------


func Transact(session *Session, payment *BlitzMessage.Payment) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `insert into PaymentTable (
            payerID,
            payeeID,
            conversationID,
            timestamp,
            paymentStatus,
            token,
            memoText,
            amount,
            currency
        ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
        payment.PayerID,
        payment.PayeeID,
        payment.ConversationID,
        payment.Timestamp,
        payment.PaymentStatus,
        payment.Token,
        payment.MemoText,
        payment.Amount,
        "usd",
    )
    error = pgsql.ResultError(result, error)
    if error != nil  { Log.LogError(error) }

    return ServerResponseForError(BlitzMessage.ResponseCode_RCSuccess, nil)
}


