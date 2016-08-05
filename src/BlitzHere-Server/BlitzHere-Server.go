

//----------------------------------------------------------------------------------------
//
//                                                  BlitzHere-Server : BlitzHere-Server.go
//                                                        The back-end server to BlitzHere
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "os"
    "fmt"
    "net"
    "time"
    "flag"
    "path"
    "errors"
    "strings"
    "reflect"
    "net/url"
    "net/http"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/ServerUtil"
    "github.com/golang/protobuf/proto"
    "github.com/stripe/stripe-go"
    "golang.org/x/net/websocket"
    "ApplePushService"
    "MessagePusher"
    "BlitzMessage"
)


type BlitzConfiguration struct {
    ServerUtil.Configuration
    StripeKey                   string
    ServiceIsFree               bool
    DailyChargeLimitDollars     float64
    ChatLimitHours              float64
}


var (
    config                      BlitzConfiguration
    PushNotificationService     ApplePushService.Service
    globalMessagePusher         *MessagePusher.MessagePusher
)


//----------------------------------------------------------------------------------------
//                                                                  ServerResponseForError
//----------------------------------------------------------------------------------------


func ServerResponseForError(code BlitzMessage.ResponseCode, error error) *BlitzMessage.ServerResponse {
    if code != BlitzMessage.ResponseCode_RCSuccess {
        Log.Errorf("%s. Error %s: %v.", Log.PrettyStackString(2), code.String(), error)
    }
    response := &BlitzMessage.ServerResponse{
        ResponseCode:   &code,
    }
    if  error != nil {
        response.ResponseMessage = proto.String(error.Error())
    }
    return response
}


func ServerResponseForCode(code BlitzMessage.ResponseCode, message *string) *BlitzMessage.ServerResponse {
    response := &BlitzMessage.ServerResponse{
        ResponseCode:       &code,
        ResponseMessage:    message,
    }
    return response
}


//----------------------------------------------------------------------------------------
//                                                                           MessageFormat
//----------------------------------------------------------------------------------------


type MessageFormat int
const (
    MFProtobuf MessageFormat = iota
    MFJSON
)

func (m MessageFormat) String() string {
    switch m {
    case MFProtobuf:
        return "MFProtobuf"
    case MFJSON:
        return "MFJSON"
    default:
        return "MFInvalid"
    }
}


//----------------------------------------------------------------------------------------
//                                                                           WriteResponse
//----------------------------------------------------------------------------------------


func WriteResponse(writer http.ResponseWriter, response *BlitzMessage.ServerResponse, messageFormat MessageFormat) {
    var error error
    var data []byte
    switch messageFormat {
    case MFProtobuf:
        writer.Header().Set("Content-Type", "application/x-protobuf")
        data, error = proto.Marshal(response)
    default:
        writer.Header().Set("Content-Type", "application/json")
        data, error = json.Marshal(response)
        data = append(data, []byte("\n")...)
        //Log.Debugf("%s", string(data))
    }
    if  error != nil {
        Log.Errorf("Error marshaling data %s: %v.", messageFormat.String(), error)
        http.Error(writer, "500 Internal Server Error", 500)
    } else {
        writer.Write(data)
    }
}


//----------------------------------------------------------------------------------------
//
//                                                                 DispatchServiceRequests
//
//----------------------------------------------------------------------------------------


