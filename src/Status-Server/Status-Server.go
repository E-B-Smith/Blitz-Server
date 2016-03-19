//  Status-Server  -  Serves an html status page for the BeingHappy-Server and HappyLabs-Server
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
    "path/filepath"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/ServerUtil"
)


var   config ServerUtil.Configuration =
            ServerUtil.Configuration {
                    SoftwareVersion:    "0.1.50",
                    ServiceName:        "Stats-Server",
                    ServicePort:        9999,
                    ServiceFilePath:    "./stats",
                    ServicePrefix:      "/beinghappy/status",
                    ServerURL:          "https://violent.blue",
                    DatabaseURI:        "postgres://happylabs:happylabs@localhost:5432/happylabs",
                    LogLevel:           Log.LogLevelAll,
                    LogFilename:        "",
            }

//  Our run-time flags --

var (
    flagUsage           bool
    flagVersion         bool
    flagVerbose         bool
    flagPID             bool
    flagInputFilename   string
)

var servers = []*ServerInfo {
    &ServerInfo{ ConfigFilename: "BeingHappyLegacy-ServerStatus.config" },
    &ServerInfo{ ConfigFilename: "BeingHappy-Server.config" },
    &ServerInfo{ ConfigFilename: "HappyLabs-Server.config" },
    &ServerInfo{ ConfigFilename: "HappyPulse-Server.config" },
    &ServerInfo{ ConfigFilename: "PulseLabs-Server.config" },
    &ServerInfo{ ConfigFilename: "Signup-Server.config" },
    &ServerInfo{ ConfigFilename: "Status-Server.config" },
}
var labsServerInfo = servers[2]



//----------------------------------------------------------------------------------------
//                                                                       SendRefreshStatus
//----------------------------------------------------------------------------------------


func SendRefreshStatus(writer http.ResponseWriter, request *http.Request) {
    //  Execute the status sql and send the result --
    Log.LogFunctionName()
    config.MessageCount++

    startTimestamp := time.Now()
    defer func() {
        error := recover();
        if error != nil {
            Log.LogStackWithError(error)
        }
        Log.Debugf("Exit SendRefreshStatus.  Elapsed: %1.4f.  Message timestamp: %v.",
            time.Since(startTimestamp).Seconds(), startTimestamp)
    } ()

    //  Server status --

    writer.Write([]byte(kWebPageHeader))
    writer.Write([]byte(kWebPageStatusHeaderBlue))

    writer.Write([]byte("<div><h3>Servers</h3><pre><code>"))
    for _, server := range servers {
        writer.Write([]byte(server.GetStatus()))
    }
    writer.Write([]byte("</code></pre></div><br>"))
    writer.Write([]byte(kWebPageStatusTrailer))

    //  Download status --

    go UpdateDownloadTableWithLogLineCount(5000)

    kDownloadsScript :=
        `\echo <div><h3>Downloads</h3><pre><code>
        select timestamp, UserNameFromIPAddress(ipaddress), filename, httpcode, totalbytes
              from appdownloadtable order by timestamp desc limit 10;
        \pset tuples_only on
        select 'Total', count(*) from appdownloadtable;
        \echo </code></pre></div>`

    var (
        error error
        statusOutput []byte
        errorOutput  []byte
    )

    if ! labsServerInfo.Config.DatabaseIsConnected() {
        labsServerInfo.Config.ConnectDatabase()
    }
    if labsServerInfo.Config.PGSQL == nil {
        error = errors.New("No configuration for HappyLabs")
    } else {
        statusOutput, errorOutput, error = labsServerInfo.Config.PGSQL.RunSQLScript(kDownloadsScript)
    }
    if error != nil || len(errorOutput) > 0 {

        writer.Write([]byte(kWebPageStatusHeaderBlue))
        writer.Write([]byte("\n<div><h3>Downloads Status Error</h3><pre><code>\n"))
        if error != nil { fmt.Fprintf(writer, "%v.\n\n", error) }
        writer.Write(errorOutput)
        writer.Write([]byte("\n\n</code></pre></div>\n"))
        writer.Write([]byte(kWebPageStatusTrailer))

    } else {

        writer.Write([]byte(kWebPageStatusHeaderBlue))
        writer.Write(statusOutput)
        writer.Write([]byte(kWebPageStatusTrailer))

    }
    //  Push notification status --
    //  < To-Do >

    //  New service request --

    var servername string
    if servernames, ok := request.URL.Query()["server"]; ok {
        if len(servernames) > 0 { servername = servernames[0] }
    }

    var currentServer *ServerInfo = nil
    for _, server := range servers {
        if strings.ToLower(server.Config.ServiceName) ==  strings.ToLower(servername) {
            currentServer = server
        }
    }
    if currentServer == nil {
        currentServer = servers[0]
    }

    writer.Write([]byte("<select onchange=\"selectServerChange()\" id=\"selectServer\" >"))
    for _, server := range servers {
        selected := ""
        if server == currentServer {
            selected = "selected='selected'"
        }
        writer.Write([]byte(fmt.Sprintf("<option value=\"%s\" %s>%s</option>",
            server.Config.ServiceName, selected, server.Config.ServiceName)))
    }
    writer.Write([]byte("</select>"))


    writer.Write([]byte(kWebPageStatusHeaderBlue))
    heading := fmt.Sprintf("\n<div><h3 style='width: 75em;'>%s Status</h3></div><br>\n", currentServer.Config.ServiceName)
    writer.Write([]byte(heading))

    if ! currentServer.Config.DatabaseIsConnected() {
        currentServer.Config.ConnectDatabase()
    }
    if currentServer.Config.PGSQL == nil {
        errorOutput = []byte(fmt.Sprintf("Database for %s is unavailable.", currentServer.Config.ServiceName))
    } else {
        statusOutput, errorOutput, error = currentServer.Config.PGSQL.RunSQLScript(kStatusQuery)
    }
    if len(errorOutput) > 0 {
        writer.Write(errorOutput)
    } else if error != nil {
        writer.Write([]byte(error.Error()))
    } else {
        writer.Write(statusOutput)
    }

    //  Done --

    writer.Write([]byte(kWebPageStatusTrailer))
    writer.Write([]byte(kWebPageTrailer))
}



