

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
    "strconv"
    "strings"
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/ServerUtil"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                        AdminFormRequest
//
//----------------------------------------------------------------------------------------


func AdminFormRequest(writer http.ResponseWriter, request *http.Request) {
    Log.LogFunctionName()

    var templateMap struct {
        AppName     string
    }
    templateMap.AppName = config.AppName

    var error error
    error = config.Template.ExecuteTemplate(writer, "Admin.html", templateMap)
    if error != nil { panic(error) }
}


//----------------------------------------------------------------------------------------
//
//                                                                             WebUserList
//
//----------------------------------------------------------------------------------------


func WebUserList(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select
            userid,
            name,
            backgroundSummary,
            editprofileid,
            isExpert,
            lastSeen
            from usertable
            where userstatus >= 5 order by lastSeen desc;`,
    )
    if error != nil {
        return
    }

    type UserStruct struct {
        LastSeen    string
        Annotation  string
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
        var editprofileid sql.NullString
        var isExpert sql.NullBool
        var lastSeen pq.NullTime
        error = rows.Scan(
            &user.UserID,
            &user.Name,
            &user.Background,
            &editprofileid,
            &isExpert,
            &lastSeen,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }
        user.LastSeen = lastSeen.Time.String()

        const bLen = 30
        if len(user.Background.String) > bLen {
            user.Background.String = user.Background.String[:bLen] + "..."
        }

        if len(editprofileid.String) > 10 {
            user.Annotation = "Approve Edit"
        } else if ! isExpert.Bool {
            row := config.DB.QueryRow(
                `select count(*) from entitytagtable
                    where userid = $1
                      and entityid = $1::uuid
                      and entitytype = 1
                      and entitytag = '.appliedexpert'`,
                user.UserID.String,
            )
            var count int64
            error = row.Scan(&count)
            if error == nil && count > 0 {
                user.Annotation = "Applied"
            }
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


    fixNullTimespan := func(ts *BlitzMessage.Timespan) *BlitzMessage.Timespan {
        if ts == nil {
            ts = &BlitzMessage.TimespanZero
        } else {
            if ts.StartTimestamp == nil {
                ts.StartTimestamp = &BlitzMessage.TimestampZero
            }
            if ts.StopTimestamp == nil {
                ts.StopTimestamp = &BlitzMessage.TimestampZero
            }
        }
        return ts
    }

    fillFormFromProfile := func() {

        //  Expertise:
        for _, tag := range updateProfile.Profile.EntityTags {
            if ! strings.HasPrefix(*tag.TagName, ".") {
                updateProfile.Expertise += fmt.Sprintf("%s, ", *tag.TagName)
            }
        }
        updateProfile.Expertise = strings.TrimRight(updateProfile.Expertise, ", ")

        //  Add extra emp & edu --
        for i := 0; i < 2; i++ {
            updateProfile.Profile.Employment = append(updateProfile.Profile.Employment, &BlitzMessage.Employment{})
            updateProfile.Profile.Education  = append(updateProfile.Profile.Education,  &BlitzMessage.Education{})
        }

        //  Fix zero dates --
        for _, emp := range updateProfile.Profile.Employment {
            emp.Timespan = fixNullTimespan(emp.Timespan)
        }
        for _, edu := range updateProfile.Profile.Education {
            edu.Timespan = fixNullTimespan(edu.Timespan)
        }
    }


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
        fillFormFromProfile()
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

    updateProfile.Profile.Name = Util.CleanStringPtrFromString(httpRequest.PostFormValue("Name"))
    updateProfile.Profile.BackgroundSummary =
        Util.CleanStringPtrFromString(httpRequest.PostFormValue("BackgroundSummary"))

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

    //  Get the employment info --

    c := httpRequest.PostFormValue("JobCount")
    count, _ := strconv.Atoi(c)
    updateProfile.Profile.Employment = make([]*BlitzMessage.Employment, 0, 10)
    for i := 0; i < count; i++ {

        c = strconv.Itoa(i)
        s := httpRequest.PostFormValue("Job-Start-"+c)
        t := ServerUtil.ParseMonthYearString(s)
        startTime := BlitzMessage.TimestampPtr(t)
        s = httpRequest.PostFormValue("Job-Stop-"+c)
        t = ServerUtil.ParseMonthYearString(s)
        stopTime  := BlitzMessage.TimestampPtr(t)

        emp := BlitzMessage.Employment {
            JobTitle:       Util.CleanStringPtrFromString(httpRequest.PostFormValue("Job-JobTitle-"+c)),
            CompanyName:    Util.CleanStringPtrFromString(httpRequest.PostFormValue("Job-CompanyName-"+c)),
            Location:       Util.CleanStringPtrFromString(httpRequest.PostFormValue("Job-Location-"+c)),
            Industry:       Util.CleanStringPtrFromString(httpRequest.PostFormValue("Job-Industry-"+c)),
            Summary:        Util.CleanStringPtrFromString(httpRequest.PostFormValue("Job-Summary-"+c)),
            Timespan:       BlitzMessage.TimespanFromTimestamps(startTime, stopTime),
        }
        updateProfile.Profile.Employment =
            append(updateProfile.Profile.Employment, &emp)
    }

    //  Get the education info --

    c = httpRequest.PostFormValue("EduCount")
    count, _ = strconv.Atoi(c)
    updateProfile.Profile.Education = make([]*BlitzMessage.Education, 0, 10)
    for i := 0; i < count; i++ {

        c = strconv.Itoa(i)
        s := httpRequest.PostFormValue("Edu-Start-"+c)
        t := ServerUtil.ParseMonthYearString(s)
        startTime := BlitzMessage.TimestampPtr(t)
        s = httpRequest.PostFormValue("Edu-Stop-"+c)
        t = ServerUtil.ParseMonthYearString(s)
        stopTime  := BlitzMessage.TimestampPtr(t)

        edu := BlitzMessage.Education {
            Degree:         Util.CleanStringPtrFromString(httpRequest.PostFormValue("Edu-Degree-"+c)),
            Emphasis:       Util.CleanStringPtrFromString(httpRequest.PostFormValue("Edu-Emphasis-"+c)),
            SchoolName:     Util.CleanStringPtrFromString(httpRequest.PostFormValue("Edu-SchoolName-"+c)),
            Summary:        Util.CleanStringPtrFromString(httpRequest.PostFormValue("Edu-Summary-"+c)),
            Timespan:       BlitzMessage.TimespanFromTimestamps(startTime, stopTime),
        }
        updateProfile.Profile.Education =
            append(updateProfile.Profile.Education, &edu)
    }

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

        updateProfile.Profile = ProfileForUserID(userID, userID)
    }

    fillFormFromProfile()
    updateProfile.ErrorMessage = "User updated."
}

