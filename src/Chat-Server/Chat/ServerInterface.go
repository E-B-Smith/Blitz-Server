

//----------------------------------------------------------------------------------------
//
//                                                                      ServerInterface.go
//                                      Chat-Server: A simple client & server chat service
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package Chat


//  Hooks for the chat server
type ServerInterface interface {

    UserMayConnect(user ChatUser) (ChatUser, error)
    UserDidConnect(user ChatUser)
    UserDidDisconnect(user ChatUser)

    UserMayEnterRoom(user ChatUser, room ChatRoom) (ChatRoom, error)
    UserDidEnterRoom(user ChatUser, room ChatRoom)
    UserDidLeaveRoom(user ChatUser, room ChatRoom)

    UserMaySendMessage(user ChatUser, message ChatMessage) (ChatMessage, error)
    UserDidSendMessage(user ChatUser, message ChatMessage)
}

