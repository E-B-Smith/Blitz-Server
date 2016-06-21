//  signup  -  Signup for an iPhone beta test - Saves the device udid.
//
//  E.B.Smith  -  March, 2015


package main


import (
    "os"
    "fmt"
    "net"
    "flag"
    "time"
    "errors"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
    "html/template"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/ServerUtil"
    "github.com/DHowett/go-plist"
    )


var   SignupFormFilename    = "form.html"
var   ProfileFormFilename   = "profile.html"
const kSMSNotificationNumber = "4156152570"
var   config ServerUtil.Configuration


//  Signup flow:
//
//  signup/form.html -> signup/profile.html -> <profile> -> signup/enroll -> signup/complete.html
//      |                      |                                 |
//      +----------------------+---------------------------------+---------> signup/error.html


var ProfileStringSkeleton =
`
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>PayloadContent</key>
        <dict>
            <key>URL</key>
            <string>%s</string>
            <key>Challenge</key>
            <string>%s</string>
            <key>DeviceAttributes</key>
            <array>
                <string>DEVICE_NAME</string>
                <string>IMEI</string>
                <string>ICCID</string>
                <string>MAC_ADDRESS_EN0</string>
                <string>MEID</string>
                <string>PRODUCT</string>
                <string>SERIAL</string>
                <string>UDID</string>
                <string>VERSION</string>
            </array>
        </dict>
        <key>PayloadOrganization</key>
        <string>BeingHappy</string>
        <key>PayloadDisplayName</key>
        <string>BeingHappy Enrollment</string>
        <key>PayloadVersion</key>
        <integer>1</integer>
        <key>PayloadUUID</key>
        <string>B5893090-A663-434F-A0E1-7CFA29F8E11F</string>
        <key>PayloadIdentifier</key>
        <string>blue.violent.happiness-enrollment</string>
        <key>PayloadDescription</key>
        <string>This is a temporary profile to enroll your device for BeingHappy ad-hoc app distribution.</string>
        <key>PayloadType</key>
        <string>Profile Service</string>
    </dict>
</plist>
`


//----------------------------------------------------------------------------------------
//                                                                    RenderTemplateString
//----------------------------------------------------------------------------------------


func RenderTemplateString(writer http.ResponseWriter, templatestring string, data interface{}) {
    var error error
    templ := template.New("template")
    templ, error = templ.Parse(templatestring)
    if  error != nil {
        http.Error(writer, error.Error(), http.StatusInternalServerError)
        return
    }
    error = templ.Execute(writer, data)
    if error != nil {
        http.Error(writer, error.Error(), http.StatusInternalServerError)
    }
}


//----------------------------------------------------------------------------------------
//                                                                       SignupFormRequest
//----------------------------------------------------------------------------------------


var SignUpFormString string = ""
var ProfileFormString string = ""


func SignupFormRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Send or process the sign-up form --
    Log.LogFunctionName()
    config.MessageCount++

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, httpRequest, "error.html", 303)
        }
    } ()

    var error error
    if SignUpFormString == "" {
        b, error := ioutil.ReadFile(SignupFormFilename)
        if error != nil {
            Log.Errorf("Can't read %s: %v.", SignupFormFilename, error)
            panic(error)
        }
        SignUpFormString = string(b)
    }
    if httpRequest.Method == "GET" {
        RenderTemplateString(writer, SignUpFormString, nil)
        return
    }

    errorMap := make(map[string]string)
    error = httpRequest.ParseForm()
    if error != nil {
        errorMap["name"] = "Bad form request."
    }

    name := strings.TrimSpace(httpRequest.PostFormValue("name"))
    if name == "" {
        errorMap["name"] = "A name is required."
    }
    email, error := Util.ValidatedEmailAddress(httpRequest.PostFormValue("email"))
    if error != nil {
        errorMap["email"] = "Please enter a valid email address."
    }
    phone, error := Util.ValidatedPhoneNumber(httpRequest.PostFormValue("phone"))
    if error != nil {
        errorMap["phone"] = "Please enter a valid 10-digit phone number."
    }

    if len(errorMap) > 0 {
        Log.Warningf("Submit error: %v.", errorMap)
        RenderTemplateString(writer, SignUpFormString, errorMap)
        return
    }

    tempID := Util.NewUUIDString()
    Log.Debugf("Validated '%s' '%s' '%s' '%s'.", name, email, phone, tempID)

    //  Insert row --

    result, error := config.DB.Exec(
        "insert into deviceudidtable "+
        "  (name, phone, email, tempid, modificationdate) values "+
        "  ($1, $2, $3, $4, $5);",
        name, phone, email, tempID, time.Now())

    var rowCount int64 = 0
    if result != nil {  rowCount, _ = result.RowsAffected() }

    if rowCount == 0 {
        result, error = config.DB.Exec(
            "update deviceudidtable set "+
            "  (name, phone, email, tempid, modificationdate) = "+
            "  ($1, $2, $3, $4, $5);",
            name, phone, email, tempID, time.Now())

        if result != nil {  rowCount, _ = result.RowsAffected() }
        if rowCount == 0 || error != nil { panic(error) }
    }

    error = Util.SendSMS(kSMSNotificationNumber, name + " started the beta test signup.")
    if error != nil { Log.LogError(error) }

    //  Send back the profile page --

    if ProfileFormString == "" {
        b, error := ioutil.ReadFile(ProfileFormFilename)
        if error != nil {
            Log.Errorf("Can't read %s: %v.", ProfileFormFilename, error)
            panic(error)
        }
        ProfileFormString = string(b)
    }

    Log.Debugf("Sending the profile info page...")

    profileMap := make(map[string]string)
    profileMap["tempid"] = tempID
    RenderTemplateString(writer, ProfileFormString, profileMap)
}


//----------------------------------------------------------------------------------------
//                                                                SignupSendProfileRequest
//----------------------------------------------------------------------------------------


func SignupSendProfileRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Send back a signed profile --
    Log.LogFunctionName()
    config.MessageCount++

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, httpRequest, "error.html", 303)
        }
    } ()

    var error error
    error = httpRequest.ParseForm()
    if error != nil { panic(error) }
    tempID := strings.TrimSpace(httpRequest.PostFormValue("tempid"))
    Log.Debugf("Tempid: %s.", tempID)
    if tempID == "" { panic(errors.New("No tempID in form")) }

    /*
    openssl smime -sign -in signup.mobileconfig.saved    \
        -out signup.mobileconfig    \
        -signer /etc/apache2/certs/violent.blue.crt     \
        -inkey /etc/apache2/certs/violent.blue.key      \
        -certfile /etc/apache2/certs/CertificatesChain.crt      \
        -outform der -nodetach
    */

    enrollURL := config.ServiceURL() + "/enroll"
    profileString := fmt.Sprintf(ProfileStringSkeleton, enrollURL, tempID)
    profileBytes := []byte(profileString)
    commandString :=
`        openssl smime -sign
            -signer /etc/apache2/certs/violent.blue.crt
            -inkey /etc/apache2/certs/violent.blue.key
            -certfile /etc/apache2/certs/CertificatesChain.crt
            -outform der -nodetach
`
    if strings.HasPrefix(config.ServerURL, "https://bh.gy") {
        commandString =
`        openssl smime -sign
            -signer   /etc/keys/bh.gy.crt
            -inkey    /etc/keys/bh.gy.pem
            -certfile /etc/keys/chain.crt
            -outform der -nodetach
`
}

    var signedProfile []byte
    commandString = Util.StringExcludingCharactersInSet(commandString, "\n\r\t")
    t := strings.Split(commandString, " ")
    commandArray := make([]string, 0, 15)
    for index := range(t) {
        if t[index] != "" { commandArray = append(commandArray, t[index]) }
    }
    Log.Debugf("Command (%d): %v.", len(commandArray), commandArray)
    var stdErrorString []byte
    signedProfile, stdErrorString, error = Util.RunShellCommand(commandArray[0], commandArray[1:], profileBytes)
    if error != nil {
        Log.Errorf("Error encoding profile: %s.", string(stdErrorString))
        panic(error)
    }

    writer.Header().Add("Content-Type", "application/octet-stream")
    writer.Header().Add("Content-Disposition", "attachment; filename=\"enroll.mobileconfig\"")
    writer.Write(signedProfile)
    Log.Debugf("Sent profile.")
}


