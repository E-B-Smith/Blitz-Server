//  Status  -  Status Messages.
//
//  Send a message via the Apple Push Notification service (APNs for short).
//
//  E.B.Smith  -  June, 2015


package ApplePushService
import "fmt"


func (status ServiceStatus) StatusString() string {
    connectionString := "Disconnected"
    if status.IsConnected {
        connectionString = "Connected"
    }
    return fmt.Sprintf("%s Queued: %d Sent: %d Errors: %d ErrorResponse: %d Feedback: %d State: %s %s",
        status.ServiceName,
        status.QueuedCount,
        status.MessageCount,
        status.ErrorCount,
        status.ErrorResponseCount,
        status.FeedbackCount,
        status.ServiceState.String(),
        connectionString,
    )
}

