//  Messages.go  -  Dispatch user messages.
//
//  E.B.Smith  -  March, 2015


package main


import (
    "fmt"
    "time"
    "errors"
    "strings"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
    "MessagePusher"
)


//----------------------------------------------------------------------------------------
//
//                                                                 ScanUserMessageTableRow
//
//----------------------------------------------------------------------------------------


const kScanUserMessageTableRow =
`   messageID,
    senderID,
    recipientID,
    creationDate,
    notificationDate,
    readDate,
    messageType,
    messageText,
    actionIcon,
    actionURL,
    conversationID,
    messageStatus
`


func ScanUserMessageTableRow(rows *sql.Rows) *BlitzMessage.UserMessage {
    var (
        messageID           string
        senderID            string
        recipientID         string
        creationDate        pq.NullTime
        notificationDate    pq.NullTime
        readDate            pq.NullTime
        messageType         int
        messageText         sql.NullString
        actionIcon          sql.NullString
        actionURL           sql.NullString
        conversationID      sql.NullString
        messageStatus       sql.NullInt64
    )
    var error error
    error = rows.Scan(
        &messageID,
        &senderID,
        &recipientID,
        &creationDate,
        &notificationDate,
        &readDate,
        &messageType,
        &messageText,
        &actionIcon,
        &actionURL,
        &conversationID,
        &messageStatus,
    )
    if error != nil {
        Log.LogError(error)
        return nil
    }
    mt := BlitzMessage.UserMessageType(messageType);
    message := BlitzMessage.UserMessage {
        MessageID:          &messageID,
        SenderID:           &senderID,
        Recipients:         []string{recipientID},
        MessageType:        &mt,
        CreationDate:       BlitzMessage.TimestampPtr(creationDate),
        NotificationDate:   BlitzMessage.TimestampPtr(notificationDate.Time),
        ReadDate:           BlitzMessage.TimestampPtr(readDate.Time),
    }
    if messageText.Valid        { message.MessageText = &messageText.String }
    if actionIcon.Valid         { message.ActionIcon = &actionIcon.String }
    if actionURL.Valid          { message.ActionURL = &actionURL.String }
    if conversationID.Valid     { message.ConversationID = &conversationID.String }

    if  messageStatus.Valid {
        message.MessageStatus = BlitzMessage.UserMessageStatus(messageStatus.Int64).Enum()
    } else {
        message.MessageStatus = BlitzMessage.UserMessageStatus(BlitzMessage.UserMessageStatus_MSNew).Enum()
    }

    return &message
}


//----------------------------------------------------------------------------------------
//
//                                                                 UserMessageFetchRequest
//
//----------------------------------------------------------------------------------------


func UserMessageFetchRequest(session *Session, fetch *BlitzMessage.UserMessageUpdate) *BlitzMessage.ServerResponse {
    //
    //  Fetch messages for the user for the given timespan --
    //

    Log.LogFunctionName()

    var startDate time.Time = pgsql.NegativeInfinityTime
    var stopDate  time.Time = pgsql.PositiveInfinityTime

    if fetch.Timespan != nil {
        if fetch.Timespan.StartTimestamp != nil {
            startDate = fetch.Timespan.StartTimestamp.Time()
        }
        if fetch.Timespan.StopTimestamp != nil {
            stopDate = fetch.Timespan.StopTimestamp.Time()
        }
    }

    warnings := 0

    rows, error := config.DB.Query(
        `select ` + kScanUserMessageTableRow +
        `   from UserMessageTable
            where recipientID = $1
              and creationDate >  $2
              and creationDate <= $3
            order by creationDate;`,
        session.UserID,
        startDate,
        stopDate,
    )
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }
    defer rows.Close()

    var messageArray []*BlitzMessage.UserMessage
    for rows.Next() {
        message := ScanUserMessageTableRow(rows)
        if message == nil {
            warnings++
        } else {
            messageArray = append(messageArray, message)
        }
    }

    Log.Debugf("Found %d message (%d warnings) in range %v to %v.", len(messageArray), warnings, startDate, stopDate)

    if len(messageArray) == 0 && warnings > 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, errors.New("Messages are not available now."))
    }
    messageUpdate := BlitzMessage.UserMessageUpdate {
        Messages:   messageArray,
    }

    if  warnings == 0 && len(messageArray) > 0 {
        var timespan BlitzMessage.Timespan
        timespan.StartTimestamp = messageArray[0].CreationDate
        timespan.StopTimestamp  = messageArray[len(messageArray)-1].CreationDate
        messageUpdate.Timespan = &timespan
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        ResponseType:   &BlitzMessage.ResponseType { UserMessageUpdate: &messageUpdate },
    }
    return response
}


//----------------------------------------------------------------------------------------
//                                                                         SendUserMessage
//----------------------------------------------------------------------------------------


