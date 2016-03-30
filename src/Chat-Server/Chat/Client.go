

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
    lock        sync.RWMutex
    userMap     map[string]*ChatUser
    roomMap     map[string]*ChatRoom
    connection  *websocket.Conn

    user        *ChatUser
    room        *ChatRoom
}


//----------------------------------------------------------------------------------------
//                                                                           NewChatClient
//----------------------------------------------------------------------------------------


func NewChatClient() *ChatClient {
    client := new(ChatClient)
    client.userMap      = make(map[string]*ChatUser)
    client.roomMap      = make(map[string]*ChatRoom)
    client.connection   = nil
    client.user         = nil
    client.room         = nil
    return client
}


//----------------------------------------------------------------------------------------
//                                                                       StringFromMessage
//----------------------------------------------------------------------------------------


func (client *ChatClient) StringFromMessage(chatMessageType *ChatMessageType) string {

    switch {

    case chatMessageType.ChatMessage != nil:
        m := chatMessageType.ChatMessage
        if m.SenderID == nil || m.RoomID == nil {
            Log.Errorf("Bad message: %+v.", m)
            return ""
        }

        client.lock.RLock()
        defer client.lock.RUnlock()

        username := *m.SenderID
        user, ok := client.userMap[username]
        if ok && user.Nickname != nil  && len(*user.Nickname) > 0 {
            username = *user.Nickname
        }

        roomname := *m.RoomID
        room, ok := client.roomMap[roomname]
        if ok && room.RoomName != nil && len(*room.RoomName) > 0 {
            roomname = *room.RoomName
        }
        return fmt.Sprintf("%15s:%15s: %s", roomname, username, *m.Message)

    case chatMessageType.ChatConnect != nil:
        m := chatMessageType.ChatConnect
        var result string
        if m.IsConnecting != nil && *m.IsConnecting {
            result = fmt.Sprintf("Connected to %s. Rooms:\n", client.connection.Config().Location.String())
            for _, room := range m.Rooms {
                result += fmt.Sprintf("%s\n", *room.RoomName)
            }
            result += fmt.Sprintf("%d rooms.", len(m.Rooms))
        } else {
            result = "Disconnected."
        }
        return result

    case chatMessageType.ChatEnterRoom != nil:
        m := chatMessageType.ChatEnterRoom
        if m.UserIsEntering != nil  && *m.UserIsEntering {
        }

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
        return fmt.Sprintf("Code %d: %s", *m.Code, *m.Message)

    default:
        return fmt.Sprintf("Unknown message type: %+v", chatMessageType)
    }

    return "Shouldn't happen."
}


//----------------------------------------------------------------------------------------
//                                                                                 Connect
//----------------------------------------------------------------------------------------


func (client *ChatClient) Connect(URL string, user ChatUser, readChannel chan <- *ChatMessageType) error {
    Log.LogFunctionName()

    client.lock.Lock()
    defer client.lock.Unlock()

    if client.connection != nil {
        return errors.New("Already connected")
    }

    parsedURL, error := url.Parse(URL)
    if error != nil {
        Log.LogError(error)
        return error
    }
    origin := "http://"+parsedURL.Host
    if parsedURL.Scheme == "wss" {
        origin = "https://"+parsedURL.Host
    }

    client.user.Format = FormatPtr(Format_FormatProtobuf)
    client.userMap = make(map[string]*ChatUser)
    client.roomMap = make(map[string]*ChatRoom)

    client.connection, error =  websocket.Dial(URL, "", origin)
    if error != nil {
        client.connection = nil
        return error
    }
    go client.ChatClientReader(readChannel)

    return nil
}


func (client *ChatClient) Disconnect() error {
    Log.LogFunctionName()

    client.lock.Lock()
    defer client.lock.Unlock()

    if client.connection == nil {
        return nil
    }

    error := client.connection.Close()
    client.connection = nil
    client.userMap = make(map[string]*ChatUser)
    client.roomMap = make(map[string]*ChatRoom)
    return error
}


func (client *ChatClient) LeaveRoom() error {
    Log.LogFunctionName()

    client.lock.RLock()
    defer client.lock.RUnlock()

    if client.connection == nil {
        return fmt.Errorf("Not connected")
    }
    if client.user == nil {
        return fmt.Errorf("Not connected as a user")
    }
    if client.room == nil {
        return fmt.Errorf("Not in a room")
    }

    chatMessage := ChatEnterRoom {
        UserIsEntering:     BoolPtr(false),
        User:               client.user,
        Room:               client.room,
    }
    return SendMessageToConnection(client.connection,
            *client.user.Format,
            &ChatMessageType { ChatEnterRoom: &chatMessage })
}


func (client *ChatClient) SendMessage(message string) error {
    Log.LogFunctionName()

    client.lock.RLock()
    defer client.lock.RUnlock()

    if client.connection == nil {
        return fmt.Errorf("Not connected")
    }
    if client.user == nil {
        return fmt.Errorf("Not connected as a user")
    }
    if client.room == nil {
        return fmt.Errorf("Not in a room")
    }

    chatMessage := ChatMessage {
        SenderID:       client.user.UserID,
        RoomID:         client.room.RoomID,
        Timestamp:      TimestampFromTime(time.Now()),
        Message:        &message,
    }
    return SendMessageToConnection(client.connection,
            *client.user.Format,
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

        case chatMessage.ChatConnect != nil:

            client.lock.Lock()
            client.room = nil
            client.roomMap = make(map[string]*ChatRoom)
            client.user = nil
            client.userMap = make(map[string]*ChatUser)
            connect := chatMessage.ChatConnect
            if connect.IsConnecting != nil  &&
               *connect.IsConnecting &&
               connect.User != nil &&
               connect.User.UserID != nil &&
               len(*connect.User.UserID) > 0 {
                client.user = connect.User
                client.userMap[*client.user.UserID] = client.user
                for _, room := range connect.Rooms {
                    if room.RoomID != nil && len(*room.RoomID) > 0 {
                        client.roomMap[*room.RoomID] = room
                    }
                }
            } else {
                client.connection.Close()
                client.connection = nil
            }
            client.lock.Unlock()


        case chatMessage.ChatMessage != nil ||
             chatMessage.ChatResponse != nil:


        case chatMessage.ChatEnterRoom != nil:

            client.lock.Lock()
            client.room = nil
            client.roomMap = make(map[string]*ChatRoom)
            enterRoom := chatMessage.ChatEnterRoom
            if enterRoom.Room != nil  &&
               enterRoom.Room.RoomID != nil &&
               len(*enterRoom.Room.RoomID) > 0 &&
               enterRoom.UserIsEntering != nil &&
               *enterRoom.UserIsEntering {
               client.room = enterRoom.Room
               client.roomMap[*enterRoom.Room.RoomID] = enterRoom.Room
            } else {
               client.room = nil
            }
            client.lock.Unlock()


        case chatMessage.ChatPresence != nil:

            client.lock.Lock()
            presence := chatMessage.ChatPresence
            client.userMap = make(map[string]*ChatUser)
            if presence.Room != nil && presence.Room.RoomID == client.room.RoomID {
                for _, user := range presence.Users {
                    client.userMap[*user.UserID] = user
                }
            }
            client.lock.Unlock()

        default:
            Log.Errorf("Error: unknown message %+v.", chatMessage)
        }

        clientReaderChannel <- &chatMessage
    }
}