//----------------------------------------------------------------------------------------
//                                                                     SignupEnrollRequest
//----------------------------------------------------------------------------------------


func SignupEnrollRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Save the data to a file --
    Log.LogFunctionName()
    config.MessageCount++

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, httpRequest, "error.html", 303)
        }
    } ()


    Log.Debugf("Request: %v.", *httpRequest)
    Log.Debugf(" Header: %v.", httpRequest.Header)

    defer httpRequest.Body.Close()
    body, error := ioutil.ReadAll(httpRequest.Body)
    if error != nil { panic(error) }

    /*
    openssl smime -verify -noverify -inform DER -in SavedData.txt
    */

    commandString := "openssl smime -verify -noverify -inform DER"

    var plistdata []byte
    commandString = Util.StringExcludingCharactersInSet(commandString, "\n\r\t")
    t := strings.Split(commandString, " ")
    commandArray := make([]string, 0, 15)
    for index := range(t) {
        if t[index] != "" { commandArray = append(commandArray, t[index]) }
    }
    plistdata, _, error = Util.RunShellCommand(commandArray[0], commandArray[1:], body)
    if error != nil { panic(error) }
    Log.Debugf("plist data: %s.", plistdata)

    //  Save to the database --

    type Profile struct {
        CHALLENGE       string
        DEVICE_NAME     string
        IMEI            string
        ICCID           string
        MAC_ADDRESS_EN0 string
        MEID            string
        PRODUCT         string
        SERIAL          string
        UDID            string
        VERSION         string
    }
    var profile Profile
    _, error = plist.Unmarshal(plistdata, &profile)
    if error != nil { panic(error) }

    //  If deviceUDID starts with FFFF then it's invalid.  Redirect to error.

    if profile.UDID == "" { panic(error) }
    udid := strings.ToLower(profile.UDID)
    if strings.HasPrefix(udid, "ffff") { panic(error) }

    tempID := strings.TrimSpace(profile.CHALLENGE)
    if tempID == "" { panic(errors.New("No tempID in profile")) }

    ipaddress := Util.IPAddressFromHTTPRequest(httpRequest)

    result, error := config.DB.Exec(
        "update deviceudidtable set ("+
            "deviceName, "+
            "imei, "+
            "iccid, "+
            "macAddress, "+
            "meid, "+
            "product, "+
            "serial, "+
            "deviceUDID, "+
            "version, "+
            "modificationdate, "+
            "notes, "+
            "IPAddress) = "+
            "($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"+
            "  where tempid = $13;",
            profile.DEVICE_NAME,
            profile.IMEI,
            profile.ICCID,
            profile.MAC_ADDRESS_EN0,
            profile.MEID,
            profile.PRODUCT,
            profile.SERIAL,
            udid,
            profile.VERSION,
            time.Now(),
            plistdata,
            ipaddress,
            tempID)
    if error != nil { panic(error) }

    var rowCount int64 = 0
    if result != nil {  rowCount, _ = result.RowsAffected() }
    if rowCount == 0 {
        Log.Errorf("No rows updated.")
        panic(errors.New("No rows updated"))
    }
    Log.Infof("Device record saved.")

    //  Get the name --

    name := "Unknown"
    rows, error := config.DB.Query("select name from deviceudidtable where tempid = $1;", tempID)
    if error != nil {
        Log.LogError(error)
    } else {
        defer rows.Close()
        for rows.Next() {
            rows.Scan(&name)
        }
    }

    //  Send a text when signed up --

    error = Util.SendSMS(kSMSNotificationNumber, name + " signed up to beta test.")
    if error != nil { Log.LogError(error) }

    //  Send back a completion message --

    http.Redirect(writer, httpRequest, "complete", 303)
}


//----------------------------------------------------------------------------------------
//                                                                          HelperHandlers
//----------------------------------------------------------------------------------------


func RedirectToFormRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    config.MessageCount++
    http.Redirect(writer, httpRequest, "signup/form", 303)
}


