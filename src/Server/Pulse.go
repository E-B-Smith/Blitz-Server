//  Pulse.go  -  Update/query HappyPulse data.
//
//  E.B.Smith  -  October, 2015.


package main


import (
    "fmt"
    "math"
    "time"
    "bytes"
    "errors"
    "strings"
    "net/url"
    "net/http"
    "database/sql"
    "html/template"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/ServerUtil"
    "github.com/golang/protobuf/proto"
    "happiness"
)


//----------------------------------------------------------------------------------------
//
//                                                                              Save Pulse
//
//----------------------------------------------------------------------------------------


func InsertPulseBeatMemberInDatabase(
        dbtransaction *sql.Tx,
        pulse *happiness.Pulse,
        beat *happiness.PulseBeat,
        member *happiness.PulseMember) error {

    Log.Debugf("Inserting (%s, %s, %v, %v).",
        *pulse.PulseID,
        happiness.TimeFromTimestamp(beat.BeatDate),
        *member.UserID,
        happiness.PulseBeatState_BSSent)

    _, error := dbtransaction.Exec(
        `insert into PulseBeatMemberTable
            (pulseID, beatDate, memberID, beatState) values ($1, $2, $3, $4);`,
        pulse.PulseID,
        happiness.NullTimeFromTimestamp(beat.BeatDate),
        member.UserID,
        happiness.PulseBeatState_BSSent)
    if error != nil {
        Log.Debugf("Error inserting member: %+v.", error)
    }
    return error
}


func InsertPulseBeatMembersInDatabase(dbtransaction *sql.Tx, pulse *happiness.Pulse, beat *happiness.PulseBeat) error {
    Log.LogFunctionName()
    var firstError error
    for _, member := range beat.Members {
        error := InsertPulseBeatMemberInDatabase(dbtransaction, pulse, beat, member)
        if error != nil && firstError == nil {
            firstError = error
        }
    }
    return firstError
}


func SaveBeatToDatabase(dbtransaction *sql.Tx, pulse *happiness.Pulse, beat *happiness.PulseBeat) error {
    Log.LogFunctionName()

    _, error := dbtransaction.Exec("savepoint SaveBeatToDatabase;")
    if error != nil {
        Log.Debugf("Savepoint error: %+v", error)
        return error
    }

    Log.Debugf("Try insert (%s, %s)...", *pulse.PulseID, happiness.TimeFromTimestamp(beat.BeatDate))
    _, error = dbtransaction.Exec(
`       insert into PulseBeatTable (
            pulseID,
            beatDate,
            expirationDate,
            updateDate,
            responseRate,
            happyScore,
            components) values ($1, $2, $3, $4, $5, $6, $7);`,
            pulse.PulseID,
            happiness.NullTimeFromTimestamp(beat.BeatDate),
            happiness.NullTimeFromTimestamp(beat.ExpirationDate),
            happiness.NullTimeFromTimestamp(beat.UpdateDate),
            beat.ResponseRate,
            beat.HappyScore,
            happiness.NullStringFromScoreComponents(beat.ScoreComponents))
    if error == nil {
        error = InsertPulseBeatMembersInDatabase(dbtransaction, pulse, beat)
        return error
    }

    Log.Debugf("Error was %+v.", error)
    _, error = dbtransaction.Exec("rollback to savepoint SaveBeatToDatabase;");
    if error != nil {
        Log.LogError(error)
        return error
    }

    Log.Debugf("Try update (%s, %s)...", *pulse.PulseID, happiness.TimeFromTimestamp(beat.BeatDate))
    var result sql.Result
    result, error = dbtransaction.Exec(
`       update PulseBeatTable set (
            pulseID,
            beatDate,
            expirationDate,
            updateDate,
            responseRate,
            happyScore,
            components) = ($1, $2, $3, $4, $5, $6, $7)
                where pulseID = $1 and beatDate = $2;`,
            pulse.PulseID,
            happiness.NullTimeFromTimestamp(beat.BeatDate),
            happiness.NullTimeFromTimestamp(beat.ExpirationDate),
            happiness.NullTimeFromTimestamp(beat.UpdateDate),
            beat.ResponseRate,
            beat.HappyScore,
            happiness.NullStringFromScoreComponents(beat.ScoreComponents))

    if pgsql.RowsUpdated(result) != 1 && error == nil {
        error = fmt.Errorf("No rows updated.")
    }
    if error == nil {
        error = InsertPulseBeatMembersInDatabase(dbtransaction, pulse, beat)
        return error
    } else {
        Log.LogError(error)
    }
    return error
}


func SaveBeatsToDatabase(dbtransaction *sql.Tx, pulse *happiness.Pulse) error {
    var firstError error
    for _, beat := range pulse.Beats {
        error := SaveBeatToDatabase(dbtransaction, pulse, beat)
        if error != nil && firstError == nil {
            firstError = error
        }
    }
    return firstError
}


