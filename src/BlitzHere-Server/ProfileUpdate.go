//  BlitzHere-Server  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "time"
    "errors"
    "strings"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func StringsAreEmptyAfterClean(strs ...*string) bool {
    result := true
    for _, s := range strs {
        if s == nil { continue }
        *s = strings.TrimSpace(*s)
        if len(*s) > 0 {
            result = false
        } else {
            s = nil
        }
    }

    return result
}


func InsertEmployment(userID *string, isHeadLineItem bool, employment *BlitzMessage.Employment) {
    Log.LogFunctionName()

    if employment == nil || userID == nil { return }

    if StringsAreEmptyAfterClean(
        employment.JobTitle,
        employment.CompanyName,
        employment.Location,
        employment.Industry,
        employment.Summary) {
        return
    }

    _, error := config.DB.Exec(
        `insert into EmploymentTable (
             userID
            ,isHeadLineItem
            ,jobTitle
            ,companyName
            ,location
            ,industry
            ,startDate
            ,stopDate
            ,summary) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
        userID,
        isHeadLineItem,
        employment.JobTitle,
        employment.CompanyName,
        employment.Location,
        employment.Industry,
        employment.Timespan.NullTimeStart(),
        employment.Timespan.NullTimeStop(),
        employment.Summary)
    if error != nil {
        Log.LogError(error)
    }
}


func InsertEducation(userID *string, education *BlitzMessage.Education) {
    Log.LogFunctionName()

    if education == nil || userID == nil { return }

    if StringsAreEmptyAfterClean(
        education.SchoolName,
        education.Degree,
        education.Emphasis,
        education.Summary) {
        return
    }

    _, error := config.DB.Exec(
        `insert into EducationTable (
             userID
            ,schoolName
            ,degree
            ,emphasis
            ,startDate
            ,stopDate
            ,summary) values ($1, $2, $3, $4, $5, $6, $7);`,
        userID,
        education.SchoolName,
        education.Degree,
        education.Emphasis,
        education.Timespan.NullTimeStart(),
        education.Timespan.NullTimeStop(),
        education.Summary)
    if error != nil {
        Log.LogError(error)
    }
}


//----------------------------------------------------------------------------------------
//                                                                           UpdateProfile
//----------------------------------------------------------------------------------------


func UpdateProfileStatusForUserID(userID string, status BlitzMessage.UserStatus) {
    _, error := config.DB.Exec("update UserTable set userStatus = $2 where userID = $1;", userID, status)
    if error != nil { Log.LogError(error) }
}


func UpdateProfile(profile *BlitzMessage.UserProfile) error {
    //  * Update the user info from the profile.
    Log.LogFunctionName()

    if profile == nil {
        return errors.New("Nil profile.")
    }

    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    userID, error := BlitzMessage.ValidateUserID(profile.UserID)
    if error != nil {
        return errors.New("Invalid user")
    }

    Log.Debugf("Updating profile %s.", userID)

    var result sql.Result
    _, error = config.DB.Exec(
        `insert into usertable (userid, creationDate)
            values ($1, current_timestamp);`, userID)
    if error == nil {
        //  New user.  Add blitz friend:
        result, error = config.DB.Exec(
            `insert into entitytagtable
                (userid, entitytype, entityid, entitytag)
                values ($1, 1, $2, '.friend');`,
            userID,
            BlitzMessage.Default_Global_BlitzUserID,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil { Log.LogError(error) }
    }

    if profile.CreationDate == nil {
        profile.CreationDate = BlitzMessage.TimestampPtr(time.Now())
    }

    //  eDebug -- Remove this:
    if profile.UserStatus != nil &&
       *profile.UserStatus == BlitzMessage.UserStatus_USConfirming {
        panic("User status confirming saved.")
    }

    _, error = config.DB.Exec(
        `update usertable set (
             name
            ,gender
            ,birthday
            ,backgroundSummary
            ,interestTags
            ,stripeAccount
            ,editProfileID
        ) = ($1, $2, $3, $4, $5, $6, $7)
                where userID = $8;`,
        profile.Name,
        profile.Gender,
        profile.Birthday.NullTime(),
        profile.BackgroundSummary,
        pgsql.NullStringFromStringArray(profile.InterestTags),
        profile.StripeAccount,
        profile.EditProfileID,
        profile.UserID,
    )
    if error != nil {
        Log.Errorf("Error updating profile %s: %+v", *profile.UserID, error)
        return error
    }

    for i := range profile.SocialIdentities {
        UpdateSocialIdentityForUserID(userID, profile.SocialIdentities[i])
    }

    UpdateContactInfoFromProfile(profile)
    UpdateUserIdentitesFromProfile(profile)

    config.DB.Exec(`delete from EmploymentTable where userID = $1;`, profile.UserID)
    InsertEmployment(profile.UserID, true, profile.HeadlineEmployment)
    for _, employment := range profile.Employment {
        InsertEmployment(profile.UserID, false, employment)
    }

    config.DB.Exec(`delete from EducationTable where userID = $1;`, profile.UserID)
    for _, education := range profile.Education {
        InsertEducation(profile.UserID, education)
    }

    SetEntityTagsWithUserID(*profile.UserID, *profile.UserID, BlitzMessage.EntityType_ETUser, profile.EntityTags)

    row := config.DB.QueryRow("select UpdateSearchIndexForUserID($1);", profile.UserID)
    var resultstring string
    error = row.Scan(&resultstring)
    if error != nil { Log.LogError(error); }

    return error
}


//----------------------------------------------------------------------------------------
//
//                                                                          UpdateProfiles
//
//----------------------------------------------------------------------------------------


func UpdateProfiles(session *Session, profiles *BlitzMessage.UserProfileUpdate,
        ) *BlitzMessage.ServerResponse {

    //  * Update each profile in the update request.
    Log.LogFunctionName()

    errorCount := 0
    var firstError error = nil

    userIDArray := make([]string, 0, len(profiles.Profiles))
    for i := 0; i < len(profiles.Profiles); i++ {
        error := UpdateProfile(profiles.Profiles[i])
        if error != nil {
            errorCount++
            if firstError == nil { firstError = error }
            Log.Errorf("Error updating profile %s: %v.", *profiles.Profiles[i].UserID, error)
        }
        if profiles.Profiles[i].UserID != nil {
            userIDArray = append(userIDArray, *profiles.Profiles[i].UserID)
        }
    }

    if errorCount > 0 {
        Log.Errorf("Found %d errors on update: %+v.", errorCount, firstError)
    }

    if  errorCount == len(profiles.Profiles) {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError,
            errors.New("No profiles to update"))
    }

    requestType := BlitzMessage.UserProfileQuery {
        UserIDs:        userIDArray,
    }
    return QueryProfiles(session, &requestType)
}


//----------------------------------------------------------------------------------------
//
//                                                             MergeProfileIDIntoProfileID
//
//----------------------------------------------------------------------------------------


func MergeProfileIDIntoProfileID(oldID string, newID string) error {
    Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()


    var (
        oldStatus       sql.NullInt64
        oldName         sql.NullString
        oldGender       sql.NullInt64
        oldBirthday     pq.NullTime
        oldSummary      sql.NullString
        oldInterests    sql.NullString
    )
    row := config.DB.QueryRow(
        `select userstatus, name, gender, birthday, backgroundSummary, interestTags
            from usertable where userid = $1;`, oldID)
    error := row.Scan(&oldStatus, &oldName, &oldGender, &oldBirthday, &oldSummary, &oldInterests)
    if error != nil {
        Log.LogError(error)
    }
    var (
        newStatus       sql.NullInt64
        newName         sql.NullString
        newGender       sql.NullInt64
        newBirthday     pq.NullTime
        newSummary      sql.NullString
        newInterests    sql.NullString
    )
    row = config.DB.QueryRow(
        `select userstatus, name, gender, birthday, backgroundSummary, interestTags
            from usertable where userid = $1;`, newID)
    error = row.Scan(&newStatus, &newName, &newGender, &newBirthday, &newSummary, &newInterests)
    if error != nil {
        Log.LogError(error)
    }
    if oldStatus.Int64 > newStatus.Int64              { newStatus = oldStatus; }
    if ! newName.Valid || len(newName.String) <= 0    { newName = oldName; }
    if ! newGender.Valid || newGender.Int64 == 0      { newGender = oldGender; }
    if ! newBirthday.Valid || newBirthday.Time.IsZero() { newBirthday = oldBirthday; }
    if ! newSummary.Valid || len(newSummary.String) <= 0     { newSummary = oldSummary }
    if ! newInterests.Valid || len(newInterests.String) <= 0 { newInterests = oldInterests }

    result, error := config.DB.Exec(
        `update usertable set (userstatus, name, gender, birthday, backgroundSummary, interestTags)
         = ($1, $2, $3, $4, $5, $6) where userid = $7;`,
        newStatus, newName, newGender, newBirthday, newSummary, newInterests, newID)
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
    }

    row = config.DB.QueryRow("select MergeUserIDIntoUserID($1, $2);", oldID, newID)
    var mergeResult string
    error = row.Scan(&mergeResult)
    if error != nil || mergeResult != "User merged" {
        Log.Errorf("Error merging userID %s into userID %s. '%s' : %v.", oldID, newID, mergeResult, error)
    } else {
        Log.Debugf("Merge success!")
    }

    Session_DeleteSessionsForUserID(oldID)

    oldidentity := IdentityStringFromString("userid") + IdentityStringFromString(oldID)
    result, error = config.DB.Exec("insert into UserIdentityTable (userid, identitystring) values ($1, $2);", newID, oldidentity)
    if error != nil {
        Log.LogError(error)
    } else {
        rowCount, _ := result.RowsAffected()
        Log.Debugf("Updated %d rows.", rowCount)
    }
    return nil
}


//----------------------------------------------------------------------------------------
//
//                                                                        StartEditProfile
//
//----------------------------------------------------------------------------------------


func StartEditProfile(session *Session, editProfile *BlitzMessage.EditProfile,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var error error
    if editProfile == nil ||
        editProfile.ProfileID == nil {
        error = fmt.Errorf("Missing fields")
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    userID := *editProfile.ProfileID
    editID := Util.NewUUIDString()

    profile := ProfileForUserID(session, userID)
    if profile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    if profile.EditProfileID != nil && len(*profile.EditProfileID) > 0 {

        if *profile.UserStatus >= BlitzMessage.UserStatus_USConfirmed {
            editID = *profile.EditProfileID
        } else {
            editID = userID
            userID = *profile.EditProfileID
        }

    } else {

        row := config.DB.QueryRow(
            `select CreateEditProfile($1, $2);`,
            userID,
            editID,
        )
        var result sql.NullString
        error = row.Scan(&result)
        if error != nil || ! result.Valid || result.String != "Profile created" {
            Log.Errorf("Can't create edit profile. Error: %v %+v.", error, result)
            if error == nil { error = fmt.Errorf("%+v", result) }
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }

    }

    editProfile.Profile = ProfileForUserID(session, userID)
    editProfile.EditProfile = ProfileForUserID(session, editID)

    return &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType {
            EditProfile:    editProfile,
        },
    }
}


