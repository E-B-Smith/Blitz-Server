//  ApplePush  -  Send an APN message.
//
//  Send a message via the Apple Push Notification service (APNs for short).
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "time"
)


type ServiceType int


const (
    ServiceTypeDevelopment  ServiceType = iota
    ServiceTypeProduction
)


type Notification struct {
    BundleID        string
    ServiceType     ServiceType
    DeviceToken     string
    MessageID       uint32
    ExpirationDate  time.Time
    MessageText     string

    //  Optional fields --

    SoundName       string
    OptionalKeys    map[string]string   //  Other keys are 'badge', 'content-available', etc.

    //  Book-keeping fields --

    DateSent        time.Time
    messageBytes    []byte
}


type ResponseStatus uint8


const (
    ResponseSuccess             = 0
    ResponseProcessingError     = 1
    ResponseMissingDeviceToken  = 2
    ResponseMissingTopic        = 3
    ResponseMissingPayload      = 4
    ResponseInvalidTokenSize    = 5
    ResponseInvalidTopicSize    = 6
    ResponseInvalidPayloadSize  = 7
    ResponseInvalidToken        = 8
    ResponseServiceShutdown     = 10
    ResponseUnkownError         = 255
)


type Response struct {
    Timestamp       time.Time
    Command         uint8
    ResponseStatus  ResponseStatus
    MessageID       uint32
    Error           error
}


type Feedback struct {
    Timestamp       time.Time
    DeviceToken     string
}


type ServiceState int


const (
    ServiceStateStopped  ServiceState = iota
    ServiceStateRunning
)
//func (state SocketState) String() string


type ServiceStatus struct {
    BundleID            string
    ServiceType         ServiceType
    ServiceName         string
    QueuedCount         int
    MessageCount        int
    ErrorCount          int
    ErrorResponseCount  int
    FeedbackCount       int
    ServiceState        ServiceState
    IsConnected         bool
    LastError           error
}
func (status ServiceStatus) String() string {
    return status.StatusString()
}


type Service interface {
    Start() error
    Status() []ServiceStatus
    Stop()
    SetFeedbackResponseFilename(filename string)
    Send(notification *Notification) error
}


func NewService() Service {
    return newService()
}


