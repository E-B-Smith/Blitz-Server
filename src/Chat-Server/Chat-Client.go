

//----------------------------------------------------------------------------------------
//
//                                                                             Chat-Client
//                                                                    A simple chat client
//
//                                                                   E.B.Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "os"
    "fmt"
    "os/exec"
    "strings"
    "violent.blue/GoKit/Log"
    "./Chat"
)


//----------------------------------------------------------------------------------------
//                                                                                 Globals
//----------------------------------------------------------------------------------------


var ChatClient *Chat.ChatClient
var ChatChannel chan *Chat.ChatMessageType
const kChatHost = "ws://blitzhere.com/blitzlabs-chat"


func RemoveEmptyStrings(a []string) []string {
    result := make([]string, 0, len(a))
    for _, s := range a {
        if len(s) > 0 {
            result = append(result, s)
        }
    }
    return result
}

const (
    kTermReset          = "\033[0m"     // Reset all custom styles
    kTermResetColor     = "\033[32m"    // Reset to default color
    kTermResetLine      = "\r\033[K"    // Return curor to start of line and clean it
    kTermClearScreen    = "\033[2J\033[;H"
)


type TermColor int32


const (
    kTermColorBlack TermColor = iota
    kTermColorRed
    kTermColorGreen
    kTermColorYellow
    kTermColorBlue
    kTermColorMagenta
    kTermColorCyan
    kTermColorWhite
    kTermColorMax
)


func SetTextColor(color TermColor) string {
    return fmt.Sprintf("\033[3%dm", color)
}


func SetBackColor(color TermColor) string {
    return fmt.Sprintf("\033[4%dm", color)
}


//----------------------------------------------------------------------------------------
//                                                                   ProcessCommandMessage
//----------------------------------------------------------------------------------------


func ProcessCommandMessage(message string) string {
    Log.LogFunctionName()

    var error error
    messageParts := strings.Split(message, " ")
    messageParts  = RemoveEmptyStrings(messageParts)
    if len(messageParts) == 0 { return "" }

    switch messageParts[0] {

    case "\\connect":
        error = ChatClient.Connect(kChatHost, ChatChannel)

    case "\\disconnect":
        error = ChatClient.Disconnect()

    case "\\enter":

    case "\\leave":
        error = ChatClient.LeaveRoom()

    }

    return error.Error()
}


//----------------------------------------------------------------------------------------
//
//                                                                                    Main
//
//----------------------------------------------------------------------------------------


func main() {
    Log.LogLevel = Log.LogLevelAll
    Log.Debugf("Howdy! Debug trace logging is on.")

    var error error

    //  Disable input buffering:
    error = exec.Command("/bin/stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
    if error != nil { Log.LogError(error) }

/*
    //  Do not display entered characters on the screen:
    error = exec.Command("/bin/stty", "-f", "/dev/tty", "-echo").Run()
    if error != nil { Log.LogError(error) }

    //  Restore the echoing state when exiting:
    defer exec.Command("/bin/stty", "-f", "/dev/tty", "echo").Run()
*/

    ChatClient = new(Chat.ChatClient)
    ChatChannel     = make(chan *Chat.ChatMessageType)
    keyboardChannel:= make(chan byte)
    messageChannel := make(chan string)

    //  Process to read from keyboard:
    go  func() {

        b := make([]byte, 1)
        reader := os.Stdin
        for {
            n, error := reader.Read(b)
            if error == nil && n > 0 {
                keyboardChannel <- b[0]
            }
        }

    } ()

    //  Process to read from chat client:
    // go func() {
    //     var msgNo int64 = 1
    //     for {
    //         time.Sleep(5.0 * time.Second)
    //         messageChannel <- fmt.Sprintf("Message %d.", msgNo)
    //         msgNo++
    //     }
    // } ()

    go func() {
        for {
            chatMessage := <- ChatChannel
            text := ChatClient.StringFromMessage(chatMessage)
            messageChannel <- text
        }
    } ()


    inputBuffer  := make([]byte, 0, 4096)

    for {

        select {
        case b := <- keyboardChannel:
            if b == '\n' {
                message := strings.TrimSpace(string(inputBuffer))
                inputBuffer = inputBuffer[:0]
                if strings.HasPrefix(message, "\\") {
                    message = ProcessCommandMessage(message)
                } else {
                    ChatClient.SendMessage(message)
                }
            } else {
                inputBuffer = append(inputBuffer, b)
            }

        case s := <- messageChannel:
            s = strings.TrimSpace(s)
            fmt.Printf(kTermResetLine+"%s\n%s", s, string(inputBuffer))
        }
    }
}

