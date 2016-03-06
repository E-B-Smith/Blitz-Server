//  DownloadLog.go  -  Reads the HTTP logs for app downloads and updates the DB.
//
//  E.B.Smith  -  June, 2015


package main


import (
    "fmt"
    "time"
    "regexp"
    "strings"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
)


func RegexFind(regexString string, targetString string) string {
    r, error := regexp.Compile(regexString)
    if error != nil {
        Log.Errorf("Regex error: %v.", error)
        return ""
    }
    return r.FindString(targetString)
}


func RegexFindSubstring(regexString string, targetString string) string {
    r, error := regexp.Compile(regexString)
    if error != nil {
        Log.Errorf("Regex error: %v.", error)
        return ""
    }
    result := r.FindStringSubmatch(targetString)
    if len(result) > 0 {
        return result[len(result)-1]
    } else {
        return ""
    }
}


func RegexFindSubstringArray(regexString string, targetString string) []string {
    r, error := regexp.Compile(regexString)
    if error != nil {
        Log.Errorf("Regex error: %v.", error)
        return nil
    }
    result := r.FindStringSubmatch(targetString)
    return result
}


func UpdateDownloadTableWithLogLine(line string) {
    //  Updates the database with the log line:
    Log.LogFunctionName()

    //  Example line:
    //  Apache:
    //  "107.3.151.67 - - [25/May/2015:11:03:16 -0700] \"GET /beinghappy/ios-labs/HappyLabs.ipa HTTP/1.1\" 200 29874597"
    //
    //  nginx
    //  "107.3.151.67 - - [12/Jul/2015:18:22:49 -0700] "GET /ios-labs/HappyLabs.ipa HTTP/1.1" 200 31821220 "-" "itunesstored/1.0 iOS/8.2 model/iPhone4,1 build/12D508 (6; dt:73)"

    if len(line) == 0 {
        return
    }

    ipaddress  := strings.TrimSpace(RegexFind("^([^\\s]*)", line))
    datestring := RegexFindSubstring("\\Q[\\E(.*?)\\Q]\\E", line)
    date, _    := time.Parse("2/Jan/2006:15:04:05 -0700", datestring)
    filename   := strings.TrimSpace(RegexFindSubstring("GET (.*?) HTTP", line))
    httpcode   := RegexFindSubstring("\"\\s([[:digit:]]+)\\s[[:digit:]]+", line)
    bytes      := RegexFindSubstring("\"\\s[[:digit:]]+\\s([[:digit:]]+)", line)

    //Log.Debugf("\n%s\n%s\n%s\n%s\n%s", ipaddress, datestring, filename, httpcode, bytes)

    _, error := config.DB.Exec(
        `insert into AppDownloadTable(
             timestamp
            ,IPAddress
            ,filename
            ,httpCode
            ,totalBytes
            ) values ($1, $2, $3, $4, $5);
        `, date, ipaddress, filename, httpcode, bytes)
    if error != nil {
        if error, ok := error.(*pq.Error); ok {
            if error.Code != "23505" {
                Log.Errorf("Error: %v.", error)
                Log.Errorf(" Line: %s.", line)
            }
        }
    }
}


func UpdateDownloadTableWithLogLineCount(linecount int) {
    //  Just update with the last linecount log lines.  if < 0 read whole Log.
    Log.LogFunctionName()

    logfilename := config.WebLog
    shellparameter := fmt.Sprintf("tail -n %d %q | grep 'GET .*\\.ipa'", linecount, logfilename)

    if linecount == 0 {
        return
    } else if linecount < 0 {
        shellparameter = fmt.Sprintf("grep 'GET .*\\.ipa' %q", logfilename)
    }

    Log.Debugf("Comand: %v.", shellparameter)
    lines, errorlines, error := Util.RunShellCommand("/bin/bash", []string{"-c", shellparameter}, nil)
    if error != nil || len(errorlines) > 0 {
        Log.Errorf("Run shell returned with error %v: %s.", error, string(errorlines))
    }

    lineStrings := strings.Split(string(lines),"\n")
    Log.Debugf("Found %d log lines.", len(lineStrings))
    for _, line := range(lineStrings) {
        UpdateDownloadTableWithLogLine(line)
    }
}