func mapAppendFromArray(stringMap map[string]bool, s []string) map[string]bool {
    for _, s := range s {
        stringMap[s] = true
    }
    return stringMap
}


func arrayFromMap(stringMap map[string]bool) []string {
    i := 0
    a := make([]string, len(stringMap))
    for key, _ := range stringMap {
        a[i] = key
        i++
    }
    return a
}


var ErrorPaymentRequired = errors.New("Payment required")


func UpdateConversationMessage(
        session *Session,
        message *BlitzMessage.UserMessage,
        ) error {
    Log.LogFunctionName()

    var error error
    conversation, error := ReadUserConversation(session.UserID, *message.ConversationID)
    if error != nil {
        Log.LogError(error)
        return error
    }

    //  Add an action to the message --

    if  message.ActionURL == nil || len(*message.ActionURL) == 0 {
        message.ActionURL = proto.String(
            fmt.Sprintf("%s?action=showchat&chatid=%s",
                config.AppLinkURL,
                *message.ConversationID,
        ))
    }

    //  Add recipients --

    if len(conversation.MemberIDs) < 2 {
        return errors.New("Not enough conversation members.")
    }

    message.Recipients = conversation.MemberIDs

    if conversation.ClosedDate != nil {
        return errors.New("This conversation is closed.")
    }

    if conversation.PaymentStatus == nil {
        conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSUnknown).Enum()
    }

    message.PaymentStatus = conversation.PaymentStatus
    switch *conversation.PaymentStatus {

    case BlitzMessage.PaymentStatus_PSIsFree,
         BlitzMessage.PaymentStatus_PSExpertAccepted:
         return nil

    case BlitzMessage.PaymentStatus_PSUnknown,
         BlitzMessage.PaymentStatus_PSTrialPeriod:
        //  Get the trial count messages:

        if  conversation.InitiatorUserID!= nil &&
            session.UserID != *conversation.InitiatorUserID {
            return nil
        }

        row := config.DB.QueryRow(
            `select count(*) from UserMessageTable
                where conversationID = $1
                  and senderID = $2;`,
            conversation.ConversationID,
            session.UserID,
        )
        var messagesSent sql.NullInt64
        error = row.Scan(&messagesSent)
        if error != nil {
            Log.LogError(error)
        }
        const trialCount = 1
        if messagesSent.Int64 < trialCount-1 {
            return nil
        }
        message.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSPaymentRequired).Enum()
        var result sql.Result
        result, error = config.DB.Exec(
            `update ConversationTable set paymentStatus = $1
                where conversationID = $2;`,
            BlitzMessage.PaymentStatus_PSPaymentRequired,
            conversation.ConversationID,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil { Log.LogError(error) }

        if messagesSent.Int64 < trialCount {
            return nil
        }
        return ErrorPaymentRequired

    case BlitzMessage.PaymentStatus_PSPaymentRequired:
        if  conversation.InitiatorUserID != nil &&
            session.UserID != *conversation.InitiatorUserID {
            return nil
        }
        return ErrorPaymentRequired

    case BlitzMessage.PaymentStatus_PSExpertNeedsAccept:
        return errors.New("Waiting for expert")

    case BlitzMessage.PaymentStatus_PSExpertRejected:
        return errors.New("Sorry, your expert is unavailable")
    }

    return nil
}


func isEmptyStringPtr(s *string) bool {
    return (s == nil || len(strings.TrimSpace(*s)) == 0)
}


func SendUserMessage(
        session *Session,
        message *BlitzMessage.UserMessage,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var error error
    if  isEmptyStringPtr(message.SenderID) ||
        isEmptyStringPtr(message.MessageID) ||
        message.MessageType == nil {
        error = errors.New("Missing fields")
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    if message.ConversationID != nil &&
        *message.MessageType == BlitzMessage.UserMessageType_MTConversation {
        error = UpdateConversationMessage(session, message)
        if error == ErrorPaymentRequired {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCPurchaseRequired, error)
        } else if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
    }

    error = WriteUserMessage(message)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }
    return ServerResponseForError(BlitzMessage.ResponseCode_RCSuccess, nil)
}


func paymentStatusForMessage(message *BlitzMessage.UserMessage) *BlitzMessage.PaymentStatus {

    var ps *BlitzMessage.PaymentStatus = nil

    if  message != nil &&
        message.ConversationID != nil &&
        len(*message.ConversationID) > 0 &&
        (message.PaymentStatus == nil ||
         *message.PaymentStatus == BlitzMessage.PaymentStatus_PSUnknown) {

        row := config.DB.QueryRow(
            `select paymentStatus from ConversationTable
                where conversationID = $1;`,
            message.ConversationID,
        )
        var paymentStatus sql.NullInt64
        error := row.Scan(&paymentStatus)
        if error != nil { Log.LogError(error) }
        if paymentStatus.Valid {
            ps = BlitzMessage.PaymentStatus(paymentStatus.Int64).Enum()
        }
    }

    return ps
}


