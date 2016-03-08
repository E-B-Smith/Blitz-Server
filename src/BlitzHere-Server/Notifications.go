//  Messages.go  -  Dispatch user messages.
//
//  E.B.Smith  -  March, 2015


package main


import (
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
//                                                                NotificationFetchRequest
//
//----------------------------------------------------------------------------------------


func NotificationFetchRequest(session *Session, fetch *BlitzMessage.NotificationUpdate) *BlitzMessage.ServerResponse {
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
        "select " +
            "messageID, "+
            "senderID, "+
            "recipientID, "+
            "creationDate, "+
            "notificationDate, "+
            "readDate, "+
            "messageType, "+
            "messageText, "+
            "actionIcon, "+
            "actionURL "+
            "from NotificationTable "+
            "  where recipientID = $1 "+
            "  and creationDate >  $2 "+
            "  and creationDate <= $3 "+
            "    order by creationDate;",
            session.UserID, startDate, stopDate)
    defer func() { if rows != nil { rows.Close(); } }()
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    var messageArray []*BlitzMessage.Notification
    for rows.Next() {
        var (
            messageID string
            senderID string
            recipientID string
            creationDate pq.NullTime
            notificationDate pq.NullTime
            readDate pq.NullTime
            messageType int
            messageText sql.NullString
            actionIcon sql.NullString
            actionURL sql.NullString
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
        )
        if error != nil {
            Log.LogError(error)
            warnings++
            continue
        }
        mt := BlitzMessage.NotificationType(messageType);
        message := BlitzMessage.Notification {
            MessageID:   &messageID,
            SenderID:    &senderID,
            Recipients:  []string{recipientID},
            MessageType: &mt,
        }
        if creationDate.Valid { message.CreationDate = BlitzMessage.TimestampFromTime(creationDate.Time) }
        if notificationDate.Valid {message.NotificationDate = BlitzMessage.TimestampFromTime(notificationDate.Time) }
        if readDate.Valid { message.ReadDate = BlitzMessage.TimestampFromTime(readDate.Time) }
        if messageText.Valid { message.MessageText = &messageText.String }
        if actionIcon.Valid { message.ActionIcon = &actionIcon.String }
        if actionURL.Valid { message.ActionURL = &actionURL.String }
        messageArray = append(messageArray, &message)
    }

    Log.Debugf("Found %d message (%d warnings) in range %v to %v.", len(messageArray), warnings, startDate, stopDate)

    if len(messageArray) == 0 && warnings > 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, errors.New("Messages are not available now."))
    }
    messageUpdate := BlitzMessage.NotificationUpdate {
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
        Response:       &BlitzMessage.ResponseType { NotificationUpdate: &messageUpdate },
    }
    return response
}



//----------------------------------------------------------------------------------------
//                                                                 SendNotificationMessage
//----------------------------------------------------------------------------------------


func SendNotificationMessage(sender string, recipients []string, message string,
                 messageType BlitzMessage.NotificationType, actionIcon string, actionURL string) {
    Log.LogFunctionName()

    for _, recipient := range recipients {
        if sender == recipient { continue; }

        _, error := config.DB.Exec("insert into NotificationTable "+
            "(messageID, " +
            " senderID, "  +
            " recipientID,"+
            " creationDate,"+
            " messageType,"+
            " messageText,"+
            " actionIcon, "+
            " actionURL  "  +
            ") values ($1, $2, $3, $4, $5, $6, $7, $8); ",
            Util.NewUUIDString(),
            sender,
            recipient,
            time.Now(),
            messageType,
            message,
            actionIcon,
            actionURL)

        if error != nil {
            Log.Errorf("Error inserting message: %+v.", error)
        }
    }
}


//----------------------------------------------------------------------------------------
//
//                                                                 NotificationSendRequest
//
//----------------------------------------------------------------------------------------


func NotificationSendRequest(session *Session,
                             sendMessage *BlitzMessage.NotificationUpdate,
                             ) *BlitzMessage.ServerResponse {
    //
    //  * Save each new message to the database.
    //

    Log.LogFunctionName()

    messagesSent := 0
    for _, message := range sendMessage.Messages {
        Log.Debugf("Message %d has %d recipients.", messagesSent+1, len(message.Recipients))
        for _, recipientID := range message.Recipients {
            _, error := config.DB.Exec("insert into NotificationTable "+
                "(messageID, " +
                " senderID, "  +
                " recipientID,"+
                " creationDate,"+
                " messageType,"+
                " messageText,"+
                " actionIcon, "+
                " actionURL  "  +
                ") values ($1, $2, $3, $4, $5, $6, $7, $8); ",
                message.MessageID,
                message.SenderID,
                recipientID,
                BlitzMessage.NullTimeFromTimestamp(message.CreationDate),
                message.MessageType,
                message.MessageText,
                message.ActionIcon,
                message.ActionURL)

            if error == nil {
                messagesSent++
            } else {
                Log.Errorf("Error inserting message: %v. MessageID: %s From: %s To: %s.",
                    error, *message.MessageID, *message.SenderID, recipientID)
            }
        }
    }

    Log.Debugf("Received %d message bundles, sent %d messages.", len(sendMessage.Messages), messagesSent)

    messageResponse := &BlitzMessage.NotificationUpdate {
        Timespan: sendMessage.Timespan,
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        Response:       &BlitzMessage.ResponseType { NotificationUpdate: messageResponse },
    }

    return response
}

