//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
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
    if error != nil { Log.Errorf("Insert UserContactInfo result: %v error: %v.", result, error) }
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
            crc32           uint32
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
             isCurrentPosition
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
            isCurrentPosition   bool
            jobTitle            sql.NullString
            companyName         sql.NullString
            location            sql.NullString
            industry            sql.NullString
            startDate           pq.NullTime
            stopDate            pq.NullTime
            summary             sql.NullString
        )
        error = rows.Scan(
            &isCurrentPosition,
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
        )
        error = rows.Scan(
            &schoolName,
            &degree,
            &emphasis,
            &startDate,
            &stopDate)
        if error != nil {
            education := BlitzMessage.Education {
                SchoolName:   StringPtrFromNullString(schoolName),
                Degree:       StringPtrFromNullString(degree),
                Emphasis:     StringPtrFromNullString(emphasis),
                Timespan:     BlitzMessage.TimespanFromNullTimes(startDate, stopDate),
            }
            educationArray = append(educationArray, &education)
        }
    }
    return educationArray
}


func ExpertiseTagsForUserID(userID string) []string {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select expertiseTag from UserExpertiseTagTable where userID = $1;`, userID)
    if error != nil {
        Log.LogError(error)
        return nil
    }
    defer rows.Close()

    tags := make([]string, 0, 10)
    for rows.Next() {
        var tag string
        error = rows.Scan(&tag)
        if error == nil {
            s := strings.TrimSpace(tag)
            if len(s) > 0 { tags = append(tags, s) }
        }
    }

    return tags
}


//----------------------------------------------------------------------------------------
//                                                                        ProfileForUserID
//----------------------------------------------------------------------------------------


func ProfileForUserID(userID string) *BlitzMessage.UserProfile {
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
    profile.ExpertiseTags = ExpertiseTagsForUserID(userID)
    profile.Education     = EducationForUserID(userID)

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

