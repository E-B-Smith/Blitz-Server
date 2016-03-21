

//----------------------------------------------------------------------------------------
//
//                                                                                    Chat
//                                                   Simple chat client & server functions
//
//                                                                   E.B.Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package Chat


import (
    "errors"
    "golang.org/x/net/websocket"
    "violent.blue/GoKit/Log"
)


type MessageFormatType int
const (
    FormatJSON      MessageFormatType = iota
    FormatProtobuf
)


type ChatClient struct {
    clientLock      sync.Mutex
    userMap         map[string]*ChatUser
    roomMap         map[string]*ChatRoom
    connection      *websocket.Conn

    currentUser     *ChatUser
    currentRoom     *ChatRoom
    MessageFormat   MessageFormatType
}


func (client *ChatClient) Connect(url net.URL) error {
    Log.LogFunctionName()

    var error error
    error = client.clientLock.Lock()
    if error != nil { return error }
    defer client.clientLock.Unlock()

    if client.connection != nil {
        return errors.New("Already connected")
    }

    MessageFormat = FormatProtobuf
    origin := "http://"+url.Host
    client.connection, error =  client.Dial(url, "", origin)
    if error != nil {
        return error
    }

    client.userMap = make(map[string]*ChatUser)
    client.roomMap = make(map[string]*ChatRoom)

    return nil
}


func (client *ChatClient) Disconnect() error {
    Log.LogFunctionName()

    if client.connection == nil {
        return nil
    }

    var error error
    error = client.connection.Close()
    client.connection = nil
    client.userMap = make(map[string]*ChatUser)
    client.roomMap = make(map[string]*ChatRoom)
}

/*
func (client *ChatClient) EnterRoom(roomName string) error {
    Log.LogFunctionName()

    roomID := client.roomMap[roomName]
}
*/


func TimestampFromTime(time time.Time) *double {
    timestamp := double(time.UnixNano()) / double(1000000000.0)
    return &timestamp
}


func TimeFromTimestamp(timestamp double) time.Time {
    i, f := math.Modf(timestamp)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    time := time.Unix(sec, nsec)
    return time
}


func (client *ChatClient) SendMessage(roomID, message string) error {
    Log.LogFunctionName()

    wireMessage := ChatMessage {
        SenderID:       currentUser.UserID,
        RoomID:         currentRoom.RoomID,
        Timestamp:      TimestampFromTime(time.Now()),
        Message:        &message,
    }

    return WriteMessage(client.connection, client.MessageFormat, wireMessage)
}


func WriteMessage(connection *websocket.Conn, format MessageFormatType, message interface{}) error {
    Log.LogFunctionName()

    var error error
    var data []byte
    switch format {
    case FormatProtobuf:
        data, error = proto.Marshal(message)
    default:
        data, error = json.Marshal(message)
        data = append(data, []byte("\n")...)
        //Log.Debugf("%s", string(data))
    }
    if error != nil {
        Log.LogError(error)
        return error
    }
    error = connection.Write(data)
    if error != nil { Log.LogError(error) }
    return error
}

