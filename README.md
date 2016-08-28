# Blitz-Server
## Server API URLs


The first API call to the Blitz service is the "Get Servers" call, which returns a list of server end points.  Do a simple GET to `https://blitzhere.com/Servers.json` and the Server.json file is returned.
        
The Server.json file is a JSON file that lists possible server API end points:

```
{
"format": "Servers-JSON-1",
"serverURLs":
    [
    "https://blitzhere.com/Servers.json"
    ],
"com.blitzhere.blitzhere-labs":
    {
    "apiURL":               "https://blitzhere.com/blitzlabs/api",
    "pushURL":                "wss://blitzhere.com/blitzlabs/push",
    "statusMessageURL":     "https://blitzhere.com/status/blitzlabs.html",
    "appUpdateURL":         "https://blitzhere.com/ios/blitz-labs/BlitzLabs.plist"
    },
"com.blitzhere.blitzhere-lab2":
    {
    "apiURL":               "https://blitzhere.com/blitzlabs/api",
    "pushURL":                "wss://blitzhere.com/blitzlabs/push",
    "statusMessageURL":     "https://blitzhere.com/status/blitzlabs.html",
    "appUpdateURL":         "https://blitzhere.com/ios/blitz-labs/BlitzLabs.plist"
    },
"com.blitzhere.blitzhere":
    {
    "apiURL":               "https://blitzhere.com/blitzhere/api",
    "pushURL":                "wss://blitzhere.com/blitzhere/push",
    "statusMessageURL":     "https://blitzhere.com/status/blitzhere.html",
    "appUpdateURL":         "https://blitzhere.com/ios/blitz/Blitz.plist"
    }
}
```

The `"format": "Servers-JSON-1",` key-value is the format, which you should check to make sure you're getting an expected format.

After the format, a dictionary of app identifiers is listed, each with a list of server end points.  The app identifier distiguishes between various production and development environments.   The app identifier `com.blitzhere.blitzhere-lab2` contains our current development server end points. 

The `apiURL` key value is the API server to which to POST protobuff (preferred) or JSON formed requests.

The `pushURL` key value is the web socket connection URL for chat messages.

The `statusMessageURL` key value is the URL of a web page that contains our current server state.  The web page is suitable for showing to the user.  This is not really used at the moment.

The `appUpdateURL` key value is the URL for iOS app to check for automatic iOS app updates.

We can add a URL for automatic Android updates too.


## Server API Calls

All of the non-chat client/server messages are passed to the 'apiURL' end point.

The API calls are POSTed to the server as protobuf (or JSON) messages.  The format POSTed determines the format returned:  post JSON in, get JSON back.

