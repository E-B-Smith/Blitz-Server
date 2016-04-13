

//----------------------------------------------------------------------------------------
//
//                                                                              Utility.go
//                                      Chat-Server: A simple client & server chat service
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package MessagePusher


import (
    "math"
    "time"
    "strings"
    "encoding/json"
    "github.com/satori/go.uuid"
    "golang.org/x/net/websocket"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                 Basic Types & Functions
//
//----------------------------------------------------------------------------------------


const (
    kReadTimeoutSeconds     time.Duration = (30 * time.Second)
    kWriteTimeoutSeconds    time.Duration = (30 * time.Second)
)


func TimestampFromTime(time time.Time) *float64 {
    timestamp := float64(time.UnixNano()) / float64(1000000000.0)
    return &timestamp
}


func TimeFromTimestamp(timestamp float64) time.Time {
    i, f := math.Modf(timestamp)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    time := time.Unix(sec, nsec)
    return time
}


func BoolPtr(b bool) *bool          { return &b }
func StringPtr(s string) *string    { return &s }
func FormatPtr(f Format) *Format    { return &f }
func NewUUIDString() string         { return uuid.NewV4().String() }


//----------------------------------------------------------------------------------------
//                                                                 SendMessageToConnection
//----------------------------------------------------------------------------------------


func SendMessageToConnection(
        connection *websocket.Conn,
        format Format,
        message *BlitzMessage.UserMessage) error {
    Log.LogFunctionName()

    var error error
    var data []byte
    switch format {
    case FormatProtobuf:
        data, error = proto.Marshal(message)
    default:
        data, error = json.Marshal(*message)
        data = append(data, []byte("\n")...)
        //Log.Debugf("%s", string(data))
    }
    if error != nil {
        Log.LogError(error)
        return error
    }

    connection.PayloadType = websocket.BinaryFrame
    connection.SetWriteDeadline(time.Now().Add(kWriteTimeoutSeconds))
    _, error = connection.Write(data)
    if error != nil { Log.LogError(error) }
    return error
}


//----------------------------------------------------------------------------------------
//                                                                           DecodeMessage
//----------------------------------------------------------------------------------------


func DecodeMessage(formatIn Format, wireMessage []byte) (
        message *BlitzMessage.ServerRequest, format Format, error error) {
    Log.LogFunctionName()

    message = &BlitzMessage.ServerRequest {}
    format = formatIn
    error = nil

    if format == FormatUnknown {
        test := strings.TrimSpace(string(wireMessage))
        if strings.HasPrefix(test, "{") {
            format = FormatJSON
        } else {
            format = FormatProtobuf
        }
    }

    switch format {
    case FormatProtobuf:
        error = proto.Unmarshal(wireMessage, message)

    default:
        error = json.Unmarshal(wireMessage, message)
        format = FormatJSON
    }

    return
}

