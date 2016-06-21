

//----------------------------------------------------------------------------------------
//
//                                                       BlitzHere-Server : DownloadApp.go
//                                             Send a text message with the download link.
//
//                                                                 E.B. Smith, March, 2015
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "net/http"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
)


func DownloadAppRequest(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    //  Save the name and phone number for later.
    //  Text a short-link that finger prints phone and redirects to the app store.
    //  Otherwise send back error.

    if request.Method == "GET" {
        if len(kSignupFormSkeleton) <= 0 {
            bytes, error := ioutil.ReadFile(serverfilespath+"/form.html")
            if error != nil {
                Log.LogError(error)
                http.Redirect(writer, request, "error.html", 303)
                return
            }
            kSignupFormSkeleton = string(bytes)
        }
        RenderTemplateString(writer, kSignupFormSkeleton, nil)
        return
    }

    errorMap := make(map[string]string)

    error := request.ParseForm()
    if error != nil {
        errorMap["name"] = "Bad form request."
    }

    name := strings.TrimSpace(request.PostFormValue("name"))
    if name == "" {
        errorMap["name"] = "A name is required."
    }
    phone := ServerUtil.StringIncludingCharactersInSet(request.PostFormValue("phone"), "0123456789")
    if len(phone) != 10 {
        errorMap["phone"] = "Please enter a 10-digit phone number."
    }

    if len(errorMap) > 0 {
        Log.Warningf("Submit error: %v.", errorMap)
        RenderTemplateString(writer, kSignupFormSkeleton, errorMap)
        return
    }

    Log.Debugf("Validated '%s' '%s'.", name, phone)

    file, error := os.OpenFile(signupfilename, os.O_APPEND|os.O_WRONLY, 0600)
    if error != nil {
        Log.LogError(error)
        http.Redirect(writer, request, "error.html", 303)
        return
    }
    defer file.Close()

    if _, error = fmt.Fprintf(file, "%s\t%s\t%s\n", name, email, phone); error != nil {
        Log.LogError(error)
        http.Redirect(writer, request, "error.html", 303)
        return
    }

    http.Redirect(writer, request, "complete.html", 303)
}


