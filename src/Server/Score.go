//  Score.go  -  Update/query score data.
//
//  E.B.Smith  -  May, 2015.


package main


import (
    "fmt"
    "strings"
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "happiness"
)


//----------------------------------------------------------------------------------------
//                                                                   UpdateScoreWithUserID
//----------------------------------------------------------------------------------------



func UpdateScoreForUserID(userID string, score *happiness.Score) error {
    Log.LogFunctionName()

    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    //Log.Debugf("Updating score %v.", score)
    //Log.Debugf("Updating score UserID %s Timestamp %v.", userID, score.Timestamp)

    var locationString sql.NullString
    if score.Location != nil {
        place := ""
        if score.Location.Placename != nil { place = strings.TrimSpace(*score.Location.Placename) }
        locationString.Valid = true
        locationString.String = fmt.Sprintf("(%f, %f, %q)",
            Float64FromFloat64Ptr(score.Location.Latitude),
            Float64FromFloat64Ptr(score.Location.Longitude),
            place,
        )
    }
    var weatherString sql.NullString
    if score.Weather != nil {
        weatherString.Valid = true
        weatherString.String = fmt.Sprintf("(%d, %f, %f, %f, %d, %f, %f, %f)",
            Int32FromInt32Ptr((*int32)(score.Weather.WeatherType)),
            Float32FromFloat32Ptr(score.Weather.Temperature),
            Float32FromFloat32Ptr(score.Weather.CloudCover),
            Float32FromFloat32Ptr(score.Weather.Precipitation),
            Int32FromInt32Ptr((*int32)(score.Weather.PrecipitationType)),
            Float32FromFloat32Ptr(score.Weather.Pressure),
            Float32FromFloat32Ptr(score.Weather.WindSpeed),
            Float32FromFloat32Ptr(score.Weather.WindBearing),
        )
    }
    var userResponseString string = "NULL"
    if score.UserResponse != nil && len(score.UserResponse) > 0 {
        response := score.UserResponse
        s := fmt.Sprintf("array[ (%d, %d, %f)::userresponse",
            *response[0].EmotionID, *response[0].EmotionCount, *response[0].EmotionValue)
        for i := 1; i < len(score.UserResponse); i++ {
            s += fmt.Sprintf(", (%d, %d, %f)::userresponse",
                *response[i].EmotionID, *response[i].EmotionCount, *response[i].EmotionValue)
        }
        s += " ]"
        //Log.Debugf("User response: %s.", s)
        userResponseString = s
    }

    query := fmt.Sprintf(
        `update ScoreTable set
            (userID,
            timestamp,
            previousTimestamp,
            previousBaseScore,
            happyScore,
            basescore,
            displayscore,
            physical,
            mental,
            vital,
            environmental,
            location,
            weather,
            testID,
            userResponse,
            userTestAssessment,
            components)
            = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, %s, $15, $16)
              where userID = $17 and timestamp = $2;`,
              userResponseString)

    result, error := config.DB.Exec(query,
            &userID,
            happiness.NullTimeFromTimestamp(score.Timestamp),
            happiness.NullTimeFromTimestamp(score.PreviousTimestamp),
            score.PreviousBaseScore,
            score.HappyScore,
            score.BaseScore,
            score.DisplayScore,
            score.Physical,
            score.Mental,
            score.Vital,
            score.Environmental,
            locationString,
            weatherString,
            score.TestID,
            score.UserTestAssessment,
            happiness.NullStringFromScoreComponents(score.ScoreComponents),
            &userID,
    )

    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }

    if error == nil && rowsUpdated > 0 {
//     Log.Debugf("Updated score row %s %v.", userID, happiness.NullTimeFromTimestamp(score.Timestamp))
    } else {
//     Log.Debugf("Inserting score %s %v: %v.", userID, happiness.NullTimeFromTimestamp(score.Timestamp), error)

        //  Insert instead --

        query = fmt.Sprintf(
            `insert into ScoreTable
                (userID,
                timestamp,
                previousTimestamp,
                previousBaseScore,
                happyScore,
                basescore,
                displayscore,
                physical,
                mental,
                vital,
                environmental,
                location,
                weather,
                testID,
                userResponse,
                userTestAssessment,
                components)
                values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, %s, $15, $16);`,
                userResponseString,
        )

        _, error = config.DB.Exec(query,
            &userID,
            happiness.NullTimeFromTimestamp(score.Timestamp),
            happiness.NullTimeFromTimestamp(score.PreviousTimestamp),
            score.PreviousBaseScore,
            score.HappyScore, // 5
            score.BaseScore,
            score.DisplayScore,
            score.Physical,
            score.Mental,
            score.Vital, // 10
            score.Environmental,
            locationString,
            weatherString,
            score.TestID,
            score.UserTestAssessment, // 15
            happiness.NullStringFromScoreComponents(score.ScoreComponents),
        )
        if error == nil {
            Log.Infof("Inserted score %s %v.", userID, happiness.NullTimeFromTimestamp(score.Timestamp))
        } else {
            Log.Errorf("Error inserting score %s %v: %v.", userID, happiness.NullTimeFromTimestamp(score.Timestamp), error)
        }
    }
    return error
}


