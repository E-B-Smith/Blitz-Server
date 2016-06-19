//  SMSReply.go  -  Replies to a Twilio command.
//
//  E.B.Smith  -  July, 2015


package main


import (
    "fmt"
    "strings"
    "net/http"
    "encoding/xml"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
)


func SendTwilioResponse(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()
    config.MessageCount++

    //  Respond to a Twilio SMS message.

    from     := httpRequest.URL.Query().Get("From")
    body     := httpRequest.URL.Query().Get("Body")

    Log.Debugf("Got SMS from '%s': '%s'.", from, body);

    body = strings.ToLower(strings.TrimSpace(body))
    from = strings.ToLower(strings.TrimSpace(from))
    from = Util.StringIncludingCharactersInSet(from, "0123456789")

    var response string

    switch body {

    case "dance":
        response =
`\O/
_O_
\O/
_O_
\O/`

    case "status":
        responsebytes, errorlines, error := Util.RunShellCommand("/bin/bash", []string{"-c", kBashStatusScript}, nil)
        if error != nil || len(errorlines) > 0 {
            Log.Errorf("Run shell returned with error %v: %s.", error, string(errorlines))
        }
        response = string(responsebytes)

    default:
        response = "Hello."
    }

    if len(response) == 0 {
        response = "Oops."
    }

    Log.Debugf("Response: %s.", response)
    writer.Header().Set("Content-Type", "text/xml")
    fmt.Fprintf(writer, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
    fmt.Fprintf(writer, "<Response>")

    lines := strings.Split(response, "\n")
    Log.Debugf("%d lines.", len(lines))
    for _, line := range(lines) {
        if len(line) > 0 {
            fmt.Fprintf(writer, "<Message>\n")
            xml.EscapeText(writer, []byte(line))
            fmt.Fprintf(writer, "</Message>\n")
        }
    }
    fmt.Fprintf(writer, "</Response>")
}


//----------------------------------------------------------------------------------------
//                                                                       kBashStatusScript
//----------------------------------------------------------------------------------------

const kBashStatusScript =
`
printf " BlitzHere-Server: "
(echo status | nc localhost 10005 | tr -d '\n' || true)

printf "\n BlitzLabs-Server: "
(echo status | nc localhost 10003 | tr -d '\n' || true)

printf "\n    Status-Server: "
(echo status | nc localhost 10007 | tr -d '\n' || true)
echo ""
`
