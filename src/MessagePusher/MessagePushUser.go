

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
    connection      *websocket.Conn
    pusher          *MessagePusher
    lock            sync.Mutex
    messageLock     sync.Mutex
    messageEvent    *sync.Cond
    messageQueue    []*BlitzMessage.UserMessage
    Format          Format
}


func NewMessagePushUser(connection *websocket.Conn, pusher *MessagePusher) *MessagePushUser {
    Log.LogFunctionName()
    user := new(MessagePushUser)
    user.pusher = pusher
    user.connection = connection
    user.messageQueue = make([]*BlitzMessage.UserMessage, 0, 10)
    user.messageEvent = sync.NewCond(&user.messageLock)
    go user.serviceMessageQueue()
    return user
}


func (user *MessagePushUser) SendMessage(message *BlitzMessage.UserMessage) {
    Log.LogFunctionName()
    user.lock.Lock()
    user.messageQueue = append(user.messageQueue, message)
    user.lock.Unlock()
    user.messageEvent.Signal()
}


func (user *MessagePushUser) serviceMessageQueue() {
    Log.LogFunctionName()

    user.lock.Lock()
    defer user.lock.Unlock()

    for user.connection != nil {

        user.lock.Unlock()
        user.messageLock.Lock()
        user.messageEvent.Wait()
        user.messageLock.Unlock()
        user.lock.Lock()

        for len(user.messageQueue) > 0 && user.connection != nil {

            format := user.Format
            connection := user.connection
            message := user.messageQueue[0]
            user.messageQueue = user.messageQueue[1:]

            user.lock.Unlock()
            error := SendMessageToConnection(connection, format, message)
            if error != nil {
                Log.LogError(error)
                user.pusher.Disconnect(user.connection)
            }
            user.lock.Lock()
        }
    }
}

