//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "time"
    "errors"
    "strings"
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                                           UpdateProfile
//----------------------------------------------------------------------------------------


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

    _, error = config.DB.Exec("insert into usertable (userid) values ($1);", userID)
    if error != nil {
        //Log.Debugf("Error inserting user '%s': %v.", userID, error)
    }

    if profile.CreationDate == nil {
        profile.CreationDate = BlitzMessage.TimestampFromTime(time.Now())
    }

    _, error = config.DB.Exec("update usertable set" +
        " (name, gender, birthday, userStatus, creationDate) =" +
        " ($1, $2, $3, $4, $5)" +
        " where userID = $6;",
        profile.Name,
        profile.Gender,
        BlitzMessage.NullTimeFromTimestamp(profile.Birthday),
        profile.UserStatus,
        BlitzMessage.NullTimeFromTimestamp(profile.CreationDate),
        &userID)
    if error != nil {
        Log.Errorf("Error updating profile %s: %+v", *profile.UserID, error)
        return error
    }

    var images []string
    for _, s := range profile.ImageURL {
        s = strings.TrimSpace(s)
        if len(s) > 0 { images = append(images, s); }
    }
    profile.ImageURL = images
    Log.Debugf("Profile in has %d images.", len(profile.ImageURL));

    if len(profile.ImageURL) > 0 {
        _, error = config.DB.Exec("update usertable set" +
            " (imageURL) = ($1) where userID = $2;",
            pgsql.NullStringFromStringArray(profile.ImageURL),
            &userID)
        if error != nil {
            Log.Errorf("Error updating profile %s: %+v", *profile.UserID, error)
            return error
        }
    }

    for i := range profile.SocialIdentities {
        UpdateSocialIdentityForUserID(userID, profile.SocialIdentities[i])
    }

    // for i := range profile.Scores {
    //     UpdateScoreForUserID(userID, profile.Scores[i])
    // }

    UpdateContactInfoFromProfile(profile)
    UpdateUserIdentitesFromProfile(profile)

    return error
}


//----------------------------------------------------------------------------------------
//
//                                                                          UpdateProfiles
//
//----------------------------------------------------------------------------------------


func UpdateProfiles(writer http.ResponseWriter, userID string, profiles *BlitzMessage.UserProfileUpdate) {
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

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
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