func DispatchServiceRequests(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Dispatch the requests based on protobuf message type.

    //Log.Debugf("Dispatch content length: %d request:\n%+v.", httpRequest.ContentLength, httpRequest)
    Log.Debugf("========================================================================== "+
        "Dispatching new message with content length %d.", httpRequest.ContentLength)

    startTimestamp := time.Now()
    var request  interface{}
    var requestType reflect.Type
    var requestTypeName = "Unknown"
    var response *BlitzMessage.ServerResponse
    defer func() {
        error := recover();
        if error != nil {
            Log.Errorf("Panic! ==============================================================")
            Log.LogStackWithError(error)
            http.Error(writer, "500 Internal Server Error", 500)
        }
        elapsed :=  time.Since(startTimestamp).Seconds()
        //Log.Debugf("Exit dispatch Nowhere.  Message timestamp: %v Response Writer: %v\nHeader: %v", startTimestamp, writer, writer.Header())
        outlength, _ := strconv.Atoi(writer.Header().Get("Content-Length"))
        outstatus, _ := strconv.Atoi(writer.Header().Get("Status-Code"))
        var (code string; message string)
        if response != nil && response.ResponseCode != nil { code = response.ResponseCode.String() }
        if response != nil && response.ResponseMessage != nil { message = *response.ResponseMessage }
        _, error = config.DB.Exec("insert into ServerStatTable "+
          "(timestamp, elapsed, messageType, bytesIn, bytesOut, statusCode, responseCode, responseMessage)"+
          " values ($1, $2, $3, $4, $5, $6, $7, $8);",
            startTimestamp,
            elapsed,
            requestTypeName,
            httpRequest.ContentLength,
            outlength,
            outstatus,
            code,
            message)
        if error != nil {
            Log.Errorf("Error writing ServerStatTable: %v.", error)
        }
        Log.Debugf("=============================================="+
            " Exit dispatch %s.  Status: %d Code: %s Elapsed: %5.3f Timestamp: %v.",
            requestTypeName, outstatus, code, elapsed, startTimestamp)
    } ()

    config.MessageCount++

    defer httpRequest.Body.Close()
    body, _ := ioutil.ReadAll(httpRequest.Body)

    //  Decode the message --

    var error error
    var messageFormat MessageFormat

    contentType := httpRequest.Header.Get("content-type")
    bodyPrefix  := string(body[:Util.Min(16, len(body))])
    bodyPrefix   = strings.TrimSpace(bodyPrefix)

    serverRequest := BlitzMessage.ServerRequest {}
    if contentType == "application/json" || strings.HasPrefix(bodyPrefix, "{") {
        Log.Debugf("JSON:\n%s\n.", string(body))
        error = json.Unmarshal(body, &serverRequest)
        messageFormat = MFJSON
    } else {
        error = proto.Unmarshal(body, &serverRequest)
        messageFormat = MFProtobuf
    }
    if error != nil || serverRequest.RequestType == nil {
        Log.Errorf("Proto decode error: %v.", error)
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCInputCorrupt, error)
        WriteResponse(writer, response, messageFormat)
        return
    }
    if httpRequest.ContentLength > 10000 {
        Log.Debugf("Request too long. Not showing decoded message.")
    } else {
        Log.Debugf("Decoded request: %+v.", serverRequest)
    }

    //  Find the message type to log it --

    requestValue := reflect.ValueOf(*serverRequest.RequestType)
    if ! requestValue.IsValid() {
        Log.Errorf("Invalid request %+v.", serverRequest.RequestType)
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCInputCorrupt, error)
        WriteResponse(writer, response, messageFormat)
        return
    }

    for i := 0; i < requestValue.NumField(); i++ {
        field := requestValue.Field(i)
        if  field.IsValid() && ! field.IsNil() {
            request = field.Interface()
            break
        }
    }
    if ! reflect.ValueOf(request).IsValid() {
        error = errors.New("Invalid request type.")
        Log.Errorf("Invalid request: %v.", error)
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCInputCorrupt, error)
        WriteResponse(writer, response, messageFormat)
        return
    }

    requestType = reflect.ValueOf(request).Elem().Type()
    if requestType != nil { requestTypeName = requestType.Name() }
    Log.Debugf("Request type '%s'.", requestTypeName)

    //  Update the session if requested --

    sessionToken := serverRequest.GetSessionToken()
    sessionRequest := serverRequest.RequestType.GetSessionRequest()
    if  sessionRequest != nil {
        ipAddress := Util.IPAddressFromHTTPRequest(httpRequest)
        response = UpdateSession(ipAddress, sessionToken, sessionRequest)
        WriteResponse(writer, response, messageFormat)
        return
    }

    //  Get the session --

    session := Session_SessionFromToken(sessionToken)
    if session == nil {
        Log.Errorf("Invalid sessionToken '%s'.  Message type: %v.", sessionToken, requestTypeName)
        error = fmt.Errorf("Not logged in.")
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCNotAuthorized, error)
        WriteResponse(writer, response, messageFormat)
        return
    }
    userID := session.UserID
    Log.Debugf("------------------------------------------ UserID %s messageType %s.", userID, requestType.Name())


    //  Dispatch the message --


    error = nil
    switch requestMessageType := request.(type) {

    case *BlitzMessage.UserProfileUpdate:
        response = UpdateProfiles(session, requestMessageType)

    case *BlitzMessage.UserProfileQuery:
        response = QueryProfiles(session, requestMessageType)

    case *BlitzMessage.UserEventBatch:
        response = UpdateUserTrackingBatch(session, requestMessageType)

    case *BlitzMessage.ConfirmationRequest:
        response = UserConfirmation(session, requestMessageType)

    case *BlitzMessage.ImageUpload:
        response = UploadImage(session, requestMessageType)

    case *BlitzMessage.FeedPostUpdateRequest:
        response = UpdateFeedPostBatch(session, requestMessageType)

    case *BlitzMessage.FeedPostFetchRequest:
        response = FetchFeedPosts(session, requestMessageType)

    case *BlitzMessage.EntityTagList:
        response = UpdateEntityTags(session, requestMessageType)

    case *BlitzMessage.AutocompleteRequest:
        response = AutocompleteRequest(session, requestMessageType)

    case *BlitzMessage.UserSearchRequest:
        response = UserSearchRequest(session, requestMessageType)

    case *BlitzMessage.ConversationRequest:
        response = StartConversation(session, requestMessageType)

    case *BlitzMessage.FetchConversations:
        response = FetchConversations(session, requestMessageType)

    case *BlitzMessage.UserMessageUpdate:
        response = UserMessageFetchRequest(session, requestMessageType)

    case *BlitzMessage.UserMessage:
        response = SendUserMessage(session, requestMessageType)

    case *BlitzMessage.UserReview:
        response = WriteReview(session, requestMessageType)

    case *BlitzMessage.UpdateConversationStatus:
        response = UpdateConversationStatus(session, requestMessageType)

    case *BlitzMessage.Charge:
        response = ChargeRequest(session, requestMessageType)

    case *BlitzMessage.UserCardInfo:
        response = UpdateCards(session, requestMessageType)

    case *BlitzMessage.FriendUpdate:
        response = SendFriendRequest(session, requestMessageType)

    case *BlitzMessage.SearchCategories:
        response = FetchSearchCategories(session, requestMessageType)

    case *BlitzMessage.EditProfile:
        response = StartEditProfile(session, requestMessageType)

    case *BlitzMessage.FetchConversationGroups:
        response = FetchConversationGroups(session, requestMessageType)

    case *BlitzMessage.LoginAsAdmin:
        response = LoginAsAdmin(session, requestMessageType)

    case *BlitzMessage.FetchPurchaseDescription:
        response = FetchPurchaseDescription(session, requestMessageType)

    case *BlitzMessage.UserInvites:
        response = SendUserInvites(session, requestMessageType)

    default:
        error = fmt.Errorf("Unrecognized request '%+v'", request)
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }


    //  Done --


    WriteResponse(writer, response, messageFormat)
}


