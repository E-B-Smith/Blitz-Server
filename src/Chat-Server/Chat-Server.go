

//----------------------------------------------------------------------------------------
//
//                                                                          Chat-Server.go
//                                      Chat-Server: A simple client & server chat service
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "net/http"
    "golang.org/x/net/websocket"
    "./Chat"
)


var globalChatServer *Chat.ChatServer = nil


func ChatServerHandler(userConnection *websocket.Conn) {
    globalChatServer.HandleChatConnection(userConnection)
}


func main() {
    //  Initialize our chat server.
    globalChatServer = Chat.NewChatServer()

    //  Set the handler, then serve the connections.
    http.Handle("/chat", websocket.Handler(ChatServerHandler))
    error := http.ListenAndServe(":12345", nil)
    if error != nil {
        panic("ListenAndServe: " + error.Error())
    }
}

