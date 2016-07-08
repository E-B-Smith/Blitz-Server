

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
    "errors"
    "strings"
    "net/http"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                             WebUserList
//
//----------------------------------------------------------------------------------------


func WebUserList(writer http.ResponseWriter, httpRequest *http.Request) {
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
        const bLen = 30
        if len(user.Background.String) > bLen {
            user.Background.String = user.Background.String[:bLen] + "..."
        }
        listUsers.Users = append(listUsers.Users, user)
    }

    error = config.Template.ExecuteTemplate(writer, "UserList.html", listUsers)
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
            return
        }
        if updateProfile.Profile.EditProfileID != nil  &&
           len(*updateProfile.Profile.EditProfileID) > 10 {
            userID = *updateProfile.Profile.EditProfileID
            updateProfile.Profile = ProfileForUserID(userID, userID)
        }
        for _, tag := range updateProfile.Profile.EntityTags {
            if ! strings.HasPrefix(*tag.TagName, ".") {
                updateProfile.Expertise += fmt.Sprintf("%s, ", *tag.TagName)
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

    //  Save the dot tags

    dotTags := make([]*BlitzMessage.EntityTag, 0, 10)
    for _, tag := range updateProfile.Profile.EntityTags {
        if strings.HasPrefix(*tag.TagName, ".") {
            dotTags = append(dotTags, tag)
        }
    }

    newTags := strings.Split(httpRequest.PostFormValue("Expertise"), ",")
    cleanTags := make([]string, 0, len(newTags))
    for _, tag := range newTags {
        cleanTag := ""
        tArray := strings.Split(tag, " ")
        for _, tagPart := range tArray {
            if len(tagPart) > 0 { cleanTag += " " + tagPart }
        }
        cleanTag = strings.ToLower(strings.TrimSpace(cleanTag))
        if len(cleanTag) > 0 {
            cleanTags = append(cleanTags, cleanTag)
        }
    }

    for _, tag := range cleanTags {
        entityTag := BlitzMessage.EntityTag {
            TagName:        proto.String(tag),
            UserHasTagged:  proto.Bool(true),
            TagCount:       proto.Int32(1),
        }
        dotTags = append(dotTags, &entityTag)
    }
    updateProfile.Profile.EntityTags = dotTags

    //  Update and send back the result --

    if true {
        error = UpdateProfile(updateProfile.Profile)
        if error != nil {
            Log.LogError(error)
            updateProfile.ErrorMessage = error.Error()
            return
        }

        //  Update expert status:

        isExpert := httpRequest.PostFormValue("IsExpert")
        Log.Debugf("IsExpert: %s.", isExpert)
        isExpertBool := false
        if isExpert == "IsExpert" {
            isExpertBool = true
        }
        var result sql.Result
        result, error = config.DB.Exec(
            `update UserTable set isExpert = $1 where userID = $2;`,
            isExpertBool,
            userID,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil {
            Log.LogError(error)
            updateProfile.ErrorMessage = error.Error()
            return
        }
        updateProfile.Profile.IsExpert = &isExpertBool

        //  Expert approve update?

        isApproved := httpRequest.PostFormValue("IsApproved")
        if isApproved == "IsApproved" {
            row := config.DB.QueryRow(
                `select ApproveEditProfile($1);`,
                userID,
            )
            var result sql.NullString
            error = row.Scan(&result)
            if error == nil && result.String != "Approved" {
                error = errors.New(result.String)
            }
            if error != nil {
                Log.LogError(error)
                updateProfile.ErrorMessage = error.Error()
                return
            }
        }
    }

    for _, tag := range updateProfile.Profile.EntityTags {
        if ! strings.HasPrefix(*tag.TagName, ".") {
            updateProfile.Expertise += fmt.Sprintf("%s, ", *tag.TagName)
        }
    }
    updateProfile.Expertise = strings.TrimRight(updateProfile.Expertise, ", ")

    updateProfile.ErrorMessage = "User updated."
}

