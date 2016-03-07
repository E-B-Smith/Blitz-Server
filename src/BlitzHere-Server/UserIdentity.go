//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "strings"
    "net/http"
    "unicode/utf8"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


func IdentityStringFromString(identity string) string {
    validCharacters := "0123456789abcdefghijklmnopqrstuvwxyz";
    identity = strings.ToLower(identity)
    identity = Util.StringIncludingCharactersInSet(identity, validCharacters)
    return identity
}


func appendID(ids []string, typeString string, valueString string) []string {
    typeString  = IdentityStringFromString(typeString)
    valueString = IdentityStringFromString(valueString)
    Log.Debugf("Maybe append %s+%s.", typeString, valueString)
    if utf8.RuneCountInString(typeString) > 0 && utf8.RuneCountInString(valueString) > 0 {
       ids = append(ids, typeString + valueString)
    }
    return ids
}


func IdentityStringsForDevice(ids []string, vendorID *string, advertisingID *string, deviceID *string) []string {
    if vendorID != nil && len(*vendorID) > 0           { ids = appendID(ids, "vendor", *vendorID) }
    if advertisingID != nil && len(*advertisingID) > 0 { ids = appendID(ids, "adid", *advertisingID) }
    if deviceID != nil && len(*deviceID) > 16          { ids = appendID(ids, "deviceid", *deviceID) }
    return ids
}


func IdentityStringsFromProfile(profile *BlitzMessage.UserProfile) []string {
    var ids []string = make([]string, 0, 10)

    ids = appendID(ids, "userid", *profile.UserID)
    for _, contact := range profile.ContactInfo {
        if contact.ContactType != nil && contact.Contact != nil {
            ids = appendID(ids, BlitzMessage.ContactType_name[int32(*contact.ContactType)], *contact.Contact)
        }
    }
    for i := range profile.SocialIdentities {
        sid := profile.SocialIdentities[i]
        if sid.SocialService != nil {
            if sid.SocialID != nil { ids = appendID(ids, *sid.SocialService, *sid.SocialID) }
            if sid.UserName != nil { ids = appendID(ids, *sid.SocialService, *sid.UserName) }
            if sid.UserURI != nil  { ids = appendID(ids, *sid.SocialService, *sid.UserURI) }
        }
    }

    rows, error := config.DB.Query("select vendorID, advertisingID, deviceUDID" +
        "  from DeviceTable where userID = $1;", *profile.UserID)
    if error != nil {
        Log.Debugf("Error getting device identity strings: %v.", error)
    }

    defer rows.Close()
    for rows.Next() {
        var (vendor string; adID string; deviceID string)
        error = rows.Scan(&vendor, &adID, &deviceID)
        if error == nil {
            ids = IdentityStringsForDevice(ids, &vendor, &adID, &deviceID)
        }
    }

    Log.Debugf("Made %d identity strings.", len(ids))
    return ids
}


func UpdateUserIdentitesFromProfile(profile *BlitzMessage.UserProfile) {
    if profile == nil { return }
    Log.Infof("UpdateUserIdentitesFromProfile %s.", *profile.UserID)
    identities := IdentityStringsFromProfile(profile)

    for i := range identities {
        insertResult, error := config.DB.Exec(
            "insert into UserIdentityTable (userID, identityString) values ($1, $2);", *profile.UserID, identities[i])

        Log.Debugf("User: '%s' Identity: '%v'.", *profile.UserID, identities[i])
        Log.Debugf("Identity insert result: %v Error: %v.", insertResult, error)

        var rowsUpdated int64 = 0
        if insertResult != nil { rowsUpdated, _ = insertResult.RowsAffected() }

        if error != nil || rowsUpdated == 0 {
            _, error := config.DB.Exec(
                "update UserIdentityTable set (userID, identityString) = ($1, $2)" +
                " where userID = $3 and identityString = $4;", *profile.UserID, identities[i], *profile.UserID, identities[i])
            if error != nil {
                Log.Errorf("Can't insert or update user profile identities for %s: %v.", *profile.UserID, error)
            }
        }
    }
}


