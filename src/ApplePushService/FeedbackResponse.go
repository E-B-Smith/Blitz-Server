//  FeedbackResponse.go  -  Write APNS feedback or response messages to a file.
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "io"
    "os"
    "fmt"
    "syscall"
    "path/filepath"
    "violent.blue/GoKit/Log"
)


type feedbackWriter struct {
    writer  io.WriteCloser
}


func (feedback *feedbackWriter) SetFilename(filename string) {
    feedback.Close()
    if len(filename) <= 0 {
        feedback.writer = os.Stderr
        return
    }
    var error error
    var flags int = syscall.O_APPEND | syscall.O_CREAT | syscall.O_WRONLY
    var mode os.FileMode = os.ModeAppend | 0700
    pathname := filepath.Dir(filename)
    if len(pathname) > 0 {
        if error = os.MkdirAll(pathname, 0700); error != nil {
            feedback.writer = os.Stderr
            Log.Errorf("Error: Can't create directory for log file '%s': %v.", filename, error)
        }
    }
    feedback.writer, error = os.OpenFile(filename, flags, mode)
    if error != nil {
        feedback.writer = os.Stderr
        Log.Errorf("Error: Can't open response file '%s' for writing: %v.", filename, error)
    }
}


func (feedback *feedbackWriter) Close() {
    _, hasClose := feedback.writer.(interface {Close()})
    if  hasClose &&
        feedback.writer != os.Stderr &&
        feedback.writer != os.Stdout {
        feedback.writer.Close()
    }
}


func (feedback *feedbackWriter) Write(response interface{}) {
    fmt.Fprintf(feedback.writer, "%+v\n", response)
}

