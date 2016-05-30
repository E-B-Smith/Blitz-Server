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
        MessageID:      &messageID,
        SenderID:       &senderID,
        Recipients:     []string{recipientID},
        MessageType:    &mt,
    }
    if creationDate.Valid       { message.CreationDate = BlitzMessage.TimestampFromTime(creationDate.Time) }
    if notificationDate.Valid   { message.NotificationDate = BlitzMessage.TimestampFromTime(notificationDate.Time) }
    if readDate.Valid           { message.ReadDate = BlitzMessage.TimestampFromTime(readDate.Time) }
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
            startDate = BlitzMessage.TimeFromTimestamp(fetch.Timespan.StartTimestamp)
        }
        if fetch.Timespan.StopTimestamp != nil {
            stopDate = BlitzMessage.TimeFromTimestamp(fetch.Timespan.StopTimestamp)
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


func updateMessageFromConversation(
        session *Session,
        message *BlitzMessage.UserMessage,
        conversation *BlitzMessage.Conversation,
        ) *BlitzMessage.ServerResponse {
    return nil
}


func UpdateConversationMessage(
        session *Session,
        message *BlitzMessage.UserMessage,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var error error
    conversation, error := ReadUserConversation(session.UserID, *message.ConversationID)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    //  Add recipients --

    message.Recipients = conversation.MemberIDs

    //  Add an action to the message --

    if  message.ActionURL == nil || len(*message.ActionURL) == 0 {
        message.ActionURL = proto.String(
            fmt.Sprintf("%s?action=showchat&chatid=%s",
                config.AppLinkURL,
                *message.ConversationID,
        ))
    }

    if conversation.ClosedDate != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid,
            errors.New("Conversation is closed."))
    }

    //  Declare MakeConversationFree function:

    makeConversationFree :=
        func (conversationID string) {
            _, error = config.DB.Exec(
                `update ConversationTable set isFree = true where conversationID = $1;`,
                *message.ConversationID,
            )
            if error != nil { Log.LogError(error) }
        }

    //  Free or paid?

    if  (conversation.IsFree != nil && *conversation.IsFree) ||
        (conversation.ChargeID != nil && len(*conversation.ChargeID) > 0) {
        return updateMessageFromConversation(session, message, conversation)
    }
    if  config.ServiceIsFree {
        _, error = config.DB.Exec(
            `update ConversationTable set isFree = true where conversationID = $1;`,
            *message.ConversationID,
        )
        makeConversationFree(*conversation.ConversationID)
        return updateMessageFromConversation(session, message, conversation)
    }

    //  Friends?

    if len(conversation.MemberIDs) < 2 {
        return ServerResponseForError(
            BlitzMessage.ResponseCode_RCInputInvalid,
            errors.New("Not enough conversation members."),
        )
    }

    tags := GetEntityTagMapForUserIDEntityIDType(
        session.UserID,
        conversation.MemberIDs[1],
        BlitzMessage.EntityType_ETUser,
    )
    if _, ok := tags[".friends"]; ok {
        makeConversationFree(*message.ConversationID)
        return updateMessageFromConversation(session, message, conversation)
    }

    //  Free for user?

    row := config.DB.QueryRow(
        `select isFree from UserTable where userID = $1;`,
        session.UserID,
    )
    var isFree sql.NullBool
    error = row.Scan(&isFree)
    if error != nil { Log.LogError(error) }
    if isFree.Bool {
        makeConversationFree(*message.ConversationID)
        return updateMessageFromConversation(session, message, conversation)
    }

    //  Less than four messages?

    if conversation.MessageCount != nil && *conversation.MessageCount <= 4 {
        return updateMessageFromConversation(session, message, conversation)
    }

    //  Make charge --

    amount := "10.00"
    memo := fmt.Sprintf("Chat with %s.",
        PrettyNameForUserID(conversation.MemberIDs[1]),
    )

    purchase := &BlitzMessage.PurchaseDescription {
        PurchaseType:           BlitzMessage.PurchaseType_PTChatConversation.Enum(),
        PurchaseTypeID:         message.ConversationID,
        //PurchaseID:           proto.String(GenerateUUID()),
        MemoText:               proto.String(memo),
        Amount:                 proto.String(amount),
        Currency:               proto.String("usd"),
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCPurchaseRequired).Enum(),
        ResponseMessage:    proto.String(memo),
        ResponseType:       &BlitzMessage.ResponseType { PurchaseDescription: purchase },
    }
    return response
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
        response := UpdateConversationMessage(session, message)
        if response != nil {
            return response
        }
    }

    error = WriteUserMessage(message)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }
    return ServerResponseForError(BlitzMessage.ResponseCode_RCSuccess, nil)
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
            ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
            on conflict do nothing;`,
            message.MessageID,
            message.SenderID,
            recipientID,
            BlitzMessage.NullTimeFromTimestamp(message.CreationDate),
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
        CreationDate:   BlitzMessage.TimestampFromTime(time.Now()),
        MessageType:    &messageType,
        MessageStatus:  &status,
        MessageText:    &message,
        ActionIcon:     &actionIcon,
        ActionURL:      &actionURL,
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
            user.SendMessage(message)
            messageCount++
        }
    }
    Log.Debugf("Sent %d catchup messages.", messageCount)
}



