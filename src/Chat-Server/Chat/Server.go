

//----------------------------------------------------------------------------------------
//
//                                                                               Server.go
//                                      Chat-Server: A simple client & server chat service
//
//                                                                   E.B.Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package Chat


import (
    "fmt"
    "sync"
    "time"
    "golang.org/x/net/websocket"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//                                                                              Data Types
//----------------------------------------------------------------------------------------


type ChatServer struct {
    lock            sync.Mutex
    connectionMap   map[*websocket.Conn]string
    userMap         map[string]*ChatServerUser
    roomMap         map[string]*ChatServerRoom
    MessageFormat   MessageFormatType
}


type ChatServerUser struct {
    ChatUser
    connection      *websocket.Conn
    roomID          *string
}


type ChatServerRoom struct {
    ChatRoom
    userIDMap       map[string]bool
}


//----------------------------------------------------------------------------------------
//                                                                           NewChatServer
//----------------------------------------------------------------------------------------


func NewChatServer() *ChatServer {
    server := new(ChatServer)
    server.connectionMap   = make(map[*websocket.Conn]string)
    server.userMap         = make(map[string]*ChatServerUser)
    server.roomMap         = make(map[string]*ChatServerRoom)
    server.MessageFormat   = FormatProtobuf
    return server
}


//----------------------------------------------------------------------------------------
//                                                                             ConnectUser
//----------------------------------------------------------------------------------------


func (server *ChatServer) ConnectUser(connection *websocket.Conn, message *ChatConnect) {
    Log.LogFunctionName()

    result := fmt.Sprintf("Connected to %s. Rooms:\n", connection.Config().Location.String())

    server.lock.Lock()
    for _, room := range server.roomMap {
        result += fmt.Sprintf("%s\n", *room.RoomName)
    }
    server.lock.Unlock()

    result += fmt.Sprintf("%d rooms.", len(msg.Rooms))
    return

    var room *ChatServerRoom
    if message.RoomID != nil {
        room, ok = server.roomMap[*message.RoomID]
    }
    if room == nil {
        server.SendResponse(connection, StatusCode_StatusInputInvalid, "No destination room for message")
        return
    }
    for roomUserID, _ := range room.userIDMap {
        if roomUserID != userID {
            roomUser := server.userMap[roomUserID]
            WriteMessage(roomUser.connection, server.MessageFormat, message)
        }
    }
}


//----------------------------------------------------------------------------------------
//                                                                          DisconnectUser
//----------------------------------------------------------------------------------------


func (server *ChatServer) DisconnectUser(connection *websocket.Conn, message *ChatConnect) {
    Log.LogFunctionName()

    server.lock.Lock()
    defer server.lock.Unlock()

    error := connection.Close()
    if error != nil  { Log.LogError(error) }

    userID, ok := server.connectionMap[connection]
    if ! ok { return }
    delete(server.connectionMap, connection)

    user, ok := server.userMap[userID]
    if ! ok { return }
    delete(server.userMap, userID)

    if user.roomID == nil { return }
    room, ok := server.roomMap[*user.roomID]
    if ! ok { return }

    delete(room.userIDMap, userID)
}


//----------------------------------------------------------------------------------------
//                                                                            SendResponse
//----------------------------------------------------------------------------------------


func (server *ChatServer) SendResponse(
        connection *websocket.Conn,
        code StatusCode,
        message string) {

    // //  Get the user --

    // var (
    //     ok          bool
    //     userID      string
    //     user        *ChatServerUser
    //     room        *ChatServerRoom
    // )

    // server.lock.Lock()
    // userID, ok = server.connectionMap[connection]
    // if ok {
    //     user, ok := server.userMap[userID]
    // }
    // if user != nil && user.roomID != nil {
    //     room, ok := server.roomMap[*user.roomID]
    // }
    // server.lock.Unlock()


    response := ChatResponse {
        Code:       &code,
        Message:    &message,
    }

    WriteMessage(connection, server.MessageFormat, &ChatMessageType { ChatResponse: &response })
}


//----------------------------------------------------------------------------------------
//
//                                                                             Chat Server
//
//----------------------------------------------------------------------------------------


func (server *ChatServer) HandleChatConnection(connection *websocket.Conn) {
    Log.LogFunctionName()

    //  Decode the message --

    var error error
    var wireMessage []byte
    connection.SetReadDeadline(time.Now().Add(60*time.Second))
    _, error = connection.Read(wireMessage)
    if error != nil {
        Log.Errorf("Disconnecting %+v because of error %+v.", connection, error)
        server.DisconnectConnection(connection)
        return
    }

    var chatMessageType ChatMessageType
    error = proto.Unmarshal(wireMessage, &chatMessageType)
    if error != nil {
        Log.LogError(error)
        server.SendResponse(connection, StatusCode_StatusInputInvalid, "Input corrupt")
        return
    }

    //  Process the message --

    switch {

    case chatMessageType.ChatMessage != nil:
        server.SendMessage(connection, chatMessageType.ChatMessage)
        return

    case chatMessageType.ChatConnect != nil:
        if chatMessageType.ChatConnect.GetIsConnecting() {
            server.ConnectUser(connection, chatMessageType.chatConnect)
        } else {
            server.DisconnectUser(connection, chatMessageType.chatConnect)
        }
        return

    case chatMessageType.ChatEnterRoom != nil:
        if chatMessageType.ChatEnterRoom.GetUserIsEntering() {
            server.EnterRoom(connection, chatMessageType.ChatEnterRoom)
        } else {
            server.LeaveRoom(connection, chatMessageType.ChatEnterRoom)
        }
        return

    default:
        Log.Errorf("Received unexpected message type %+v.", chatMessageType)
        server.SendResponse(connection, Chat.StatusInputInvalid, "Unknown message type")
    }

}

