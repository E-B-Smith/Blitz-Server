//  AppOptions.go  -  Options for the app.
//
//  E.B.Smith  -  February, 2016.


package main


import (
    "violent.blue/GoKit/Log"
    "happiness"
)


//----------------------------------------------------------------------------------------
//
//                                                                         PulseAppOptions
//
//----------------------------------------------------------------------------------------


func AppOptionsForSession(session *Session) *happiness.AppOptions {

    var options *happiness.AppOptions

    if session == nil || session.Device.AppID == nil {
        return nil
    }

    switch (*session.Device.AppID) {
    case "io.beinghappy.happypulse":
    case "io.beinghappy.happypulse-d":
        pulseOptions := happiness.PulseOptions {
            MinimumEmotionsToScore:         Int32PtrFromInt32(3),
            MinimumMembersForPulse:         Int32PtrFromInt32(2),
            MinimumResponsesForPulseStats:  Int32PtrFromInt32(1),
        }
        options = &happiness.AppOptions {
            Options:        &happiness.AppOptions_PulseOptions { PulseOptions: &pulseOptions },
        }
    case "io.beinghappy.beinghappy":
    case "io.beinghappy.beinghappy-d":
        beingHappyOptions := happiness.BeingHappyOptions {
            MinimumEmotionsToScore:         Int32PtrFromInt32(3),
        }
        options = &happiness.AppOptions {
            Options:        &happiness.AppOptions_BeingHappyOptions { BeingHappyOptions: &beingHappyOptions },
        }
    default:
        Log.Errorf("Invalid app for options: '%s'.", session.Device.AppID)
    }

    return options
}
