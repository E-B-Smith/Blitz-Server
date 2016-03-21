

//----------------------------------------------------------------------------------------
//
//                                                                      ChatMessageType.go
//                                                   Simple chat client & server functions
//
//                                                                   E.B.Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package Chat


import (
    "fmt"
)


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
