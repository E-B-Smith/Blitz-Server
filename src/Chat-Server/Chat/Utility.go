

//----------------------------------------------------------------------------------------
//
//                                                                               Client.go
//                                      Chat-Server: A simple client & server chat service
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package Chat


import (
    "math"
    "time"
    "encoding/json"
    "github.com/satori/go.uuid"
    "golang.org/x/net/websocket"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//
//                                                                 Basic Types & Functions
//
//----------------------------------------------------------------------------------------


type MessageFormatType int
const (
    FormatJSON      MessageFormatType = iota
    FormatProtobuf
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
func NewUUIDString() string         { return uuid.NewV4().String() }


//----------------------------------------------------------------------------------------
//                                                                            WriteMessage
//----------------------------------------------------------------------------------------


func WriteMessage(connection *websocket.Conn, format MessageFormatType, message *ChatMessageType) error {
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

    connection.SetReadDeadline(time.Now().Add(10 * time.Second))
    _, error = connection.Write(data)
    if error != nil { Log.LogError(error) }
    return error
}

