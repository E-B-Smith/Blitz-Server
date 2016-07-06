

//----------------------------------------------------------------------------------------
//
//                                                          BlitzHere-Server : WebForms.go
//                                                                   Handle HTTP Web Forms
//
//                                                                  E.B. Smith, July, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    // "errors"
    "strings"
    "net/http"
    "database/sql"
    "violent.blue/GoKit/Log"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                            WebListUsers
//
//----------------------------------------------------------------------------------------


func WebListUsers(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select userid, name, backgroundSummary from usertable
            where userstatus >= 5 order by lastSeen desc;`,
    )
    if error != nil {
        return
    }

    type UserStruct struct {
        UserID      sql.NullString
        Name        sql.NullString
        Background  sql.NullString
    }
    type ListUsers struct {
        AppName     string
        Users       []UserStruct
    }
    listUsers := ListUsers {
        AppName:        config.AppName,
    }

    for rows.Next() {
        var user UserStruct
        error = rows.Scan(&user.UserID, &user.Name, &user.Background)
        if error != nil {
            Log.LogError(error)
            continue
        }
        if len(user.Background.String) > 20 {
            user.Background.String = user.Background.String[:20] + "..."
        }
        listUsers.Users = append(listUsers.Users, user)
    }

    error = config.Template.ExecuteTemplate(writer, "ListUsers.html", listUsers)
    if error != nil {
        Log.LogError(error)
    }
}




//----------------------------------------------------------------------------------------
//
//                                                                        WebUpdateProfile
//
//----------------------------------------------------------------------------------------


func WebUpdateProfile(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    updateProfile := struct {
        AppName         string
        ErrorMessage    string
        Profile         *BlitzMessage.UserProfile
        Expertise       string
    } {
        AppName: config.AppName,
    }

    defer func() {
        up := &updateProfile
        error := config.Template.ExecuteTemplate(writer, "UpdateProfile.html", up)
        if error != nil {
            Log.LogError(error)
        }
    } ()

    var error error
    if httpRequest.Method == "GET" {
        userID := httpRequest.URL.Query().Get("uid")
        updateProfile.Profile = ProfileForUserID(userID, userID)
        if updateProfile.Profile == nil {
            updateProfile.ErrorMessage = fmt.Sprintf("Invalid UserID '%s'.", userID)
        }
        for _, tag := range updateProfile.Profile.EntityTags {
            if ! strings.HasPrefix(*tag.TagName, ".") {
                updateProfile.Expertise += fmt.Sprintf("'%s', ", *tag.TagName)
            }
        }
        updateProfile.Expertise = strings.TrimRight(updateProfile.Expertise, ", ")
        return
    }


    error = httpRequest.ParseForm()
    if error != nil {
        Log.LogError(error)
        updateProfile.ErrorMessage = error.Error()
        return
    }

    userID := httpRequest.PostFormValue("UserID")
    updateProfile.Profile = ProfileForUserID(userID, userID)
    if updateProfile.Profile == nil {
        updateProfile.ErrorMessage = fmt.Sprintf("Invalid UserID '%s'.", userID)
    }

    updateProfile.Profile.Name = proto.String(httpRequest.PostFormValue("Name"))
    updateProfile.Profile.BackgroundSummary =
        proto.String(httpRequest.PostFormValue("BackgroundSummary"))


    //  Update and send back the result --

    if false {
        error = UpdateProfile(updateProfile.Profile)
        if error != nil {
            Log.LogError(error)
            updateProfile.ErrorMessage = error.Error()
        }
    }

    updateProfile.ErrorMessage = "User updated."
}

