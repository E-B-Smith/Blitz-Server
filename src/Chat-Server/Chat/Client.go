

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
    "fmt"
    "net"
    "sync"
    "time"
    "errors"
    "net/url"
    "golang.org/x/net/websocket"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//
//                                                                             Chat Client
//
//----------------------------------------------------------------------------------------


type ChatClient struct {
    clientLock      sync.Mutex
    userMap         map[string]*ChatUser
    roomMap         map[string]*ChatRoom
    connection      *websocket.Conn

    currentUser     *ChatUser
    currentRoom     *ChatRoom
    MessageFormat   MessageFormatType
}


//----------------------------------------------------------------------------------------
//                                                                           NewChatClient
//----------------------------------------------------------------------------------------


func NewChatClient() *ChatClient {
    client := new(ChatClient)
    client.userMap         = make(map[string]*ChatUser)
    client.roomMap         = make(map[string]*ChatRoom)
    client.connection      = nil
    client.currentUser     = nil
    client.currentRoom     = nil
    client.MessageFormat   = FormatProtobuf
    return client
}


//----------------------------------------------------------------------------------------
//                                                                       StringFromMessage
//----------------------------------------------------------------------------------------


func (client *ChatClient) StringFromMessage(chatMessageType *ChatMessageType) string {

    switch {
    case chatMessageType.ChatMessage != nil:
        m := chatMessageType.ChatMessage
        room := client.roomMap[*m.RoomID]
        return fmt.Sprintf("%15s: %s", *room.RoomName, *m.Message)

    case chatMessageType.ChatConnect != nil:
        m := chatMessageType.ChatConnect
        result := fmt.Sprintf("Connected to %s. Rooms:\n", client.connection.Config().Location.String())
        for _, room := range m.Rooms {
            result += fmt.Sprintf("%s\n", *room.RoomName)
        }
        result += fmt.Sprintf("%d rooms.", len(m.Rooms))
        return result

    case chatMessageType.ChatEnterRoom != nil:

    case chatMessageType.ChatPresence != nil:
        m := chatMessageType.ChatPresence
        result := fmt.Sprintf("Room %s occupants:\n", *m.Room.RoomName)
        for _, user := range m.Users {
            result += fmt.Sprintf("%s\n", *user.Nickname)
        }
        result += fmt.Sprintf("%d occupants.", len(m.Users))
        return result

    case chatMessageType.ChatResponse != nil:
        m := chatMessageType.ChatResponse
        return fmt.Sprintf("Code %d: %s", m.Code, m.Message)
    }

    return ""
}


//----------------------------------------------------------------------------------------
//                                                                                 Connect
//----------------------------------------------------------------------------------------


func (client *ChatClient) Connect(URL string, readChannel chan <- *ChatMessageType) error {
    Log.LogFunctionName()

    client.clientLock.Lock()
    defer client.clientLock.Unlock()

    if client.connection != nil {
        return errors.New("Already connected")
    }

    parsedURL, error := url.Parse(URL)
    if error != nil {
        Log.LogError(error)
        return error
    }
    origin := "http://"+parsedURL.Host
    client.connection, error =  websocket.Dial(URL, "", origin)
    if error != nil {
        return error
    }

    client.MessageFormat = FormatProtobuf
    client.userMap = make(map[string]*ChatUser)
    client.roomMap = make(map[string]*ChatRoom)

    go client.ChatClientReader(readChannel)

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
    return error
}


func (client *ChatClient) LeaveRoom() error {
    Log.LogFunctionName()

    if client.currentRoom == nil {
        return nil
    }
    chatMessage := ChatEnterRoom {
        UserIsEntering:     BoolPtr(false),
        User:               client.currentUser,
        RoomID:             client.currentRoom.RoomID,
    }
    return SendMessageToConnection(client.connection,
            client.MessageFormat,
            &ChatMessageType { ChatEnterRoom: &chatMessage })
}


func (client *ChatClient) SendMessage(message string) error {
    Log.LogFunctionName()

    chatMessage := ChatMessage {
        SenderID:       client.currentUser.UserID,
        RoomID:         client.currentRoom.RoomID,
        Timestamp:      TimestampFromTime(time.Now()),
        Message:        &message,
    }

    return SendMessageToConnection(client.connection,
            client.MessageFormat,
            &ChatMessageType { ChatMessage: &chatMessage })
}


func (client *ChatClient) ChatClientReader(clientReaderChannel chan <- *ChatMessageType) {
    Log.LogFunctionName()

    for client.connection != nil {

        var wireMessage []byte
        client.connection.SetReadDeadline(time.Now().Add(60*time.Second))
        _, error := client.connection.Read(wireMessage)
        if error, ok := error.(net.Error); ok && error.Timeout() { continue }
        if error != nil {
            Log.LogError(error)
            return
        }

        var chatMessage ChatMessageType
        error = proto.Unmarshal(wireMessage, &chatMessage)
        if error != nil {
            Log.LogError(error)
            continue
        }

        switch {

        case chatMessage.ChatMessage != nil:
            clientReaderChannel <- &chatMessage

        case chatMessage.ChatResponse != nil:
            clientReaderChannel <- &chatMessage

        }
    }
}