func ShowRequest(writer http.ResponseWriter, request *http.Request) {
    config.MessageCount++
    Log.Debugf("Request:\n%v\nServer File Path:\n%s", request, config.ServiceFilePath)
    fmt.Fprintf(writer,
        "<html><p>Hi!\n<br>\n<br>Request:\n<br>%v\n<br>\n<br>File Path:  %s\n<br></p></html>",
        *request, config.ServiceFilePath)
}


func SendHello(writer http.ResponseWriter, request *http.Request) {
    config.MessageCount++
    Log.Debugf("Request:\n%v\n", *request)
    fmt.Fprintf(writer, "<html><p>Hello!\n\n<br><br>Request:\n<br>%v\n<br></p></html>", *request)
}


//----------------------------------------------------------------------------------------
//                                                                           Signup-Server
//----------------------------------------------------------------------------------------


func SignupServer() int {
    Log.LogLevel = Log.LogLevelAll
    commandLine := strings.Trim(fmt.Sprint(os.Args), "[]")

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
        fmt.Fprintf(os.Stdout, "Version %s.\n", Util.CompileVersion())
        return 0
    }

    if len(flagInputFilename) > 0 {
        flagInputFile, error := os.Open(flagInputFilename)
        if error != nil {
            Log.Errorf("Error: Can't open file '%s' for reading: %v.", flagInputFilename, error)
            return 1
        }
        defer flagInputFile.Close()
        error = ServerUtil.ParseConfigFile(&config, flagInputFile)
        if error != nil {
            Log.Errorf("Error: %v.", error)
            return 1
        }
        //Log.Debugf("Parsed configuration: %v.", *globalConfiguration)
    }
    if flagPID {
        fmt.Fprintf(os.Stdout, "%s\n", config.PIDFileName())
        return 0
    }
    if flagVerbose {
        config.LogLevel = Log.LogLevelDebug
    }

    Log.SetFilename(config.LogFilename);
    Log.Startf("Signup version %s pid %d.", Util.CompileVersion(), os.Getpid())
    Log.Infof("Command line: %s.",  commandLine)
    Log.Debugf("Configuration: %v.", config)

    //  Lock our PID file --

    error := config.CreatePIDFile()
    if error != nil {
        Log.Errorf("%v", error)
        return 1
    }
    defer config.RemovePIDFile()

    //  Set our working path --

    if error = os.Chdir(config.ServiceFilePath); error != nil {
        Log.Errorf("Error setting the home path '%s': %v.", config.ServiceFilePath, error)
        return 1
    } else {
        config.ServiceFilePath, _ = os.Getwd()
        Log.Debugf("Working directory: '%s'", config.ServiceFilePath)
    }

    //  Connect to database --

    error = config.ConnectDatabase()
    if error != nil { return 1 }
    defer config.DisconnectDatabase();

    //  Make our listener --

    httpListener, error := net.Listen("tcp", ":"+strconv.Itoa(config.ServicePort))
    if error != nil {
        Log.Errorf("Can't listen on port %d: %v.", config.ServicePort, error)
        return 1
    }

    //  Set up an interrupt handler --

    config.AttachToInterrupts(httpListener)

    //  Set up & start our http handlers --

    http.HandleFunc(config.ServicePrefix, RedirectToFormRequest)
    http.HandleFunc(config.ServicePrefix+"/form", SignupFormRequest)
    http.HandleFunc(config.ServicePrefix+"/profile", SignupSendProfileRequest)
    http.HandleFunc(config.ServicePrefix+"/enroll", SignupEnrollRequest)
    http.HandleFunc(config.ServicePrefix+"/hello", SendHello)

//  http.HandleFunc("/", ShowRequest)
    http.Handle("/",
        http.StripPrefix(config.ServicePrefix,
        http.FileServer(http.Dir(config.ServiceFilePath))))

    Log.Infof("Listening on %d:%s.", config.ServicePort, config.ServicePrefix)
    http.Serve(httpListener, nil)
    Log.Exitf("EOJ")
    return 0
}


func main() {
    os.Exit(SignupServer())
}


