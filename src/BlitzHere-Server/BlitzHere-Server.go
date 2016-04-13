

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
    "golang.org/x/net/websocket"
    "ApplePushService"
    "MessagePusher"
    "BlitzMessage"
)


var (
    config                      ServerUtil.Configuration
    PushNotificationService     ApplePushService.Service
    globalMessagePusher         *MessagePusher.MessagePusher
)


//----------------------------------------------------------------------------------------
//                                                                  ServerResponseForError
//----------------------------------------------------------------------------------------


func ServerResponseForError(code BlitzMessage.ResponseCode, error error) *BlitzMessage.ServerResponse {
    Log.Errorf("%s. Error %s: %v.", Log.PrettyStackString(2), code.String(), error)
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
        response = UpdateFeedPost(session, requestMessageType)

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

    case *BlitzMessage.UserMessageUpdate:
        if serverRequest.RequestType.MessageSendRequest != nil {
            response =  UserMessageSendRequest(session, requestMessageType)
        }

    default:
        error = fmt.Errorf("Unrecognized request '%+v'", request)
        response = ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }


    //  Done --


    WriteResponse(writer, response, messageFormat)
}

/*
    friendUpdate := clientRequest.GetFriendUpdate()
    if friendUpdate != nil {
        UpdateFriends(writer, session, friendUpdate)
        return
    }
    friendRequest := clientRequest.GetFriendRequest()
    if friendRequest != nil {
        FriendRequest(writer, session, friendRequest)
        return
    }
    messageSendRequest := clientRequest.GetMessageSendRequest()
    if messageSendRequest != nil {
        SendMessages(writer, userID, messageSendRequest)
        return
    }
    messageFetchRequest := clientRequest.GetMessageFetchRequest()
    if messageFetchRequest != nil {
        FetchMessages(writer, userID, messageFetchRequest)
        return
    }
    debugRequest := clientRequest.GetDebugMessage()
    if debugRequest != nil {
        SaveDebugMessages(writer, userID, debugRequest)
        return
    }
    storyFetchRequest := clientRequest.GetStoryFetchRequest()
    if storyFetchRequest != nil {
        FetchStories(writer, userID, storyFetchRequest)
        return
    }
    imageUpload := clientRequest.GetImageUpload()
    if imageUpload != nil {
        UploadImage(writer, userID, imageUpload)
        return
    }
    acceptInviteRequest := clientRequest.GetAcceptInviteRequest()
    if acceptInviteRequest != nil {
        AcceptInviteRequest(writer, session, acceptInviteRequest)
        return
    }
    profilesFromContactInfo := clientRequest.GetProfilesFromContactInfo()
    if profilesFromContactInfo != nil {
        ProfilesFromContactInfoRequest(writer, session, profilesFromContactInfo)
        return
    }
    sendPulse := clientRequest.GetSendNewPulseBeat()
    if sendPulse != nil {
        SendNewPulseBeat(writer, session, sendPulse)
        return
    }
    scorePulseBeat := clientRequest.GetScorePulseBeat()
    if scorePulseBeat != nil {
        ScorePulseBeatRequest(writer, session, scorePulseBeat)
        return
    }
    pulsesForUser := clientRequest.GetPulsesForUser()
    if pulsesForUser != nil {
        GetPulsesForUser(writer, session, pulsesForUser)
        return
    }
    pulseStatusUpdate := clientRequest.GetPulseStatusUpdate()
    if pulseStatusUpdate != nil {
        UpdatePulseStatus(writer, session, pulseStatusUpdate)
        return
    }
    fetchScoresRequest := clientRequest.GetFetchScoresRequest()
    if fetchScoresRequest != nil {
        FetchScoresRequest(writer, session, fetchScoresRequest)
        return
    }
    validatePulseRequest := clientRequest.GetValidatePulseRequest()
    if validatePulseRequest != nil {
        ValidatePulseRequest(writer, session, validatePulseRequest)
        return
    }
    updateScoreRequest := clientRequest.GetScoreUpdate()
    if updateScoreRequest != nil {
        UpdateScoresRequest(writer, session, updateScoreRequest)
    }
*/


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


func AdminFormRequest(writer http.ResponseWriter, request *http.Request) {
    Log.LogFunctionName()

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, request, "error.html", 303)
        }
    } ()

    var templateMap struct {
        AppName     string
    }
    templateMap.AppName = config.AppName

    var error error
    error = config.Template.ExecuteTemplate(writer, "Admin.html", templateMap)
    if error != nil { panic(error) }
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
    if (flagVersion) {
        fmt.Fprintf(os.Stdout, "Version %s compiled %s.\n", ServerUtil.CompileVersion(), ServerUtil.CompileTime())
        return 0
    }
    if len(flagConfigFilename) > 0 {
        error = config.ParseConfigFileNamed(flagConfigFilename)
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

    //  Start --

    //  Apply configuration paramaters --


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

    //  Set up & start our http handlers --

    Session_InitializeSessions()
    http.HandleFunc(config.ServicePrefix+"/api", DispatchServiceRequests)
    http.HandleFunc(config.ServicePrefix+"/hello", sendHello)
    http.HandleFunc(config.ServicePrefix+"/image", GetImage)
    http.HandleFunc(config.ServicePrefix+"/admin", AdminFormRequest)
    http.HandleFunc(config.ServicePrefix+"/message", SystemMessageFormRequest)
    http.HandleFunc(config.ServicePrefix+"/shortlink", LinkShortnerFormRequest)

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

    //  Push messages --

    globalMessagePusher = MessagePusher.NewMessagePusher()
    http.Handle(config.ServicePrefix+"/push", websocket.Handler(PushMessageHandler))


    Log.Infof("Server listening at %d:%s.", config.ServicePort, config.ServicePrefix)
    http.Serve(httpListener, nil)
    return 0
}


func main () {
    exitStatus := Server()
    Log.Exitf("Exit status %d.", exitStatus)
    Log.FlushMessages();
    os.Exit(exitStatus)
}


