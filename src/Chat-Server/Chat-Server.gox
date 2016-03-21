//  Chat-Server  -  A simple chat server.
//
//  E.B.Smith  -  March, 2016


package main


import (
    "sync/Mutex"
    "violent.blue/GoKit/Log"
)


/*
message ChatMessage {
  optional string     senderID    = 1;
  optional string     roomID      = 2;
  optional Timestamp  timestamp   = 3;
  optional string     message     = 4;
}


message ChatUser {
  optional string     userID      = 1;
  optional string     nickname    = 2;
}


message ChatRoom {
  optional string     roomID      = 1;
  optional string     roomName    = 2;
}


message ChatConnect {
  optional bool       isConnecting      = 1;
  optional ChatUser   user              = 2;
  repeated ChatRoom   rooms             = 3;    //  <= Reply
}


message ChatRoomEntrance {
  optional ChatUser   user            = 1;
  optional string     roomID          = 2;
  optional bool       userIsEntering  = 3;
}


message ChatPresence {
  optional ChatRoom   room        = 1;
  repeated ChatUser   users       = 2;
}
*/


type ChatServer struct {
    mapLock     sync.Mutex
    userMap     map[string]*ChatUser
    roomMap     map[string]ChatRoom

    Status() []ServiceStatus
    Stop()
    Send(notification *Notification) error
}


type ChatMessage struct {
    Message     interface {}
}


func (server *ChatServer) ConnectUser(newUser ChatUser) ChatMessage {
    Log.LogFunctionName()

    if newUser.userID == nil || newUser.nickname == nil
        return ChatMessage{ Message: ChatError{ Code InputInvalid Message: "Invalid user." }
    }

    server.mapLock.Lock();
    defer server.mapLock.Unlock()

    user, ok = userMap[newUser.userID]
    if ! ok {
        userMap[newUser.userID] = &newUser
    }
    user.nickname = newUser.nickname


}