//----------------------------------------------------------------------------------------
//                                                                            HelpHandlers
//----------------------------------------------------------------------------------------


func RedirectToSendRefreshStatus(writer http.ResponseWriter, httpRequest *http.Request) {
    http.Redirect(writer, httpRequest, "status/refresh", 303)
}


func SendHello(writer http.ResponseWriter, request *http.Request) {
   Log.Debugf("Request:\n%v\n", *request)
   fmt.Fprintf(writer, "<html><body><p>Hello Stats Server!</p></body></html>")
}


func ShowRequest(writer http.ResponseWriter, request *http.Request) {
   Log.Debugf("Request:\n%v\nServer File Path:\n%s", request, config.ServiceFilePath)
   fmt.Fprintf(writer, "<html><p>Hi!\n<br>\n<br>Request:\n<br>%v\n<br>\n<br>File Path:  %s\n<br></p></html>",
        *request, config.ServiceFilePath)
}



//----------------------------------------------------------------------------------------
//                                                                            Stats-Server
//----------------------------------------------------------------------------------------


func StatsServer() int {
    var error error
    Log.LogLevel = Log.LogLevelAll
    commandLine := strings.Trim(fmt.Sprint(os.Args), "[]")

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
        fmt.Fprintf(os.Stdout, "Version %s.\n", config.SoftwareVersion)
        return 0
    }

    //  Parse the config file --

    flagInputFilename, error = filepath.Abs(flagInputFilename)
    if len(flagInputFilename) > 0 {
        error = config.ParseFilename(flagInputFilename)
        if error != nil {
            return 1
        }
    }
    if flagPID {
        fmt.Fprintf(os.Stdout, "%s\n", config.PIDFileName())
        return 0
    }
    if flagVerbose {
        config.LogLevel = Log.LogLevelDebug
    }

    Log.SetFilename(config.LogFilename);
    Log.Startf("Stats-Server version %s pid %d.", config.SoftwareVersion, os.Getpid())
    Log.Infof ("Command line: %s.",  commandLine)
    Log.Debugf("Configuration: %v.", config)

    //  Lock our PID file --

    error = config.CreatePIDFile()
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

    //  Start database --

    error = config.ConnectDatabase()
    if error != nil { return 1 }
    defer config.DisconnectDatabase();

    //  Load configurations --

    for _, server := range servers {
        server.LoadConfiguration()
        defer func() { server.Config.DisconnectDatabase() } ()
    }

    // //  Get the prod and labs configs --

    // filename := filepath.Dir(flagInputFilename)
    // filename  = filepath.Join(filename, "BeingHappy-Server.config")
    // error = prodConfig.ParseFilename(filename)
    // if error != nil {
    //     Log.Warningf("Can't open '%s': %v.", filename, error)
    // } else {
    //     // eDebug
    //     prodConfig.DatabaseURI = "postgres://happinessadmin:happinessadmin@localhost:5432/happinessdatabase"
    //     prodConfig.ConnectDatabase()
    // }
    // defer prodConfig.DisconnectDatabase()

    // filename  = filepath.Dir(flagInputFilename)
    // filename  = filepath.Join(filename, "HappyLabs-Server.config")
    // error = labsConfig.ParseFilename(filename)
    // if error != nil {
    //     Log.Warningf("Can't open '%s'.", filename)
    // } else {
    //     labsConfig.ConnectDatabase()
    // }
    // defer labsConfig.DisconnectDatabase()

    //  Make our listener --

    httpListener, error := net.Listen("tcp", ":"+strconv.Itoa(config.ServicePort))
    if error != nil {
        Log.Errorf("Can't listen on port %d: %v.", config.ServicePort, error)
        return 1
    }

    //  Set up an interrupt handler --

    config.AttachToInterrupts(httpListener)

    //  Update the app download tables --

    go UpdateDownloadTableWithLogLineCount(-1)

    //  Set up & start our http handlers --

    http.HandleFunc(config.ServicePrefix, RedirectToSendRefreshStatus)
    http.HandleFunc(config.ServicePrefix+"/refresh", SendRefreshStatus)
    http.HandleFunc(config.ServicePrefix+"/hello", SendHello)
    http.HandleFunc(config.ServicePrefix+"/sms", SendTwilioResponse)

    http.Handle("/",
        http.StripPrefix(config.ServicePrefix,
        http.FileServer(http.Dir(config.ServiceFilePath))))

    Log.Infof("Listening for http at %d:%s.", config.ServicePort, config.ServicePrefix)
    http.Serve(httpListener, nil)
    Log.Exitf("EOJ")
    return 0
}


