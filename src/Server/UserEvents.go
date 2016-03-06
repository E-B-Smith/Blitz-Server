//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "net/http"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "happiness"
    )


//----------------------------------------------------------------------------------------
//                                                                         UpdateUserEvent
//----------------------------------------------------------------------------------------


func UpdateUserEvent(userID string, event *happiness.UserEvent) error {
    //Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    result, error := config.DB.Exec("update UserEventTable set " +
        " (userID, timestamp, location, event, eventData) = " +
        " ($1, $2, row($3, $4, $5), $6, $7)" +
        " where userID = $8 and timestamp = $9;",
        &userID,
        happiness.NullTimeFromTimestamp(event.Timestamp),
        event.Location.Latitude,
        event.Location.Longitude,
        event.Location.Placename,
        event.Event,
        pgsql.NullStringFromStringArray(event.EventData),
        &userID,
        happiness.NullTimeFromTimestamp(event.Timestamp));

    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }

    if error == nil && rowsUpdated > 0 {
//     Log.Debugf("Updated event row %s %v.",
//         userID, happiness.NullTimeFromTimestamp(event.Timestamp))
    } else {
        Log.Debugf("Inserting event %s %v: %v.",
           userID, happiness.NullTimeFromTimestamp(event.Timestamp), error)

        //  Insert instead --

        _, error := config.DB.Exec("insert into UserEventTable " +
            " (userID, timestamp, location, event, eventdata) values " +
            " ($1, $2, row($3, $4, $5), $6, $7);",
            &userID,
            happiness.NullTimeFromTimestamp(event.Timestamp),
            event.Location.Latitude,
            event.Location.Longitude,
            event.Location.Placename,
            event.Event,
            pgsql.NullStringFromStringArray(event.EventData));

        if error != nil {
            Log.Errorf("Error inserting event %s %v: %v.", userID, happiness.NullTimeFromTimestamp(event.Timestamp), error)
        }
    }
    return error
}


//----------------------------------------------------------------------------------------
//                                                                        UpdateUserEvents
//----------------------------------------------------------------------------------------


func UpdateUserEvents(writer http.ResponseWriter, userID string, userEvents *happiness.UserEventsRequest) {
    //  * Update each user event in the update request.
    Log.LogFunctionName()

    errorCount := 0
    var firstError error = nil
    var lastTimestamp *happiness.Timestamp = nil

    for i := 0; i < len(userEvents.UserEvents); i++ {
        error := UpdateUserEvent(userID, userEvents.UserEvents[i])
        if error == nil {
            lastTimestamp = userEvents.UserEvents[i].Timestamp
        } else {
            errorCount++
            if firstError == nil { firstError = error }
            Log.Errorf("Error updating event %s %v: %v.",
                userID, happiness.NullTimeFromTimestamp(userEvents.UserEvents[i].Timestamp), error)
        }
    }

    code := happiness.ResponseCode_RCSuccess
    var message string

    if errorCount > 0 {
        Log.Errorf("Found %d errors on update.", errorCount)
        code = happiness.ResponseCode_RCServerWarning
        message = firstError.Error()
    }

    if  errorCount == len(userEvents.UserEvents) {
        code = happiness.ResponseCode_RCServerError
        if (message == "") { message = "No events to update" }
    }

    userEventResponse := &happiness.UserEventsResponse {
        LatestEventUpdate: lastTimestamp,
    }

    response := &happiness.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
        Response:           &happiness.ServerResponse_UserEventResponse { UserEventResponse: userEventResponse},
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}


//----------------------------------------------------------------------------------------
//                                                                       SaveDebugMessages
//----------------------------------------------------------------------------------------


func SaveDebugMessages(writer http.ResponseWriter, userID string, debugMessage *happiness.DebugMessage) {
    Log.LogFunctionName()

    var responseMessage string
    code := happiness.ResponseCode_RCSuccess

    for _, text := range(debugMessage.DebugText) {
        if text == "erase-profile" {
            row := config.DB.QueryRow("select EraseUserID($1);", userID)
            var message string
            error := row.Scan(&message)
            if error != nil {
                message = fmt.Sprintf("Error erasing userID %s: %v.", userID, error)
                code = happiness.ResponseCode_RCServerError
            } else {
                message = fmt.Sprintf("Erased userID %s. Message: '%s'.", userID, message)
            }
            Log.Infof("%s", message)
            responseMessage = message
        }
    }

    response := &happiness.ServerResponse {
        ResponseCode: &code,
        ResponseMessage: &responseMessage,
    }
    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}


