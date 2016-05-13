//  BlitzHere-Server  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "strings"
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


func AddContactInfoToUserID(userID string, contact *BlitzMessage.ContactInfo) {
    result, error := config.DB.Exec("insert into UserContactTable " +
        " (userID, contactType, contact, isverified) values " +
        " ($1, $2, $3, $4) ;",
        userID,
        contact.ContactType,
        Util.CleanStringPtr(contact.Contact),
        contact.IsVerified);
    if error != nil {
        Log.Errorf("Insert UserContactInfo result: %v error: %v.", result, error)
    } else {
        Log.Debugf("Added %s.", *Util.CleanStringPtr(contact.Contact))
    }
}


func UpdateContactInfoFromProfile(profile *BlitzMessage.UserProfile) {
    Log.LogFunctionName()

    if profile.UserID == nil { return }

    result, error := config.DB.Exec("delete from UserContactTable where userID = $1;", profile.UserID)
    if error != nil { Log.Debugf("Delete UserContactInfo result: %v error: %v.", result, error) }

    for _, contact := range(profile.ContactInfo) {
        AddContactInfoToUserID(*profile.UserID, contact)
    }
}


func ContactInfoForUserID(userID string) []*BlitzMessage.ContactInfo {
    Log.LogFunctionName();

    rows, error := config.DB.Query("select contactType, contact, isverified " +
        "  from UserContactTable where userid = $1", userID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.Errorf("Error getting contacts: %v.", error)
        return nil
    }

    contactArray := make([]*BlitzMessage.ContactInfo, 0, 5)
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
            contactArray = append(contactArray, &contactStruct)
        } else {
            Log.LogError(error);
        }
    }
    return contactArray
}


func ImagesForUserID(userID string) []*BlitzMessage.ImageData {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select dateAdded, imageContent, contentType, crc32
            from ImageTable
            where userID = $1
              and (deleted is null or deleted = false)
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
            crc32           int64
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
            Crc32:          &crc32,
        }
        imageData.ImageURL = StringPtr(ImageURLForImageData(userID, &imageData))
        imageArray = append(imageArray, &imageData)
    }

    Log.Debugf("Profile has %d images.", len(imageArray))
    return imageArray
}


func EmploymentForUserID(userID string) []*BlitzMessage.Employment {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select
             isHeadlineItem
            ,jobTitle
            ,companyName
            ,location
            ,industry
            ,startDate
            ,stopDate
            ,summary
            from EmploymentTable where userID = $1
            order by stopDate desc;`, userID);
    if error != nil {
        Log.LogError(error)
        return nil
    }
    defer rows.Close()

    employmentArray := make([]*BlitzMessage.Employment, 0, 5)
    for rows.Next() {
        var (
            isHeadlineItem      bool
            jobTitle            sql.NullString
            companyName         sql.NullString
            location            sql.NullString
            industry            sql.NullString
            startDate           pq.NullTime
            stopDate            pq.NullTime
            summary             sql.NullString
        )
        error = rows.Scan(
            &isHeadlineItem,
            &jobTitle,
            &companyName,
            &location,
            &industry,
            &startDate,
            &stopDate,
            &summary,
        )
        if error != nil {
            Log.LogError(error)
        } else {
            employment := BlitzMessage.Employment {
                IsHeadlineItem: BoolPtr(isHeadlineItem),
                JobTitle:       StringPtrFromNullString(jobTitle),
                CompanyName:    StringPtrFromNullString(companyName),
                Location:       StringPtrFromNullString(location),
                Industry:       StringPtrFromNullString(industry),
                Summary:        StringPtrFromNullString(summary),
            }
            employment.Timespan = BlitzMessage.TimespanFromNullTimes(startDate, stopDate)
            employmentArray = append(employmentArray, &employment)
        }
    }
    return employmentArray
}


func EducationForUserID(userID string) []*BlitzMessage.Education {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select
             schoolName
            ,degree
            ,emphasis
            ,startDate
            ,stopDate
            ,summary
                from EducationTable where userID = $1
                order by stopDate desc;`, userID)
    if error != nil {
        Log.LogError(error)
        return nil
    }
    defer rows.Close()

    educationArray := make([]*BlitzMessage.Education, 0, 5)
    for rows.Next() {
        var (
            schoolName         sql.NullString
            degree             sql.NullString
            emphasis           sql.NullString
            startDate          pq.NullTime
            stopDate           pq.NullTime
            summary            sql.NullString
        )
        error = rows.Scan(
            &schoolName,
            &degree,
            &emphasis,
            &startDate,
            &stopDate,
            &summary)
        if error != nil {
            Log.LogError(error)
        } else {
            education := BlitzMessage.Education {
                SchoolName:   StringPtrFromNullString(schoolName),
                Degree:       StringPtrFromNullString(degree),
                Emphasis:     StringPtrFromNullString(emphasis),
                Timespan:     BlitzMessage.TimespanFromNullTimes(startDate, stopDate),
                Summary:      StringPtrFromNullString(summary),
            }
            educationArray = append(educationArray, &education)
        }
    }
    return educationArray
}


