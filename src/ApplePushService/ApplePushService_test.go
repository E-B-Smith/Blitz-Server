//  ApplePush_test  -  Test sending an APN message.
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "fmt"
    "time"
    "sync"
    "testing"
    "math/rand"
    "violent.blue/GoKit/Log"
)


func notificationTester(t *testing.T, pushService Service, tagString string) {
    notification := Notification{

        BundleID:       "io.beinghappy.happypulse-d",
        ServiceType:    ServiceTypeProduction,
        DeviceToken:    "0409DB9F393F323BF080F49A3F4C2CDDE104B6793E3B846DBB087E53AAA7028A",
/*
        BundleID:       "io.beinghappy.beinghappy-d",
        ServiceType:    ServiceTypeDevelopment,
        DeviceToken:    "YgIuMMNzNGMZ1GxU/kNDgHByWbw7PZp7ZqliVxH+DBg=",
*/
/*
        BundleID:       "io.beinghappy.beinghappy",
        ServiceType:    ServiceTypeProduction,
        DeviceToken:    "BkeEMGn3LHED457CKAHHKb45mZ6GTVU3Xh0zEC2DXmU=",
*/
        MessageText:    "Hey1",
    }

    for i := 1; i < 3; i++ {
        notification.MessageID = 0
        notification.MessageText = fmt.Sprintf("%s: Hey %d!", tagString, i)
        error := pushService.Send(&notification)
        if error == nil {
            Log.Debugf("Tried push: '%s'.", notification.MessageText)
        } else  {
            t.Errorf("Unexpected error: %v.", error)
        }
        var waittime time.Duration = time.Second*time.Duration(rand.Float64()*10.0)
        time.Sleep(waittime)
    }
}


func TestApplePush(t *testing.T) {
    Log.LogLevel = Log.LevelAll

    pushService := NewService()
    error := pushService.Start()
    if error != nil { t.Errorf("Unexpected error: %v.", error) }


    var waiter sync.WaitGroup
    waiter.Add(2)

    walltime := time.Now().Format("15:04")
    go func(startTime string) {
        notificationTester(t, pushService, fmt.Sprintf("%s #1", startTime))
        waiter.Done()
    } (walltime)

    go func(startTime string) {
        notificationTester(t, pushService, fmt.Sprintf("%s #2", startTime))
        waiter.Done()
    } (walltime)

    Log.Debugf("\nWaiting before sending again...")
    time.Sleep(30 * time.Second)
    waiter.Add(2)

    walltime = time.Now().Format("15:04")
    go func(startTime string) {
        notificationTester(t, pushService, fmt.Sprintf("%s #3", startTime))
        waiter.Done()
    } (walltime)

    go func(startTime string) {
        notificationTester(t, pushService, fmt.Sprintf("%s #4", startTime))
        waiter.Done()
    } (walltime)

    waiter.Wait()

    statusArray := pushService.Status()
    Log.Debugf("\n\n>>> All notifications sent?  Status# (%d)", len(statusArray))
    for _, status := range statusArray {
        Log.Debugf("Status: %+v.", status)
    }

    Log.Debugf("\n\n>>> Waiting before stopping.")
    time.Sleep(15 * time.Second)
    Log.Debugf("\n\n>>> Stopping...")
    pushService.Stop()
    statusArray = pushService.Status()
    for _, status := range statusArray {
        Log.Debugf("Status: %+v.", status)
    }

    if error != nil { t.Errorf("Unexpected error: %v.", error) }
    Log.Debugf("Done with test.")
}


