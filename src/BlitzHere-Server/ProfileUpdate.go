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


func UpdateEmployment(userID *string, isCurrentPosition bool, employment *BlitzMessage.Employment) {
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
            ,isCurrentPosition
            ,jobTitle
            ,companyName
            ,location
            ,industry
            ,startDate
            ,stopDate
            ,summary) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
        userID,
        isCurrentPosition,
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


func UpdateEducation(userID *string, education *BlitzMessage.Education) {
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
    UpdateEmployment(profile.UserID, true, profile.HeadlineEmployment)
    for _, employment := range profile.Employment {
        UpdateEmployment(profile.UserID, false, employment)
    }

    config.DB.Exec(`delete from EducationTable where userID = $1;`, profile.UserID)
    for _, education := range profile.Education {
        UpdateEducation(profile.UserID, education)
    }

    SetEntityTags(*profile.UserID, profile.ExpertiseTags)

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

    for i := 0; i < len(profiles.Profiles); i++ {
        error := UpdateProfile(profiles.Profiles[i])
        if error != nil {
            errorCount++
            if firstError == nil { firstError = error }
            Log.Errorf("Error updating profile %s: %v.", *profiles.Profiles[i].UserID, error)
        }
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    var message string

    if errorCount > 0 {
        Log.Errorf("Found %d errors on update.", errorCount)
        code = BlitzMessage.ResponseCode_RCServerWarning
        message = firstError.Error()
    }

    if  errorCount == len(profiles.Profiles) {
        code = BlitzMessage.ResponseCode_RCServerError
        if (message == "") { message = "No profiles to update" }
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
        ResponseMessage: &message,
    }
    return response
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


    var (oldStatus sql.NullInt64; oldName sql.NullString; oldGender sql.NullInt64; oldBirthday pq.NullTime; oldImage sql.NullString)
    row := config.DB.QueryRow("select userstatus, name, gender, birthday, imageURL[1] from usertable where userid = $1;", oldID)
    error := row.Scan(&oldStatus, &oldName, &oldGender, &oldBirthday, &oldImage)
    if error != nil {
        Log.LogError(error)
    }
    var (newStatus sql.NullInt64; newName sql.NullString; newGender sql.NullInt64; newBirthday pq.NullTime; newImage sql.NullString)
    row = config.DB.QueryRow("select userstatus, name, gender, birthday, imageURL[1] from usertable where userid = $1;", newID)
    error = row.Scan(&newStatus, &newName, &newGender, &newBirthday, &newImage)
    if error != nil {
        Log.LogError(error)
    }
    if oldStatus.Int64 > newStatus.Int64 { newStatus = oldStatus; }
    if ! newName.Valid || len(newName.String) <= 0 { newName = oldName; }
    if ! newGender.Valid || newGender.Int64 == 0 { newGender = oldGender; }
    if ! newBirthday.Valid || newBirthday.Time.IsZero() { newBirthday = oldBirthday; }
    if ! newImage.Valid || len(newImage.String) <= 0 { newImage = oldImage; }

    result, error := config.DB.Exec(
        `update usertable set (userstatus, name, gender, birthday, imageURL[1])
         = ($1, $2, $3, $4, $5) where userid = $6;`,
        newStatus, newName, newGender, newBirthday, newImage, newID)
    if error != nil {
        Log.LogError(error)
    } else {
        rowCount, _ := result.RowsAffected()
        if rowCount != 1 { Log.Errorf("Update row count not 1!: %d.", rowCount); }
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
        Log.Debugf("Updated %ld rows.", rowCount)
    }
    return nil
}