func InsertPulse(dbtransaction *sql.Tx, pulse *happiness.Pulse) error {
    _, error := dbtransaction.Exec(
`       insert into PulseTable (
             pulseID
            ,senderID
            ,title
            ,body
            ,color
            ,teamIsVisible
            ,creationDate
            ,updateDate
            ,testID
            ,pulseStatus
            ) values
            ($1, $2, $3, $4, ($5, $6, $7)::ColorRGB256, $8, $9, $10, $11, $12);`,
            pulse.PulseID,
            pulse.SenderID,
            pulse.Title,
            pulse.Body,
            pulse.Color.Red, pulse.Color.Green, pulse.Color.Blue,
            pulse.TeamIsVisible,
            happiness.NullTimeFromTimestamp(pulse.CreationDate),
            happiness.NullTimeFromTimestamp(pulse.UpdateDate),
            pulse.TestID,
            pulse.PulseStatus)
    return error
}


func UpdatePulse(dbtransaction *sql.Tx, pulse *happiness.Pulse) error {
    _, error := dbtransaction.Exec(
`       update PulseTable set (
             pulseID
            ,senderID
            ,title
            ,body
            ,color
            ,teamIsVisible
            ,creationDate
            ,updateDate
            ,testID
            ,pulseStatus
            ) =
            ($1, $2, $3, $4, ($5, $6, $7)::ColorRGB256, $8, $9, $10, $11, $12)
            where pulseID = $1;`,
            pulse.PulseID,
            pulse.SenderID,
            pulse.Title,
            pulse.Body,
            pulse.Color.Red, pulse.Color.Green, pulse.Color.Blue,
            pulse.TeamIsVisible,
            happiness.NullTimeFromTimestamp(pulse.CreationDate),
            happiness.NullTimeFromTimestamp(pulse.UpdateDate),
            pulse.TestID,
            pulse.PulseStatus)
    return error
}


func SavePulseToDatabase(pulse *happiness.Pulse) error {
    Log.LogFunctionName()

    dbtransaction, error := config.DB.Begin()
    if error != nil {
        Log.LogError(error)
        return error
    }
    defer func() {
        if dbtransaction != nil {
            error := dbtransaction.Rollback()
            if error != nil {
                Log.Errorf("Error rolling back: %+v.", error)
            }
        }
    } ()

    _, error = dbtransaction.Exec("savepoint SavePulseToDatabase;")
    if error != nil {
        Log.Debugf("Savepoint error: %+v", error)
        return error
    }

    error = InsertPulse(dbtransaction, pulse)
    if error != nil {
        Log.Debugf("Insert pulse error: %+v.", error)
        _, error = dbtransaction.Exec("rollback to savepoint SavePulseToDatabase;");
        if error != nil {
            Log.LogError(error)
            return error
        }
        error = UpdatePulse(dbtransaction, pulse)
        if error != nil {
            Log.LogError(error)
            return error
        }
    }

    // error = SaveBeatsToDatabase(dbtransaction, pulse)
    // if error != nil {
    //     Log.LogError(error)
    //     return error
    // }

    //  Only save the last beat --

    if len(pulse.Beats) == 0 {
        error = fmt.Errorf("No Pulse Beats to send.")
        Log.LogError(error)
        return error
    }
    beat := pulse.Beats[len(pulse.Beats)-1]
    error = SaveBeatToDatabase(dbtransaction, pulse, beat)
    if error != nil {
        Log.LogError(error)
        return error
    }

    error = dbtransaction.Commit()
    if error == nil {
        dbtransaction = nil
    } else {
        Log.LogError(error)
    }
    return error
}


//----------------------------------------------------------------------------------------
//
//                                                                              Load Pulse
//
//----------------------------------------------------------------------------------------


func MembersForPulseBeat(pulseID string, beatDate time.Time) ([]*happiness.PulseMember, error) {
    //  Load member info for a pulse beat --

    rows, error := config.DB.Query(
        "select memberID, scoreDate from PulseBeatMemberTable where pulseID = $1 and beatDate = $2;",
            pulseID, beatDate)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    members := make([]*happiness.PulseMember, 0)
    for rows.Next() {
        var (
            memberID         string
            scoreDate        pq.NullTime
        )
        error = rows.Scan(&memberID, &scoreDate)
        if error != nil {
            Log.LogError(error)
            return nil, error
        }
        member := happiness.PulseMember {
            UserID:         &memberID,
            ScoreDate:      happiness.TimestampPtrFromNullTime(scoreDate),
        }
        members = append(members, &member)
    }
    return members, nil
}


