//  Messages.go  -  Dispatch user messages.
//
//  E.B.Smith  -  March, 2015


package main


import (
    "fmt"
    "time"
    "errors"
    "strings"
    "strconv"
    "net/http"
    "math/rand"
    "database/sql"
    "encoding/xml"
    "encoding/json"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "happiness"
)


//----------------------------------------------------------------------------------------
//
//                                                                            FetchStories
//
//----------------------------------------------------------------------------------------


func FetchStories(httpWriter http.ResponseWriter, userID string, fetch *happiness.StoryUpdate) {
    //
    //  Fetch stories for the user for the given timespan --
    //

    Log.LogFunctionName()

    var startDate time.Time = pgsql.NegativeInfinityTime
    var stopDate  time.Time = pgsql.PositiveInfinityTime

    if fetch.Timespan != nil {
        if fetch.Timespan.StartTimestamp != nil {
            startDate = happiness.TimeFromTimestamp(fetch.Timespan.StartTimestamp)
        }
        if fetch.Timespan.StopTimestamp != nil {
            stopDate = happiness.TimeFromTimestamp(fetch.Timespan.StopTimestamp)
        }
    }

    warnings := 0

    rows, error := config.DB.Query(
        "select " +
            "storyID, "+
            "storyType, "+
            "creationDate, "+
            "happyScore, "+
            "storyText, "+
            "storyAttribution "+
            "from StoryTable "+
            "where creationDate >  $1 "+
            "  and creationDate <= $2 "+
            "    order by creationDate;",
            startDate, stopDate)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        SendError(httpWriter, happiness.ResponseCode_RCServerError, error)
        return
    }

    var storyArray []*happiness.Story
    for rows.Next() {
        var (
            storyID string
            storyType int
            creationDate pq.NullTime
            happyScore float64
            storyText sql.NullString
            attribution sql.NullString
        )
        error = rows.Scan(
            &storyID,
            &storyType,
            &creationDate,
            &happyScore,
            &storyText,
            &attribution)
        if error != nil {
            Log.LogError(error)
            warnings++
            continue
        }
        storyTypeType := happiness.StoryType(storyType);
        story := happiness.Story {
            StoryID:     &storyID,
            StoryType:   &storyTypeType,
            HappyScore:  &happyScore,
            StoryText:   &storyText.String,
            StoryAttribution: &attribution.String,
        }
        if creationDate.Valid { story.CreationDate = happiness.TimestampFromTime(creationDate.Time) }
        storyArray = append(storyArray, &story)
    }

    Log.Debugf("Found %d stories in range %v to %v.", len(storyArray), startDate, stopDate)

    if len(storyArray) == 0 && warnings > 0 {
        SendError(httpWriter, happiness.ResponseCode_RCServerError, errors.New("Stories are not available now."))
        return
    }

    storyUpdate := happiness.StoryUpdate { Stories: storyArray }
    if  warnings == 0 && len(storyArray) > 0 {
        var timespan happiness.Timespan
        timespan.StartTimestamp = storyArray[0].CreationDate
        timespan.StopTimestamp  = storyArray[len(storyArray)-1].CreationDate
        storyUpdate.Timespan = &timespan
    }

    code := happiness.ResponseCode_RCSuccess
    response := &happiness.ServerResponse {
        ResponseCode:   &code,
        Response:       &happiness.ServerResponse_StoryUpdate { StoryUpdate: &storyUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshalling data: %v.", error)
        SendError(httpWriter, happiness.ResponseCode_RCServerError, error)
        return
    }

    httpWriter.Write(data)
}



//----------------------------------------------------------------------------------------
//
//                                                                           FetchMessages
//
//----------------------------------------------------------------------------------------


