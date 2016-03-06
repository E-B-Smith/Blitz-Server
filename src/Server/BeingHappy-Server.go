//  Happiness-Server  -  The server back-end to BeingHappy.
//
//  E.B.Smith  -  November, 2014


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
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/ServerUtil"
    "github.com/golang/protobuf/proto"
    "ApplePushService"
    "happiness"
)


var globalVersion            string = "0.0.0"
var globalCompileTime        string = "Wed Nov 11 09:01:25 PST 2015"
var PushNotificationService  ApplePushService.Service;

var config ServerUtil.Configuration =
    ServerUtil.Configuration {
        ServiceName:        "HappyLabs-Server",
        ServicePort:        9797,
        ServiceFilePath:    "./HappyLabs",
        ServicePrefix:      "/beinghappy/service-d",
        ServerURL:          "https://violent.blue",
        DatabaseURI:        "psql://happylabsadmin:happylabsadmin@localhost:5432/happylabsdatabase",
        LogLevel:           Log.LevelAll,
        LogFilename:        "",
    }


//----------------------------------------------------------------------------------------
//                                                                               SendError
//----------------------------------------------------------------------------------------


func SendError(writer http.ResponseWriter, code happiness.ResponseCode, error error) {
    //  Write a server error to http body --
    Log.Errorf("%s. Error %s: %v.", Log.PrettyStackString(2), code.String(), error)
    response := &happiness.ServerResponse{}
    response.ResponseCode = &code
    if  error != nil {
        response.ResponseMessage = proto.String(error.Error())
    }
    data, _ := proto.Marshal(response)
    writer.Write(data)
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

    var (
        messageType     string
        responseCode    int
        responseMessage string
    )
    startTimestamp := time.Now()
    defer func() {
        error := recover();
        if error != nil {
            Log.LogStackWithError(error)
            if messageType == "" { messageType = "Panic" }
        }
        //Log.Debugf("Exit dispatch message.  Message timestamp: %v Response Writer: %v\nHeader: %v", startTimestamp, writer, writer.Header())
        elapsed := time.Since(startTimestamp).Seconds()
        Log.Debugf("Exit dispatch message.  Elapsed: %5.3f Timestamp: %v.", elapsed, startTimestamp)
        Log.Debugf("==========================================================================")
        outlength, _ := strconv.Atoi(writer.Header().Get("Content-Length"))
        outstatus, _ := strconv.Atoi(writer.Header().Get("Status-Code"))
        _, error = config.DB.Exec("insert into MessageStatTable "+
          "(timestamp, elapsed, message, bytesIn, bytesOut, statusCode, responseCode, responseMessage)"+
          " values ($1, $2, $3, $4, $5, $6, $7, $8);",
            startTimestamp,
            elapsed,
            messageType,
            httpRequest.ContentLength,
            outlength,
            outstatus,
            responseCode,
            responseMessage)
        if error != nil {
            Log.Errorf("Error writing MessageStatTable: %v.", error)
        }
    } ()

    config.MessageCount++

    defer httpRequest.Body.Close()
    body, _ := ioutil.ReadAll(httpRequest.Body)

    var error error
    clientRequest := happiness.ClientRequest{}
    error = proto.Unmarshal(body, &clientRequest)
    if error != nil {
        Log.Errorf("Proto decode error: %v.", error)
        SendError(writer, happiness.ResponseCode_RCInputCorrupt, error)
        return
    }
    Log.Debugf("Message:\n%+v.", clientRequest)  //  eDebug - remove


    //  Find the message type to log it --

    messageType = "Unknown"
    clientRequestMessage := clientRequest.GetClientRequestMessage()
    if clientRequestMessage == nil {
        messageType = "None"
    } else {
        messageValue := reflect.ValueOf(clientRequestMessage)
        if messageValue.Elem().IsValid() {
           messageType = messageValue.Elem().Type().Name()
        }
    }

    //  Update the session if requested --

    sessionToken := clientRequest.GetSessionToken()
    sessionRequest := clientRequest.GetSessionRequest()
    if  sessionRequest != nil {
        ipAddress := Util.IPAddressFromHTTPRequest(httpRequest)
        UpdateSession(writer, ipAddress, sessionToken, sessionRequest)
        return
    }

    //  Get the session --

    session := Session_SessionFromToken(sessionToken)
    if session == nil {
        Log.Errorf("Invalid sessionToken '%s'.  Message type: %v.", sessionToken, messageType)
        SendError(writer, happiness.ResponseCode_RCNotAuthorized, error)
        return
    }
    userID := session.UserID
    Log.Debugf("------------------------------------------ UserID %s messageType %s.", userID, messageType)

    eventRequest := clientRequest.GetUserEventsRequest()
    if eventRequest != nil {
        UpdateUserEvents(writer, userID, eventRequest)
        return
    }
    profileUpdate := clientRequest.GetProfileUpdate()
    if profileUpdate != nil {
        UpdateProfiles(writer, userID, profileUpdate)
        return
    }
    profileQuery := clientRequest.GetProfileQuery()
    if profileQuery != nil {
        QueryProfiles(writer, userID, profileQuery)
        return
    }
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
    profileConfirmation := clientRequest.GetConfirmationRequest()
    if profileConfirmation != nil {
        UserConfirmation(writer, session, profileConfirmation)
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

    error = errors.New("Unrecognized request")
    Log.Errorf("Unknown client request: %+v.", clientRequest)
    SendError(writer, happiness.ResponseCode_RCInputCorrupt, error)
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


//----------------------------------------------------------------------------------------
//
//                                                                   Main & HappinesServer
//
//----------------------------------------------------------------------------------------


func HappinesServer() (returnValue int) {
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

    Log.LogLevel = Log.LevelAll
    commandLine := strings.Trim(fmt.Sprint(os.Args), "[]")

    //  Do config params --

    var (flagUsage bool; flagVersion bool; flagVerbose bool; flagPID bool; flagInputFilename string)

    flag.BoolVar(&flagUsage,   "h", false, "Help.  Print usage and exit.")
    flag.BoolVar(&flagUsage,   "?", false, "Help.  Print usage and exit.")
    flag.BoolVar(&flagVersion, "v", false, "Version.  Print version and exit.")
    flag.BoolVar(&flagVerbose, "V", false, "Verbose.  Verbose output.")
    flag.BoolVar(&flagPID,     "p", false, "PID filename.  Print the pid filename and exit.")
    flag.StringVar(&flagInputFilename, "c", "", "Configuration.  The file from which to read the configuration.")
    flag.Parse()

    if (flagUsage) {
        flag.Usage()
        return 0
    }
    if (flagVersion) {
        fmt.Fprintf(os.Stdout, "Version %s compiled %s.\n", globalVersion, globalCompileTime)
        return 0
    }
    if len(flagInputFilename) > 0 {
        flagInputFile, error := os.Open(flagInputFilename)
        if error != nil {
            Log.Errorf("Error: Can't open file '%s' for reading: %v.", flagInputFilename, error)
            return 1
        }
        defer flagInputFile.Close()
        error = config.ParseFile(flagInputFile)
        if error != nil {
            Log.Errorf("Error: %v.", error)
            return 1
        }
        //Log.Debugf("Parsed configuration file")
    }
    if flagVerbose {
        config.LogLevel = Log.LevelDebug
    }
    if flagPID {
        fmt.Fprintf(os.Stdout, "%s\n", config.PIDFileName())
        return 0
    }

    //  Start --

    Log.SetFilename(config.LogFilename);
    Log.Startf("Happiness-Server version %s pid %d compiled %s.", globalVersion, os.Getpid(), globalCompileTime)
     Log.Infof("Command line: %s.", commandLine)
    Log.Debugf("Configuration: %+v.", config)

    //  Lock our PID file --

    error := config.CreatePIDFile()
    if error != nil {
        Log.Errorf("%v", error)
        return 1
    }
    defer config.RemovePIDFile()

    //  Set our path --

    if error = os.Chdir(config.ServiceFilePath); error != nil {
        Log.Errorf("Error setting the home path '%s': %v.", config.ServiceFilePath, error)
        return 1
    } else {
        config.ServiceFilePath, _ = os.Getwd()
        Log.Debugf("Working directory: '%s'", config.ServiceFilePath)
    }

    //  Apply configuration paramaters --

    if error = config.ApplyConfiguration(); error != nil {
        Log.Errorf("Configuration error: %v", error)
        return 1
    }

    //  Start database --

    Log.Infof("Starting database %s.", config.DatabaseURI)
    error = config.ConnectDatabase()
    if error != nil { return 1 }

    _, error = config.DB.Exec("insert into MessageStatTable "+
       "  (timestamp, message) values ($1, 'Started');", time.Now());
    if error != nil {
        Log.Errorf("Error writing MessageStatTable: %v.", error)
    }

    //  Defer closing --

    defer func() {
        error := recover();
        if error != nil {
            message := fmt.Sprintf("%v", error)
            config.DB.Exec("insert into MessageStatTable "+
                "  (timestamp, message, responseMessage) values ($1, 'Fatal', $2);", time.Now(), message);
        }
        _, error = config.DB.Exec("insert into MessageStatTable "+
            "  (timestamp, message) values ($1, 'Terminated');", time.Now())
        if error != nil {
            Log.Errorf("Error writing MessageStatTable: %v.", error)
        }
        config.DisconnectDatabase()
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
    http.HandleFunc(config.ServicePrefix+"/send",  SendMessageGetMethod)
    http.HandleFunc(config.ServicePrefix+"/story", FetchStoryGetMethod)
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

    Log.Infof("Server listening at %d:%s.", config.ServicePort, config.ServicePrefix)
    http.Serve(httpListener, nil)
    return 0
}


func main () {
    exitStatus := HappinesServer()
    Log.Exitf("Exit status %d.", exitStatus)
    Log.FlushMessages();
    os.Exit(exitStatus)
}