func BeatsWithPulseID(pulseID string) ([]*happiness.PulseBeat, error) {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select
            beatdate,
            expirationdate,
            updatedate,
            responserate,
            happyscore,
            components::text
            from PulseBeatTable where pulseID = $1
             order by beatDate;`,
                pulseID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil,  error
    }

    beats := make([]*happiness.PulseBeat, 0)
    for rows.Next() {
        var (
            beatDate           pq.NullTime
            expirationDate     pq.NullTime
            updateDate         pq.NullTime
            responseRate       sql.NullFloat64
            happyScore         sql.NullFloat64
            componentsText     sql.NullString
        )
        error = rows.Scan(
            &beatDate,
            &expirationDate,
            &updateDate,
            &responseRate,
            &happyScore,
            &componentsText)
        if error != nil {
            Log.LogError(error)
            return nil, error
        }
        components := happiness.ScoreComponentsFromNullString(componentsText)
        beat := &happiness.PulseBeat {
            BeatDate:          happiness.TimestampPtrFromNullTime(beatDate),
            ExpirationDate:    happiness.TimestampPtrFromNullTime(expirationDate),
            UpdateDate:        happiness.TimestampPtrFromNullTime(updateDate),
            ResponseRate:      Float64PtrFromNullFloat(responseRate),
            HappyScore:        Float64PtrFromNullFloat(happyScore),
            ScoreComponents:   components,
        }
        beat.Members, error = MembersForPulseBeat(pulseID, beatDate.Time)
        if error != nil { return nil, error }
        Log.Debugf("Found %d beat members.", len(beat.Members))
        beats = append(beats, beat)
    }

    return beats, nil
}


func PulseWithPulseID(pulseID string) (*happiness.Pulse, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
            pulseid,
            senderid,
            title,
            body,
            (color).red,
            (color).green,
            (color).blue,
            teamisvisible,
            creationdate,
            updatedate,
            testid,
            pulseStatus
            from PulseTable where pulseID = $1;`,
        pulseID)
    var (
        PulseID     string
        SenderID    string
        Title       sql.NullString
        Body        sql.NullString
        ColorR      sql.NullInt64
        ColorG      sql.NullInt64
        ColorB      sql.NullInt64
        TeamIsVisible bool
        CreationDate pq.NullTime
        UpdateDate   pq.NullTime
        TestID      string
        PulseStatusInt sql.NullInt64
    )
    error := row.Scan(
        &PulseID,
        &SenderID,
        &Title,
        &Body,
        &ColorR,
        &ColorG,
        &ColorB,
        &TeamIsVisible,
        &CreationDate,
        &UpdateDate,
        &TestID,
        &PulseStatusInt)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    var PulseStatus happiness.PulseStatus = happiness.PulseStatus_PSUpdated
    if PulseStatusInt.Valid { PulseStatus = happiness.PulseStatus(PulseStatusInt.Int64) }

    pulse := happiness.Pulse {
        PulseID:        &PulseID,
        SenderID:       &SenderID,
        Title:          StringPtrFromNullString(Title),
        Body:           StringPtrFromNullString(Body),
        Color:          happiness.ColorRGB256Ptr(int32(ColorR.Int64), int32(ColorG.Int64), int32(ColorB.Int64)),
        TeamIsVisible:  &TeamIsVisible,
        CreationDate:   happiness.TimestampPtrFromNullTime(CreationDate),
        UpdateDate:     happiness.TimestampPtrFromNullTime(UpdateDate),
        TestID:         &TestID,
        PulseStatus:    &PulseStatus,
    }

    pulse.Beats, error = BeatsWithPulseID(pulseID)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    Log.Debugf("Found %d beats.", len(pulse.Beats))
    return &pulse, nil
}


//----------------------------------------------------------------------------------------
//
//                                                                      Pulse Maintainence
//
//----------------------------------------------------------------------------------------


