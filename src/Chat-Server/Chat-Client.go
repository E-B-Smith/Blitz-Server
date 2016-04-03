

//----------------------------------------------------------------------------------------
//
//                                                                          Chat-Client.go
//                                      Chat-Server: A simple client & server chat service
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


const kChatHost = "ws://localhost:12345/chat"


//----------------------------------------------------------------------------------------
//                                                                       Terminal Handling
//----------------------------------------------------------------------------------------


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


func SetBackgroundColor(color TermColor) string {
    return fmt.Sprintf("\033[4%dm", color)
}


func DisableTerminalInputBuffer() error {
    //  Disable input buffering:
    error := exec.Command("/bin/stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
    if error != nil { Log.LogError(error) }
    return error
}


func SetTerminalEcho(enable bool) error {
    //  Do not display entered characters on the screen:
    s := "-echo"
    if enable { s = "echo"}
    error := exec.Command("/bin/stty", "-f", "/dev/tty", s).Run()
    if error != nil { Log.LogError(error) }
    return error
}


func RestoreTerminal() error {
    SetTerminalEcho(true)
    error := exec.Command("/bin/stty", "-f", "/dev/tty", "sane").Run()
    if error != nil { Log.LogError(error) }
    fmt.Printf("%s\n", kTermReset)
    return error
}


//----------------------------------------------------------------------------------------
//                                                                   ProcessCommandMessage
//----------------------------------------------------------------------------------------


var ChatClient *Chat.ChatClient
var ChatChannel chan *Chat.ChatMessageType


func ProcessCommandMessage(message string) string {
    Log.LogFunctionName()

    var error error
    messageParts := strings.Split(message, " ")
    messageParts  = RemoveEmptyStrings(messageParts)
    if len(messageParts) == 0 { return "" }

    switch messageParts[0] {

    case "\\connect":
        if len(messageParts) < 3 {
            error = fmt.Errorf("Expected: \\connect <user-id> <nickname>")
        } else {
            user := Chat.ChatUser {
                UserID:         &messageParts[1],
                Nickname:       &messageParts[2],
                Format:         Chat.FormatPtr(Chat.Format_FormatProtobuf),
            }
            error = ChatClient.Connect(kChatHost, user, ChatChannel)
        }

    case "\\disconnect":
        error = ChatClient.Disconnect()

    case "\\enter":

    case "\\leave":
        error = ChatClient.LeaveRoom()

    default:
        error = fmt.Errorf("Unknown command '%s'.", messageParts[0])
    }

    return error.Error()
}


func RemoveEmptyStrings(a []string) []string {
    result := make([]string, 0, len(a))
    for _, s := range a {
        s = strings.TrimSpace(s)
        if len(s) > 0 {
            result = append(result, s)
        }
    }
    return result
}


//----------------------------------------------------------------------------------------
//
//                                                                                    Main
//
//----------------------------------------------------------------------------------------


func main() {
    Log.LogLevel = Log.LogLevelAll
    Log.Debugf("Howdy! Debug trace logging is on.")

    ChatClient      = Chat.NewChatClient()
    ChatChannel     = make(chan *Chat.ChatMessageType)
    keyboardChannel:= make(chan byte)
    messageChannel := make(chan string)

    DisableTerminalInputBuffer()
    defer RestoreTerminal()

    //  Read from the keyboard:
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

    // For testing, pretend we're reading the chat client:
    // go func() {
    //     var msgNo int64 = 1
    //     for {
    //         time.Sleep(5.0 * time.Second)
    //         messageChannel <- fmt.Sprintf("Message %d.", msgNo)
    //         msgNo++
    //     }
    // } ()

    // Read from the chat client:
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
                if message == "\\exit" {
                    fmt.Printf(kTermResetLine+"Goodbye\n")
                    break
                }
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