func UpdateScoresRequest(writer http.ResponseWriter,
            session *Session,
            scoreUpdate *happiness.ScoreUpdate) {

    var firstError error
    for _, score := range scoreUpdate.Scores {
        error := UpdateScoreForUserID(session.UserID, score)
        if error != nil && firstError != nil {
            firstError = error
        }
    }
    code := happiness.ResponseCode_RCSuccess
    response := happiness.ServerResponse {
        ResponseCode:       &code,
    }
    if firstError != nil {
        response.ResponseCode = happiness.ResponseCodePtr(happiness.ResponseCode_RCServerError)
        response.ResponseMessage = StringPtrFromString(firstError.Error())
    }

    data, error := proto.Marshal(&response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}


//----------------------------------------------------------------------------------------
//                                                                            ScanScoreRow
//----------------------------------------------------------------------------------------


const kScoreTableFields =
`   userID,
    timestamp,
    previousTimestamp,
    previousBaseScore,
    happyScore,
    basescore,
    displayscore,
    physical,
    mental,
    vital,
    environmental,
    (location).latitude,
    (location).longitude,
    (location).placename,
    (weather).weatherType,
    (weather).temperature,
    (weather).cloudCover,
    (weather).precipitation,
    (weather).precipitationType,
    testID,
    userResponse,
    userTestAssessment,
    components
`


type RowScanner interface {
    Scan(dest ...interface{}) error
}


func ScanScoreFromRow(rows RowScanner) (*happiness.Score, error) {
    var (
        error error;
        userID sql.NullString;
        timestamp pq.NullTime;
        previousTimestamp pq.NullTime;
        previousBaseScore sql.NullFloat64;
        happyscore sql.NullFloat64;
        basescore sql.NullFloat64;
        displayscore sql.NullFloat64;
        physical sql.NullFloat64;
        mental sql.NullFloat64;
        vital sql.NullFloat64;
        environmental sql.NullFloat64;
        locationLat sql.NullFloat64;
        locationLng sql.NullFloat64;
        locationPlc sql.NullString;
        weatherType sql.NullInt64;
        weatherTemp sql.NullFloat64;
        weatherCloud sql.NullFloat64;
        weatherP sql.NullFloat64;
        weatherPT sql.NullInt64;
        testID sql.NullString;
        userResponse sql.NullString;
        userTestAssessment sql.NullFloat64;
        components sql.NullString;
    )
    error = rows.Scan(
        &userID,
        &timestamp,
        &previousTimestamp,
        &previousBaseScore,
        &happyscore,
        &basescore,
        &displayscore,
        &physical,
        &mental,
        &vital,
        &environmental,
        &locationLat,
        &locationLng,
        &locationPlc,
        &weatherType,
        &weatherTemp,
        &weatherCloud,
        &weatherP,
        &weatherPT,
        &testID,
        &userResponse,
        &userTestAssessment,
        &components,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    score := happiness.Score {
        Timestamp: happiness.TimestampFromTime(timestamp.Time),
        PreviousTimestamp: happiness.TimestampFromTime(previousTimestamp.Time),
        PreviousBaseScore: &previousBaseScore.Float64,
        HappyScore: &happyscore.Float64,
        BaseScore: &basescore.Float64,
        DisplayScore: &displayscore.Float64,
        Physical: &physical.Float64,
        Mental: &mental.Float64,
        Vital: &vital.Float64,
        Location: &happiness.Location{ Latitude: &locationLat.Float64, Longitude: &locationLng.Float64, Placename: &locationPlc.String },
        Environmental: &environmental.Float64,
        TestID: &testID.String,
        UserTestAssessment: &userTestAssessment.Float64,
        ScoreComponents: happiness.ScoreComponentsFromNullString(components),
    }

    //Log.Debugf("UserResponse: %s.", userResponse.String)
    responses := make([]*happiness.UserResponse, 0, 6)
    recArray := strings.Split(userResponse.String, "\",\"")
    for _, s := range(recArray) {
        s = strings.TrimLeft(s, "{\"")
        s = strings.Trim(s, "\"")
        var (eID int32; eCnt int32; eVal float32)
        //Log.Debugf("%s", s)
        fmt.Sscanf(s, "(%d,%d,%f)", &eID, &eCnt, &eVal)
        response := happiness.UserResponse { EmotionID: &eID, EmotionCount: &eCnt, EmotionValue: &eVal }
        responses = append(responses, &response)
    }
    //Log.Debugf("Responses: %v.", responses)
    score.UserResponse = responses

    weather := happiness.Weather{
        WeatherType: (*happiness.WeatherType)(Int32PtrFromNullInt64(weatherType)),
        Temperature: Float32PtrFromNullFloat64(weatherTemp),
        CloudCover: Float32PtrFromNullFloat64(weatherCloud),
        Precipitation: Float32PtrFromNullFloat64(weatherP),
        PrecipitationType: (*happiness.PrecipitationType)(Int32PtrFromNullInt64(weatherPT)),
    }
    score.Weather = &weather

    return &score, nil
}


//----------------------------------------------------------------------------------------
//                                                                    LatestScoreForUserID
//----------------------------------------------------------------------------------------


func LatestScoreForUserID(userIDIn string) *happiness.Score {
    Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    query :=
        `select `+kScoreTableFields+`
            from ScoreTable where userID = $1
            order by timestamp desc limit 1;`
    row := config.DB.QueryRow(query, userIDIn)
    score, _ := ScanScoreFromRow(row)
    return score
}


//----------------------------------------------------------------------------------------
//                                                                         ScoresForUserID
//----------------------------------------------------------------------------------------


func ScoresForUserID(userIDIn string) []*happiness.Score {
    Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    query :=
        `select `+kScoreTableFields+`
            from ScoreTable where userID = $1
            order by timestamp;`

    rows, error := config.DB.Query(query, userIDIn)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.Errorf("Error reading score: %v.", error)
        return nil
    }

    array := make([]*happiness.Score, 0, 10)
    for rows.Next() {
        score, error := ScanScoreFromRow(rows)
        if error == nil {
            array = append(array, score)
        }
    }
    return array
}


//----------------------------------------------------------------------------------------
//                                                                      FetchScoresRequest
//----------------------------------------------------------------------------------------


func FetchScoresRequest(writer http.ResponseWriter,
            session *Session,
            fetchScoresRequest *happiness.FetchScoresRequest) {
    Log.LogFunctionName()

    scores := ScoresForUserID(session.UserID)
    Log.Debugf("Found %d scores for user %s.", len(scores), session.UserID)
    var scoreUpdate = happiness.ScoreUpdate {
        Scores: scores,
    }

    code := happiness.ResponseCode_RCSuccess
    response := &happiness.ServerResponse {
        ResponseCode:       &code,
        Response:           &happiness.ServerResponse_ScoreUpdate { ScoreUpdate: &scoreUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}

