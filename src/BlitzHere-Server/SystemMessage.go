//  Status-Server  -  Serves an html status page for the BeingHappy-Server and HappyLabs-Server
//
//  E.B.Smith  -  March, 2015


package main


import (
    "fmt"
    "time"
    "strings"
    "net/http"
    "html/template"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                SystemMessageFormRequest
//
//----------------------------------------------------------------------------------------


func SystemMessageFormRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Send or process the sign-up form --
    Log.LogFunctionName()
    config.MessageCount++

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, httpRequest, "error.html", 303)
        }
    } ()

    var templateMap struct {
        AppName     string
        MessageText string
        ActionIcon  string
        ActionURL   template.URL
    }
    templateMap.AppName = config.AppName

    var error error
    if httpRequest.Method == "GET" {
        error = config.Template.ExecuteTemplate(writer, "SendMessage.html", templateMap)
        if error != nil { Log.LogError(error) }
        return
    }

    var wasError bool = false
    error = httpRequest.ParseForm()
    if error != nil {
        wasError = true
        templateMap.MessageText = "Bad form request."
    }

    messageText := strings.TrimSpace(httpRequest.PostFormValue("messageText"))
    if messageText == "" {
        wasError = true
        templateMap.MessageText = "A message body is required."
    }
    actionIcon := strings.TrimSpace(httpRequest.PostFormValue("actionIcon"))
    if actionIcon == "" {
        actionIcon = "Icon"
    }
    actionURL := strings.TrimSpace(httpRequest.PostFormValue("actionURL"))

    if wasError {
        Log.Warningf("Submit error: %v.", templateMap)
        error = config.Template.ExecuteTemplate(writer, "SendMessage.html", templateMap)
        if error != nil { Log.LogError(error) }
        return
    }

    //  Insert rows --

    rows, error := config.DB.Query("select userID from UserTable where userStatus >= $1;", BlitzMessage.UserStatus_USActive);
    defer pgsql.CloseRows(rows);
    if error != nil {
        Log.LogError(error)
        templateMap.MessageText = fmt.Sprintf("System error: %v.", error)
        error = config.Template.ExecuteTemplate(writer, "SendMessage.html", templateMap)
        if error != nil { Log.LogError(error) }
        return
    }

    messageID   := Util.NewUUIDString();
    messageDate := time.Now();

    rowCount := 0
    errorCount := 0

    for rows.Next() {
        rowCount++
        var recipientID string
        rows.Scan(&recipientID)
        _, error := config.DB.Exec(
            "insert into messagetable "+
            "(messageID        "+
            ",senderID         "+
            ",recipientID      "+
            ",creationDate     "+
            ",messageType      "+
            ",messageText      "+
            ",actionIcon       "+
            ",actionURL        "+
            ") values ($1, $2, $3, $4, $5, $6, $7, $8);",
            messageID,
            BlitzMessage.Default_Globals_SystemUserID,
            recipientID,
            messageDate,
            BlitzMessage.NotificationType_NTSystem,
            messageText,
            actionIcon,
            actionURL,
        )
        if error != nil {
            Log.LogError(error)
            errorCount++
        }
    }

    //  Send back the result --

    templateMap.MessageText = fmt.Sprintf("Sent %d messages with %d errors.", rowCount, errorCount)
    error = config.Template.ExecuteTemplate(writer, "SendMessage.html", templateMap)
    if error != nil { Log.LogError(error) }
}