//----------------------------------------------------------------------------------------
//                                                                         Request Helpers
//----------------------------------------------------------------------------------------


func showRequest(writer http.ResponseWriter, request *http.Request) {
   Log.Debugf("Request:\n%v\nServer File Path:\n%s", request, config.ServiceFilePath)
   fmt.Fprintf(writer, "<html><p>Hi!\n<br>\n<br>Request:\n<br>%v\n<br>\n<br>File Path:  %s\n<br></p></html>",
        request, config.ServiceFilePath)
}


func sendHello(writer http.ResponseWriter, request *http.Request) {
   Log.Debugf("Request:\n%v\n", request)
   fmt.Fprintf(writer, "<html><p>Hello!\n\n<br><br>Request:\n<br>%v\n<br></p></html>", request)
}


func PushMessageHandler(userConnection *websocket.Conn) {
    globalMessagePusher.HandlePushConnection(userConnection)
}


//----------------------------------------------------------------------------------------
//
//                                                                      Main & BlitzServer
//
//----------------------------------------------------------------------------------------


func Server() (returnValue int) {
    //  The body of the work --

    defer func() {
        //  Catch panics --
        error := recover()
        if error != nil {
            Log.Errorf("Panic!")
            Log.LogStackWithError(error)
            returnValue = 1
        }
    } ()

    var error error
    Log.LogLevel = Log.LogLevelAll

    //  Do config params --

    var (flagUsage bool; flagVersion bool; flagVerbose bool; flagPID bool; flagConfigFilename string)

    flag.BoolVar(&flagUsage,   "h", false, "Help.  Print usage and exit.")
    flag.BoolVar(&flagUsage,   "?", false, "Help.  Print usage and exit.")
    flag.BoolVar(&flagVersion, "v", false, "Version.  Print version and exit.")
    flag.BoolVar(&flagVerbose, "V", false, "Verbose.  Verbose output.")
    flag.BoolVar(&flagPID,     "p", false, "PID filename.  Print the pid filename and exit.")
    flag.StringVar(&flagConfigFilename, "c", "", "Configuration.  The file from which to read the configuration.")
    flag.Parse()

    if (flagUsage) {
        flag.Usage()
        return 0
    }
    if flagVerbose {
        config.LogLevel = Log.LogLevelDebug
    }
    config.LogLevel = Log.LogLevelDebug
    if (flagVersion) {
        fmt.Fprintf(os.Stdout, "Version %s compiled %s.\n", Util.CompileVersion(), Util.CompileTime())
        return 0
    }
    if len(flagConfigFilename) > 0 {
        error = ServerUtil.ParseConfigFileNamed(&config, flagConfigFilename)
        if error != nil {
            Log.Errorf("Error: %v.", error)
            return 1
        }
    }
    if flagPID {
        fmt.Fprintf(os.Stdout, "%s\n", config.PIDFileName())
        return 0
    }
    if error = config.OpenConfig(); error != nil {
        Log.Errorf("Configuration error: %v", error)
        return 1
    }
    if flagVerbose {
        config.LogLevel = Log.LogLevelDebug
    }

    //  Check the configuration --

    if len(config.StripeKey) == 0 {
        Log.Errorf("No Stripe key found.")
        return 1
    }
    stripe.Key = config.StripeKey
    if config.ChatLimitHours <= 0 {
        config.ChatLimitHours = 24
    }

    //  Start --

    //  Add a start time to the database --

    _, error = config.DB.Exec("insert into ServerStatTable "+
       "  (timestamp, messageType) values ($1, 'Started');", time.Now());
    if error != nil {
        Log.Errorf("Error writing ServerStatTable: %v.", error)
    }

    //  Defer closing --

    defer func() {
        error := recover();
        if error != nil {
            message := fmt.Sprintf("%v", error)
            config.DB.Exec("insert into ServerStatTable "+
                "  (timestamp, messageType, responseMessage) values ($1, 'Fatal', $2);", time.Now(), message);
        }
        _, error = config.DB.Exec("insert into ServerStatTable "+
            "  (timestamp, messageType) values ($1, 'Terminated');", time.Now())
        if error != nil {
            Log.Errorf("Error writing ServerStatTable: %v.", error)
        }
        config.CloseConfig()
    } ()

    //  Make our listener --

    httpListener, error := net.Listen("tcp", ":"+strconv.Itoa(config.ServicePort))
    if error != nil {
        Log.Errorf("Can't listen on port %d: %v.", config.ServicePort, error)
        return 1
    }

    //  Set up an interrupt handler --

    config.AttachToInterrupts(httpListener)

    //  Start the messenger services --

    apnsfilename := path.Dir(config.LogFilename)
    apnsfilename  = path.Join(apnsfilename, "APNS.log")
    PushNotificationService = ApplePushService.NewService();
    PushNotificationService.SetFeedbackResponseFilename(apnsfilename)
    PushNotificationService.Start()
    defer PushNotificationService.Stop()

    StartNotifier()
    defer StopNotifier()

    StartScheduler()
    defer StopScheduler()
    ScheduleTask(time.Minute, ConversationCloser)

    //  Set up & start our http handlers --

    Session_InitializeSessions()
    http.HandleFunc(config.ServicePrefix+"/api", DispatchServiceRequests)
    http.HandleFunc(config.ServicePrefix+"/hello", sendHello)
    http.HandleFunc(config.ServicePrefix+"/image", GetImage)
    http.HandleFunc(config.ServicePrefix+"/downloadapp", DownloadAppRequest)
    http.HandleFunc(config.ServicePrefix+"/admin", AdminFormRequest)
    http.HandleFunc(config.ServicePrefix+"/admin/message", SystemMessageFormRequest)
    http.HandleFunc(config.ServicePrefix+"/admin/shortlink", LinkShortnerFormRequest)
    http.HandleFunc(config.ServicePrefix+"/admin/userlist", WebUserList)
    http.HandleFunc(config.ServicePrefix+"/admin/updateprofile", WebUpdateProfile)

    //  Set up short links --

    url, _ := url.Parse(config.ShortLinkURL)
    if url != nil {
        path := strings.TrimRight(url.Path, "/ ")
        if len(path) > 0 {
            Log.Infof("Shortlink redirection at %s.", path)
            http.HandleFunc(path+"/", RedirectShortLink)
        }
    }

    //  Set up our app deep-link handler --

    appLinkURL, _  := url.Parse(config.AppLinkURL)
    if appLinkURL != nil {
        Log.Infof("Starting HTTPAppDeepLink handler at '%s'.", appLinkURL.Path)
        http.Handle(appLinkURL.Path, RedirectWithQueryHandler(appLinkURL.Path+"/", 308))
        http.HandleFunc(appLinkURL.Path+"/", HTTPAppDeepLink)
    }

    //  Catch all --

    http.Handle("/",
        http.StripPrefix(config.ServicePrefix,
        http.FileServer(http.Dir(config.ServiceFilePath))))

    //  Handle push messages --

    globalMessagePusher = MessagePusher.NewMessagePusher()
    globalMessagePusher.UserDidConnect = UserDidConnectToPusher
    http.Handle(config.ServicePrefix+"/push", websocket.Handler(PushMessageHandler))

    server := &http.Server{
        ReadTimeout:   5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    Log.Infof("Server listening at %d:%s.", config.ServicePort, config.ServicePrefix)
    server.Serve(httpListener)
    return 0
}


func main () {
    exitStatus := Server()
    Log.Exitf("Exit status %d.", exitStatus)
    Log.FlushMessages();
    os.Exit(exitStatus)
}