func main() {
    os.Exit(StatsServer())
}



//----------------------------------------------------------------------------------------
//
//                                                                        Status SQL Query
//
//----------------------------------------------------------------------------------------


var kStatusQuery =
`

--  ServerStats.sql
--
--  EB Smith, March 2015


\echo <div id="timestamp"><h3>Time</h3><pre><code>
select to_char(Now(), 'FMDay FMMonth FMDD, FMHH12:MI am') as "Time";
\echo </code></pre></div>


\echo <div><h3>Status</h3><pre><code>
select messageType as "Message", StringFromTimeInterval(Now(), timestamp) as "Elapsed"
    from ServerStatTable
    where messageType in ( 'Started', 'Terminated' )
    order by timestamp desc limit 1;
\echo </code></pre></div>


\echo <div><h3>Messages</h3><pre><code>
select count(*) as "Count", sum(bytesin) as "Bytes In", sum(bytesout) as "Bytes Out"
    from ServerStatTable;
\echo </code></pre></div>


\echo <div><h3>Users</h3><pre><code>
select sum(1) as "Total",
    sum(case when userstatus > 1 then 1 else 0 end) as "Active"
    from usertable;
\echo </code></pre></div>


\echo <div><h3>Visitors by Month</h3><pre><code>
with newusers as (
select
    date_trunc('month', CreationDate) as timestamp,
    count(*) as usercount
    from UserTable
    where CreationDate is not null
    group by date_trunc('month', CreationDate)
    order by date_trunc('month', CreationDate)
),
uniqueusers as (
select date_trunc('month', timestamp) as timestamp,
    count(distinct userid) as usercount
    from usereventtable
    group by 1 order by 1
)
select
    to_char(date_trunc('month', newusers.timestamp), 'Mon YYYY') as "Month",
    newusers.usercount as "New",
    uniqueusers.usercount as "Unique",
    rpad(lpad('', (newusers.usercount/10)::int, '#'),
        ((uniqueusers.usercount - newusers.usercount)/10)::int,
        '+') as "New / Unique by Month"
        from newusers
        full outer join uniqueusers on newusers.timestamp = uniqueusers.timestamp
;
\echo </code></pre></div>


\echo <div><h3>Returning Visits per Month</h3><pre><code>
with userVisitDays as (
select
    date_trunc('day', timestamp) as timestamp,
    userid
        from usereventtable
        group by 1, 2
        order by 2, 1
),
userVisitsPerMonth as (
select
    userid,
    date_trunc('month', timestamp) as timestamp,
    count(*) as visitsPerMonth
        from userVisitDays
        group by 1, 2
        order by 1, 2
)
select
    to_char(date_trunc('month', timestamp), 'Mon YYYY') as "Month",
    to_char(avg(visitsPerMonth), 'FM99.00') as "avg",
    max(visitsPerMonth),
    rpad(
        lpad('', (avg(visitsPerMonth))::int, '#'),
        (max(visitsPerMonth) - avg(visitsPerMonth))::int,
        '+') as "Returning user visit days per month"
        from userVisitsPerMonth
        group by timestamp
        order by timestamp
;
\echo </code></pre></div>


\echo <div><h3>Activty by Month</h3><pre><code>
with monthlyusers as (
select
    date_trunc('month', timestamp) as timestamp,
    userid,
    count(*) as events
        from usereventtable
        group by 1,2
        order by 1,2
)
select
    to_char(date_trunc('month', timestamp), 'Mon YYYY') as "Month",
    count(*) as "Unique Users",
    sum(monthlyusers.events) as "User Events",
    rpad(lpad('', (count(*)/10)::int, '#'),
        (sum(monthlyusers.events)/100 - (count(*)/10)::int)::int,
        '+') as "Users / Activity by Month"
    from monthlyusers
        group by timestamp
        order by timestamp
;
\echo </code></pre></div>


\echo <div><h3>Recent Activity</h3><pre><code>
select distinct usereventtable.userid,
        StringFromTimeInterval(Now(), max(usereventtable.timestamp)) || ' ago' as "Last Active",
        max(usereventtable.timestamp) as "When",
        usertable.name,
        devicetable.modelName
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    join devicetable on usereventtable.userid = devicetable.userid
    group by usereventtable.userid, usertable.name, devicetable.modelName
    order by max(usereventtable.timestamp) desc
    limit 20;
\echo </code></pre></div>


\echo <div><h3>Most Active Users</h3><pre><code>
select count(*) as "events", usereventtable.userid, usertable.name
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    group by usereventtable.userid, usertable.name
union
select count(*), 'Total', ' '
    from usereventtable
order by events desc
limit 20;
\echo </code></pre></div>


\echo <div><h3>Most Connected</h3><pre><code>
select friendtable.userid, usertable.name, count(*) as "friends"
    from friendtable
    left join usertable on usertable.userid = friendtable.userid
    group by friendtable.userid, usertable.name
    order by "friends" desc
    limit 20;
\echo </code></pre></div>


\echo <div><h3>Network Messages</h3><pre><code>
select messageType as "Message", count(*) as "Count",
    to_char(sum(elapsed), '99990.9990') as "Tot. Sec.",
    to_char(avg(elapsed), '99990.9990') as "Avg Response Sec.",
    sum(bytesin) as "Bytes In", sum(bytesout) as "Bytes Out"
    from ServerStatTable
    group by messageType
    order by "Tot. Sec." desc;
\echo </code></pre></div>
`



//----------------------------------------------------------------------------------------
//
//                                                                         Status Web Page
//
//----------------------------------------------------------------------------------------


var kWebPageHeader =
`
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-type" content="text/html; charset=UTF-8">
<meta name="apple-mobile-web-app-capable" content="yes">
<meta name="apple-mobile-web-app-status-bar-style" content="black">
<meta name="viewport" content="width=device-width, minimal-ui">
<link href="./favicon.png"  type="image/png" rel="shortcut icon">
<title>BeingHappy : Server Status</title>
<link rel="stylesheet" type="text/css" href="style.css">
<script type="text/javascript" src="stats.js"></script>
</head>

<body>
<div id="header">
<div>
<img id="logo" src="Logo.png" alt="Logo">
<div>
<span>being</span>
<span>happy</span>
</div>
</div>
</div>
<div id="stats-body">
`

var kWebPageStatusHeaderGrey =
`
<div class="stats-grey">
`

var kWebPageStatusHeaderRed =
`
<div class="stats-red">
`

var kWebPageStatusHeaderBlue =
`
<div class="stats-blue">
`

var kWebPageStatusTrailer =
`
</div>
`

var kWebPageTrailer =
`
</div>
</body>
</html>
`

