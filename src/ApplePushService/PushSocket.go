//  PushSocket  -  The push service TCP socket handler.
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "io"
    "fmt"
    "time"
    "sync"
    "bytes"
    "errors"
    "strings"
    "crypto/tls"
    "violent.blue/GoKit/Log"
)


type PushSocket struct {
    Status              ServiceStatus
    tlsConnection       *tls.Conn
    ReaderCompleted     sync.WaitGroup
    ServiceStateChannel chan ServiceState
    NotificationChannel chan Notification
    Feedback            *feedbackWriter
}



//----------------------------------------------------------------------------------------
//                                                                            ReadFeedback
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) ReadFeedback() error {
    Log.LogFunctionName()

    var error error
    socketName := pushSocket.Status.ServiceName
    pemdata := Resource.ResourceBytesNamed(socketName+".pem")
    if pemdata == nil {
        return fmt.Errorf("Fatal: PEM for '%s' not found", socketName)
    }
    keydata := Resource.ResourceBytesNamed(socketName+".key")
    if keydata == nil {
        return fmt.Errorf("Fatal: Key for '%s' not found", socketName)
    }
    certificate, error := tls.X509KeyPair(*pemdata, *keydata)
    if error != nil {
        Log.LogError(error)
        return error
    }

    hostname := "feedback.push.apple.com"
    if pushSocket.Status.ServiceType == ServiceTypeDevelopment {
        hostname = "feedback.sandbox.push.apple.com"
    }
    hostportname := hostname+":2196"

    //  eDebug - Make sure InsecureSkipVerify is false

    config := tls.Config {
        ServerName:                 hostname,
        Certificates:               []tls.Certificate{certificate},
        SessionTicketsDisabled:     false,
        InsecureSkipVerify:         false,
        ClientAuth:                 tls.RequireAndVerifyClientCert,
    }

    var messagesThisRun int
    var feedbackConnection *tls.Conn
    feedbackConnection, error = tls.Dial("tcp", hostportname, &config)
    if error != nil {
        Log.Errorf("Open error: %v.", error)
        return error
    }
    defer func() {
        feedbackConnection.Close()
        Log.Debugf("Read %d feedback messages.", messagesThisRun)
    } ()

    error = feedbackConnection.Handshake()
    if error != nil {
        Log.Errorf("TLS Handshake error: %v.", error)
        return error
    }

    costate := feedbackConnection.ConnectionState()
    Log.Debugf("Open connection state: %+v.", costate)
    Log.Debugf("Handshake: %v Resumed: %+v.", costate.HandshakeComplete, costate.DidResume)
    Log.Debugf("Opened socket.")

    var n int
    var readBuffer bytes.Buffer
    buffer := make([]byte, DecodeFeedbackByteLength)

    for error == nil {
        Log.Debugf("Waiting read...")
        feedbackConnection.SetReadDeadline(time.Now().Add(time.Second * 30))
        n, error = feedbackConnection.Read(buffer)
        Log.Debugf("Read %d bytes.  Error type is '%T'; error struct is '%+v'.", n, error, error)
        readBuffer.Write(buffer[:n])
        feedback := DecodeFeedback(&readBuffer)
        for feedback != nil {
            messagesThisRun++
            pushSocket.Status.FeedbackCount++
            pushSocket.Feedback.Write(feedback)
            feedback = DecodeFeedback(&readBuffer)
        }
    }
    return nil
}



//----------------------------------------------------------------------------------------
//                                                                                    Open
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) Open() error {
    Log.LogFunctionName()

    pushSocket.Close()

    var error error
    socketName := pushSocket.Status.ServiceName
    pemdata := Resource.ResourceBytesNamed(socketName+".pem")
    if pemdata == nil {
        return fmt.Errorf("Fatal: PEM for '%s' not found", socketName)
    }
    keydata := Resource.ResourceBytesNamed(socketName+".key")
    if keydata == nil {
        return fmt.Errorf("Fatal: Key for '%s' not found", socketName)
    }
    certificate, error := tls.X509KeyPair(*pemdata, *keydata)
    if error != nil {
        Log.LogError(error)
        return error
    }

    hostname := "gateway.push.apple.com"
    if pushSocket.Status.ServiceType == ServiceTypeDevelopment {
        hostname = "gateway.sandbox.push.apple.com"
    }
    hostportname := hostname+":2195"

    //  eDebug - Make sure InsecureSkipVerify is false

    config := tls.Config {
        ServerName:                 hostname,
        Certificates:               []tls.Certificate{certificate},
        SessionTicketsDisabled:     false,
        InsecureSkipVerify:         false,
        ClientAuth:                 tls.RequireAndVerifyClientCert,
    }
    Log.Debugf("TLS Config: %+v.", config)

    pushSocket.tlsConnection, error = tls.Dial("tcp", hostportname, &config)
    if error != nil || pushSocket.tlsConnection == nil {
        Log.Errorf("Open error: %+v.", error)
        return error
    }

    error = pushSocket.tlsConnection.Handshake()
    if error != nil {
        Log.Errorf("TLS Handshake error: %+v.", error)
        return error
    }
    time.Sleep(time.Second)     //  Wait for connection to settle down.
    costate := pushSocket.tlsConnection.ConnectionState()
    Log.Debugf("Open connection state: %+v.", costate)
    Log.Debugf("Handshake: %+v Resumed: %+v.", costate.HandshakeComplete, costate.DidResume)
    Log.Debugf("Opened socket.")
    pushSocket.Status.IsConnected = true

    go pushSocket.Read()
    time.Sleep(time.Second)     //  Wait for connection to settle down.
    go pushSocket.ReadFeedback()
    return nil
}



