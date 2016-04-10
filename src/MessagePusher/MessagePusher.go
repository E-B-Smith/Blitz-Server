

//----------------------------------------------------------------------------------------
//
//                                                        MessagePusher : MessagePusher.go
//                                                             Push messages via websocket
//
//                                                                 E.B. Smith, April, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package MessagePusher


import (
    "sync"
    "time"
    "strings"
    "golang.org/x/net/websocket"
    "violent.blue/GoKit/Log"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                                           MessagePusher
//----------------------------------------------------------------------------------------


type MessagePusher struct {
    lock            sync.RWMutex
    connectionMap   map[*websocket.Conn]string  //  Connection -> UserID
    userMap         map[string]*MessagePushUser //  UserID -> MessagePushUser
}


func NewMessagePusher() *MessagePusher {
    Log.LogFunctionName()
    pusher := new(MessagePusher)
    pusher.connectionMap   = make(map[*websocket.Conn]string)
    pusher.userMap         = make(map[string]*MessagePushUser)
    return pusher
}


func (pusher *MessagePusher) PushMessage(message *BlitzMessage.UserMessage) {
    Log.LogFunctionName()

    pusher.lock.RLock()
    defer pusher.lock.RUnlock()
    for _, recipientID := range message.Recipients {
        if user, ok := pusher.userMap[recipientID]; ok {
            user.SendMessage(message)
        }
    }
}


func (pusher *MessagePusher) Disconnect(connection *websocket.Conn) {
    Log.LogFunctionName()

    pusher.lock.Lock()
    defer pusher.lock.Unlock()

    userID, ok := pusher.connectionMap[connection]
    if ! ok { return }

    delete(pusher.connectionMap, connection)

    user, ok := pusher.userMap[userID]
    if ok {
        user.lock.Lock()
        user.connection = nil
        user.pusher = nil
        user.messageEvent = nil
        user.messageQueue = make([]*BlitzMessage.UserMessage, 0, 0)
        user.lock.Unlock()
        delete(pusher.userMap, userID)
    }
}


func (pusher *MessagePusher) Connect(connection *websocket.Conn, message *BlitzMessage.ServerRequest) {
    Log.LogFunctionName()

    userID := "10101"   //  eDebug UserIDFromConnectionMessage(message)
    userID = strings.TrimSpace(userID)
    if len(userID) == 0 { return }

    pusher.lock.Lock()
    defer pusher.lock.Unlock()

    user := NewMessagePushUser(connection, pusher)
    pusher.connectionMap[connection] = userID
    pusher.userMap[userID] = user
}


//----------------------------------------------------------------------------------------
//
//                                                                    HandlePushConnection
//
//----------------------------------------------------------------------------------------


func (pusher *MessagePusher) HandlePushConnection(connection *websocket.Conn) {
    Log.LogFunctionName()

    //  Read the message --

    var error error
    var wireMessage []byte
    connection.SetReadDeadline(time.Now().Add(kReadTimeoutSeconds))
    _, error = connection.Read(wireMessage)
    if error != nil {
        Log.Errorf("Disconnecting %+v because of error %+v.", *connection, error)
        pusher.Disconnect(connection)
        return
    }

    //  Decode the message --

    format := FormatUnknown
    var request *BlitzMessage.ServerRequest
    request, format, error = DecodeMessage(format, wireMessage)
    if error != nil {
        Log.LogError(error)
        return
    }

    //  Route the message --

    if request.RequestType.PushConnect != nil {
        pusher.Connect(connection, request)
        return
    }

    if request.RequestType.PushDisconnect != nil {
        pusher.Disconnect(connection)
        return
    }

    Log.Errorf("Received unexpected message type %+v.", *request.RequestType)
}

