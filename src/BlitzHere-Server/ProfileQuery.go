

//----------------------------------------------------------------------------------------
//
//                                                      BlitzHere-Server : ProfileQuery.go
//                                                                    User profile queries
//
//                                                               E.B. Smith, November 2014
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "errors"
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


func CleanContactInfo(contact *BlitzMessage.ContactInfo) error {

    var BadContactInfo = errors.New("Invalid contact info")

    if contact == nil ||
       contact.ContactType == nil ||
       contact.Contact == nil {
        return BadContactInfo
    }

    switch *contact.ContactType {
    case BlitzMessage.ContactType_CTEmail:
        email, error := Util.ValidatedEmailAddress(*contact.Contact)
        if error != nil {
            contact.Contact = nil
        } else {
            contact.Contact = &email
        }

    case BlitzMessage.ContactType_CTPhoneSMS:
        phone, error := Util.ValidatedPhoneNumber(*contact.Contact)
        if error != nil {
            contact.Contact = nil
        } else {
            contact.Contact = &phone
        }

    default:
        contact.Contact = Util.CleanStringPtr(contact.Contact)

    }

    if contact.Contact == nil {
        return BadContactInfo
    }

    return nil
}


func AddContactInfoToUserID(userID string, contact *BlitzMessage.ContactInfo) {

    var error error
    error = CleanContactInfo(contact)
    if error != nil {
        return
    }

    result, error := config.DB.Exec(
        `insert into UserContactTable
        (userID, contactType, contact)
        values ($1, $2, $3);`,
        userID,
        contact.ContactType,
        Util.CleanStringPtr(contact.Contact),
    )
    if error != nil {
        Log.Errorf("Insert UserContactInfo result: %v error: %v.", result, error)
    } else {
        Log.Debugf("Added %s.", *Util.CleanStringPtr(contact.Contact))
    }
}


func UpdateContactInfoFromProfile(profile *BlitzMessage.UserProfile) {
    Log.LogFunctionName()

    if profile.UserID == nil { return }

    // result, error := config.DB.Exec("delete from UserContactTable where userID = $1;", profile.UserID)
    // if error != nil { Log.Debugf("Delete UserContactInfo result: %v error: %v.", result, error) }

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
        var (
            contactType int
            contact     string
            verified    sql.NullBool
        )
        error = rows.Scan(&contactType, &contact, &verified)
        if error == nil {
            ct := BlitzMessage.ContactType(contactType)
            contactStruct := BlitzMessage.ContactInfo {
                ContactType: &ct,
                Contact: &contact,
                IsVerified: &verified.Bool,
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
            DateAdded:      BlitzMessage.TimestampPtr(dateAdded),
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


func ProfileForUserID(sessionUserID string, userID string) *BlitzMessage.UserProfile {
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
            ,isFree
            ,editProfileID
            ,isAdmin
            ,chatCharge
            ,callCharge
            ,shortQACharge
            ,longQACharge
            ,charityPercent
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
        isFree          sql.NullBool
        editProfileID   sql.NullString
        isAdmin         sql.NullBool
        chatCharge      sql.NullString
        callCharge      sql.NullString
        shortQACharge   sql.NullString
        longQACharge    sql.NullString
        charityPercent  sql.NullFloat64
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
        &isFree,
        &editProfileID,
        &isAdmin,
        &chatCharge,
        &callCharge,
        &shortQACharge,
        &longQACharge,
        &charityPercent,
    )
    if error != nil {
        Log.Errorf("Error scanning row: %v.", error)
        return nil
    }

    profile := new(BlitzMessage.UserProfile)
    profile.UserID      = proto.String(profileID)
    profile.UserStatus  = BlitzMessage.UserStatus(userStatus.Int64).Enum()
    profile.CreationDate= BlitzMessage.TimestampPtr(creationDate.Time)
    profile.LastSeen    = BlitzMessage.TimestampPtr(lastSeen.Time)
    profile.Name        = proto.String(name.String)
    profile.Gender      = BlitzMessage.Gender(gender.Int64).Enum()
    profile.Birthday    = BlitzMessage.TimestampPtr(birthday.Time)
    profile.BackgroundSummary = proto.String(background.String)
    profile.InterestTags = pgsql.StringArrayFromNullString(interestTags)
    profile.Images       = ImagesForUserID(userID)
    profile.SocialIdentities = SocialIdentitiesWithUserID(userID)
    profile.ContactInfo   = ContactInfoForUserID(userID)
    if len(sessionUserID) > 0 {
        profile.EntityTags= GetEntityTagsWithUserID(sessionUserID, userID, BlitzMessage.EntityType_ETUser)
    }
    profile.Education     = EducationForUserID(userID)
    profile.Employment    = EmploymentForUserID(userID)

    profile.IsExpert      = proto.Bool(isExpert.Bool)
    profile.IsAdmin       = proto.Bool(isAdmin.Bool)
    profile.ChatFee       = proto.String(chatCharge.String)
    profile.CallFeePerHour= proto.String(callCharge.String)
    profile.ShortQAFee    = proto.String(shortQACharge.String)
    profile.LongQAFee     = proto.String(longQACharge.String)
    profile.CharityPercent= proto.Float64(charityPercent.Float64)
    profile.StripeAccount = proto.String(stripeAccount.String)
    if config.ServiceIsFree {
        profile.ServiceIsFreeForUser = proto.Bool(true)
    } else {
        profile.ServiceIsFreeForUser = proto.Bool(isFree.Bool)
    }
    profile.EditProfileID = proto.String(editProfileID.String)

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
            profile := ProfileForUserID(session.UserID, userID)
            if profile != nil && profile.UserStatus != nil &&
                *profile.UserStatus >= BlitzMessage.UserStatus_USConfirmed {
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
            BlitzMessage.UserStatus_USConfirmed,
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
        profile := ProfileForUserID(session.UserID, userID)
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