func FetchMessages(httpWriter http.ResponseWriter, userID string, fetch *happiness.MessageUpdate) {
    //
    //  Fetch messages for the user for the given timespan --
    //

    Log.LogFunctionName()

    var startDate time.Time = pgsql.NegativeInfinityTime
    var stopDate  time.Time = pgsql.PositiveInfinityTime

    if fetch.Timespan != nil {
        if fetch.Timespan.StartTimestamp != nil {
            startDate = happiness.TimeFromTimestamp(fetch.Timespan.StartTimestamp)
        }
        if fetch.Timespan.StopTimestamp != nil {
            stopDate = happiness.TimeFromTimestamp(fetch.Timespan.StopTimestamp)
        }
    }

    warnings := 0

    rows, error := config.DB.Query(
        "select " +
            "messageID, "+
            "senderID, "+
            "recipientID, "+
            "creationDate, "+
            "notificationDate, "+
            "readDate, "+
            "messageType, "+
            "messageText, "+
            "actionIcon, "+
            "actionURL "+
            "from MessageTable "+
            "  where recipientID = $1 "+
            "  and creationDate >  $2 "+
            "  and creationDate <= $3 "+
            "    order by creationDate;",
            userID, startDate, stopDate)
    defer func() { if rows != nil { rows.Close(); } }()
    if error != nil {
        Log.LogError(error)
        SendError(httpWriter, happiness.ResponseCode_RCServerError, error)
        return
    }

    var messageArray []*happiness.Message
    for rows.Next() {
        var (
            messageID string
            senderID string
            recipientID string
            creationDate pq.NullTime
            notificationDate pq.NullTime
            readDate pq.NullTime
            messageType int
            messageText sql.NullString
            actionIcon sql.NullString
            actionURL sql.NullString
        )
        error = rows.Scan(
            &messageID,
            &senderID,
            &recipientID,
            &creationDate,
            &notificationDate,
            &readDate,
            &messageType,
            &messageText,
            &actionIcon,
            &actionURL,
        )
        if error != nil {
            Log.LogError(error)
            warnings++
            continue
        }
        mt := happiness.MessageType(messageType);
        message := happiness.Message {
            MessageID:   &messageID,
            SenderID:    &senderID,
            Recipients:  []string{recipientID},
            MessageType: &mt,
        }
        if creationDate.Valid { message.CreationDate = happiness.TimestampFromTime(creationDate.Time) }
        if notificationDate.Valid {message.NotificationDate = happiness.TimestampFromTime(notificationDate.Time) }
        if readDate.Valid { message.ReadDate = happiness.TimestampFromTime(readDate.Time) }
        if messageText.Valid { message.MessageText = &messageText.String }
        if actionIcon.Valid { message.ActionIcon = &actionIcon.String }
        if actionURL.Valid { message.ActionURL = &actionURL.String }
        messageArray = append(messageArray, &message)
    }

    Log.Debugf("Found %d message (%d warnings) in range %v to %v.", len(messageArray), warnings, startDate, stopDate)

    if len(messageArray) == 0 && warnings > 0 {
        SendError(httpWriter, happiness.ResponseCode_RCServerError, errors.New("Messages are not available now."))
        return
    }
    messageUpdate := happiness.MessageUpdate {
        Messages:   messageArray,
    }

    if  warnings == 0 && len(messageArray) > 0 {
        var timespan happiness.Timespan
        timespan.StartTimestamp = messageArray[0].CreationDate
        timespan.StopTimestamp  = messageArray[len(messageArray)-1].CreationDate
        messageUpdate.Timespan = &timespan
    }

    code := happiness.ResponseCode_RCSuccess
    response := &happiness.ServerResponse {
        ResponseCode:   &code,
        Response:       &happiness.ServerResponse_MessageResponse{ MessageResponse: &messageUpdate },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshalling data: %v.", error)
        SendError(httpWriter, happiness.ResponseCode_RCServerError, error)
        return
    }

    httpWriter.Write(data)
}



//----------------------------------------------------------------------------------------
//                                                                          SendAppMessage
//----------------------------------------------------------------------------------------


func SendAppMessage(sender string, recipients []string, message string,
                 messageType happiness.MessageType, actionIcon string, actionURL string) {
    Log.LogFunctionName()

    for _, recipient := range recipients {
        if sender == recipient { continue; }

        _, error := config.DB.Exec("insert into MessageTable "+
            "(messageID, " +
            " senderID, "  +
            " recipientID,"+
            " creationDate,"+
            " messageType,"+
            " messageText,"+
            " actionIcon, "+
            " actionURL  "  +
            ") values ($1, $2, $3, $4, $5, $6, $7, $8); ",
            Util.NewUUIDString(),
            sender,
            recipient,
            time.Now(),
            messageType,
            message,
            actionIcon,
            actionURL)

        if error != nil {
            Log.Errorf("Error inserting message: %+v.", error)
        }
    }
}



//----------------------------------------------------------------------------------------
//
//                                                                            SendMessages
//
//----------------------------------------------------------------------------------------