func ExistingProfileFromIdentities(identities []string) *BlitzMessage.UserProfile {
    Log.LogFunctionName()

    if len(identities) == 0 { return nil }
    var profile *BlitzMessage.Profile = nil

    queryP1 :=
        "select useridentitytable.userid from useridentitytable" +
        "  left join usertable on usertable.userid = useridentitytable.userid" +
        "  where identitystring in ($1"
    queryP3 := ")" +
        "  group by usertable.userid, useridentitytable.userid" +
        "  order by usertable.creationdate" +
        "  limit 1;"

    queryP2 := ""
    for i:= 2; i <= len(identities); i++ {
        queryP2 += fmt.Sprintf(", $%d", i)
    }

    query := queryP1+queryP2+queryP3
    Log.Debugf("Query is %s.", query)
    ii := make([]interface{}, len(identities))
    for i := range identities { ii[i] = identities[i] }
    rows, error := config.DB.Query(query, ii...)
    if error != nil {
        Log.Errorf("Error finding identity: %v.", error)
        return nil;
    }

    var rowCount int = 0
    defer rows.Close()
    for rows.Next() {
        var userID string
        error = rows.Scan(&userID)
        if error != nil {
            Log.Errorf("Error scanning row: %v.", error)
            return nil
        }
        rowCount++
        profile = ProfileForUserID(userID)
        if profile != nil {
            Log.Debugf("Found row count %d: %s.", rowCount, *profile.UserID);
            return profile;
        }
    }
    Log.Debugf("Found row count %d but returning nil.", rowCount);
    return nil;
}



//----------------------------------------------------------------------------------------
//
//                                                                 ProfilesFromContactInfo
//
//----------------------------------------------------------------------------------------


func ProfilesFromContactInfo(profilesIn []*BlitzMessage.UserProfile) []*BlitzMessage.UserProfile {
    Log.LogFunctionName()

    profileMap := make(map[string]*BlitzMessage.Profile)
    for _, profile := range profilesIn {
        name := "<Anon>"
        if profile.Name != nil { name = *profile.Name }
        Log.Debugf("Looking for %s.", name)

        var rowCount int = 0
        for _, contactInfo := range profile.ContactInfo {
            if contactInfo == nil || contactInfo.Contact == nil || contactInfo.ContactType == nil {
                continue
            }
            rows, error := config.DB.Query("select userid from UserContactTable where contacttype = $1 and contact = $2;",
                contactInfo.ContactType, contactInfo.Contact)
            defer pgsql.CloseRows(rows)
            if error != nil {
                Log.LogError(error)
                continue
            }
            for rows.Next() {
                var userID string
                error = rows.Scan(&userID)
                if error != nil {
                    Log.LogError(error)
                    continue
                }
                memberProfile := ProfileForUserID(userID)
                if memberProfile != nil {
                    profileMap[*memberProfile.UserID] = memberProfile
                    rowCount++
                }
            }
        }
        if rowCount == 0 {
            if profile.UserID == nil {
                profile.UserID = StringPtrFromString(Util.NewUUIDString())
            }
            status := BlitzMessage.UserStatus_USInvited
            profile.UserStatus = &status
            error := UpdateProfile(profile)
            if error != nil {
                Log.LogError(error)
                continue
            }
            profileMap[*profile.UserID] = profile
        }
    }

    profilesOut := make([]*BlitzMessage.Profile, 0, len(profileMap))
    for _, profile := range profileMap {
        profilesOut = append(profilesOut, profile)
    }

    return profilesOut
}


func ProfilesFromContactInfoRequest(writer http.ResponseWriter,
            session *Session,
            profilesFromContactInfo *BlitzMessage.ProfilesFromContactInfo) {
    Log.LogFunctionName()

    profileUpdate := BlitzMessage.ProfileUpdate { }
    profileUpdate.Profiles = ProfilesFromContactInfo(profilesFromContactInfo.Profiles)

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
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