func NameForUserID(userID string) (string, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(`select name from UserTable where userID = $1;`, userID)
    var name sql.NullString
    error := row.Scan(&name)
    return name.String, error
}


func PrettyNameForUserID(userID string) string {
    name, _ := NameForUserID(userID)
    name = strings.TrimSpace(name)
    if len(name) == 0 {
        name = "Someone"
    }
    return name
}


//----------------------------------------------------------------------------------------
//                                                                        ProfileForUserID
//----------------------------------------------------------------------------------------


func ProfileForUserID(session *Session, userID string) *BlitzMessage.UserProfile {
    Log.Infof("ProfileForUserId (%T) %s.", userID, userID)

    rows, error := config.DB.Query(
        `select
             userID
            ,userStatus
            ,creationDate
            ,lastSeen
            ,name
            ,gender
            ,birthday
            ,backgroundSummary
            ,interestTags
            ,isExpert
            ,stripeAccount
        from UserTable where userID = $1;`, userID)
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
        profileID       string
        userStatus      sql.NullInt64
        creationDate    pq.NullTime
        lastSeen        pq.NullTime
        name            sql.NullString
        gender          sql.NullInt64
        birthday        pq.NullTime
        background      sql.NullString
        interestTags    sql.NullString
        isExpert        sql.NullBool
        stripeAccount   sql.NullString
    )
    error = rows.Scan(
        &profileID,
        &userStatus,
        &creationDate,
        &lastSeen,
        &name,
        &gender,
        &birthday,
        &background,
        &interestTags,
        &isExpert,
        &stripeAccount,
    )
    if error != nil {
        Log.Errorf("Error scanning row: %v.", error)
        return nil
    }

    profile := new(BlitzMessage.UserProfile)
    profile.UserID      = proto.String(profileID)
    profile.UserStatus  = BlitzMessage.UserStatus(userStatus.Int64).Enum()
    profile.CreationDate= BlitzMessage.TimestampFromTime(creationDate.Time)
    profile.LastSeen    = BlitzMessage.TimestampFromTime(lastSeen.Time)
    profile.Name        = proto.String(name.String)
    profile.Gender      = BlitzMessage.Gender(gender.Int64).Enum()
    profile.Birthday    = BlitzMessage.TimestampFromTime(birthday.Time)
    profile.BackgroundSummary = proto.String(background.String)
    profile.InterestTags = pgsql.StringArrayFromNullString(interestTags)

    profile.Images        = ImagesForUserID(userID)
    profile.SocialIdentities = SocialIdentitiesWithUserID(userID)
    profile.ContactInfo   = ContactInfoForUserID(userID)
    profile.EntityTags    = GetEntityTagsWithUserID(session.UserID, userID, BlitzMessage.EntityType_ETUser)
    profile.Education     = EducationForUserID(userID)
    profile.Employment    = EmploymentForUserID(userID)

    profile.IsExpert      = proto.Bool(isExpert.Bool)
    profile.StripeAccount = proto.String(stripeAccount.String)

    //  Fix up the 'headline' employment --

    for index, emp := range profile.Employment {
        if emp.IsHeadlineItem != nil && *emp.IsHeadlineItem {
            profile.HeadlineEmployment = emp
            profile.Employment = append(profile.Employment[:index], profile.Employment[index+1:]...)
            break
        }
    }

    AddReviewsToProfile(profile)

    return profile
}