func SendMessages(httpWriter http.ResponseWriter, userID string, sendMessage *happiness.MessageUpdate) {
    //
    //  * Save each new message to the database.
    //  * Kick off a task to send any un-sent new messages.
    //

    Log.LogFunctionName()

    messagesSent := 0
    for _, message := range sendMessage.Messages {
        Log.Debugf("Message %d has %d recipients.", messagesSent+1, len(message.Recipients))
        for _, recipientID := range message.Recipients {
            _, error := config.DB.Exec("insert into MessageTable "+
                "(messageID, " +
                " senderID, "  +
                " recipientID,"+
                " creationDate,"+
                " messageType,"+
                " messageText,"+
                " actionIcon, "+
                " actionURL  "  +
                ") values ($1, $2, $3, $4, $5, $6, $7, $8); ",
                message.MessageID,
                message.SenderID,
                recipientID,
                happiness.NullTimeFromTimestamp(message.CreationDate),
                message.MessageType,
                message.MessageText,
                message.ActionIcon,
                message.ActionURL)

            if error == nil {
                messagesSent++
            } else {
                Log.Errorf("Error inserting message: %v. MessageID: %s From: %s To: %s.",
                    error, *message.MessageID, *message.SenderID, recipientID)
            }
        }
    }

    Log.Debugf("Received %d message bundles, sent %d messages.", len(sendMessage.Messages), messagesSent)

    messageResponse := &happiness.MessageUpdate {
        Timespan: sendMessage.Timespan,
    }

    code := happiness.ResponseCode_RCSuccess
    response := &happiness.ServerResponse {
        ResponseCode:   &code,
        Response:       &happiness.ServerResponse_MessageResponse { MessageResponse: messageResponse },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(httpWriter, happiness.ResponseCode_RCServerError, error)
        return
    }

    httpWriter.Write(data)
}



//----------------------------------------------------------------------------------------
//
//                                                                    SendMessageGetMethod
//
//----------------------------------------------------------------------------------------


func SendMessageGetMethod(writer http.ResponseWriter, httpRequest *http.Request) {

    //
    //  HTTP 'GET' interface to send a message.
    //
    //  send?from=<>&to=<>&body=<>&type=<>&icon=<>&url=<>

    Log.LogFunctionName()

    if httpRequest.URL == nil {
        http.Error(writer, "Not Found", 404)
        return
    }
    var error error
    from := httpRequest.URL.Query().Get("from")
    from, error = happiness.ValidateUserID(&from)
    if error != nil {
        http.Error(writer, "Not Found", 404)
        return
    }
    to := httpRequest.URL.Query().Get("to")
    to, error = happiness.ValidateUserID(&to)
    if error != nil {
        http.Error(writer, "Not Found", 404)
        return
    }
    body := httpRequest.URL.Query().Get("body")
    if len(body) == 0 {
        http.Error(writer, "Not Found", 404)
        return
    }
    mtype := httpRequest.URL.Query().Get("type")
    if len(mtype) == 0 {
        http.Error(writer, "Not Found", 404)
        return
    }
    var (icon, url sql.NullString)
    icon.String = httpRequest.URL.Query().Get("icon")
    if len(icon.String) > 0 {
        icon.Valid = true
    }
    url.String = httpRequest.URL.Query().Get("url")
    if len(url.String) > 0 {
        url.Valid = true
    }

    messageID := Util.NewUUIDString()

    _, error = config.DB.Exec("insert into MessageTable "+
        "(messageID, " +
        " senderID, "  +
        " recipientID,"+
        " creationDate,"+
        " messageType,"+
        " messageText,"+
        " actionIcon, "+
        " actionURL  "  +
        ") values ($1, $2, $3, $4, $5, $6, $7, $8); ",
        messageID,
        from,
        to,
        time.Now(),
        mtype,
        body,
        icon,
        url)

    if error != nil {
        Log.Errorf("Error inserting message: %v. MessageID: %s From: %s To: %s.",
            error, messageID, from, to)
        http.Error(writer, "Not Found", 404)
        return
    }

    Log.Debugf("Sent message.")
    fmt.Fprintf(writer, "OK")
}



//----------------------------------------------------------------------------------------
//
//                                                                     FetchStoryGetMethod
//
//----------------------------------------------------------------------------------------