func PulsesForUser(userID string) ([]*happiness.Pulse, error) {
    Log.LogFunctionName()

    //  Get sent pulses --

    rows, error := config.DB.Query("select pulseID from PulseTable where senderID = $1;", userID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    Log.Debugf("Getting sent pulses...")
    var pulses []*happiness.Pulse
    for rows.Next() && error == nil {

        var pulseID string
        error = rows.Scan(&pulseID)
        if error != nil { return nil, error }

        var pulse *happiness.Pulse
        pulse, error = PulseWithPulseID(pulseID)
        if error != nil { return nil, error }
        pulses = append(pulses, pulse)
    }

    Log.Debugf("Getting received pulses...")
    rows, error = config.DB.Query(
        "select pulseID, memberPulseStatus from PulseBeatMemberTable where memberID = $1;", userID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    for rows.Next() && error == nil {

        var (pulseID string; memberPulseStatus sql.NullInt64)
        error = rows.Scan(&pulseID, &memberPulseStatus)
        if error != nil { return nil, error }

        var pulse *happiness.Pulse
        pulse, error = PulseWithPulseID(pulseID)
        if error != nil { return nil, error }

        var status happiness.PulseStatus = happiness.PulseStatus_PSUpdated
        if pulse.PulseStatus != nil {
            status = *pulse.PulseStatus
        }
        if memberPulseStatus.Valid &&
           status != happiness.PulseStatus_PSDeleted {
           status = happiness.PulseStatus(memberPulseStatus.Int64)
        }
        pulse.PulseStatus = &status
        Log.Debugf("Member status: %s.", happiness.PulseStatus_name[int32(*pulse.PulseStatus)])
        pulses = append(pulses, pulse)

    }
    rows.Close()

    return pulses, nil
}


func GetPulsesForUser(writer http.ResponseWriter,
            session *Session,
            pulsesForUser *happiness.PulseUpdate) {
    Log.LogFunctionName()

    //  Get the pulses --

    pulses, error := PulsesForUser(session.UserID)
    if error != nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Add the profiles for all the members --

    memberMap := make(map[string]bool)
    for _, pulse := range pulses {
        memberMap[*pulse.SenderID] = true
        for _, beat := range pulse.Beats {
            for _, member := range beat.Members {
                memberMap[*member.UserID] = true
            }
        }
    }

    profiles := make([]*happiness.Profile, 0, len(memberMap))
    for userID, _ := range memberMap {
        profile := ProfileForUserID(userID)
        if profile != nil {
            profiles = append(profiles, profile)
        }
    }

    //  Marshal the response --

    pulseUpdate := happiness.PulseUpdate {
        Pulses:         pulses,
        MemberProfiles: profiles,
    }
    response := &happiness.ServerResponse {
        ResponseCode:   happiness.ResponseCodePtr(happiness.ResponseCode_RCSuccess),
        Response:       &happiness.ServerResponse_PulseUpdate { PulseUpdate:    &pulseUpdate},
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}


func UpdatePulseStatus(writer http.ResponseWriter,
            session *Session,
            pulseStatusUpdate *happiness.PulseStatusUpdate) {
    Log.LogFunctionName()

    for _, status := range pulseStatusUpdate.Status {
        if  status.PulseStatus == nil  || status.PulseID == nil {
            continue
        }

        if *status.PulseStatus == happiness.PulseStatus_PSDeleted {
            result, error := config.DB.Exec(
                `update PulseTable set
                    pulseStatus = $1,
                    updateDate = now()
                    where pulseID = $2;`,
                status.PulseStatus, status.PulseID)
            Log.Debugf("Updated %d pulses.", pgsql.RowsUpdated(result))
            if error != nil || pgsql.RowsUpdated(result) != 1 {
                Log.LogError(error)
            }
        } else {
            result, error := config.DB.Exec(
                `update PulseBeatMemberTable set
                    memberPulseStatus = $1
                    where pulseID = $2 and memberID = $3;`,
                status.PulseStatus, status.PulseID, session.UserID)
            Log.Debugf("Updated %d pulses.", pgsql.RowsUpdated(result))
            if error != nil || pgsql.RowsUpdated(result) == 0 {
                Log.LogError(error)
            }
        }
    }

    response := &happiness.ServerResponse {
        ResponseCode:   happiness.ResponseCodePtr(happiness.ResponseCode_RCSuccess),
    }
    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}



//----------------------------------------------------------------------------------------
//
//                                                                               SendPulse
//
//----------------------------------------------------------------------------------------


func SendSinglePulseBeat(pulse happiness.Pulse) error {

    //  * Assumes the pulse has been validated.
    //  * Save the pulse to the database.
    //  * For each user:
    //    - If existing user and has notifications on, send a notification.
    //    - Otherwise send an email.

    Log.LogFunctionName()
    var error error

    //  Check the pulse --

    if len(pulse.Beats) == 0 {
        error = fmt.Errorf("No beat to send.")
        return error
    }
    beat := pulse.Beats[len(pulse.Beats)-1]

    //  Fix the dates --

    if pulse.CreationDate == nil {
        pulse.CreationDate = happiness.TimestampFromTime(time.Now())
    }
    pulse.UpdateDate = happiness.TimestampFromTime(time.Now())

    //  Save the pulse --

    error = SavePulseToDatabase(&pulse)
    if error != nil { return error }

    //  Build app-link URL --

    senderProfile := ProfileForUserID(*pulse.SenderID)
    if senderProfile == nil {
        error = fmt.Errorf("The Pulse sender doesn't exist.")
        return error
    }

    senderName := "Someone"
    if senderProfile.Name != nil {
        senderName = *senderProfile.Name
    }
    senderEmail := "< No email address >"
    senderContact := senderProfile.ContactInfoOfType(happiness.ContactType_CTEmail)
    if senderContact != nil && senderContact.Contact != nil {
        senderEmail = *senderContact.Contact
    }

    //
    //  Send emails & messages --
    //

    replyDate := happiness.TimeFromTimestamp(beat.ExpirationDate)
    replyDateString := replyDate.Format("Monday January 2, 3:04:05 pm MST")

    for _, member := range beat.Members {

        message := fmt.Sprintf("%s sent you a new HappyPulse.", senderName)
        longURL := fmt.Sprintf("%s/?action=showpulse&pulse=%s&beat=%f&message=%s&senderid=%s&userid=%s",
                config.AppLinkURL,
                *pulse.PulseID,
                *beat.BeatDate.Epoch,
                url.QueryEscape(message),
                *pulse.SenderID,
                *member.UserID)
        var shortURL string
        shortURL, error = LinkShortner_ShortLinkFromLink(longURL)
        if error != nil {
            Log.LogError(error)
            continue
        }

        SendAppMessage(*pulse.SenderID, []string{ *member.UserID }, message, happiness.MessageType_MTPulse, "Pulse", longURL)

        var templateMap = struct {
            SenderName      string
            SenderEmail     string
            PulseTitle      string
            ReplyDate       string
            AppDeepLink     template.HTML
            TeamMemberName  string
            TeamMemberID    string
        }{
            senderName,
            senderEmail,
            *pulse.Title,
            replyDateString,
            template.HTML(shortURL),
            *member.Name,
            *member.UserID,
        }

        var emailBuffer bytes.Buffer
        subject := fmt.Sprintf("%s sent you a HappyPulse", senderName)
        error = config.Template.ExecuteTemplate(&emailBuffer, "NewPulseEmail.html", templateMap)
        if error != nil {
            Log.LogError(error)
        } else {
            Log.Debugf("Email: %s\n%s\n%s.", *member.ContactInfo.Contact, subject, emailBuffer.String())
            error = config.SendEmail(*member.ContactInfo.Contact, subject, emailBuffer.String())
            if error != nil { Log.LogError(error) }
        }
    }

    return nil
}


func SendNewPulseBeat(writer http.ResponseWriter,
            session *Session,
            sendPulse *happiness.SendPulse) {
    Log.LogFunctionName()

    //  * Re-validate the pulse.
    //  * Make sure the product ID is correct.
    //  * Validate the receipt with Apple.
    //  * Send the pulse.

    if  sendPulse.ValidatedPulse == nil ||
        sendPulse.ValidatedPulse.Pulse == nil ||
        sendPulse.ValidatedPulse.Beat == nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("Invalid Pulse"))
        return
    }

    var error error
    code, message, productID := ValidatePulseBeat(session, sendPulse.ValidatedPulse.Pulse, sendPulse.ValidatedPulse.Beat)
    if code != happiness.ResponseCode_RCSuccess {
        SendError(writer, code, errors.New(message))
        return
    }

    if session.Device.ModelName != nil {
        model := strings.ToLower(*session.Device.ModelName)
        if strings.Contains(model, "simulator") {
            //  Don't validate in-app purchases made in the simulator --
            sendPulse.ValidatedPulse.ProductID = nil
            productID = ""
        }
    }

    if sendPulse.ValidatedPulse.ProductID != nil && len(*sendPulse.ValidatedPulse.ProductID) == 0 {
        sendPulse.ValidatedPulse.ProductID = nil
    }

    if sendPulse.TransactionID != nil && len(*sendPulse.TransactionID) == 0 {
        sendPulse.TransactionID = nil
    }

    if (sendPulse.ValidatedPulse.ProductID == nil && len(productID) > 0) {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("Invalid product ID"))
        return
    }

    if  sendPulse.ValidatedPulse.ProductID != nil && sendPulse.TransactionID == nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("Invalid transaction ID"))
        return
    }

    var storeReceipt *ServerUtil.AppleInAppReceipt = nil
    if sendPulse.ValidatedPulse.ProductID != nil && sendPulse.TransactionID != nil {
        storeTransaction, error := StoreTransactionWithTransactionID("Apple", *sendPulse.TransactionID)
        if storeTransaction != nil {
            SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("Pulse already sent"))
            return
        }
        storeReceipt, error = ServerUtil.ValidateAppleReceiptTransaction(sendPulse.ReceiptData, *sendPulse.TransactionID)
        if error != nil {
            SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
            return
        }
    }

    pulse := sendPulse.ValidatedPulse.Pulse
    pulse.Beats = append(pulse.Beats, sendPulse.ValidatedPulse.Beat)
    error = SendSinglePulseBeat(*pulse)

    code = happiness.ResponseCode_RCSuccess
    if error != nil {
        message = error.Error()
        code = happiness.ResponseCode_RCInputInvalid
    }
    response := &happiness.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
    }
    if error == nil {
        updatedPulse, _ := PulseWithPulseID(*pulse.PulseID)
        pulseUpdate := happiness.PulseUpdate {
            Pulses: []*happiness.Pulse{updatedPulse},
        }
        response.Response = &happiness.ServerResponse_PulseUpdate { PulseUpdate: &pulseUpdate }
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }
    writer.Write(data)


    if sendPulse.TransactionID != nil {
        storeTransaction := StoreTransaction {
            StoreID:            "Apple",
            StoreTransactionID: *sendPulse.TransactionID,
            UserID:             *sendPulse.ValidatedPulse.Pulse.SenderID,
            Purchase:           *sendPulse.ValidatedPulse.ProductID,
            PurchaseDate:       TimeFromRFC3339(&storeReceipt.PurchaseDate),
            Locale:             StringFromPtr(sendPulse.Locale),
            LocalizedPrice:     StringFromPtr(sendPulse.LocalizedPrice),
        }
        StoreTransactionInsert(storeTransaction)
    }
}


