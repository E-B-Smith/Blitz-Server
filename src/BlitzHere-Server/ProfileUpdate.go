//  BlitzHere-Server  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "time"
    "errors"
    "strings"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func StringsAreEmptyWithClean(strs ...*string) bool {
    result := true

    for _, s := range strs {
        if s == nil { continue }
        *s = strings.TrimSpace(*s)
        if len(*s) > 0 {
            result = false
        }
    }

    return result
}


func InsertEmployment(userID *string, isHeadLineItem bool, employment *BlitzMessage.Employment) {
    Log.LogFunctionName()

    if employment == nil || userID == nil { return }

    if StringsAreEmptyWithClean(
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
        BlitzMessage.NullTimeFromTimespanStart(employment.Timespan),
        BlitzMessage.NullTimeFromTimespanStop(employment.Timespan),
        employment.Summary)
    if error != nil {
        Log.LogError(error)
    }
}


func InsertEducation(userID *string, education *BlitzMessage.Education) {
    Log.LogFunctionName()

    if education == nil || userID == nil { return }

    if StringsAreEmptyWithClean(
        education.SchoolName,
        education.Degree,
        education.Emphasis) {
        return
    }

    _, error := config.DB.Exec(
        `insert into EducationTable (
             userID
            ,schoolName
            ,degree
            ,emphasis
            ,startDate
            ,stopDate) values ($1, $2, $3, $4, $5, $6);`,
        userID,
        education.SchoolName,
        education.Degree,
        education.Emphasis,
        BlitzMessage.NullTimeFromTimespanStart(education.Timespan),
        BlitzMessage.NullTimeFromTimespanStop(education.Timespan))
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

    _, error = config.DB.Exec(
        `insert into usertable (userid, creationDate)
            values ($1, current_timestamp);`, userID)
    if error != nil {
        //Log.Debugf("Error inserting user '%s': %v.", userID, error)
    }

    if profile.CreationDate == nil {
        profile.CreationDate = BlitzMessage.TimestampFromTime(time.Now())
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
            ,interestTags) = ($1, $2, $3, $4, $5)
                where userID = $6;`,
        profile.Name,
        profile.Gender,
        BlitzMessage.NullTimeFromTimestamp(profile.Birthday),
        profile.BackgroundSummary,
        pgsql.NullStringFromStringArray(profile.InterestTags),
        profile.UserID)
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

    SetEntityTagsWithUserID(*profile.UserID, *profile.UserID, BlitzMessage.EntityType_ETUser, profile.ExpertiseTags)

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
    error = pgsql.RowUpdateError(result, error)
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


