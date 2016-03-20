

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


type ChatClient struct {
    clientLock  sync.Mutex
    userMap     map[string]*ChatUser
    roomMap     map[string]*ChatRoom
    connection  *websocket.Conn

    currentUser *ChatUser
    currentRoom *ChatRoom
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


func (client *ChatClient) EnterRoom(roomName string) error {
    Log.LogFunctionName()

    roomID := client.roomMap[roomName]
}


func (client *ChatClient) SendMessage(roomID, message string) error {
    Log.LogFunctionName()

    timestamp := double(time.Now().UnixNano()) / double(1000000000.0)
    wireMessage := ChatMessage {
        SenderID:       currentUser.UserID,
        RoomID:         currentRoom.RoomID,
        Timestamp:      &timestamp,
        Message:        &message,
    }



}
