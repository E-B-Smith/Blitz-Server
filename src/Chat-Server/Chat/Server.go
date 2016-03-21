

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
//                                                                          ChatServerUser
//----------------------------------------------------------------------------------------


type ChatServerUser struct {
    ChatUser
    Connection  *websocket.Conn
}


type ChatServerRoom struct {
    ChatRoom
    userIDArray       []string
}


//----------------------------------------------------------------------------------------
//
//                                                                             Chat Server
//
//----------------------------------------------------------------------------------------


type ChatServer struct {
    lock            sync.Mutex
    connectionMap   map[*websocket.Conn]string
    userMap         map[string]*ChatServerUser
    roomMap         map[string]*ChatServerRoom
}


func (server *ChatServer) HandleUserConnection(connection *websocket.Conn) {
    Log.LogFunctionName()

    var error error
    var wireMessage []byte
    connection.SetReadDeadline(time.Now().Add(60*time.Second))
    _, error = connection.Read(wireMessage)
    if error != nil {
        Log.LogError(error)
        return
    }

    var chatMessage ChatMessageType
    error = proto.Unmarshal(wireMessage, &chatMessage)
    if error != nil {
        Log.LogError(error)
        return
    }

    userID, ok := server.connectionMap[connection]
    if ! ok { userID = "" }

    user, ok := server.userMap[userID]
    if ! ok {
        user = new(ChatServerUser)
        user.UserID = &userID
        user.Nickname = StringPtr("Buddy")
        user.Connection = connection
        server.userMap[userID] = user
        server.connectionMap[connection] = userID
    }

}


var globalChatServer *ChatServer = nil


func ChatServerHandler(userConnection *websocket.Conn) {
    globalChatServer.HandleUserConnection(userConnection)
}


