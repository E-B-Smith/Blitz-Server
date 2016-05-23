

//----------------------------------------------------------------------------------------
//
//                                                      MessagePusher : MessagePushUser.go
//                                                             Push messages via websocket
//
//                                                                   E.B.Smith, April 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package MessagePusher


import (
    "sync"
    "time"
    "golang.org/x/net/websocket"
    "violent.blue/GoKit/Log"
    "BlitzMessage"
)


type Format int
const (
    FormatUnknown Format = iota
    FormatJSON
    FormatProtobuf
)


type MessagePushUser struct {
    lock            sync.Mutex
    connection      *websocket.Conn
    writeChannel    chan *BlitzMessage.UserMessage
    userID          string
    LastMessageTime *time.Time
    Format          Format
}


func NewMessagePushUser(connection *websocket.Conn) *MessagePushUser {
    Log.LogFunctionName()
    user := new(MessagePushUser)
    user.connection = connection
    user.writeChannel = make(chan *BlitzMessage.UserMessage)
    return user
}


func (user *MessagePushUser) UserID() string {
    return user.userID
}


func (user *MessagePushUser) Disconnect() {
    Log.LogFunctionName()
    user.lock.Lock()
    if user.connection != nil {
        user.connection.Close()
        user.connection = nil
    }
    user.lock.Unlock()
}


func (user *MessagePushUser) SendMessage(message *BlitzMessage.UserMessage) {
    Log.LogFunctionName()
    user.writeChannel <- message
}


func (user *MessagePushUser) IsConnected() bool {
    user.lock.Lock()
    defer user.lock.Unlock()
    return (user.connection != nil)
}

