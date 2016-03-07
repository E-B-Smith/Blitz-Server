//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                            UpdateContactInfoFromProfile
//----------------------------------------------------------------------------------------


func UpdateContactInfoFromProfile(profile *BlitzMessage.UserProfile) {
    Log.LogFunctionName()

    if profile.UserID == nil { return }

    result, error := config.DB.Exec("delete from UserContactTable where userID = $1;", profile.UserID)
    if error != nil { Log.Debugf("Delete UserContactInfo result: %v error: %v.", result, error) }

    for _, contact := range(profile.ContactInfo) {
        result, error = config.DB.Exec("insert into UserContactTable " +
            " (userID, contactType, contact, isverified) values " +
            " ($1, $2, $3, $4) ;",
            profile.UserID,
            contact.ContactType,
            Util.CleanStringPtr(contact.Contact),
            contact.IsVerified);
    if error != nil { Log.Errorf("Insert UserContactInfo result: %v error: %v.", result, error) }
    }
}


func AddContactInfoToProfile(profile *BlitzMessage.UserProfile) {
    Log.LogFunctionName();

    rows, error := config.DB.Query("select contactType, contact, isverified " +
        "  from UserContactTable where userid = $1", profile.UserID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.Errorf("Error getting contacts: %v.", error)
        return
    }

    for rows.Next() {
        var (contactType int; contact string; verified bool)
        error = rows.Scan(&contactType, &contact, &verified)
        if error == nil {
            ct := BlitzMessage.ContactType(contactType)
            contactStruct := BlitzMessage.ContactInfo {
                ContactType: &ct,
                Contact: &contact,
                IsVerified: &verified,
            }
            if profile.ContactInfo == nil { profile.ContactInfo = make([]*BlitzMessage.ContactInfo, 0, 5) }
            profile.ContactInfo = append(profile.ContactInfo, &contactStruct)
        } else {
            Log.LogError(error);
        }
    }
}



//----------------------------------------------------------------------------------------
//                                                                        ProfileForUserID
//----------------------------------------------------------------------------------------


func ProfileForUserID(userID string) *BlitzMessage.UserProfile {
    Log.Infof("ProfileForUserId (%T) %s.", userID, userID)

    rows, error := config.DB.Query(
        "select userID, userStatus, name, gender, birthday, imageURL," +
        "  creationDate from UserTable where userID = $1;", userID)
    defer pgsql.CloseRows(rows)

    if error != nil {
        Log.Debugf("Error finding user for %s: %v.", userID, error)
        return nil
    }

    if !rows.Next() {
        Log.Debugf("No rows.")
        return nil;
    }

    var (
        profileID   string;
        userStatus  sql.NullInt64;
        name        sql.NullString;
        gender      sql.NullInt64;
        birthday    pq.NullTime;
        imageURLs   sql.NullString;
        creationDate pq.NullTime;
    )
    error = rows.Scan(
        &profileID,
        &userStatus,
        &name,
        &gender,
        &birthday,
        &imageURLs,
        &creationDate,
    )
    if error != nil {
        Log.Errorf("Error scanning row: %v.", error)
        return nil
    }

    profile := new(BlitzMessage.Profile)
    profile.UserID      = proto.String(profileID)
    profile.UserStatus  = BlitzMessage.UserStatus(userStatus.Int64).Enum()
    profile.Name        = proto.String(name.String)
    profile.Gender      = BlitzMessage.Gender(gender.Int64).Enum()
    profile.Birthday    = BlitzMessage.TimestampFromTime(birthday.Time)
    profile.ImageURL    = pgsql.StringArrayFromNullString(imageURLs)
    Log.Debugf("Profile has %d images: %v.", len(profile.ImageURL), profile.ImageURL)
    profile.SocialIdentities = SocialIdentitiesWithUserID(userID)
    profile.CreationDate   = BlitzMessage.TimestampFromTime(creationDate.Time)
    profile.UserSummary    = UserStatsSummary(userID)
    profile.CircleSummary  = CircleStatsSummary(userID)
    profile.GlobalSummary  = GlobalStatsSummary(userID)
    profile.WeatherSummary = WeatherSummary(userID)
    profile.HeartsSent     = HeartsSent(userID)
    profile.HeartsReceived = HeartsReceived(userID)
    profile.LatestScore    = LatestScoreForUserID(userID)
    if profile.LatestScore != nil {
        profile.LastHappyScore = profile.LatestScore.HappyScore
    }
    AddContactInfoToProfile(profile)

    return profile
}


func QueryProfiles(writer http.ResponseWriter, userID string, profileQuery *BlitzMessage.UserProfileQuery) {
    Log.LogFunctionName()

    var profileUpdate BlitzMessage.ProfileUpdate
    for i := range profileQuery.UserIDs {
        profile := ProfileForUserID(profileQuery.UserIDs[i])
        if profile != nil {
            profileUpdate.Profiles = append(profileUpdate.Profiles, profile)
        }
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    var message string

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
        Response:           &BlitzMessage.ServerResponse_ProfileUpdate { ProfileUpdate: &profileUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}

