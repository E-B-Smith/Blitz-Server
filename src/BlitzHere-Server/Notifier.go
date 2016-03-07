//  Notifier  -  A daemon that wakes every so often to send any new push notifications
//               that need to sent.
//
//  E.B.Smith  -  March, 2014


package main


import (
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
    MessageTable.recipientID,
    MessageTable.messageText,
    DeviceTable.appID,
    DeviceTable.notificationToken,
    DeviceTable.appIsReleaseVersion
      from MessageTable
      join DeviceTable on DeviceTable.userID = MessageTable.recipientID
        where MessageTable.notificationDate is null
          and DeviceTable.notificationToken is not null
          and DeviceTable.appID is not null
        order by MessageTable.recipientID, MessageTable.creationDate

    and MessageTable.senderID <> MessageTable.recipientID

*/


func notifyTask() {
    //  If the user has an outstanding message send a notification.

    Log.LogFunctionName()

    defer func() {
        if panicInfo := recover(); panicInfo != nil {
            Log.LogStackWithError(panicInfo)
        }
        Log.Debugf("Exit NotifyTask.")
    } ()

    SendPulseReminders()

    notificationsToSend := 0
    notificationErrors := 0

    rows, error := config.DB.Query(
`       select distinct on (1)
            MessageTable.recipientID,
            MessageTable.messageText,
            MessageTable.actionURL,
            DeviceTable.appID,
            DeviceTable.notificationToken,
            DeviceTable.appIsReleaseVersion
              from MessageTable
              join DeviceTable on DeviceTable.userID = MessageTable.recipientID
                where MessageTable.notificationDate is null
                  and DeviceTable.notificationToken is not null
                  and DeviceTable.appID is not null
                order by MessageTable.recipientID, MessageTable.creationDate;`)
    defer rows.Close()
    if error != nil {
        Log.LogError(error)
        return
    }

    for rows.Next() {
        //  Send a notification for each pending message --

        var (
            recipientID     string
            messageText     string
            actionURL       sql.NullString
            appID           string
            notificationToken string
            appIsReleaseVersion bool
        )

        notificationsToSend++
        error = rows.Scan(&recipientID, &messageText, &actionURL, &appID, &notificationToken, &appIsReleaseVersion)
        if error != nil {
            notificationErrors++
            Log.LogError(error)
            continue
        }

        serviceType := ApplePushService.ServiceTypeDevelopment
        if appIsReleaseVersion {
            serviceType = ApplePushService.ServiceTypeProduction
        }

        notification := ApplePushService.Notification {
            BundleID:       appID,
            ServiceType:    serviceType,
            DeviceToken:    notificationToken,
            MessageText:    messageText,
            SoundName:      "Pulse.caf",
        }
        if actionURL.Valid && len(actionURL.String) > 0 {
            notification.OptionalKeys = map[string]string {
                "url":  actionURL.String,
            }
        }
        error := PushNotificationService.Send(&notification)
        if error != nil {
            notificationErrors++
            Log.LogError(error)
            continue
        }

        //  Mark the recipient as having been sent a message --

        _, error = config.DB.Exec(
            `update MessageTable set notificationDate = $1
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
    defer Log.Debugf("Exit Notifier")

    var shouldContinue bool = true
    for shouldContinue {
        var timer *time.Timer = time.NewTimer(time.Second*30)
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

