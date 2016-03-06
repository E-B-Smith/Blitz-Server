//  APSRunLoop  -  The push service run loop guts.
//
//  Send a message via the Apple Push Notification service (APNs for short).
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "sync"
    "errors"
    "sync/atomic"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//                                                                              APNService
//----------------------------------------------------------------------------------------


type APNService struct {
    notificationChannel         chan Notification
    serviceStatusChannel        chan ServiceState
    socketMap                   map[string](*PushSocket)
    serviceState                ServiceState
    serviceLock                 sync.Mutex
    nextMessageID               uint32
    errorCount                  int64
    feedbackWriter              feedbackWriter
}


//----------------------------------------------------------------------------------------
//                                                                                 RunLoop
//----------------------------------------------------------------------------------------


func (service *APNService) runLoop() {
    Log.LogFunctionName()

    defer func() {
        error := recover();
        if error != nil {
            Log.LogStackWithError(error)
        }
    } ()

    for service.serviceState == ServiceStateRunning {
        select {
        case notification  := <- service.notificationChannel:
            Log.Debugf("Routing notification %+v.", notification)

            servicename := notification.ServiceName()
            pushSocket, found := service.socketMap[servicename];
            if !found {
                pushSocket = NewPushSocketForNotification(notification)
                pushSocket.Feedback = &service.feedbackWriter
                service.socketMap[servicename] = pushSocket
            }
            if pushSocket.Status.ServiceState == ServiceStateStopped {
                pushSocket.Status.ServiceState = ServiceStateRunning
                go pushSocket.RunLoop()
            }
            pushSocket.NotificationChannel <- notification

        case service.serviceState = <- service.serviceStatusChannel:
        }
    }

    Log.Debugf("Shutting down.")
    for _, pushSocket := range(service.socketMap) {
        pushSocket.ServiceStateChannel <- ServiceStateStopped
        <- pushSocket.ServiceStateChannel
    }
    Log.Debugf("Service stopped.")
    service.serviceState = ServiceStateStopped
    service.serviceStatusChannel <- service.serviceState
}



//----------------------------------------------------------------------------------------
//                                                                                   Start
//----------------------------------------------------------------------------------------


func (service *APNService) Start() error {
    Log.LogFunctionName()
    service.serviceLock.Lock()
    defer service.serviceLock.Unlock()
    if  service.serviceState == ServiceStateRunning {
        return nil;
    }

    if  service.socketMap == nil {
        service.notificationChannel     = make(chan Notification)
        service.serviceStatusChannel    = make(chan ServiceState)
        service.socketMap               = make(map[string](*PushSocket))
    }

    service.serviceState = ServiceStateRunning
    go service.runLoop()

    return nil
}



//----------------------------------------------------------------------------------------
//                                                                                    Stop
//----------------------------------------------------------------------------------------


func (service *APNService) Stop() {
    Log.LogFunctionName()
    service.serviceLock.Lock()
    defer service.serviceLock.Unlock()
    service.serviceStatusChannel <- ServiceStateStopped
    <- service.serviceStatusChannel
    service.feedbackWriter.Close()
    Log.Debugf("Stop exit.")
}


//----------------------------------------------------------------------------------------
//                                                                                  Status
//----------------------------------------------------------------------------------------


func (service *APNService) Status() []ServiceStatus {
    Log.LogFunctionName()
    service.serviceLock.Lock()
    defer service.serviceLock.Unlock()

    var i int = 0
    status := make([]ServiceStatus, len(service.socketMap))
    for _, socket := range service.socketMap {
        status[i] = socket.Status
        i++
    }
    return status
}


//----------------------------------------------------------------------------------------
//                                                                                    Send
//----------------------------------------------------------------------------------------


func (service *APNService) Send(notification *Notification) error {
    Log.LogFunctionName()
    if notification == nil { return errors.New("Nil notification"); }
    if service.serviceState != ServiceStateRunning {
        service.errorCount++
        return errors.New("Service not started")
    }

    socketname := notification.ServiceName()
    if socketname == "" {
        service.errorCount++
        return errors.New("Invalid bundleID or service type")
    }
    var error error
    notification.MessageID = atomic.AddUint32(&service.nextMessageID, 1)
    notification.messageBytes, error = EncodeNotification(notification)
    if error != nil  {
        service.errorCount++
        return error
    }

    service.notificationChannel <- *notification
    Log.Debugf("Exit send.")
    return nil
}


//----------------------------------------------------------------------------------------
//                                                             SetFeedbackResponseFilename
//----------------------------------------------------------------------------------------


func (service *APNService) SetFeedbackResponseFilename(filename string) {
    service.feedbackWriter.SetFilename(filename)
}


//----------------------------------------------------------------------------------------
//                                                                              newService
//----------------------------------------------------------------------------------------


type ApplePushServiceWrapper struct {
    apnservice *APNService
}

func (apsw ApplePushServiceWrapper) Start() error                   { return apsw.apnservice.Start() }
func (apsw ApplePushServiceWrapper) Status() []ServiceStatus        { return apsw.apnservice.Status() }
func (apsw ApplePushServiceWrapper) Stop()                          { apsw.apnservice.Stop() }
func (apsw ApplePushServiceWrapper) SetFeedbackResponseFilename(filename string)    { apsw.apnservice.SetFeedbackResponseFilename(filename) }
func (apsw ApplePushServiceWrapper) Send(notification *Notification) error          { return apsw.apnservice.Send(notification) }

func newService() Service {
    var service ApplePushServiceWrapper
    service.apnservice = &APNService { nextMessageID: 1000 }
    service.SetFeedbackResponseFilename("")
    return service
}


