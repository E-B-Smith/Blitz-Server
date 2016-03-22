

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
}


type ChatServerUser struct {
    ChatUser
    connection      *websocket.Conn
    roomID          *string
    MessageFormat   MessageFormatType
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
    return server
}

//----------------------------------------------------------------------------------------
//                                                                      UserFromConnection
//----------------------------------------------------------------------------------------


func (server *ChatServer) UserFromConnection(connection *websocket.Conn) (
        userID string,
        user *ChatServerUser,
        room *ChatServerRoom) {

    userID = ""
    user = nil
    room = nil

    if connection == nil { return }

    var ok bool
    userID, ok = server.connectionMap[connection]
    if ! ok { return }

    user, ok = server.userMap[userID]
    if ! ok { return }

    if user.roomID != nil {
        room, ok = server.roomMap[*user.roomID]
    }

    return
}

//----------------------------------------------------------------------------------------
//                                                                             ConnectUser
//----------------------------------------------------------------------------------------


func (server *ChatServer) ConnectUser(connection *websocket.Conn, newUser *ChatUser) {
    Log.LogFunctionName()
/*
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
*/
}


//----------------------------------------------------------------------------------------
//                                                                              Disconnect
//----------------------------------------------------------------------------------------


func (server *ChatServer) Disconnect(connection *websocket.Conn) {
    Log.LogFunctionName()

    server.lock.Lock()
    defer server.lock.Unlock()

    userID, user, room := server.UserFromConnection(connection)

    if room != nil {
        server.LeaveRoom(connection)
    }

    if user != nil {

        chatDisconnect := ChatConnect {
            IsConnecting:   BoolPtr(false),
            User:           &user.ChatUser,
        }
        chatMessage := ChatMessageType {
            ChatConnect:    &chatDisconnect,
        }
        SendMessageToConnection(connection, user.MessageFormat, &chatMessage)

    }

    if len(userID) > 0 {
        delete(server.userMap, userID)
    }

    delete(server.connectionMap, connection)
}


//----------------------------------------------------------------------------------------
//                                                                       SendMessageToRoom
//----------------------------------------------------------------------------------------


func (server *ChatServer) SendMessageToRoom(room *ChatServerRoom, message *ChatMessageType) {
    Log.LogFunctionName()
    for userID, _ := range room.userIDMap {
        user, ok := server.userMap[userID]
        if ok {
            SendMessageToConnection(user.connection, user.MessageFormat, message)
        }
    }
}


//----------------------------------------------------------------------------------------
//                                                                               EnterRoom
//----------------------------------------------------------------------------------------


func (server *ChatServer) EnterRoom(connection *websocket.Conn, roomID string) {
    Log.LogFunctionName()

    server.lock.Lock()
    defer server.lock.Unlock()

    userID, user, room := server.UserFromConnection(connection)
    if user == nil { return }

    if room != nil {
        server.LeaveRoom(connection)
    }

    room = server.roomMap[roomID]
    room.userIDMap[userID] = true
    user.roomID = &roomID

    enterMessage := ChatEnterRoom {
        User:           &user.ChatUser,
        RoomID:         room.RoomID,
        UserIsEntering: BoolPtr(true),
    }
    chatMessage := ChatMessageType { ChatEnterRoom: &enterMessage }
    server.SendMessageToRoom(room, &chatMessage)

    //  Send room presence to user --

    userArray := make([]*ChatUser, 0, len(room.userIDMap))
    for userID, _ := range room.userIDMap {
        roomUser := server.userMap[userID]
        userArray = append(userArray, &roomUser.ChatUser)
    }

    presenceMessage := ChatPresence {
        Room:       &room.ChatRoom,
        Users:      userArray,
    }

    chatMessage = ChatMessageType {
        ChatPresence: &presenceMessage,
    }

    SendMessageToConnection(connection, user.MessageFormat, &chatMessage)
}


//----------------------------------------------------------------------------------------
//                                                                               LeaveRoom
//----------------------------------------------------------------------------------------


func (server *ChatServer) LeaveRoom(connection *websocket.Conn) {
    Log.LogFunctionName()

    server.lock.Lock()
    defer server.lock.Unlock()

    userID, user, room := server.UserFromConnection(connection)
    if user == nil || room == nil { return }

    leaveMessage := ChatEnterRoom {
        User:           &user.ChatUser,
        RoomID:         room.RoomID,
        UserIsEntering: BoolPtr(false),
    }
    chatMessage := ChatMessageType { ChatEnterRoom: &leaveMessage }

    server.SendMessageToRoom(room, &chatMessage)

    //  Delete user from room --

    delete(room.userIDMap, userID)
    if len(room.userIDMap) == 0 {
        delete(server.roomMap, *room.RoomID)
    }

    user.roomID = nil
}


//----------------------------------------------------------------------------------------
//                                                                            SendResponse
//----------------------------------------------------------------------------------------


func (server *ChatServer) SendResponse(
        connection *websocket.Conn,
        code StatusCode,
        message string) {
    Log.LogFunctionName()

    response := ChatResponse {
        Code:       &code,
        Message:    &message,
    }

    server.lock.Lock()
    defer server.lock.Unlock()
    _, user, _ := server.UserFromConnection(connection)

    format := FormatJSON
    if user != nil {
        format = user.MessageFormat
    }

    SendMessageToConnection(connection,
        format,
        &ChatMessageType { ChatResponse: &response })
}


//----------------------------------------------------------------------------------------
//                                                                         SendChatMessage
//----------------------------------------------------------------------------------------


func (server *ChatServer) SendChatMessage(message *ChatMessage) {
    Log.LogFunctionName()

    if message.SenderID == nil ||
       message.RoomID == nil ||
       message.Message == nil {
        return
    }
    message.Timestamp = TimestampFromTime(time.Now())

    server.lock.Lock()
    defer server.lock.Unlock()

    room, ok := server.roomMap[*message.RoomID]
    if ! ok { return }

    server.SendMessageToRoom(room, &ChatMessageType{ChatMessage: message})
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
        server.Disconnect(connection)
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
        server.SendChatMessage(chatMessageType.ChatMessage)
        return

    case chatMessageType.ChatConnect != nil:
        if chatMessageType.ChatConnect.GetIsConnecting() {
            server.ConnectUser(connection, chatMessageType.ChatConnect.User)
        } else {
            server.Disconnect(connection)
        }
        return

    case chatMessageType.ChatEnterRoom != nil:
        if chatMessageType.ChatEnterRoom.GetUserIsEntering() {
            if chatMessageType.ChatEnterRoom.RoomID == nil {
                server.SendResponse(connection, StatusCode_StatusInputInvalid, "No room ID")
            } else {
                server.EnterRoom(connection, *chatMessageType.ChatEnterRoom.RoomID)
            }
        } else {
            server.LeaveRoom(connection)
        }
        return

    default:
        Log.Errorf("Received unexpected message type %+v.", chatMessageType)
        server.SendResponse(connection, StatusCode_StatusInputInvalid, "Unknown message type")
        return
    }
}

