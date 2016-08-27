

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
    "errors"
    "strings"
    "strconv"
    "database/sql"
    "github.com/lib/pq"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/card"
    "github.com/stripe/stripe-go/charge"
    "github.com/stripe/stripe-go/customer"
    "github.com/stripe/stripe-go/refund"
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
    error = pgsql.UpdateResultError(result, error)
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
        if stripeCard.TokenizationMethod == "apple_pay" {
            continue
        }

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
//                                                  PayReferralIfMemberWasReferralToExpert
//----------------------------------------------------------------------------------------


func PayReferralIfMemberWasReferralToExpert(memberID, expertID string) {
    Log.LogFunctionName()

    var error error
    row := config.DB.QueryRow(
        `select rt.referrerID, rt.referralCode from ReferralTable rt
             join FeedPostTable fp on postID = rt.referenceID::uuid
            where rt.referreeID = $1
              and rt.validFromDate <= transaction_timestamp()
              and rt.validToDate >= transaction_timestamp()
              and fp.userID = $2
         order by creationDate limit 1;`,
         expertID,
         memberID,
    )
    var referrerID, referralCode string
    error = row.Scan(&referrerID, &referralCode)
    if error != nil {
        Log.LogError(error)
        return
    }

    var result sql.Result
    result, error = config.DB.Exec(
        `update ReferralTable set
            redemptionDate = transaction_timestamp()
            where referralCode = $1;`,
        referralCode,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
    }

    memberName := PrettyNameForUserID(memberID)
    expertName := PrettyNameForUserID(expertID)

    message := fmt.Sprintf(
        "Bounty!  %s accepted your referral of %s.",
        memberName,
        expertName,
    )

    error = SendUserMessageInternal(
        BlitzMessage.Default_Global_SystemUserID,
        [] string { referrerID },
        "",
        message,
        BlitzMessage.UserMessageType_MTNotification,
        "",
        "",
    )
    if error != nil {
        Log.LogError(error)
    }

}


//----------------------------------------------------------------------------------------
//
//                                                                           ChargeRequest
//
//----------------------------------------------------------------------------------------