func ValidatePulseBeat(
    session *Session,
    pulse *happiness.Pulse,
    beat *happiness.PulseBeat) (
        code happiness.ResponseCode,
        message string,
        productID string) {

    //  Check the pulse --

    productID = ""
    message = ""
    code = happiness.ResponseCode_RCInputInvalid

    if pulse.PulseID == nil {
        message = "Invalid pulse ID."
        return
    }
    if pulse.Title == nil || len(*pulse.Title) == 0 {
        message = "The pulse needs a topic."
        return
    }

    var senderProfile *happiness.Profile = nil
    if  pulse.SenderID != nil {
        senderProfile = ProfileForUserID(*pulse.SenderID)
    }
    if senderProfile == nil {
        message = "The Pulse sender doesn't exist."
        return
    }

    //  Add the sender too --

    senderMember := happiness.PulseMember {
        UserID:         pulse.SenderID,
        Name:           senderProfile.Name,
        ContactInfo:    senderProfile.ContactInfoOfType(happiness.ContactType_CTEmail),
        Profile:        senderProfile,
    }
    beat.Members = append(beat.Members, &senderMember)
    Log.Debugf("Beat has %d members.", len(beat.Members))

    //  Make sure that the userid / contact match.
    //  - If they don't match, fix it.
    //  - If no user exists, create one.

    memberMap := make(map[string]*happiness.PulseMember)
    for _, member := range beat.Members {

        profilesIn := happiness.Profile {
            UserID:         member.UserID,
            Name:           member.Name,
            ContactInfo:    []*happiness.ContactInfo { member.ContactInfo },
        }
        profilesOut := ProfilesFromContactInfo([]*happiness.Profile { &profilesIn })
        if len(profilesOut) == 0 {
            Log.Debugf("Nil profile.")
            continue
        }

        memberName := ""
        if member.Name != nil {
            memberName = strings.TrimSpace(*member.Name)
        }
        profileName := ""
        if profilesOut[0].Name != nil {
            profileName = strings.TrimSpace(*profilesOut[0].Name)
        }

        var namePtr *string = nil
        if len(profileName) == 0 {
            if len(memberName) != 0 {
                namePtr = &memberName
                profilesOut[0].Name = namePtr
                UpdateProfile(profilesOut[0])
            }
        } else {
            namePtr = &profileName
        }

        fixedMember := happiness.PulseMember {
            UserID:         profilesOut[0].UserID,
            Name:           namePtr,
            ContactInfo:    profilesOut[0].ContactInfoOfType(happiness.ContactType_CTEmail),
            Profile:        profilesOut[0],
        }

        Log.Debugf("---")
        if fixedMember.UserID != nil { Log.Debugf(" Adding: %s.", *fixedMember.UserID) }
        if fixedMember.Name   != nil { Log.Debugf("   Name: %s.", *fixedMember.Name) }
        if fixedMember.ContactInfo != nil && fixedMember.ContactInfo.Contact != nil {
            Log.Debugf("Contact: %s.", *fixedMember.ContactInfo.Contact)
        }
        memberMap[*fixedMember.UserID] = &fixedMember
    }

    var missingName bool = false
    var missingContact bool = false

    beat.Members = make([]*happiness.PulseMember, 0, len(memberMap))
    for _, member := range memberMap {
        if member.Name == nil || len(*member.Name) == 0 { missingName = true }
        if member.ContactInfo == nil ||
           member.ContactInfo.Contact == nil ||
           len(*member.ContactInfo.Contact) == 0 {
            missingContact = true
        }
        beat.Members = append(beat.Members, member)
    }
    Log.Debugf("Final beat has %d members.", len(beat.Members))

    memberCount := session.AppOptions.GetPulseOptions().MinimumMembersForPulse
    if memberCount == nil { memberCount = Int32PtrFromInt32(3) }

    if len(beat.Members) < int(*memberCount) {
        message = fmt.Sprintf("At least %d team members are needed to send a pulse.", *memberCount)
        return
    }

    if missingContact {
        message = "A team member is missing an email address."
        return
    }

    if missingName {
        message = "A team member is missing a name."
        return
    }

    if ! config.PulsesAreFree {
        if session.Device.AppID == nil {
            message = "App ID is missing from session."
            return
        }
        productID = "1011"  //  *session.Device.AppID + ".singlepulse"
    }

    code = happiness.ResponseCode_RCSuccess
    return
}