func WriteUserMessage(message *BlitzMessage.UserMessage) error {
    Log.LogFunctionName()

    var firstError error
    recipients := make(map[string]bool)
    recipients = mapAppendFromArray(recipients, message.Recipients)
    recipients = mapAppendFromArray(recipients, []string {*message.SenderID} )
    message.Recipients = arrayFromMap(recipients)

    if len(message.Recipients) == 0 {
        return errors.New("No recipients")
    }

    message.PaymentStatus = paymentStatusForMessage(message)

    for _, recipientID := range message.Recipients {
        _, error := config.DB.Exec(
            `insert into UserMessageTable(
                messageID,
                senderID,
                recipientID,
                creationDate,
                messageType,
                messageText,
                actionIcon,
                actionURL,
                conversationID
            ) values ($1, $2, $3, transaction_timestamp(), $4, $5, $6, $7, $8)
            on conflict do nothing;`,
            message.MessageID,
            message.SenderID,
            recipientID,
            //BlitzMessage.NullTimeFromTimestamp(message.CreationDate),
            message.MessageType,
            message.MessageText,
            message.ActionIcon,
            message.ActionURL,
            message.ConversationID,
        )

        if error != nil {
            Log.Errorf("Error inserting message: %v. MessageID: %s From: %s To: %s.",
                error, *message.MessageID, *message.SenderID, recipientID)
            if firstError == nil { firstError = error }
        }
    }

    globalMessagePusher.PushMessage(message)
    return firstError
}


//----------------------------------------------------------------------------------------
//                                                                 SendUserMessageInternal
//----------------------------------------------------------------------------------------


func SendUserMessageInternal(
        sender string,
        recipients []string,
        conversationID string,
        message string,
        messageType BlitzMessage.UserMessageType,
        actionIcon string,
        actionURL string,
        ) error {
    Log.LogFunctionName()

    if  len(sender) == 0 ||
        len(recipients) == 0 ||
        len(message) == 0 {
        error := errors.New("Bad parameters")
        Log.LogError(error)
        return error
    }

    status := BlitzMessage.UserMessageStatus_MSNew
    blitzMessage := &BlitzMessage.UserMessage {
        MessageID:      StringPtr(Util.NewUUIDString()),
        SenderID:       &sender,
        Recipients:     recipients,
        CreationDate:   BlitzMessage.TimestampPtr(time.Now()),
        MessageType:    &messageType,
        MessageStatus:  &status,
        MessageText:    &message,
        ActionIcon:     &actionIcon,
        ActionURL:      &actionURL,
    }
    if len(conversationID) > 0 {
        blitzMessage.ConversationID = proto.String(conversationID)
    }

    WriteUserMessage(blitzMessage)
    return nil
}


//----------------------------------------------------------------------------------------
//
//                                                                  UserMessageSendRequest
//
//----------------------------------------------------------------------------------------


/*
func UserMessageSendUserMessage(
        session *Session,
        sendMessage *BlitzMessage.UserMessageUpdate,
        ) *BlitzMessage.ServerResponse {

    //
    //  * Save each new message to the database.
    //

    Log.LogFunctionName()

    messagesSent := 0
    for _, message := range sendMessage.Messages {
        Log.Debugf("Message %d has %d recipients.", messagesSent+1, len(message.Recipients))
        SendBlitzUserMessage(message)
        messagesSent += len(message.Recipients)
    }

    Log.Debugf("Received %d message bundles, sent %d messages.", len(sendMessage.Messages), messagesSent)

    messageResponse := &BlitzMessage.UserMessageUpdate {
        Timespan: sendMessage.Timespan,
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        ResponseType:   &BlitzMessage.ResponseType { UserMessageUpdate: messageResponse },
    }

    return response
}
*/


//----------------------------------------------------------------------------------------
//
//                                                                  UserDidConnectToPusher
//
//----------------------------------------------------------------------------------------


func UserDidConnectToPusher(
        pusher *MessagePusher.MessagePusher,
        user *MessagePusher.MessagePushUser,
        ) {
    Log.LogFunctionName()

    if  user.LastMessageTime == nil {
        return
    }

    rows, error := config.DB.Query(
        `select ` + kScanUserMessageTableRow +
        `   from UserMessageTable
           where recipientID = $1
             and creationDate > $2
           order by creationDate
           limit 50;`,
        user.UserID(),
        user.LastMessageTime,
    )
    if error != nil {
        Log.LogError(error)
        return
    }
    defer rows.Close()

    var messageCount int = 0
    for rows.Next() {
        message := ScanUserMessageTableRow(rows)
        if message != nil {
            message.PaymentStatus = paymentStatusForMessage(message)
            user.SendMessage(message)
            messageCount++
        }
    }
    Log.Debugf("Sent %d catchup messages.", messageCount)
}



