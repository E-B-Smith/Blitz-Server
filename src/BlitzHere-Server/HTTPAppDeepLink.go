

//----------------------------------------------------------------------------------------
//
//                                                   BlitzHere-Server : HTTPAppDeepLink.go
//                                               Save an app deep-link for later execution
//
//                                                                   E.B. Smith, May, 2015
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "time"
    "bytes"
    "regexp"
    "strconv"
    "strings"
    "net/url"
    "net/http"
    "html/template"
    "encoding/json"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)



//----------------------------------------------------------------------------------------
//                                                                     SignatureFromDevice
//----------------------------------------------------------------------------------------


type DeviceInfo struct {
    Width           float64
    Height          float64
    ColorDepth      float64
    TimeZone        float64
    Scale           float64
    RPM             float64
    IPAddress       string
    Language        string
    Model           string
    BuildVersion    string
}


func SignatureFromDeviceInfo(deviceInfo DeviceInfo) string {
    language := "en"
    if deviceInfo.Language != "" {
        language = strings.ToLower(Util.FirstNRunes(deviceInfo.Language, 2))
    }
    model := ModelNameFromString(deviceInfo.Model)

    return fmt.Sprintf("%d-%d-%d-%d-%05.2f-%s-%s-%s-%s",
        int(deviceInfo.Width),
        int(deviceInfo.Height),
        int(deviceInfo.ColorDepth),
        int(deviceInfo.TimeZone),
        deviceInfo.Scale,
        deviceInfo.IPAddress,
        language,
        model,
        deviceInfo.BuildVersion)
}


func ModelNameFromString(s string) string {
    switch {
        case strings.Contains(s, "iPod"):
            return "iPod"
        case strings.Contains(s, "iPad"):
            return "iPod"
        case strings.Contains(s, "iPhone"):
            return "iPhone"
    }
    return ""
}



//----------------------------------------------------------------------------------------
//
//                                                                         HTTPAppDeepLink
//
//----------------------------------------------------------------------------------------


func HTTPAppDeepLink(writer http.ResponseWriter, httpRequest *http.Request) {

    //  Save the request / device signature and the deeplink for later retrieval.
    //  Saves:
    //      IP Address
    //      Language
    //      Screen Size & Scale
    //      OS Version
    //      Device Type / model (iPhone / iPad / iPod) (ModelName?)
    //      timezone (?)
    //      rpm
    //      Time  -  For expiration
    //      mark as used ?

    //  Two cases:
    //      Is iPhone:  Save data and send back page with deep link.
    //      Otherwise:  Send back web page.

    //  User-Agent:
    //  [Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1]

    defer func() {
        error := recover();
        if error != nil {
            Log.LogStackWithError(error)
        }
        Log.Debugf("Exit HTTPAppDeepLink.")
    } ()

    var error error
    Log.LogFunctionName()
    Log.Debugf("Request:\n%+v\n", httpRequest)

    userAgent := httpRequest.Header.Get("User-Agent")
    model := ModelNameFromString(userAgent)
    referrer := httpRequest.Referer()
    Log.Debugf("Referrer: %s.", referrer)

    deepLinkURL, error := url.Parse(config.AppLinkURL)
    if error != nil { Log.LogError(error) }
    deepLinkPath := deepLinkURL.Path+"/"
    deepLinkURL.RawQuery = httpRequest.URL.RawQuery
    deepLinkURL.Scheme = config.AppLinkScheme
    deepLinkURLString := deepLinkURL.String()
    Log.Debugf("Deeplink URL: %s.", deepLinkURLString)

    senderID := deepLinkURL.Query().Get("senderid")
    userID   := deepLinkURL.Query().Get("userid")
    message  := deepLinkURL.Query().Get("message")

    //  Get the template map ready --

    type TemplateMap struct {
        AppName         string
        AppDeepLink     template.URL
        AutoOpenDeepLink bool
        AppStoreLink    template.URL
        Message         string
    }

    templateMap := TemplateMap {
        AppName:        config.AppName,
        Message:        message,
        AppDeepLink:    template.URL(deepLinkURLString),
        AppStoreLink:   template.URL(config.AppStoreURL),
        AutoOpenDeepLink: false,
    }

    if httpRequest.Method == "GET" {
        Log.Debugf("Path: %s", httpRequest.URL.Path)
        if httpRequest.URL.Path == deepLinkPath {

            if len(model) == 0 {
                //  Not iPhone.  Send a page with the iPhone method.
                writer.Header().Set("Referer", referrer)
                error = config.Template.ExecuteTemplate(writer, "HTTPDeepLink.html", templateMap)
                if error != nil { Log.LogError(error) }
                return
            } else {
                writer.Header().Set("Referer", referrer)
                error = config.Template.ExecuteTemplate(writer, "Sniffer.html", templateMap)
                if error != nil { Log.LogError(error) }
                return
            }

        } else {
            Log.Debugf("urlPath: %s.", deepLinkPath)
            http.ServeFile(writer, httpRequest, strings.TrimPrefix(httpRequest.URL.Path, deepLinkPath))
            return
        }
    }

    if httpRequest.Method != "POST" {
        http.Error(writer, "Invalid Method", 405)
        return
    }

    error = httpRequest.ParseForm()
    if error != nil { Log.LogError(error) }
    //Log.Debugf("Form: %+v\n.", httpRequest.Form)

    deviceJSON := httpRequest.PostFormValue("DeviceInfo")
    var deviceInfo DeviceInfo
    error = json.Unmarshal([]byte(deviceJSON), &deviceInfo)
    if error != nil { Log.LogError(error) }

    var rexp *regexp.Regexp
    rexp, error = regexp.Compile("Mobile\\/(.*?)\\s|$")
    if error != nil {
        Log.Errorf("Regex error: %v.", error)
    }
    buildArray := rexp.FindStringSubmatch(userAgent)
    deviceInfo.BuildVersion = ""
    if len(buildArray) > 0 {
        deviceInfo.BuildVersion = buildArray[len(buildArray)-1]
    }
    deviceInfo.IPAddress = httpRequest.Header.Get("X-Real-Ip")
    deviceInfo.Language  = httpRequest.Header.Get("Accept-Language")
    deviceInfo.Model     = model
    deviceInfo.TimeZone  *= -60.0
    Log.Debugf("DeviceInfo:\n%+v\n.", deviceInfo)

    //  Save the signature & deeplink for later --

    invite := BlitzMessage.UserInvite {
        UserID:     StringPtrFromString(senderID),
        FriendID:   StringPtrFromString(userID),
        Message:    StringPtrFromString(message),
    }
    Log.Debugf("Invite: %s.", invite.String())

    var inviteData []byte
    inviteData, error = proto.Marshal(&invite)
    if error != nil { Log.LogError(error) }

    signature := SignatureFromDeviceInfo(deviceInfo)
    _, error = config.DB.Exec(
        `insert into HTTPDeepLinkTable
            (deviceSignature, deviceRPM, creationDate, inviteData, referrer) values
            ($1, $2, $3, $4, $5);`,
            signature, deviceInfo.RPM, time.Now(), inviteData, referrer)
    if error != nil { Log.LogError(error) }

    //  Send back the web page --

    var buffer bytes.Buffer
    templateMap.AutoOpenDeepLink = true
    error = config.Template.ExecuteTemplate(&buffer, "HTTPDeepLink.html", templateMap)
    if error != nil { Log.LogError(error) }

    //Log.Debugf("Doc:\n%s\n.", buffer.String())
    writer.Write(buffer.Bytes())
}