func ValidatePulseRequest(writer http.ResponseWriter,
            session *Session,
            validatePulse *happiness.ValidatedPulse) {
    Log.LogFunctionName()

    code, message, productID := ValidatePulseBeat(session, validatePulse.Pulse, validatePulse.Beat)
    validatePulse.ProductID = &productID

    response := &happiness.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    &message,
        Response:           &happiness.ServerResponse_ValidatePulseResponse { ValidatePulseResponse: validatePulse },
    }
    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}


//----------------------------------------------------------------------------------------
//
//                                                                              ScorePulse
//
//----------------------------------------------------------------------------------------


//------------------------------------------------------------------ ScorePulseBeatRequest


func ScorePulseBeatRequest(writer http.ResponseWriter,
            session *Session,
            beatScore *happiness.ScorePulseBeat) {
    Log.LogFunctionName()

    dbtransaction, error := config.DB.Begin()
    if error != nil {
        Log.LogError(error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }
    defer func() {
        if dbtransaction == nil { return }
        error := dbtransaction.Rollback()
        if error != nil {
            Log.Errorf("Error rolling back: %+v.", error)
        }
    } ()

    //  Get the test --

    var testID string
    row := dbtransaction.QueryRow("select testID from PulseTable where pulseID = $1;", beatScore.PulseID)
    error = row.Scan(&testID)
    if error != nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Score the test --

    score, error := ScoreForTest(session.UserID, testID, beatScore.Responses, beatScore.Location, beatScore.Weather)
    if error != nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }
    score.Timestamp = beatScore.ScoreDate

    if score.Timestamp == nil || beatScore.PulseID == nil || beatScore.BeatDate == nil {
        error = fmt.Errorf("Beat error (%v, %v, %v).", score.Timestamp, beatScore.PulseID, beatScore.BeatDate)
        Log.LogError(error)
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Write the score --

    Log.Debugf("Score: %+v.", score)
    //error = UpdateScoreForUserIDTx(dbtransaction, session.UserID, score)
    error = UpdateScoreForUserID(session.UserID, score)
    if error != nil {
        Log.LogError(error)
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Update the pulse beat member table --

    var result sql.Result
    result, error = dbtransaction.Exec(
        `update PulseBeatMemberTable set
            (scoreDate, beatState) = ($1, $2)
            where pulseID = $3 and beatDate = $4 and memberID = $5;`,
            happiness.TimeFromTimestamp(score.Timestamp),
            happiness.PulseBeatState_BSComplete,
            beatScore.PulseID,
            happiness.TimeFromTimestamp(beatScore.BeatDate),
            session.UserID)
    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }
    if error != nil || rowsUpdated != 1 {
        if error == nil { error = fmt.Errorf("Member not in pulse beat.") }
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Update the pulse date --

    result, error = dbtransaction.Exec(
        `update PulseTable set updateDate = $1 where pulseID = $2;`,
            happiness.TimeFromTimestamp(score.Timestamp),
            beatScore.PulseID)
    if error != nil || pgsql.RowsUpdated(result) != 1 {
        Log.LogError(error)
        SendError(writer, happiness.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid pulseID"))
        return
    }

    //  Commit the db --

    error = dbtransaction.Commit()
    if error == nil {
        dbtransaction = nil
    } else {
        Log.LogError(error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    //  Update the aggregate the beat data --

    error = AggregatePulseBeatScores(*beatScore.PulseID, happiness.TimeFromTimestamp(beatScore.BeatDate))
    if error != nil {
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    //  Start marshaling the response --

    pulse, error := PulseWithPulseID(*beatScore.PulseID)
    if error != nil {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, error)
        return
    }

    //  Send a message if the pulse is fulfilled --

    var pulseBeat *happiness.PulseBeat
    for _, pulseBeat = range pulse.Beats {
        if *pulseBeat.BeatDate.Epoch == *beatScore.BeatDate.Epoch { break }
    }

    if pulseBeat != nil && pulseBeat.ResponseRate != nil  && *pulseBeat.ResponseRate >= 1.0 {
        message := fmt.Sprintf("All team members have responded to your pulse '%s'.",
            Util.TruncateStringToLength(*pulse.Title, 20))
        url := fmt.Sprintf("%s/?action=showpulse&pulse=%s&beat=%f",
            config.AppLinkURL,
            *pulse.PulseID,
            *beatScore.BeatDate.Epoch)
        SendAppMessage(happiness.Default_Globals_SystemUserID, []string{ *pulse.SenderID }, message, happiness.MessageType_MTPulse, "Pulse", url)
    }

    updatedPulse, error := PulseWithPulseID(*pulse.PulseID)
    if error != nil {
        Log.LogError(error)
        updatedPulse = pulse
    }

    pulseUpdate := happiness.PulseUpdate {
        Pulses: []*happiness.Pulse{updatedPulse},
        Scores: []*happiness.Score{score},
    }
    response := &happiness.ServerResponse {
        ResponseCode:   happiness.ResponseCodePtr(happiness.ResponseCode_RCSuccess),
        Response:       &happiness.ServerResponse_PulseUpdate { PulseUpdate: &pulseUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)

}


//--------------------------------------------------------------------------- ScoreForTest


//  Avg.  Male  Female Team Trust  Motiv  Satis

var globalTestTable = [9][7]float64 {
    {  4.5,  5,  4,  2,  5,  5,  4 },
    { -3.5, -3, -4, -2,  0,  2, -2 },
    {  3.0,  3,  3,  5,  2,  5,  4 },
    {  4.0,  5,  3,  4,  3,  4,  3 },
    { -3.5, -3, -4, -3, -1, -5, -3 },
    { -4.5, -5, -4,  0, -2,  0,  0 },
    { -3.5, -4, -3, -4, -1, -5, -2 },
    { -4.0, -4, -4, -5, -4, -1, -5 },
    {  3.5,  3,  4, -3,  2,  4,  4 },
}


func ScoreForTest(
        userID string,
        testID string,
        responses []*happiness.UserResponse,
        location *happiness.Location,
        weather *happiness.Weather) (*happiness.Score, error) {

    if (len(responses) == 0) {
        return nil, fmt.Errorf("Not enough responses")
    }

    //  Score the user's response --

    Log.LogFunctionName()
    components := []*happiness.ScoreComponent {
        &happiness.ScoreComponent { Label: StringPtrFromString("average"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("male"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("female"), Score: Float64PtrFromFloat64(0.0) },

        &happiness.ScoreComponent { Label: StringPtrFromString("team"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("trust"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("motivation"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("satisfaction"), Score: Float64PtrFromFloat64(0.0) },
    }

    for _, response := range responses {
        for index, _ := range components {
            if response.EmotionID != nil && *response.EmotionID >= 0 && *response.EmotionID < 9 {
                *components[index].Score += globalTestTable[*response.EmotionID][index]
            }
        }
    }

    for index, _ := range components {
        *components[index].Score /= float64(len(responses))
        *components[index].Score  = (*components[index].Score + 5.0) / 10.0
    }

    baseScore := *components[0].Score
    happyScore := baseScore
    displayScore := math.Max(math.Min(happyScore, 0.910), 0.310)

    score := happiness.Score {
        TestID:         &testID,
        Location:       location,
        Weather:        weather,
        UserResponse:   responses,
        HappyScore:     Float64PtrFromFloat64(happyScore),
        BaseScore:      Float64PtrFromFloat64(baseScore),
        DisplayScore:   Float64PtrFromFloat64(displayScore),
        ScoreComponents:components[3:],
    }

    return &score, nil
}


//--------------------------------------------------------------- AggregatePulseBeatScores


func AggregatePulseBeatScores(pulseID string, beatDate time.Time) error {

    Log.LogFunctionName()
    rows, error := config.DB.Query(
        `select
            PulseBeatMemberTable.memberID,
            PulseBeatMemberTable.scoreDate,
            ScoreTable.happyScore,
            ScoreTable.components::text
            from PulseBeatMemberTable
            left join ScoreTable
            on PulseBeatMemberTable.memberID = ScoreTable.userID and
               PulseBeatMemberTable.scoreDate = ScoreTable.timestamp
            where PulseBeatMemberTable.pulseID = $1
              and PulseBeatMemberTable.beatDate = $2;`,
        pulseID,
        beatDate)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return error
    }

    var (
        members int
        responders int
        totalHappy float64
    )

    totals := []*happiness.ScoreComponent {
        &happiness.ScoreComponent { Label: StringPtrFromString("team"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("trust"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("motivation"), Score: Float64PtrFromFloat64(0.0) },
        &happiness.ScoreComponent { Label: StringPtrFromString("satisfaction"), Score: Float64PtrFromFloat64(0.0) },
    }

    for rows.Next() {
        var (
            memberID    string
            scoreDate   pq.NullTime
            happyScore  sql.NullFloat64
            componentsText sql.NullString
        )
        error = rows.Scan(&memberID, &scoreDate, &happyScore, &componentsText)
        if error != nil {
            Log.LogError(error)
            return error
        }
        members++
        if scoreDate.Valid {
            responders++
            totalHappy += happyScore.Float64
            components := happiness.ScoreComponentsFromNullString(componentsText)
            count := MinInt(len(components), len(totals))
            Log.Debugf("Summing %d components:\n%+v\n%+v.", count, components, componentsText)
            for i := 0; i < count; i++ {
                *totals[i].Score += *components[i].Score
            }
        }
    }

    if members == 0 { return nil }
    if responders > 0 {
        totalHappy /= float64(responders)
        for i, _ := range totals {
            *totals[i].Score /= float64(responders)
        }
    } else {
        totalHappy = 0
        for i, _ := range totals {
            *totals[i].Score = 0
        }
    }

    responseRate := float64(responders) / float64(members)

    result, error := config.DB.Exec(
`       update PulseBeatTable set (
            updateDate,
            responseRate,
            happyScore,
            components) = ($1, $2, $3, $4)
                where pulseID = $5 and beatDate = $6;`,
        time.Now(),
        responseRate,
        totalHappy,
        happiness.NullStringFromScoreComponents(totals[:]),
        pulseID,
        beatDate)

    if error != nil || pgsql.RowsUpdated(result) == 0 {
        Log.Errorf("No rows updated. Error: %+v.", error)
    }

    return error
}


//----------------------------------------------------------------------------------------
//
//                                                                      SendPulseReminders
//
//----------------------------------------------------------------------------------------


func SendPulseReminders() {
    Log.LogFunctionName()
    rows, error := config.DB.Query(
        `select
                PulseBeatTable.pulseID,
                PulseBeatMemberTable.beatDate,
                PulseBeatMemberTable.memberID
            from PulseBeatTable
            join PulseBeatMemberTable
              on PulseBeatMemberTable.pulseID = PulseBeatTable.pulseID
            where PulseBeatTable.expirationDate > now()
            and   PulseBeatTable.expirationDate < now() + '1 hour'::interval
            and   PulseBeatMemberTable.beatState < $1;`,
            happiness.PulseBeatState_BSReminded)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return
    }

    var remindersSent int = 0
    message := "You have a pulse that needs a response soon."

    for rows.Next() {
        var (
            pulseID     string
            beatDate    time.Time
            memberID    string
        )
        error = rows.Scan(&pulseID, &beatDate, &memberID)
        if error != nil {
            Log.LogError(error)
            continue
        }
        url := fmt.Sprintf("%s/?action=showpulse&pulse=%s&beat=%d",
            config.AppLinkURL,
            pulseID,
            beatDate.Unix())

        remindersSent++
        SendAppMessage(happiness.Default_Globals_SystemUserID,
            []string{ memberID }, message, happiness.MessageType_MTPulse, "Pulse", url)

        var result sql.Result
        result, error = config.DB.Exec(
            `update PulseBeatMemberTable set beatState = $1
                where pulseID = $2 and beatDate = $3 and memberID = $4;`,
            happiness.PulseBeatState_BSReminded, pulseID, beatDate, memberID)
        if error != nil || pgsql.RowsUpdated(result) != 1 {
            Log.LogError(error)
        }
    }
    Log.Debugf("Sent %d reminders.", remindersSent)
}