func QueryProfilesByEntity(session *Session, query *BlitzMessage.UserProfileQuery,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    queryString := "select distinct entityID from EntityTagTable where entityType = $1 "
    paramArray := make([]string, 0)

    if query.EntityID != nil && len(*query.EntityID) > 0 {

        queryString = "select distinct userID from EntityTagTable where entityType = $1 "
        queryString += fmt.Sprintf(" and entityID = $%d ", len(paramArray) + 2)
        paramArray = append(paramArray, *query.EntityID)

    }

    if len(query.EntityTags) > 0 {
        queryString += fmt.Sprintf(" and entityTag = any ($%d) ", len(paramArray) + 2)
        nullstring := pgsql.NullStringFromStringArray(query.EntityTags)
        paramArray = append(paramArray, nullstring.String)
    }

    if query.EntityUserID != nil && len(*query.EntityUserID) > 0 {
        queryString += fmt.Sprintf(" and userID = $%d ", len(paramArray) + 2)
        paramArray = append(paramArray, *query.EntityUserID)
    }

    queryString += ";"

    var error error
    var rows *sql.Rows

    switch len(paramArray) {
    case 1:
        rows, error = config.DB.Query(
            queryString,
            BlitzMessage.EntityType_ETUser,
            paramArray[0],
        )

    case 2:
        rows, error = config.DB.Query(
            queryString,
            BlitzMessage.EntityType_ETUser,
            paramArray[0],
            paramArray[1],
        )

    case 3:
        rows, error = config.DB.Query(
            queryString,
            BlitzMessage.EntityType_ETUser,
            paramArray[0],
            paramArray[1],
            paramArray[2],
        )

    }
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }
    defer rows.Close()

    var profileUpdate BlitzMessage.UserProfileUpdate

    for rows.Next() {
        var userID string
        error = rows.Scan(&userID)
        if error != nil {
            Log.LogError(error)
        } else {
            profile := ProfileForUserID(session, userID)
            if profile != nil {
                profileUpdate.Profiles = append(profileUpdate.Profiles, profile)
            }
        }
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseType:       &BlitzMessage.ResponseType { UserProfileUpdate: &profileUpdate },
    }

    return response
}


func QueryProfiles(session *Session, query *BlitzMessage.UserProfileQuery,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  len(query.EntityTags) > 0 ||
        query.EntityUserID != nil ||
        query.EntityID != nil {
        return QueryProfilesByEntity(session, query)
    }

    var profileList []string

    if query.FetchDemoProfiles != nil &&
       *query.FetchDemoProfiles {
        profileList = make([]string, 0, 10)
        rows, error := config.DB.Query(
            `select userID from UserTable
                where userStatus >= $1
                limit 10;`,
            BlitzMessage.UserStatus_USConfirming,
        )
        if error != nil {
            Log.LogError(error)
            return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
        }
        defer rows.Close()
        for rows.Next() {
            var userID string
            error = rows.Scan(&userID)
            if error != nil {
                Log.LogError(error)
            } else {
                profileList = append(profileList, userID)
            }
        }
    } else {
        profileList = query.UserIDs
    }

    var profileUpdate BlitzMessage.UserProfileUpdate
    for _, userID := range profileList {
        profile := ProfileForUserID(session, userID)
        if profile != nil {
            profileUpdate.Profiles = append(profileUpdate.Profiles, profile)
        }
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseType:       &BlitzMessage.ResponseType { UserProfileUpdate: &profileUpdate },
    }
    return response
}

