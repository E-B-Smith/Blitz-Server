//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "time"
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "happiness"
)


//----------------------------------------------------------------------------------------
//                                                            UpdateContactInfoFromProfile
//----------------------------------------------------------------------------------------


func UpdateContactInfoFromProfile(profile *happiness.Profile) {
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


func AddContactInfoToProfile(profile *happiness.Profile) {
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
            ct := happiness.ContactType(contactType)
            contactStruct := happiness.ContactInfo {
                ContactType: &ct,
                Contact: &contact,
                IsVerified: &verified,
            }
            if profile.ContactInfo == nil { profile.ContactInfo = make([]*happiness.ContactInfo, 0, 5) }
            profile.ContactInfo = append(profile.ContactInfo, &contactStruct)
        } else {
            Log.LogError(error);
        }
    }
}



//----------------------------------------------------------------------------------------
//
//
//                                                                              Statistics
//
//
//----------------------------------------------------------------------------------------

//----------------------------------------------------------------------------------------
//                                                                         SummaryForQuery
//----------------------------------------------------------------------------------------


func SummaryForQuery(userID string, query string) []*happiness.Score {
    Log.LogFunctionName()

    var (rows *sql.Rows; error error)
    if userID == "" {
        rows, error = config.DB.Query(query)
    } else {
        rows, error = config.DB.Query(query, userID)
    }
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil
    }

    scoreArray := make([]*happiness.Score, 0, 10)
    for rows.Next() {
        var (
            interval time.Time
            happyscore sql.NullFloat64
            basescore sql.NullFloat64
            displayscore sql.NullFloat64
            physical sql.NullFloat64
            mental sql.NullFloat64
            vital sql.NullFloat64
            environmental sql.NullFloat64
        )
        error = rows.Scan(
            &interval,
            &happyscore,
            &basescore,
            &displayscore,
            &physical,
            &mental,
            &vital,
            &environmental,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }
        score := happiness.Score {
            Timestamp: happiness.TimestampFromTime(interval),
            HappyScore: &happyscore.Float64,
            BaseScore: &basescore.Float64,
            DisplayScore: &displayscore.Float64,
            Physical: &physical.Float64,
            Mental: &mental.Float64,
            Vital: &vital.Float64,
            Environmental: &environmental.Float64,
        }
        scoreArray = append(scoreArray, &score)
    }
    return scoreArray
}



//----------------------------------------------------------------------------------------
//                                                                        UserStatsSummary
//----------------------------------------------------------------------------------------


func UserResponseStats(userID string, startTime time.Time, stopTime time.Time) []*happiness.UserResponse {
    Log.LogFunctionName()

    result := make([]*happiness.UserResponse, 0, 10)
    rows, error :=  config.DB.Query(
        `select (r).emotionid, sum((r).emotioncount), sum((r).emotionvalue) from (
            select unnest(userresponse)::userresponse as r from scoretable
              where userid = $1 and timestamp <= $2 and timestamp > $3
        ) s
        group by (r).emotionid order by (r).emotionid;`,
        userID, startTime, stopTime,
    )
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return result
    }

    for rows.Next() {
        var (ecount int32; eid int32; evalue float32)
        error = rows.Scan(&eid, &ecount, &evalue)
        if error != nil {
            Log.LogError(error)
            continue
        }
        response := happiness.UserResponse{
            EmotionID:    &eid,
            EmotionCount: &ecount,
            EmotionValue: &evalue,
        }
        result = append(result, &response)
    }
    return result
}


func UserStatsSummary(userID string) *happiness.ScoreSummary {
    Log.LogFunctionName()

    //  Day --

    var summary *happiness.ScoreSummary
    summary = new(happiness.ScoreSummary)
    query :=
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '7 days' and userid = $1
              group by day order by day;
        `
    summary.Week = SummaryForQuery(userID, query)
    weekagodur, _ := time.ParseDuration("-168h")
    weekago := time.Now().Add(weekagodur)
    timenow := time.Now()
    if len(summary.Week) > 0 {
        summary.Week[0].UserResponse = UserResponseStats(userID, timenow, weekago)
    }

    //  Month --

    query =
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '30 days' and userid = $1
              group by day order by day;
        `
    summary.Month = SummaryForQuery(userID, query)

    //  Year ---

    query =
        `select current_timestamp - date_trunc('month', age(timestamp)) as month,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('month', age(timestamp)) < interval '12 months' and userid = $1
              group by month order by month;
        `
    summary.Year = SummaryForQuery(userID, query)
    return summary;
}


