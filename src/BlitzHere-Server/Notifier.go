//  Notifier  -  A daemon that wakes every so often to send any new push notifications
//               that need to sent.
//
//  E.B.Smith  -  March, 2014


package main


import (
    "fmt"
    "time"
    "database/sql"
    "ApplePushService"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//
//                                                                                Notifier
//
//----------------------------------------------------------------------------------------


/*

select distinct on (1)
    UserMessageTable.recipientID,
    UserMessageTable.messageText,
    DeviceTable.appID,
    DeviceTable.notificationToken,
    DeviceTable.appIsReleaseVersion
      from UserMessageTable
      join DeviceTable on DeviceTable.userID = UserMessageTable.recipientID
        where UserMessageTable.notificationDate is null
          and DeviceTable.notificationToken is not null
          and DeviceTable.appID is not null
        order by UserMessageTable.recipientID, UserMessageTable.creationDate

    and UserMessageTable.senderID <> UserMessageTable.recipientID

*/


func UnreadCountForUserID(userID string) int64 {
    Log.LogFunctionName()
    row := config.DB.QueryRow(
        `select count(*) from usermessagetable
            where recipientid = $1
              and senderID <> $1
              and readdate is null
              and conversationid is not null;`,
        userID,
    )
    var count int64
    error := row.Scan(&count)
    if error != nil { Log.LogError(error) }
    return count
}


func notifyTask() {
    //  If the user has an outstanding message send a notification.

    Log.LogFunctionName()

    defer func() {
        if panicInfo := recover(); panicInfo != nil {
            Log.LogStackWithError(panicInfo)
        }
        Log.Debugf("Exit NotifyTask.")
    } ()

    notificationsToSend := 0
    notificationErrors := 0

    rows, error := config.DB.Query(
`       select distinct on (1)
            UserMessageTable.recipientID,
            UserMessageTable.messageText,
            UserMessageTable.actionURL,
            UserMessageTable.conversationID,
            UserMessageTable.senderID,
            UserMessageTable.messageID,
            DeviceTable.appID,
            DeviceTable.notificationToken,
            DeviceTable.appIsReleaseVersion
              from UserMessageTable
              join DeviceTable on DeviceTable.userID = UserMessageTable.recipientID
                where UserMessageTable.notificationDate is null
                  and UserMessageTable.recipientID <> UserMessageTable.senderID
                  and DeviceTable.notificationToken is not null
                  and DeviceTable.appID is not null
                order by UserMessageTable.recipientID, UserMessageTable.creationDate;`)
    if error != nil {
        Log.LogError(error)
        return
    }
    defer rows.Close()

    for rows.Next() {
        //  Send a notification for each pending message --

        var (
            recipientID     string
            messageText     string
            actionURL       sql.NullString
            conversationID  sql.NullString
            senderID        sql.NullString
            messageID        sql.NullString
            appID           string
            notificationToken string
            appIsReleaseVersion bool
        )

        notificationsToSend++
        error = rows.Scan(
            &recipientID,
            &messageText,
            &actionURL,
            &conversationID,
            &senderID,
            &messageID,
            &appID,
            &notificationToken,
            &appIsReleaseVersion,
        )
        if error != nil {
            notificationErrors++
            Log.LogError(error)
            continue
        }

        serviceType := ApplePushService.ServiceTypeDevelopment
        if appIsReleaseVersion {
            serviceType = ApplePushService.ServiceTypeProduction
        }

        name, error := NameForUserID(senderID.String)
        if len(name) > 0 {
            messageText = fmt.Sprintf("%s says:\n%s", name ,messageText)
        }

        notification := ApplePushService.Notification {
            BundleID:       appID,
            ServiceType:    serviceType,
            DeviceToken:    notificationToken,
            MessageText:    messageText,
            OptionalKeys:   map[string]string { "sound": "NewMessage.caf"},
        }
        if len(conversationID.String) > 0 {
            notification.OptionalKeys["conversationID"] = conversationID.String
        }
        if len(senderID.String) > 0 {
            notification.OptionalKeys["senderID"] = senderID.String
        }
        if actionURL.Valid && len(actionURL.String) > 0 {
            notification.OptionalKeys["url"] = actionURL.String
        }
        if messageID.Valid && len(messageID.String) > 0 {
            notification.OptionalKeys["messageID"] = messageID.String
        }
        badgeCount := UnreadCountForUserID(recipientID)
        notification.OptionalKeys["badge"] = fmt.Sprintf("%d", badgeCount)

        error = PushNotificationService.Send(&notification)
        if error != nil {
            notificationErrors++
            Log.LogError(error)
            continue
        }

        //  Mark the recipient as having been sent a message --

        _, error = config.DB.Exec(
            `update UserMessageTable set notificationDate = $1
                where notificationDate is null and recipientID = $2;`,
            time.Now(), recipientID)
        if error != nil { Log.LogError(error) }
    }

    Log.Debugf("Notifications to send: %d notification errors: %d.",
        notificationsToSend, notificationErrors)

    status := PushNotificationService.Status()
    Log.Debugf("ApplePushService Status:")
    for _, stat := range(status) {
        Log.Debugf("%s", stat.String())
    }
}


var notifierChannel chan bool


func notifier() {
    Log.LogFunctionName()
    defer Log.Debugf("=> Exit Notifier <=")

    var shouldContinue bool = true
    for shouldContinue {
        var timer *time.Timer = time.NewTimer(time.Second*2)
        select {
            case shouldContinue = <- notifierChannel:
                Log.Debugf("Notifier should continue: %v.", shouldContinue)

            case <- timer.C:
                notifyTask()
        }
    }
}


func StartNotifier() {
    Log.LogFunctionName()
    notifierChannel = make(chan bool)
    go notifier()
}


func StopNotifier() {
    Log.LogFunctionName()
    notifierChannel <- false
}