//----------------------------------------------------------------------------------------
//
//                                                                  InviteRequestForDevice
//
//----------------------------------------------------------------------------------------


func InviteRequestForDevice(device *BlitzMessage.DeviceInfo) *BlitzMessage.UserInvite {
    Log.LogFunctionName()
    Log.Debugf("Getting invite for: %+v.", device)

    timezone, _ := strconv.ParseFloat(*device.Timezone, 64)
    deviceInfo := DeviceInfo{
        Width:          *device.ScreenSize.Width,
        Height:         *device.ScreenSize.Height,
        ColorDepth:     float64(*device.ColorDepth),
        TimeZone:       timezone,
        Scale:          float64(*device.ScreenScale),
        IPAddress:      *device.IPAddress,
        Language:       *device.Language,
        Model:          *device.ModelName,
        BuildVersion:   *device.SystemBuildVersion,
    }

/*
    create table HTTPDeepLinkTable
        (
         deviceSignature    text        not null
        ,deviceRPM          real        not null
        ,creationDate       timestamptz not null
        ,claimDate          timestamptz
        ,deepLink           text        not null
        );
*/

    signature := SignatureFromDeviceInfo(deviceInfo)
    rows, error := config.DB.Query(
        `select inviteData, creationDate, deviceRPM from HTTPDeepLinkTable
            where deviceSignature = $1
              and claimDate is null
            order by creationDate, deviceRPM desc;`,
            signature)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return nil
    }

    for rows.Next() {
        var (
            inviteData      []byte
            creationDate    time.Time
            deviceRPM       float64
        )
        error = rows.Scan(&inviteData, &creationDate, &deviceRPM)
        if error != nil { Log.LogError(error); continue; }

        invite := BlitzMessage.UserInvite {}
        error = proto.Unmarshal(inviteData, &invite)
        if error != nil { Log.LogError(error); continue; }
        if invite.FriendID == nil { continue; }

        Log.Debugf("Found: %s.", invite.String())

        result, error := config.DB.Exec(
            `update HTTPDeepLinkTable set claimDate = $1, deviceType = $2
                where deviceSignature = $3 and creationDate = $4;`,
            time.Now(), device.ModelName, signature, creationDate)

        if error != nil || pgsql.RowsUpdated(result) != 1 {
            Log.Errorf("Didn't HTTPDeepLinkTable. Error: %v.", error)
        }

        if invite.FriendID == nil || invite.UserID == nil {
            return nil
        } else {
            return &invite
        }
    }
    return nil
}