func CircleStatsSummary(userID string) *happiness.ScoreSummary {
    Log.LogFunctionName()

    //  Day --

    var summary *happiness.ScoreSummary
    summary = new(happiness.ScoreSummary)
    query :=
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '7 days'
                and (userid = $1
                or userid in
                (select friendid from friendtable where friendstatus = 5 and userid = $1))
              group by day order by day;
        `
    summary.Week = SummaryForQuery(userID, query)

    //  Month --

    query =
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '30 days'
                and (userid = $1
                or userid in
                (select friendid from friendtable where friendstatus = 5 and userid = $1))
              group by day order by day;
        `
    summary.Month = SummaryForQuery(userID, query)

    //  Year ---

    query =
        `select current_timestamp - date_trunc('month', age(timestamp)) as month,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('month', age(timestamp)) < interval '12 months'
                and (userid = $1
                or userid in
                (select friendid from friendtable where friendstatus = 5 and userid = $1))
              group by month order by month;
        `
    summary.Year = SummaryForQuery(userID, query)
    return summary;
}


func GlobalStatsSummary(userID string) *happiness.ScoreSummary {
    Log.LogFunctionName()

    //  Day --

    var summary *happiness.ScoreSummary
    summary = new(happiness.ScoreSummary)
    query :=
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '7 days'
              group by day order by day;
        `
    summary.Week = SummaryForQuery("", query)

    //  Month --

    query =
        `select current_timestamp - date_trunc('day', age(timestamp)) as day,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('day', age(timestamp)) < interval '30 days'
              group by day order by day;
        `
    summary.Month = SummaryForQuery("", query)

    //  Year ---

    query =
        `select current_timestamp - date_trunc('month', age(timestamp)) as month,
            avg(happyscore),
            avg(basescore),
            avg(displayscore),
            avg(physical),
            avg(mental),
            avg(vital),
            avg(environmental)
              from scoretable
              where date_trunc('month', age(timestamp)) < interval '12 months'
              group by month order by month;
        `
    summary.Year = SummaryForQuery("", query)
    return summary;
}


func WeatherSummary(userID string) []*happiness.WeatherSummary {
    Log.LogFunctionName()

    query :=
        `select count(*), (weather).weatherType as "wt", avg(happyscore)
            from scoretable where userid = $1
            group by "wt" order by "wt";`

    rows, error := config.DB.Query(query, userID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil
    }

    array := make([]*happiness.WeatherSummary, 0, 10)
    for rows.Next() {
        var (
            weatherCount sql.NullInt64
            weatherType  sql.NullInt64
            happyscore   sql.NullFloat64
        )
        error = rows.Scan(
            &weatherCount,
            &weatherType,
            &happyscore,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }
        var ct = int32(weatherCount.Int64)
        var hs = float32(happyscore.Float64)
        var wt = happiness.WeatherType(weatherType.Int64)
        weather := happiness.WeatherSummary {
            WeatherCount: &ct,
            HappyScore: &hs,
            WeatherType: &wt,
        }
        array = append(array, &weather)
    }
    return array
}


func HeartsSent(userID string) *int32 {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select sum(case when messagetype=8 or messagetype=9 then 1 else 0 end) as sent
            from messagetable where senderid = $1;`, userID)

    var sent sql.NullInt64
    error := row.Scan(&sent)
    if error != nil { Log.LogError(error) }
    sent32 := int32(sent.Int64)
    return &sent32
}


func HeartsReceived(userID string) *int32 {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select sum(case when messagetype=8 or messagetype=9 then 1 else 0 end) as received
            from messagetable where recipientid = $1;`, userID)

    var received sql.NullInt64
    error := row.Scan(&received)
    if error != nil { Log.LogError(error) }
    r32 := int32(received.Int64)
    return &r32
}



//----------------------------------------------------------------------------------------
//                                                                        ProfileForUserID
//----------------------------------------------------------------------------------------


func ProfileForUserID(userID string) *happiness.Profile {
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

    profile := new(happiness.Profile)
    profile.UserID      = proto.String(profileID)
    profile.UserStatus  = happiness.UserStatus(userStatus.Int64).Enum()
    profile.Name        = proto.String(name.String)
    profile.Gender      = happiness.Gender(gender.Int64).Enum()
    profile.Birthday    = happiness.TimestampFromTime(birthday.Time)
    profile.ImageURL    = pgsql.StringArrayFromNullString(imageURLs)
    Log.Debugf("Profile has %d images: %v.", len(profile.ImageURL), profile.ImageURL)
    profile.SocialIdentities = SocialIdentitiesWithUserID(userID)
    profile.CreationDate   = happiness.TimestampFromTime(creationDate.Time)
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


func QueryProfiles(writer http.ResponseWriter, userID string, profileQuery *happiness.ProfileQuery) {
    Log.LogFunctionName()

    var profileUpdate happiness.ProfileUpdate
    for i := range profileQuery.UserIDs {
        profile := ProfileForUserID(profileQuery.UserIDs[i])
        if profile != nil {
            profileUpdate.Profiles = append(profileUpdate.Profiles, profile)
        }
    }

    code := happiness.ResponseCode_RCSuccess
    var message string

    response := &happiness.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
        Response:           &happiness.ServerResponse_ProfileUpdate { ProfileUpdate: &profileUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}

