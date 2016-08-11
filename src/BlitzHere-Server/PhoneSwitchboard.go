

//----------------------------------------------------------------------------------------
//
//                                                  BlitzHere-Server : PhoneSwitchboard.go
//                                                            Maintain the Twilio numbers.
//
//                                                                 E.B. Smith, August 2016
//                        -©- Copyright © 2014-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "time"
    "strings"
    "net/http"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "BlitzMessage"
)


func ConnectTwilioCall(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    //  Respond to a Twilio voice call --

    twilio   := httpRequest.URL.Query().Get("To")
    from     := httpRequest.URL.Query().Get("From")
    body     := httpRequest.URL.Query().Get("Body")


    Log.Debugf("Got voice call from '%s' to: '%s' body: '%s'.",
        from,
        twilio,
        body,
    )

    body = strings.ToLower(strings.TrimSpace(body))
    from = strings.ToLower(strings.TrimSpace(from))
    from = Util.StringIncludingCharactersInSet(from, "0123456789")
    from = strings.TrimLeft(from, "1")

    twilio = Util.StringIncludingCharactersInSet(twilio, "0123456789")
    twilio = strings.TrimLeft(twilio, "1")

    row := config.DB.QueryRow(
        `select
            conversationID,
            expertPhoneNumber,
            clientPhoneNumber,
            startDate,
            stopDate
                from PhoneNumberTable
                where phoneNumber = $1;`,
        twilio,
    )
    var (
        conversationID          sql.NullString
        expertPhoneNumber       sql.NullString
        clientPhoneNumber       sql.NullString
        startDate                pq.NullTime
        stopDate                 pq.NullTime
    )
    error := row.Scan(
        &conversationID,
        &expertPhoneNumber,
        &clientPhoneNumber,
        &startDate,
        &stopDate,
    )
    if error != nil {
        Log.LogError(error)
    }

    if error != nil ||
        !(from == expertPhoneNumber.String ||
         from == clientPhoneNumber.String) ||
        ! stopDate.Valid || time.Since(stopDate.Time) > 0 {
        tml :=
`<Response>
    <Say>Welcome to Blitz Experts.  There is no call scheduled at this time.</Say>
    <Say>Welcome to Blitz Experts.  There is no call scheduled at this time.</Say>
    <Say>Goodbye</Say>
    <Hangup/>
</Response>`
        fmt.Fprintf(writer, tml)
        return
    }

    numberToCall := expertPhoneNumber.String
    if from == numberToCall {
        numberToCall = clientPhoneNumber.String
    }
    Log.Debugf("Number to call: '%s'.", numberToCall)

    if len(numberToCall) == 0 {
        tml :=
`<Response>
    <Say>Welcome to Blitz Experts.  The other party has not configured their phone number.</Say>
    <Say>Welcome to Blitz Experts.  The other party has not configured their phone number.</Say>
    <Say>Goodbye</Say>
    <Hangup/>
</Response>`
        fmt.Fprintf(writer, tml)
        return
    }

    writer.Header().Set("Content-Type", "text/xml")
    fmt.Fprintf(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
    tml := fmt.Sprintf(
`<Response>
    <Say>You are connecting to your expert through Blitz</Say>
    <Dial>+1%s</Dial>
</Response>`,
        numberToCall,
    )
    fmt.Fprintf(writer, tml)
}


//----------------------------------------------------------------------------------------
//                                                                MaintainPhoneSwitchboard
//----------------------------------------------------------------------------------------


func SendCallNotificationToConversationID(conversationID string) {

    actionURL := fmt.Sprintf(
        "%s?action=showchat&chatid=%s",
        config.AppLinkURL,
        conversationID,
    )
    error := SendUserMessageInternal(
        BlitzMessage.Default_Global_SystemUserID,
        MembersForConversationID(conversationID),
        conversationID,
        "You have a Blitz call right now.",
        BlitzMessage.UserMessageType_MTConversation,
        "",
        actionURL,
    )
    if error != nil {
        Log.LogError(error)
    }
}


func CloseCallForConversationID(conversationID string) {
    Log.LogFunctionName()

    _, error := config.DB.Exec(
        `update PhoneNumberTable set
            expertPhoneNumber = null,
            clientPhoneNumber = null,
            conversationID = null,
            startDate = null,
            stopDate = null
                where conversationID = $1;`,
        conversationID,
    )
    if error != nil { Log.LogError(error) }
}


func MaintainPhoneSwitchboard() {
    Log.LogFunctionName()

    //  First disconnect any old phone connections --

    Log.Debugf("Disconnect old connections.")
    _, error := config.DB.Exec(
        `update PhoneNumberTable set
            expertPhoneNumber = null,
            clientPhoneNumber = null,
            conversationID = null,
            startDate = null,
            stopDate = null
                where stopDate < transaction_timestamp();`,
    )
    if error != nil {
        Log.LogError(error)
    }

    //  Now hook up any calls --

    Log.Debugf("Make new connections.")
    var rows *sql.Rows
    rows, error = config.DB.Query(
        `select
            conversationID,
            c.contact,
            e.contact,
            acceptDate,
            callDate,
            extract(epoch from suggestedDuration)
         from conversationTable
            join userContactTable c on
                (conversationTable.initiatorID = c.userID
                and c.contactType = $1
                and c.isVerified = true)
            join userContactTable e on
                (conversationTable.expertID = e.userID
                and e.contactType = $1
                and e.isVerified = true)
            where callPhoneNumber is null
              and conversationType = $2
              and closedDate is null
              and acceptDate is not null;`,
        BlitzMessage.ContactType_CTPhoneSMS,
        BlitzMessage.ConversationType_CTCall,
    )
    if error != nil {
        Log.LogError(error)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var (
            conversationID      string
            clientContact       sql.NullString
            expertContact       sql.NullString
            acceptDate          pq.NullTime
            callDate            pq.NullTime
            suggestedDuration   sql.NullFloat64
        )
        error = rows.Scan(
            &conversationID,
            &clientContact,
            &expertContact,
            &acceptDate,
            &callDate,
            &suggestedDuration,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }

        Log.Debugf("Getting Phone# for conversationID %s.", conversationID)

        startTime := acceptDate.Time
        stopTime  := callDate.Time.Add(time.Duration(suggestedDuration.Float64) + 60*time.Minute)

        var row *sql.Row
        row = config.DB.QueryRow(
            `update PhoneNumberTable set
                conversationID = $1,
                clientPhoneNumber = $2,
                expertPhoneNumber = $3,
                startDate = $4,
                stopDate  = $5
                    where phonenumber =
                        (select phonenumber from phonenumbertable
                                where conversationID is null limit 1)
                returning phonenumber;`,
            conversationID,
            clientContact,
            expertContact,
            startTime,
            stopTime,
        )
        var phoneNumber sql.NullString
        error = row.Scan(&phoneNumber)
        if error != nil && error.Error() != "sql: no rows in result set" {
            Log.Errorf("SQL Error was: %+v.", error)
            continue
        }
        if phoneNumber.Valid {
            var result sql.Result
            result, error = config.DB.Exec(
                `update conversationTable set
                    callPhoneNumber = $1
                    where conversationID = $2;`,
                phoneNumber,
                conversationID,
            )
            error = pgsql.UpdateResultError(result, error)
            if error != nil {
                Log.LogError(error)
            }
            continue
        }
        //  Else we're out of phone numbers.
        Log.Errorf(">>> Out of Twilio Numbers <<<")
    }
}

