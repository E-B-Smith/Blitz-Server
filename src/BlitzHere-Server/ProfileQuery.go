//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
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


func ImagesForUserID(userID string) []*BlitzMessage.ImageData {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select dateAdded, imageContent, contentType, crc32
            from ImageTable
            where userID = $1
            order by dateAdded desc;`, userID)

    if error != nil {
        Log.LogError(error)
        return nil
    }
    defer rows.Close()

    imageArray := make([]*BlitzMessage.ImageData, 0, 5)
    for rows.Next() {
        var (
            dateAdded       pq.NullTime
            imageContent    sql.NullInt64
            contentType     sql.NullString
            crc32           sql.NullInt64
        )
        error = rows.Scan(&dateAdded, &imageContent, &contentType, &crc32)
        if error != nil {
            Log.LogError(error)
            continue
        }
        content := BlitzMessage.ImageContent(*Int32PtrFromNullInt64(imageContent))
        imageData := BlitzMessage.ImageData {
            DateAdded:      BlitzMessage.TimestampPtrFromNullTime(dateAdded),
            ImageContent:   &content,
            ContentType:    StringPtrFromNullString(contentType),
        }
        imageData.ImageURL = StringPtr(ImageURLForImageData(userID, &imageData))
        imageArray = append(imageArray, &imageData)
    }

    Log.Debugf("Profile has %d images: %v.", len(imageArray))
    return imageArray
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

    profile := new(BlitzMessage.UserProfile)
    profile.UserID      = proto.String(profileID)
    profile.UserStatus  = BlitzMessage.UserStatus(userStatus.Int64).Enum()
    profile.Name        = proto.String(name.String)
    profile.Gender      = BlitzMessage.Gender(gender.Int64).Enum()
    profile.Birthday    = BlitzMessage.TimestampFromTime(birthday.Time)
    profile.Images      = ImagesForUserID(userID)
    profile.SocialIdentities = SocialIdentitiesWithUserID(userID)
    profile.CreationDate   = BlitzMessage.TimestampFromTime(creationDate.Time)
    AddContactInfoToProfile(profile)

    return profile
}


func QueryProfiles(session *Session, profileQuery *BlitzMessage.UserProfileQuery,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var profileUpdate BlitzMessage.UserProfileUpdate
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
        ResponseType:       &BlitzMessage.ResponseType { ProfileUpdate: &profileUpdate },
    }
    return response
}

