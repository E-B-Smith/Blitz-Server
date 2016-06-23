

//----------------------------------------------------------------------------------------
//
//                                                       BlitzHere-Server : DownloadApp.go
//                                             Send a text message with the download link.
//
//                                                                 E.B. Smith, March, 2015
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "time"
    "strings"
    "strconv"
    "net/http"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


func DownloadAppRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    //  Save the name and phone number for later.
    //  Text a short-link that finger prints phone and redirects to the app store.
    //  Otherwise send back error.

    respondWithErrorType := func(errorType string) {
        Log.Errorf("Error type '%s'.", errorType)
        url := fmt.Sprintf("/index.html?action=%s#download", errorType)
        http.Redirect(writer, httpRequest, url, 303)
    }

    if httpRequest.Method == "GET" {
        respondWithErrorType("")
        return
    }

    error := httpRequest.ParseForm()
    if error != nil {
        respondWithErrorType("error-error")
        return
    }

    name := strings.TrimSpace(httpRequest.PostFormValue("name"))
    if name == "" {
        respondWithErrorType("error-name")
        return
    }
    phone := Util.StringIncludingCharactersInSet(httpRequest.PostFormValue("phone"), "0123456789")
    if len(phone) != 10 {
        respondWithErrorType("error-phone")
        return
    }

    Log.Debugf("Validated '%s' '%s'.", name, phone)

    //  Send url like eksprt://blitzhere.com/blitzhere?action=confirm&redirect=<appstore>&code=<code>&contact=<phone>

    row := config.DB.QueryRow(
        `select userID from UserContactTable
            where ContactType = $1 and contact = $2 and isVerified = true;`,
        BlitzMessage.ContactType_CTPhoneSMS,
        phone,
    )
    var userID sql.NullString
    error = row.Scan(&userID)
    if error != nil || ! userID.Valid {

        //  Create a new user --

        contactInfo := BlitzMessage.ContactInfo {
            ContactType:    BlitzMessage.ContactType_CTPhoneSMS.Enum(),
            Contact:        proto.String(phone),
            IsVerified:     proto.Bool(false),
        }

        userID.String = Util.NewUUIDString()
        userID.Valid  = true

        userProfile := BlitzMessage.UserProfile {
            UserID:         proto.String(userID.String),
            UserStatus:     BlitzMessage.UserStatus_USInvited.Enum(),
            CreationDate:   BlitzMessage.TimestampPtr(time.Now()),
            LastSeen:       BlitzMessage.TimestampPtr(time.Now()),
            Name:           proto.String(name),
            ContactInfo:    []*BlitzMessage.ContactInfo { &contactInfo },
        }

        error = UpdateProfile(&userProfile)
        if error != nil {
            respondWithErrorType("error-error")
            return
        }

    }

    longSecret := Util.NewUUIDString()
    i, _ := strconv.ParseInt(longSecret[0:4], 16, 32)
    confirmCode := fmt.Sprintf("%05d", i)
    var result sql.Result
    result, error = config.DB.Exec(
        `update UserContactTable set
            code = $1, codeDate = current_timestamp
            where userid = $2
              and contacttype = $3
              and contact = $4`,
        confirmCode,
        userID.String,
        BlitzMessage.ContactType_CTPhoneSMS,
        phone,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
    }

    fullURL := fmt.Sprintf("%s/?action=confirm&code=%s&contact=%s&redirect=%s",
        config.AppLinkURL,
        confirmCode,
        phone,
        config.AppStoreURL,
    )
    Log.Debugf("Full URL is '%s'.", fullURL)
    shortURL, error := LinkShortner_ShortLinkFromLink(fullURL)
    if error != nil {
        respondWithErrorType("error-error")
        return
    }

    message := fmt.Sprintf("Welcome! Download %s:\n%s",
        config.AppName,
        shortURL,
    )
    Util.SendSMS(phone, message)

    respondWithErrorType("complete")
}


