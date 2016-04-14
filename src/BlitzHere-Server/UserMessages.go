//  Messages.go  -  Dispatch user messages.
//
//  E.B.Smith  -  March, 2015


package main


import (
    "fmt"
    "time"
    "errors"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "BlitzMessage"
)


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
        `select
            messageID,
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
                from UserMessageTable
                  where recipientID = $1
                  and creationDate >  $2
                  and creationDate <= $3
                    order by creationDate;`,
            session.UserID, startDate, stopDate)
    defer func() { if rows != nil { rows.Close(); } }()
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    var messageArray []*BlitzMessage.UserMessage
    for rows.Next() {
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
            warnings++
            continue
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


        messageArray = append(messageArray, &message)
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


func SendBlitzUserMessage(message *BlitzMessage.UserMessage) error {
    Log.LogFunctionName()

    if message.SenderID == nil {
        return fmt.Errorf("No sender ID")
    }

    var recipients []string
    if message.ConversationID == nil {
        recipients = message.Recipients
        recipients = append(recipients, *message.SenderID)
    } else {
        recipients = MembersForConversationID(*message.ConversationID)
    }
    message.Recipients = recipients

    for _, recipientID := range recipients {
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
                conversationID)
                values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
            message.MessageID,
            message.SenderID,
            recipientID,
            BlitzMessage.NullTimeFromTimestamp(message.CreationDate),
            message.MessageType,
            message.MessageText,
            message.ActionIcon,
            message.ActionURL,
            message.ConversationID)

        if error != nil {
            Log.Errorf("Error inserting message: %v. MessageID: %s From: %s To: %s.",
                error, *message.MessageID, *message.SenderID, recipientID)
        }
    }

    globalMessagePusher.PushMessage(message)
    return nil
}


//----------------------------------------------------------------------------------------
//                                                                         SendUserMessage
//----------------------------------------------------------------------------------------


func SendUserMessage(
        sender string,
        recipients []string,
        message string,
        messageType BlitzMessage.UserMessageType,
        actionIcon string,
        actionURL string) {
    Log.LogFunctionName()

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

    SendBlitzUserMessage(blitzMessage)
}


//----------------------------------------------------------------------------------------
//
//                                                                  UserMessageSendRequest
//
//----------------------------------------------------------------------------------------


func UserMessageSendRequest(
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

