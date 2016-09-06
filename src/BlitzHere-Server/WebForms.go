

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
    "io/ioutil"
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

        if len(user.Name.String) == 0 {
            user.Name.String = "< no name given >"
            user.Name.Valid = true
        }

        const bLen = 30
        if len(user.Background.String) > bLen {
            user.Background.String = user.Background.String[:bLen] + "..."
        }

        tagMap := GetEntityTagMapForUserIDEntityIDType(
            user.UserID.String,
            user.UserID.String,
            BlitzMessage.EntityType_ETUser,
        )

        if len(editprofileid.String) > 10 {
            user.Annotation = "Approve Edit"
        } else if ! isExpert.Bool {
            if _, ok := tagMap[".appliedexpert"]; ok {
                user.Annotation = "Applied"
            }
        }

        if isExpert.Bool {
            user.Annotation += "*"
        }

        if _, ok := tagMap[".expertimporthelp"]; ok {
            user.Annotation += " help"
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
        ProfileImage    string
        BackgroundImage string
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

        //  Images --
        for _, img := range updateProfile.Profile.Images {
            switch *img.ImageContent {
            case BlitzMessage.ImageContent_ICUserProfile:
                if len(updateProfile.ProfileImage) == 0 {
                    updateProfile.ProfileImage =
                        ImageURLForImageData(*updateProfile.Profile.UserID, img)
                }
            case BlitzMessage.ImageContent_ICUserBackground:
                if len(updateProfile.BackgroundImage) == 0 {
                    updateProfile.BackgroundImage =
                        ImageURLForImageData(*updateProfile.Profile.UserID, img)
                }
            }
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
        return
    }

    //  Erase user?

    eraseUser := httpRequest.PostFormValue("EraseUser")
    if  eraseUser == "EraseUser" {
        result, error := config.DB.Exec(
            `select eraseUserID($1);`,
            userID,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil {
            Log.LogError(error)
            updateProfile.ErrorMessage = error.Error()
            return
        }
        if updateProfile.Profile.EditProfileID != nil &&
           len(*updateProfile.Profile.EditProfileID) > 0 {
            result, error := config.DB.Exec(
                `select eraseUserID($1);`,
                updateProfile.Profile.EditProfileID,
            )
            error = pgsql.UpdateResultError(result, error)
            if error != nil {
                Log.LogError(error)
                updateProfile.ErrorMessage = error.Error()
                return
            }
        }
        updateProfile.ErrorMessage = "User erased."
        return
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
        if  isApproved == "IsApproved" &&
            updateProfile.Profile.EditProfileID != nil &&
            len(*updateProfile.Profile.EditProfileID) > 0 {

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
        userID = *updateProfile.Profile.EditProfileID
        }
        updateProfile.Profile = ProfileForUserID(userID, userID)
    }

    fillFormFromProfile()
    updateProfile.ErrorMessage = "User updated."
}


//----------------------------------------------------------------------------------------
//
//                                                                            WebImageEdit
//
//----------------------------------------------------------------------------------------


func WebImageEdit(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    type ImageType struct {
        Caption     string
        CRC         int64
        URL         string
    }
    imageEdit := struct {
        AppName         string
        Name            string
        ErrorMessage    string
        UserID          string
        Images          []ImageType
    } {
        AppName: config.AppName,
    }

    defer func() {
        ie := &imageEdit
        error := config.Template.ExecuteTemplate(writer, "ImageEdit.html", ie)
        if error != nil {
            Log.LogError(error)
        }
    } ()

    Log.Debugf("Request: %+v.", httpRequest)

    var error error
    userID := httpRequest.URL.Query().Get("uid")
    if len(userID) == 0 {
        imageEdit.ErrorMessage = "User ID required."
    }
    imageEdit.UserID = userID
    imageEdit.Name = PrettyNameForUserID(userID)

    deleteCRC := httpRequest.URL.Query().Get("delete")
    if len(deleteCRC) > 0 {
        var result sql.Result
        result, error = config.DB.Exec(
            `update ImageTable set deleted = true
                where userID = $1
                  and crc32 = $2;`,
            userID,
            deleteCRC,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil {
            Log.LogError(error)
            imageEdit.ErrorMessage = error.Error()
        } else {
            imageEdit.ErrorMessage = "Deleted"
        }
    }


    //  Parse a form upload --
    if httpRequest.Method == "POST" {
        error = httpRequest.ParseMultipartForm(10000000)
        if error != nil {
            Log.LogError(error)
            imageEdit.ErrorMessage = error.Error()
            return
        }

        //Log.Debugf("Form: %+v.", httpRequest.Form)
        //Log.Debugf("Post: %+v.", httpRequest.PostForm)

        file, header, error := httpRequest.FormFile("pic")
        if error != nil {
            Log.LogError(error)
            imageEdit.ErrorMessage = error.Error()
            return
        }
        defer file.Close()
        imageBytes, error := ioutil.ReadAll(file)
        if error != nil {
            Log.LogError(error)
            imageEdit.ErrorMessage = error.Error()
            return
        }
        contentType := header.Header.Get("Content-Type")
        imageContent, _ := strconv.Atoi(httpRequest.FormValue("imageContent"))

        Log.Debugf("Content Type: %s Image Type: %d.", contentType, imageContent);

        error = SaveImage(
            userID,
            BlitzMessage.ImageContent(imageContent),
            contentType,
            imageBytes,
        )
        if error != nil {
            imageEdit.ErrorMessage = error.Error()
        }
    }


    images := ImagesForUserID(userID)
    for idx, img := range images {

        s := "Unknown"
        switch {
        case img.ImageContent == nil:
        case *img.ImageContent == BlitzMessage.ImageContent_ICUserProfile:
            s = "Profile"
        case *img.ImageContent == BlitzMessage.ImageContent_ICUserBackground:
            s = "Background"
        }
        s = fmt.Sprintf("%s %d / %d", s, idx+1, len(images))

        newImg := ImageType {
            Caption:    s,
            CRC:        *img.Crc32,
            URL:        ImageURLForImageData(userID, img),
        }

        imageEdit.Images = append(imageEdit.Images, newImg)
    }
}

