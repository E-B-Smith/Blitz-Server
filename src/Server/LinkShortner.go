//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "errors"
    "strings"
    "net/url"
    "net/http"
    "violent.blue/GoKit/Log"
)


//----------------------------------------------------------------------------------------
//                                                           LinkShortner_ShortLinkForLink
//----------------------------------------------------------------------------------------


func LinkShortner_ShortLinkFromLink(url string) (result string, error error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow("select InsertStringGivingShortHash($1);", url)
    var hash string
    error = row.Scan(&hash)
    if error != nil || hash == "" {
        Log.LogError(error)
        return "", error
    }
    return (config.ShortLinkURL + "/" + hash), nil
}



//----------------------------------------------------------------------------------------
//                                                          LinkShortner_LinkFromShortLink
//----------------------------------------------------------------------------------------


func LinkShortner_LinkFromShortLink(shortLink string) (url string, error error) {
    Log.LogFunctionName()

    var linkhash string
    parts := strings.Split(shortLink, "/")
    for i := len(parts)-1; i >= 0; i-- {
        linkhash = strings.TrimSpace(parts[i])
        if len(linkhash) > 3 { break }
    }
    if len(linkhash) <= 3 {
        return "", errors.New("Bad link")
    }
    row := config.DB.QueryRow(
        "select dataBlob from shortnertable where datahash = $1;", linkhash)
    error = row.Scan(&url)
    if error != nil {
        Log.LogError(error)
        return "", error
    } else {
        Log.Debugf("Expanded %s to %s.", linkhash, url)
        return url, error
    }
}



//----------------------------------------------------------------------------------------
//                                                                       RedirectShortLink
//----------------------------------------------------------------------------------------


func RedirectShortLink(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()
    Log.Debugf("Request: %+v.", httpRequest)

    //  Look up the short link and re-direct to long link --

    kLinkErrorRedirectURL := config.Localizef("kLinkErrorRedirectURL", "http://beinghappy.io/")

    if httpRequest.URL == nil {
        http.Redirect(writer, httpRequest, kLinkErrorRedirectURL, 301)
        return
    }
    Log.Debugf("Expanding %s.", httpRequest.URL.Path)

    urlstring, error := LinkShortner_LinkFromShortLink(httpRequest.URL.Path)
    Log.Debugf("Expanded %s to %s.", httpRequest.URL.Path, urlstring)
    if urlstring == "" || error != nil {
        http.Redirect(writer, httpRequest, kLinkErrorRedirectURL, 301)
        return
    }

    if strings.HasPrefix(urlstring, "http:") ||
       strings.HasPrefix(urlstring, "https:") ||
       strings.HasPrefix(urlstring, "mailto:") {
       http.Redirect(writer, httpRequest, urlstring, 301)
    } else {
        //  Else it's an app link:

        userid  := ""
        message := ""
        u, error := url.Parse(urlstring)
        if error == nil {
            message = u.Query().Get("message")
            userid  = u.Query().Get("userid")
        }
        messageCookie := http.Cookie {
            Name:   "message",
            Value:  message,
            Path:   "/",
            Domain: "bh.gy",
        }
        http.SetCookie(writer, &messageCookie)
        useridCookie := http.Cookie {
            Name:   "userid",
            Value:  userid,
            Path:   "/",
            Domain: "bh.gy",
        }
        http.SetCookie(writer, &useridCookie)
        urlCookie := http.Cookie {
            Name:   "appURL",
            Value:  urlstring,
            Path:   "/",
            Domain: "bh.gy",
        }
        http.SetCookie(writer, &urlCookie)
/*
        //  Un-comment for debugging --

        appNotInstalledCookie := http.Cookie {
            Name:   "appNotInstalled",
            Value:  "false",
            Path:   "/",
            Domain: "bh.gy",
        }
        http.SetCookie(writer, &appNotInstalledCookie)
*/
        http.Redirect(writer, httpRequest, config.AppLinkURL, 307)
    }
}



//----------------------------------------------------------------------------------------
//
//                                                                 LinkShortnerFormRequest
//
//----------------------------------------------------------------------------------------


func LinkShortnerFormRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    //  Send or process the short-link form --
    Log.LogFunctionName()
    Log.Debugf("Request: %+v.", httpRequest)

    defer func() {
        if error := recover(); error != nil {
            Log.LogStackWithError(error)
            http.Redirect(writer, httpRequest, "error.html", 303)
        }
    } ()

    fieldMap := make(map[string]string)
    fieldMap["AppName"]  = config.AppName
    fieldMap["longLink"] = ""
    fieldMap["shortLink"] = ""

    var error error
    if httpRequest.Method == "GET" {
        error = config.Template.ExecuteTemplate(writer, "ShortLink.html", fieldMap)
        if error != nil { panic(error) }
        return
    }

    error = httpRequest.ParseForm()
    if error != nil {
        fieldMap["longLinkError"] = "Bad form request."
    }

    longLink := strings.TrimSpace(httpRequest.PostFormValue("longLink"))
    shortLink := strings.TrimSpace(httpRequest.PostFormValue("shortLink"))
    fieldMap["longLink"] = longLink
    fieldMap["shortLink"] = shortLink
    if longLink == "" && shortLink == "" {
        fieldMap["longLinkError"] = "A long or short link is required."
        Log.Warningf("Submit error: %v.", fieldMap)
        config.Template.ExecuteTemplate(writer, "ShortLink.html", fieldMap)
        return
    }

    //  Get the links --

    if longLink != "" {
        shortLink, error = LinkShortner_ShortLinkFromLink(longLink)
        if shortLink == "" || error != nil {
            if error == nil { error = errors.New("No short link available.") }
            fieldMap["longLinkError"] = error.Error()
            Log.Warningf("Submit error: %v.", fieldMap)
            config.Template.ExecuteTemplate(writer, "ShortLink.html", fieldMap)
            return
        }
        fieldMap["shortLink"] = shortLink
    } else {
        longLink, error = LinkShortner_LinkFromShortLink(shortLink)
        if longLink == "" || error != nil {
            if error == nil { error = errors.New("No long link available.") }
            fieldMap["shortLinkError"] = error.Error()
            Log.Warningf("Submit error: %v.", fieldMap)
            config.Template.ExecuteTemplate(writer, "ShortLink.html", fieldMap)
            return
        }
        fieldMap["longLink"] = longLink
    }

    //  Send back the result --

    config.Template.ExecuteTemplate(writer, "ShortLink.html", fieldMap)
}


