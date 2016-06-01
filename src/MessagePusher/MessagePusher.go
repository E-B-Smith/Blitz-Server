

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
    "fmt"
    "net"
    "sync"
    "time"
    "strings"
    "golang.org/x/net/websocket"
    "violent.blue/GoKit/Log"
    "BlitzMessage"
)


type UserMayConnectFunction func(pusher *MessagePusher, message *BlitzMessage.ServerRequest) error
type UserDidConnectFunction func(pusher *MessagePusher, user *MessagePushUser)


//----------------------------------------------------------------------------------------
//                                                                           MessagePusher
//----------------------------------------------------------------------------------------


type MessagePusher struct {
    lock            sync.RWMutex
    connectionMap   map[*websocket.Conn]string  //  Connection -> UserID
    userMap         map[string]*MessagePushUser //  UserID -> MessagePushUser
    UserMayConnect  UserMayConnectFunction      //  Optional call back function
    UserDidConnect  UserDidConnectFunction      //  Optional call back function
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


func (pusher *MessagePusher) Connect(connection *websocket.Conn) (*MessagePushUser, error) {
    Log.LogFunctionName()

    //  Read the connect message --

    var n int
    var error error
    var wireMessage []byte = make([]byte, 3200)

    connection.SetReadDeadline(time.Now().Add(kReadTimeoutSeconds))
    for n == 0 {
        n, error = connection.Read(wireMessage)
        Log.Debugf("Read %d/%d bytes.", n, len(wireMessage))
        if error != nil {
            Log.Errorf("Disconnecting %+v because of error %+v.", *connection, error)
            return nil, error
        }
    }

    //  Decode the message --

    format := FormatUnknown
    var request *BlitzMessage.ServerRequest
    request, format, error = DecodeMessage(format, wireMessage[:n])
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    Log.Debugf("Connect message: %+v.", request)
    if  request == nil ||
        request.SessionToken == nil ||
        request.RequestType  == nil ||
        request.RequestType.PushConnect == nil ||
        request.RequestType.PushConnect.UserID == nil {
        return nil, fmt.Errorf("Bad connect message")
    }

    pusher.lock.Lock()
    defer pusher.lock.Unlock()

    //  Connect user?

    if pusher.UserMayConnect != nil {
        error = pusher.UserMayConnect(pusher, request)
        if error != nil {
            Log.LogError(error)
            return nil, error
        }
    }

    userID := strings.TrimSpace(*request.RequestType.PushConnect.UserID)

    user := NewMessagePushUser(connection)
    user.Format = format
    user.userID = userID
    user.LastMessageTime = request.RequestType.PushConnect.LastMessageTimestamp.TimePtr()

    pusher.connectionMap[connection] = userID
    pusher.userMap[userID] = user

    Log.Debugf("Connected user %s", user.userID)
    return user, nil
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
        user.Disconnect()
        delete(pusher.userMap, userID)
    }
}


//----------------------------------------------------------------------------------------
//                                                                          readConnection
//----------------------------------------------------------------------------------------


func IsTemporaryTimeoutError(err error) bool {
    if err, ok := err.(net.Error); ok && err.Timeout() && err.Temporary() {
        return true
    }
    return false
}


func (pusher *MessagePusher) readConnection(connection *websocket.Conn, readChannel chan *BlitzMessage.ServerRequest) {
    Log.LogFunctionName()
    defer Log.Debugf("Exit readConnection")

    var n int
    var error error
    var wireMessage []byte = make([]byte, 3200)
    var message *BlitzMessage.ServerRequest
    var format Format = FormatUnknown

    for {

        connection.SetReadDeadline(time.Now().Add(kReadTimeoutSeconds))
        n, error = connection.Read(wireMessage)
        if  IsTemporaryTimeoutError(error) {
            //  Send a ping keep alive message --
            Log.Debugf("Sending ping.")
            connection.PayloadType = websocket.PingFrame
            connection.Write(wireMessage[:0])
            continue
        }

        if error != nil {
            Log.Errorf("Disconnecting %+v because of error %+v.", *connection, error)
            pusher.Disconnect(connection)
            return
        }

        Log.Debugf("Read %d/%d bytes.", n, len(wireMessage))
        message, format, error = DecodeMessage(format, wireMessage[:n])
        if error != nil {
            Log.LogError(error)
            pusher.Disconnect(connection)
            return
        }

        if message != nil {
            readChannel <- message
        }
    }
}


//----------------------------------------------------------------------------------------
//
//                                                                    HandlePushConnection
//
//----------------------------------------------------------------------------------------


func (pusher *MessagePusher) HandlePushConnection(connection *websocket.Conn) {
    Log.LogFunctionName()
    defer Log.Debugf("Exit HandlePushConnection.")

    var error error
    var user *MessagePushUser

    user, error = pusher.Connect(connection)
    if error != nil {
        Log.Errorf("Disconnecting %+v because of error %+v.", *connection, error)
        return
    }

    //  Loop, proccessing the messages --

    messageQueue := make([]*BlitzMessage.UserMessage, 0, 10)
    readChannel  := make(chan *BlitzMessage.ServerRequest)
    go pusher.readConnection(connection, readChannel)

    //  Call back --

    if pusher.UserDidConnect != nil {
        go pusher.UserDidConnect(pusher, user)
    }

    for user.IsConnected() {

        Log.Debugf("Waiting for messages...")
        select {
            case request := <- readChannel:
                if request != nil &&
                   request.RequestType != nil &&
                   request.RequestType.PushDisconnect != nil {
                   pusher.Disconnect(connection)
                }

            case message := <- user.writeChannel:
                messageQueue = append(messageQueue, message)
        }
        Log.Debugf("Sending messages...")

        for len(messageQueue) > 0 && user.IsConnected() {

            message := messageQueue[0]
            messageQueue = messageQueue[1:]

            error = SendMessageToConnection(connection, user.Format, message)
            if IsTemporaryTimeoutError(error) {
                error = SendMessageToConnection(connection, user.Format, message)
            }
            if error != nil {
                Log.LogError(error)
                pusher.Disconnect(connection)
            } else {
                Log.Debugf("Sent one message.")
            }

        }
    }
}