//----------------------------------------------------------------------------------------
//                                                                                  IsOpen
//----------------------------------------------------------------------------------------



func (pushSocket *PushSocket) IsOpen() bool {
    return (pushSocket != nil && pushSocket.tlsConnection != nil && pushSocket.Status.IsConnected)
}


//----------------------------------------------------------------------------------------
//                                                                                   Close
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) Close() {
    if  pushSocket == nil { return }
    if  pushSocket.tlsConnection != nil {
        pushSocket.tlsConnection.Close()
        pushSocket.tlsConnection = nil
    }
}


//----------------------------------------------------------------------------------------
//                                                                                   Write
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) Write(notification Notification) error {
    Log.LogFunctionName()
    if pushSocket.tlsConnection == nil {
        return errors.New("Error: Socket closed")
    }

    Log.Debugf("Start write...")
    pushSocket.tlsConnection.SetWriteDeadline(time.Now().Add(time.Second * 10.0))
    n, error := pushSocket.tlsConnection.Write(notification.messageBytes)
    if n != len(notification.messageBytes) && error == nil {
        error = errors.New("Error: Partial write")
    }
    if error != nil {
        Log.Errorf("Wrote %d of %d bytes with error %v.", n, len(notification.messageBytes), error)
        pushSocket.Status.ErrorCount++
    } else {
        Log.Debugf("Wrote full %d bytes.", n)
        notification.DateSent = time.Now()
        pushSocket.Status.MessageCount++
    }
    return error
}



//----------------------------------------------------------------------------------------
//                                                                                    Read
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) Read() {
    Log.LogFunctionName()

    var readError error

    defer func() {
        panicInfo := recover()

        switch {
        case panicInfo != nil:
            Log.LogStackWithError(panicInfo)
            pushSocket.Status.ErrorCount++
            pushSocket.Status.LastError = fmt.Errorf("Panic error: %+v", panicInfo)

        case readError == io.EOF || strings.Contains(readError.Error(), "use of closed network connection"):
            Log.Debugf("Read socket closed.")

        default:
            pushSocket.Status.ErrorCount++
            pushSocket.Status.LastError = readError
        }

        Log.Debugf("Exit Read.")
        pushSocket.Close()
        pushSocket.ReaderCompleted.Done()
    } ()

    var n int
    var readBuffer bytes.Buffer
    buffer := make([]byte, DecodeResponseByteLength)
    pushSocket.ReaderCompleted.Add(1)

    for readError == nil {
        Log.Debugf("Waiting read...")
        pushSocket.tlsConnection.SetReadDeadline(time.Time{})
        n, readError = pushSocket.tlsConnection.Read(buffer)
        Log.Debugf("Read %d bytes.  Error type is '%T'; error struct is '%+v'.", n, readError, readError)
        readBuffer.Write(buffer[:n])
        response := DecodeResponse(&readBuffer)
        Log.Debugf("Decoded response is %+v.", response)
        for response != nil {
            pushSocket.Status.ErrorResponseCount++
            pushSocket.Feedback.Write(response)
            response = DecodeResponse(&readBuffer)
        }
    }
}



//----------------------------------------------------------------------------------------
//                                                            NewPushSocketForNotification
//----------------------------------------------------------------------------------------


func NewPushSocketForNotification(notification Notification) *PushSocket {
    Log.LogFunctionName()
    socket := new(PushSocket)
    socket.Status.BundleID      = notification.BundleID
    socket.Status.ServiceType   = notification.ServiceType
    socket.Status.ServiceName   = notification.ServiceName()
    socket.ServiceStateChannel  = make(chan ServiceState)
    socket.NotificationChannel  = make(chan Notification)
    return socket
}



//----------------------------------------------------------------------------------------
//                                                                                 RunLoop
//----------------------------------------------------------------------------------------


func (pushSocket *PushSocket) RunLoop() {
    Log.LogFunctionName()

    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
        Log.Debugf("Waiting for reader to complete...")
        pushSocket.Status.ServiceState = ServiceStateStopped
        pushSocket.Close()
        pushSocket.ReaderCompleted.Wait()
        Log.Debugf("Exit RunLoop.")
        pushSocket.Status.IsConnected = false
        pushSocket.ServiceStateChannel <- pushSocket.Status.ServiceState
    } ()

    for pushSocket.Status.ServiceState >= ServiceStateRunning {
        select {
        case pushSocket.Status.ServiceState = <- pushSocket.ServiceStateChannel:
            Log.Debugf("Service state now %v.", pushSocket.Status.ServiceState)

        case notification := <- pushSocket.NotificationChannel:
            if ! pushSocket.IsOpen() {
                pushSocket.Status.LastError = nil
                pushSocket.Status.LastError = pushSocket.Open()
                if pushSocket.Status.LastError != nil {
                    pushSocket.Status.ServiceState = ServiceStateStopped
                    pushSocket.Close()
                }
            }
            pushSocket.Status.LastError = pushSocket.Write(notification)
        }
    }
}

