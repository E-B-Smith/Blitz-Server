

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


    Log.Debugf("Got voice call from '%s': '%s'.", from, body);

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
            callDate,
            extract(epoch from suggestedDuration)
                from PhoneNumberTable
                where phoneNumber = $1;`,
        twilio,
    )
    var (
        conversationID          sql.NullString
        expertPhoneNumber       sql.NullString
        clientPhoneNumber       sql.NullString
        callDate                pq.NullTime
        suggestedDuration       sql.NullFloat64
    )
    error := row.Scan(
        &conversationID,
        &expertPhoneNumber,
        &clientPhoneNumber,
        &callDate,
        &suggestedDuration,
    )
    if error != nil {
        Log.LogError(error)
        http.Error(writer, "Forbidden", 403)
        return
    }

    numberToCall := clientPhoneNumber.String
    if from == numberToCall {
        numberToCall = expertPhoneNumber.String
    }

    if len(numberToCall) == 0 {
        http.Error(writer, "Not found", 404)
        return
    }

    Log.Debugf("Number: %s.", numberToCall)
    writer.Header().Set("Content-Type", "text/xml")
    fmt.Fprintf(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
    tml := fmt.Sprintf(
`<Response>
    <Say>You are connect via Blitz, Inc.</Say>
    <Dial>+1%s</Dial>
</Response>`,
        numberToCall,
    )
    fmt.Fprintf(writer, tml)
}


//----------------------------------------------------------------------------------------
//                                                                MaintainPhoneSwitchboard
//----------------------------------------------------------------------------------------


func MaintainPhoneSwitchboard() {
    Log.LogFunctionName()

    //  First disconnect any old phone connections --

    Log.Debugf("Disconnect old connections.")
    _, error := config.DB.Exec(
        `update PhoneNumberTable set
            expertPhoneNumber = null,
            clientPhoneNumber = null,
            conversationID = null,
            callDate = null,
            callDuration = null
                where callDate + callDuration < transaction_timestamp();`,
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
            where callDate - '1 minute'::interval < transaction_timestamp()
              and callPhoneNumber is null;`,
        BlitzMessage.ContactType_CTPhoneSMS,
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
            callDate            pq.NullTime
            suggestedDuration   sql.NullFloat64
        )
        error = rows.Scan(
            &conversationID,
            &clientContact,
            &expertContact,
            &callDate,
            &suggestedDuration,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }

        Log.Debugf("Getting Phone# for conversationID %s.", conversationID)

        var row *sql.Row
        row = config.DB.QueryRow(
            `update PhoneNumberTable set
                conversationID = $1,
                clientPhoneNumber = $2,
                expertPhoneNumber = $3,
                callDate = $4,
                callDuration = $5
                    where phonenumber =
                        (select phonenumber from phonenumbertable
                                where conversationID is null limit 1)
                returning phonenumber;`,
            &conversationID,
            &clientContact,
            &expertContact,
            &callDate,
            &suggestedDuration,
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
            //  Send a "Call is now message here"
            continue
        }
        //  Else we're out of phone numbers.
        Log.Errorf(">>> Out of Twilio Numbers <<<")
    }
}
