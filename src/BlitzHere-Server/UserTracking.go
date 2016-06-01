//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
    )


//----------------------------------------------------------------------------------------
//                                                                 UpdateUserTrackingEvent
//----------------------------------------------------------------------------------------


func UpdateUserTrackingEvent(userID string, event *BlitzMessage.UserEvent) error {
    //Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    result, error := config.DB.Exec(
        "update UserEventTable set " +
        " (userID, timestamp, location, event, eventData) = " +
        " ($1, $2, row($3, $4, $5), $6, $7)" +
        " where userID = $8 and timestamp = $9;",
        &userID,
        event.Timestamp.NullTime(),
        event.Location.Coordinate.Latitude,
        event.Location.Coordinate.Longitude,
        event.Location.Placename,
        event.Event,
        pgsql.NullStringFromStringArray(event.EventData),
        &userID,
        event.Timestamp.NullTime(),
    )

    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }

    if error == nil && rowsUpdated > 0 {
//     Log.Debugf("Updated event row %s %v.",
//         userID, BlitzMessage.NullTimeFromTimestamp(event.Timestamp))
    } else {
        Log.Debugf("Inserting event %s %v: %v.",
           userID, event.Timestamp.NullTime(),
           error,
        )

        //  Insert instead --

        _, error := config.DB.Exec("insert into UserEventTable " +
            " (userID, timestamp, location, event, eventdata) values " +
            " ($1, $2, row($3, $4, $5), $6, $7);",
            &userID,
            event.Timestamp.NullTime(),
            event.Location.Coordinate.Latitude,
            event.Location.Coordinate.Longitude,
            event.Location.Placename,
            event.Event,
            pgsql.NullStringFromStringArray(event.EventData),
        )

        if error != nil {
            Log.Errorf("Error inserting event %s %v: %v.", userID, event.Timestamp.NullTime(), error)
        }
    }
    return error
}


//----------------------------------------------------------------------------------------
//                                                                          UserEventBatch
//----------------------------------------------------------------------------------------


func UpdateUserTrackingBatch(session *Session, userEvents *BlitzMessage.UserEventBatch,
    ) *BlitzMessage.ServerResponse {

    //  * Update each user event in the update request.
    Log.LogFunctionName()

    errorCount := 0
    var firstError error = nil
    var lastTimestamp *BlitzMessage.Timestamp = nil

    for i := 0; i < len(userEvents.UserEvents); i++ {
        error := UpdateUserTrackingEvent(session.UserID, userEvents.UserEvents[i])
        if error == nil {
            lastTimestamp = userEvents.UserEvents[i].Timestamp
        } else {
            errorCount++
            if firstError == nil { firstError = error }
            Log.Errorf("Error updating event %s %v: %v.",
                session.UserID,
                userEvents.UserEvents[i].Timestamp.NullTime(),
                error,
            )
        }
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    var message string

    if errorCount > 0 {
        Log.Errorf("Found %d errors on update.", errorCount)
        code = BlitzMessage.ResponseCode_RCServerWarning
        message = firstError.Error()
    }

    if  errorCount == len(userEvents.UserEvents) {
        code = BlitzMessage.ResponseCode_RCServerError
        if (message == "") { message = "No events to update" }
    }

    userTrackingResponse := &BlitzMessage.UserEventBatchResponse {
        LatestEventUpdate: lastTimestamp,
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
        ResponseType:       &BlitzMessage.ResponseType { UserEventBatchResponse: userTrackingResponse},
    }
    return response
}


//----------------------------------------------------------------------------------------
//                                                                       SaveDebugMessages
//----------------------------------------------------------------------------------------


func SaveDebugMessages(session *Session, debugMessage *BlitzMessage.DebugMessage,
    ) *BlitzMessage.ServerResponse {

    Log.LogFunctionName()

    var responseMessage string
    code := BlitzMessage.ResponseCode_RCSuccess

    for _, text := range(debugMessage.DebugText) {
        if text == "erase-profile" {
            row := config.DB.QueryRow("select EraseUserID($1);", session.UserID)
            var message string
            error := row.Scan(&message)
            if error != nil {
                message = fmt.Sprintf("Error erasing userID %s: %v.", session.UserID, error)
                code = BlitzMessage.ResponseCode_RCServerError
            } else {
                message = fmt.Sprintf("Erased userID %s. Message: '%s'.", session.UserID, message)
            }
            Log.Infof("%s", message)
            responseMessage = message
        }
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
        ResponseMessage: &responseMessage,
    }
    return response
}