func FetchStoryGetMethod(writer http.ResponseWriter, httpRequest *http.Request) {

    //  HTTP 'GET' interface to send a message.
    //
    //  http://bh.gy/service/story?happyscore=xxx&quotetype=<quote|action>&ttl=<seconds>&form=<text|json|jsonp|xml>&xid=XXX&jsonp='xx'
    //  curl https://bh.gy/service/story?happyscore=0.5\&quotetype=quote\&ttl=60\&xid=123\&form=xml

    //  Required: None.
    //
    //  Returns:
    //
    //    xid:
    //    quotetype:
    //    happyscore:
    //    ttl:
    //    text:
    //    byline:
    //    jsonp:

    Log.LogFunctionName()
    if  httpRequest.URL == nil {
        http.Error(writer, "Bad Request", 400)
        return
    }
    Log.Debugf("URL: %+v.", httpRequest.URL)

    var error error
    score, _ := strconv.ParseFloat(httpRequest.URL.Query().Get("happyscore"), 64)
    storyType:= httpRequest.URL.Query().Get("quotetype")
    ttl, _   := strconv.ParseInt(httpRequest.URL.Query().Get("ttl"), 10, 64)
    form     := httpRequest.URL.Query().Get("form")
    xid      := httpRequest.URL.Query().Get("xid")
    jsonp    := httpRequest.URL.Query().Get("jsonp")

    if form == "jsonp" && jsonp == "" {
        http.Error(writer, "Bad Request", 400)
        return
    }

    storyTypeLow  := 0
    storyTypeHigh := 255

    switch storyType {
    case "quote":
        storyTypeLow  = int(happiness.StoryType_STQuote)
        storyTypeHigh = int(happiness.StoryType_STQuote)

    case "action":
        storyTypeLow  = int(happiness.StoryType_STAction)
        storyTypeHigh = int(happiness.StoryType_STAction)

    default:
    }

    scoreLow  := 0.0
    scoreHigh := 1.0
    if score > 0.0 {
        if score > 1.0 { score /= 10.0 }
        if score > 1.0 { score = 1.0 }
        scoreLow  = score - 0.125
        scoreHigh = score + 0.125
    }

    interval := int64(rand.Float64() * 10000000.0)
    if ttl > 0 {
        interval = time.Now().Unix() / ttl;
    }

    row := config.DB.QueryRow(
        `with stories as (
          select storyType, storyText, storyAttribution from StoryTable
          where happyScore >= $1
            and happyScore <= $2
            and storyType >=  $3
            and storyType <=  $4
        )
        select storyType, storyText, storyAttribution
          from stories limit 1
          offset ($5 % (select count(*) from stories));`,
          scoreLow, scoreHigh, storyTypeLow, storyTypeHigh, interval)

    var (storyTypeV int; text sql.NullString; byline sql.NullString)
    error = row.Scan(&storyTypeV, &text, &byline)
    if error != nil {
        Log.LogError(error)
        http.Error(writer, "Bad Request", 400)
        return
    }

    result := make(map[string]string)
    if xid != "" { result["xid"] = xid; }
    if score > 0 { result["happyscore"] = fmt.Sprintf("%f", score); }
    switch storyTypeV {
        case int(happiness.StoryType_STQuote):   result["quotetype"] = "quote"
        case int(happiness.StoryType_STAction):  result["quotetype"] = "action"
        default:                                 result["quotetype"] = "unknown"
    }
    if ttl != 0 { result["ttl"] = fmt.Sprintf("%d", ttl); }
    result["text"] = text.String
    if byline.Valid { result["byline"] = byline.String; }

    var resultBytes []byte
    switch form {
    case "json":
        writer.Header().Set("Content-Type", "application/json")
        resultBytes, error = json.Marshal(result)

    case "jsonp":
        writer.Header().Set("Content-Type", "application/javascript")
        var jsonBytes []byte
        jsonBytes, error = json.Marshal(result)
        if error == nil {
            jsonp = strings.TrimLeft(jsonp, "\"'")
            jsonp = strings.TrimRight(jsonp, "\"'")
            resultBytes = []byte(strings.Replace(jsonp, "?", string(jsonBytes), -1))
        }

    case "xml":
        writer.Header().Set("Content-Type", "application/xml")
        type HappyQuote struct {
            XMLName  xml.Name   `xml:"happyquote"`
            XID      string     `xml:"xid,omitempty"`
            QuoteType string    `xml:"quotetype,omitempty"`
            HappyScore float64  `xml:"happyscore,omitempty"`
            TTL      int64      `xml:"ttl,omitempty"`
            Text     string     `xml:"text,omitempty"`
            Byline   string     `xml:"byline,omitempty"`
        }
        xmlQuote := &HappyQuote{
            XID:        xid,
            QuoteType:  storyType,
            HappyScore: score,
            TTL:        ttl,
            Text:       text.String,
            Byline:     byline.String,
        }
        Log.Debugf("XML: %+v.", xmlQuote)
        resultBytes, error = xml.MarshalIndent(xmlQuote, "", "  ")

    default:
        writer.Header().Set("Content-Type", "text/plain")
        if byline.String == "" {
            resultBytes = []byte(text.String)
        } else {
            resultBytes = []byte(fmt.Sprintf("%s\n\n— %s", text.String, byline.String))
        }
    }

    if len(resultBytes) == 0 {
        Log.LogError(error)
        http.Error(writer, "Internal Server Error", 500)
        return
    }

    resultBytes = append(resultBytes, byte('\n'))
    Log.Debugf("%s", resultBytes)
    fmt.Fprintf(writer, "%s", resultBytes)
}


