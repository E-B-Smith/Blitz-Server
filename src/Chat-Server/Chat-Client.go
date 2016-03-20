//  Chat-Client  -  A simple chat server.
//
//  E.B.Smith  -  March, 2016


package main


import (
    "os"
    "fmt"
    "time"
    "os/exec"
    "strings"
    "violent.blue/GoKit/Log"
)


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

    inputChannel := make(chan byte)
    messageChannel := make(chan string)

    //  Process to read from keyboard:
    go  func() {

        b := make([]byte, 1)
        reader := os.Stdin
        for {
            n, error := reader.Read(b)
            if error == nil && n > 0 {
                inputChannel <- b[0]
            }
        }

    } ()

    //  Process to read from chat client:
    go func() {

        var msgNo int64 = 1
        for {
            time.Sleep(5.0 * time.Second)
            messageChannel <- fmt.Sprintf("Message %d.", msgNo)
            msgNo++
        }

    } ()

    inputBuffer  := make([]byte, 0, 4096)

    for {

        select {
        case b := <- inputChannel:
            if b == '\n' {
                message := strings.TrimSpace(string(inputBuffer))
                inputBuffer = inputBuffer[:0]
                if strings.HasPrefix(message, "\\") {
                    message = ProcessCommandMessage(message)
                } else {
                    SendMessage(message)
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