func ChargeRequest(session *Session, chargeReq *BlitzMessage.Charge) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var amountI int = 0
    var amountS = "0.00"
    if chargeReq.Amount != nil {
        amountF, _ := strconv.ParseFloat(*chargeReq.Amount, 64)
        amountS = fmt.Sprintf("%1.00f", amountF)
        amountI = int(amountF * 100)
    }

    if (chargeReq.ChargeStatus == nil ||
        *chargeReq.ChargeStatus != BlitzMessage.ChargeStatus_CSChargeRequest ||
        chargeReq.PayerID == nil ||
        len(*chargeReq.PayerID) == 0 ||
        chargeReq.ChargeToken == nil ||
        len(*chargeReq.ChargeToken) == 0 ||
        chargeReq.TokenType == nil ||
        chargeReq.PurchaseType == nil ||
        chargeReq.PurchaseTypeID == nil ||
        len(*chargeReq.PurchaseTypeID) == 0 ||
        amountI < 0) {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Missing fields"))
    }

    //  Get the Stripe customerID --

    //stripeCID, error := StripeCIDFromUserID(*chargeReq.PayerID)
    stripeCID, error := CreateStripeCIDFromUserIDToken(*chargeReq.PayerID, "")
    if error != nil {
        Log.Errorf("StripeCIDFromUserID returned '%+v': %+v.", stripeCID, error)
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
            if stripeError, ok := error.(*stripe.Error); ok {
                error = errors.New(stripeError.Msg)
            }
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
        chargeReq.ChargeToken = &newCard.ID
    }

    //  First check if we are over our daily fail-safe limit:

    row := config.DB.QueryRow(
        `select sum(amount) from ChargeTable
            where (now() - timestamp) < '24 hours'
              and chargeStatus = $1;`,
        BlitzMessage.ChargeStatus_CSCharged,
    )
    var total sql.NullFloat64
    error = row.Scan(&total)
    if error != nil || total.Float64 >= config.DailyChargeLimitDollars {
        Log.Errorf("Charge limit reached! Total: %1.2f Error: %v.", total.Float64, error)
        error = fmt.Errorf("Sorry, we aren't able to submit charges at the moment.")
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    //  Check to see if conversation is still open --

    if *chargeReq.PurchaseType == BlitzMessage.PurchaseType_PTChatConversation {

        row := config.DB.QueryRow(
            `select
                closedDate,
                paymentStatus
                    from ConversationTable
                    where conversationID = $1;`,
            chargeReq.PurchaseTypeID,
        )
        var (closedDate pq.NullTime; paymentStatus sql.NullInt64)
        error = row.Scan(&closedDate, &paymentStatus)
        if error != nil { Log.LogError(error) }

        if error != nil || closedDate.Valid {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Conversation is closed"))
        }
        if paymentStatus.Int64 > int64(BlitzMessage.PaymentStatus_PSPaymentRequired) {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Already paid"))
        }
    }

    //  Insert charge status:

    var result sql.Result
    result, error = config.DB.Exec(
        `insert into ChargeTable (
            chargeID,
            timestamp,
            chargeStatus,
            payerID,
            purchaseType,
            purchaseTypeID,
            memoText,
            amount,
            currency,
            chargeToken
        ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
        chargeReq.ChargeID,
        chargeReq.Timestamp.NullTime(),
        chargeReq.ChargeStatus,
        chargeReq.PayerID,
        chargeReq.PurchaseType,
        chargeReq.PurchaseTypeID,
        chargeReq.MemoText,
        amountS,
        chargeReq.Currency,
        chargeReq.ChargeToken,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil  {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    stripeReason := "Charged"
    stripeChargeID := ""
    chargeReq.ChargeStatus = BlitzMessage.ChargeStatus(BlitzMessage.ChargeStatus_CSCharged).Enum()
    responseCode := BlitzMessage.ResponseCode_RCSuccess

    //  Charge stripe --

    if amountI > 0 {
        chargeParams := &stripe.ChargeParams{
          Amount:           uint64(amountI),
          Currency:         "usd",
          Customer:         stripeCID,
        }
        error = chargeParams.SetSource(*chargeReq.ChargeToken)
        if error != nil {
            Log.LogError(error)
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }

        if chargeReq.MemoText != nil {
            chargeParams.Desc = *chargeReq.MemoText
        }

        //  We're charging the customer rather than the card.
        //  Not Needed:
        // error = chargeParams.SetSource(stripeCID)
        // if error != nil { Log.LogError(error) }

        Log.Debugf("Charge params: %+v", *chargeParams)
        stripeCharge, stripeError := charge.New(chargeParams)
        if stripeError != nil {
            Log.LogError(stripeError)
            Log.Debugf("Charge: %+v.", stripeCharge)
            responseCode = BlitzMessage.ResponseCode_RCPaymentError
            chargeReq.ChargeStatus = BlitzMessage.ChargeStatus(BlitzMessage.ChargeStatus_CSDeclined).Enum()

            chargeReq.ProcessorReason = proto.String("Your card was declined (1).")
            if stripeCharge != nil && len(stripeCharge.FailMsg) > 0 {
                chargeReq.ProcessorReason = &stripeCharge.FailMsg
            } else {
                if stripeErrorType, ok := stripeError.(*stripe.Error); ok {
                    chargeReq.ProcessorReason = &stripeErrorType.Msg
                }
            }

        } else {
            stripeChargeID = stripeCharge.ID
            chargeReq.ProcessorReason = &stripeReason
        }
    }

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
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    if responseCode == BlitzMessage.ResponseCode_RCSuccess &&
       (*chargeReq.PurchaseType == BlitzMessage.PurchaseType_PTChatConversation ||
        *chargeReq.PurchaseType == BlitzMessage.PurchaseType_PTCall) {

        var conversationID string = *chargeReq.PurchaseTypeID

        result, error = config.DB.Exec(
            `update ConversationTable set
                chargeID = $1,
                paymentStatus = $2
                where conversationID = $3;`,
            chargeReq.ChargeID,
            BlitzMessage.PaymentStatus_PSExpertNeedsAccept,
            conversationID,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil {
            Log.LogError(error)
            return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
        }

        members := MembersForConversationID(conversationID)
        var otherMember string
        for _, member := range members {
            if member != session.UserID {
                otherMember = member
                break
            }
        }

        memberName := PrettyNameForUserID(session.UserID)
        expertName := PrettyNameForUserID(otherMember)

        expertMessage := fmt.Sprintf(
            "Congratulations, %s\nhas requested your expertise.\nPlease reply immediately" +
            " to ensure the best service experience. This chat window will be open for the" +
            " next 24 hours, unless your requester may be satisfied earlier.",
            memberName,
        )

        actionURL := fmt.Sprintf("%s?action=showchat&chatid=%s",
            config.AppLinkURL,
            conversationID,
        )

        error = SendUserMessageInternal(
            BlitzMessage.Default_Global_SystemUserID,
            []string{ otherMember },
            conversationID,
            expertMessage,
            BlitzMessage.UserMessageType_MTConversation,
            "",
            actionURL,
        )

        userMessage := fmt.Sprintf(
            "Thank you for your payment!\n"+
            "%s has 24 hours to accept your request.\n"+
            "Otherwise your money will be automatically refunded to you.",
            expertName,
        )

        error = SendUserMessageInternal(
            BlitzMessage.Default_Global_SystemUserID,
            []string{ session.UserID } ,
            conversationID,
            userMessage,
            BlitzMessage.UserMessageType_MTConversation,
            "",
            actionURL,
        )

        //  See if we need to pay a referral fee --

        PayReferralIfMemberWasReferralToExpert(session.UserID, otherMember)
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseType:       &BlitzMessage.ResponseType { ChargeResponse: chargeReq },
    }
    return response
}


//----------------------------------------------------------------------------------------
//
//                                                                FetchPurchaseDescription
//
//----------------------------------------------------------------------------------------


func FetchPurchaseDescription(session *Session,
                              fetch *BlitzMessage.FetchPurchaseDescription,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if fetch == nil ||
       fetch.Purchase == nil ||
       fetch.Purchase.PurchaseType == nil ||
       fetch.Purchase.PurchaseTypeID == nil ||
       len(*fetch.Purchase.PurchaseTypeID) == 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Missing fields"))
    }

    var error error
    purchase := fetch.Purchase

    switch *purchase.PurchaseType {

    case BlitzMessage.PurchaseType_PTChatConversation:
        error = UpdatePurchaseDescriptionForConversation(session, purchase)

    default:
        error = errors.New("Purchase not supported")
    }

    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FetchPurchaseDescription: fetch },
    }

    return response
}


//----------------------------------------------------------------------------------------
//
//                                                                          RefundChargeID
//
//----------------------------------------------------------------------------------------


func RefundChargeID(chargeID string, reason string) error {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select chargeStatus, processorChargeID
            from ChargeTable
            where chargeID = $1
            for update;`,
        chargeID,
    )
    var ( chargeStatus sql.NullInt64; processorChargeID sql.NullString)
    error := row.Scan(&chargeStatus, &processorChargeID)
    if error != nil {
        Log.LogError(error)
        return error
    }
    if  BlitzMessage.ChargeStatus(chargeStatus.Int64) == BlitzMessage.ChargeStatus_CSRefunded {
        return errors.New("Already refunded")
    }
    if  BlitzMessage.ChargeStatus(chargeStatus.Int64) != BlitzMessage.ChargeStatus_CSCharged ||
        ! processorChargeID.Valid {
        return errors.New("No charges")
    }

    params := &stripe.RefundParams{
        Charge:     processorChargeID.String,
        Reason:     "requested_by_customer",
    }
    params.Meta = map[string]string { "MemoText": reason }

    var resp *stripe.Refund
    resp, error = refund.New(params)
    if error != nil {
        Log.LogError(error)
        return error
    }
    var result sql.Result
    result, error = config.DB.Exec(
        `update ChargeTable set
            chargeStatus = $1,
            refundDate = transaction_timestamp(),
            refundProcessorID = $2,
            refundMemo = $3
                where chargeID = $4;`,
        BlitzMessage.ChargeStatus_CSRefunded,
        resp.ID,
        reason,
        chargeID,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
    }
    return error
}

