//  Message.go  -  Encode/Decode APN messages.
//
//  E.B.Smith  -  June, 2015


package ApplePushService


import (
    "fmt"
    "time"
    "bytes"
    "errors"
    "strings"
    "encoding/hex"
    "encoding/json"
    "encoding/base64"
    "encoding/binary"
    "violent.blue/GoKit/Log"
)


func (state ServiceState) String() string {
    var stateString []string = []string{
        "Stopped",
        "Running",
        "Connected",
    }
    return stateString[state]
}



//----------------------------------------------------------------------------------------
//                                                                      EncodeNotification
//----------------------------------------------------------------------------------------


//    For 'new' style APNs --


func AddAPNItem(buffer *bytes.Buffer, itemID uint8, item interface{}) {
    //  Add an APNS item packet to the buffer --
    binary.Write(buffer, binary.BigEndian, itemID)
    switch t := item.(type) {
        case uint8, int8:
            binary.Write(buffer, binary.BigEndian, uint16(1))
            binary.Write(buffer, binary.BigEndian, item)
        case uint16, int16:
            binary.Write(buffer, binary.BigEndian, uint16(2))
            binary.Write(buffer, binary.BigEndian, item)
        case uint32, int32:
            binary.Write(buffer, binary.BigEndian, uint16(4))
            binary.Write(buffer, binary.BigEndian, item)
        case string, []uint8, []int8:
            stringitem, _ := item.([]byte)
            var strlen uint16 = uint16(len(stringitem))
            binary.Write(buffer, binary.BigEndian, strlen)
            buffer.Write(stringitem)
        default:
            error := fmt.Errorf("Unexpected type %T", t)
            Log.LogError(error)
            panic(error)
    }
}


func EncodeNotification(notification *Notification) ([]byte, error) {
    Log.LogFunctionName()

    if  notification.BundleID == "" {
        return nil, errors.New("Missing bundle ID")
    }

    if notification.DeviceToken == "" {
        return nil, errors.New("Missing device token")
    }

    payloadvalues := make(map[string]string)
    for key, val := range notification.OptionalKeys {
        payloadvalues[key] = val
    }
    message := strings.TrimSpace(notification.MessageText);
    if len(message) > 0 {
        payloadvalues["alert"] = message
    }

    if _, ok := payloadvalues["sound"]; !ok {
        if notification.SoundName == "" {
            payloadvalues["sound"] = "default"
        } else {
            payloadvalues["sound"] = notification.SoundName
        }
    }

    if len(payloadvalues) == 0 {
        return nil, errors.New("Error: Empty message")
    }

    apsmap := make(map[string](map[string]string))
    apsmap["aps"] = payloadvalues

    var error error
    var payload []byte
    payload, error = json.Marshal(apsmap)
    Log.Debugf("Payload: %s.", string(payload))

    var tokenBytes []byte
    tokenBytes, error = hex.DecodeString(notification.DeviceToken)
    if error != nil || len(tokenBytes) != 32 {
        return nil, fmt.Errorf("Invalid device token '%s'", notification.DeviceToken)
    }
    Log.Debugf("Device token: '%s'.", hex.EncodeToString(tokenBytes))

    var messageID uint32 = notification.MessageID

    var epochExpire uint32 = uint32(time.Now().Unix() + 12.0 * 60.0 * 60.0);  //  12 hours.
    if !notification.ExpirationDate.IsZero() {
        epochExpire = uint32(notification.ExpirationDate.Unix())
    }

    buffer := new(bytes.Buffer)

    binary.Write(buffer, binary.BigEndian, uint8(2))            //  Command
    binary.Write(buffer, binary.BigEndian, uint32(0))           //  Frame length.  Fix up later.
    AddAPNItem(buffer, 1, tokenBytes)
    AddAPNItem(buffer, 2, payload)
    AddAPNItem(buffer, 3, uint32(messageID))
    AddAPNItem(buffer, 4, uint32(epochExpire))
    AddAPNItem(buffer, 5, uint8(10))                            //  Priority

    //  Fix up length --

    binary.BigEndian.PutUint32(buffer.Bytes()[1:5], uint32(buffer.Len()-5))

    Log.Debugf("Notification #%d: %d bytes:\n%s.", messageID, buffer.Len(), hex.Dump(buffer.Bytes()));

    //  Done --

    return buffer.Bytes(), nil
}


func (notification Notification) ServiceName() string {
    if notification.BundleID == "" ||
        !(notification.ServiceType == ServiceTypeDevelopment ||
          notification.ServiceType == ServiceTypeProduction) {
          return ""
    }
    var socketName string
    if notification.ServiceType == ServiceTypeDevelopment {
        socketName = notification.BundleID + ".development"
    } else {
        socketName = notification.BundleID + ".production"
    }
    return socketName
}



//----------------------------------------------------------------------------------------
//                                                                          DecodeResponse
//----------------------------------------------------------------------------------------


const DecodeResponseByteLength = 6


func DecodeResponse(responseBuffer *bytes.Buffer) *Response {
    if responseBuffer.Len() < DecodeResponseByteLength { return nil; }

    errorStrings := []string{
        "No errors encountered",// 0
        "Processing error",
        "Missing device token",
        "Missing topic",
        "Missing payload",
        "Invalid token size",   // 5
        "Invalid topic size",
        "Invalid payload size",
        "Invalid token",
        "Unknown error",
        "Service shutdown",     // 10
    }

    var response Response
    binary.Read(responseBuffer, binary.BigEndian, &response.Command)
    binary.Read(responseBuffer, binary.BigEndian, &response.ResponseStatus)
    binary.Read(responseBuffer, binary.BigEndian, &response.MessageID)

    index := int(response.ResponseStatus)
    if index >= len(errorStrings) { index = 9; }
    response.Error = errors.New(errorStrings[index])
    response.Timestamp = time.Now()

    return &response
}



//----------------------------------------------------------------------------------------
//                                                                          DecodeFeedback
//----------------------------------------------------------------------------------------


const DecodeFeedbackByteLength = 38


func DecodeFeedback(feedbackBuffer *bytes.Buffer) *Feedback {
    if feedbackBuffer.Len() < DecodeFeedbackByteLength { return nil; }

    var (
        feedback    Feedback
        epoch       uint32
        tsize       uint16
    )
    binary.Read(feedbackBuffer, binary.BigEndian, &epoch)
    binary.Read(feedbackBuffer, binary.BigEndian, &tsize)

    tokenBytes := make([]byte, tsize)
    feedbackBuffer.Read(tokenBytes)
    feedback.DeviceToken = base64.StdEncoding.EncodeToString(tokenBytes)
    feedback.Timestamp = time.Unix(int64(epoch), 0)

    return &feedback
}


