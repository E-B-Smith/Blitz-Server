# README.md

## Blitz-Server

* [Git and Github](#Git-and-Github)
* [Accessing the Server](#accessing-the-blitzhere-server)
* [Server Users and Files](#Server-Users-and-Files)
* [Building the Server App](#Building-the-Server-App)
* [Server API Endpoints](#Server-API-Endpoints)
* [Server API Calls](#Server-API-Calls)

### Git and Github

#### Git

Git is a version control system: It tracks the changes to a directory of files on your computer. The history of all the changes is kept, and changes can be reverted or fast forwarded as needed. Also, git enables you to share the changes with collaborators.  The collaborators can make browse and modify the files and share back their changes.

Read the git introduction at the git website:  [Git: About Version Control](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control)

#### Github

Github is a company that hosts git repositories. The advantage of paying github to host our repository is that it reduces the amount of administration time we have to spend hosting the service ourselfs, plus they have some nice value added tools to manage repositories.

[About Github](https://github.com/business)

[Github Help](https://help.github.com/)


#### BlitzHere at Github

All the source files for building the apps, the server app, the web site, and the design and documentation are on Github at the [BlitzHere Github site.](https://github.com/BlitzHere/)

The project repositories are split by function and their names are usually self explanitory.

* Blitz-Android Private  -  The Android version of Blitz
* Blitz-Web  -  The BlitzHere website.
* Blitz-iOS  -  The BlitzHere iOS App.
* Blitz-Server  -  The BlitzHere server.
* Blitz-Design  -  Blitz documentation and design documents.
* Blitz-Legacy  -  The original Java server code.

#### Inviting new develpoers to the Github project

Invite a new user to the Blitz Github project [here, at the Blitz Github project page](https://github.com/BlitzHere/)

#### Setting up Github access on your local computer

Instructions for setting up `git` for Github can be found [here, at the Github documentation.](https://help.github.com/articles/set-up-git/)

### Accessing the BlitzHere Server

The BlitzHere server contains all the public facing files for BlitzHere.  It hosts the website, the database, and the server apps that comminucate with the iOS and Android mobile apps.

The server is hosted by the Amazon AWS service: [Amazon AWS.](https://aws.amazon.com)

#### Configuring your local computer to access the server

Access the server via ssh.  Make sure your In your `~/.ssh/` directory.

```
authorized_keys                     config                                  known_hosts
smith.ed.b@gmail.com@github.com     smith.ed.b@gmail.com@github.com.pub
```

`~/.ssh/config`
```
ServerAliveInterval 60

Host blitzhere
    HostName blitzhere.com
    User blitzhere
    IdentityFile ~/.ssh/smith.ed.b@gmail.com@github.com
    Port 22
```

You should be able to access the server by typing `ssh blitzhere`:
```
Clarity:Blitz-Server Edward$ ssh blitzhere
Welcome to Ubuntu 16.04.2 LTS (GNU/Linux 3.13.0-100-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Sun Apr  2 13:38:27 PDT 2017

  System load:  0.02              Processes:           122
  Usage of /:   28.1% of 7.74GB   Users logged in:     1
  Memory usage: 9%                IP address for eth0: 172.31.0.28
  Swap usage:   0%

  Graph this data and manage this system at:
    https://landscape.canonical.com/

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.


Last login: Sun Apr  2 13:38:27 2017 from 24.5.77.27
```

 The `whoami` command will tell you the user that you are logged in as:
```
blitzhere@blitzhere:~$ whoami
blitzhere
blitzhere@blitzhere:~$
```

Log out with the `exit` command:

```
blitzhere@blitzhere:~$ exit
logout
Connection to blitzhere.com closed.
```

### Server Users and Files

There are three users on the server:

* `ubuntu` - The super user account. Don't use this account.
* `sysadmin` - The sysadmin account for managing the computer. This account is `sudo` capable and is mostly used for upgrading system software.
* `blitzhere` - The is a basic, regular user account that contains all the blitzhere files.

#### BlitzHere Files

* backups - Database backups
* bin - The BlitzHere executable files
* database - Database maintenance scripts
* log - Server log files
* www - The BlitzHere website files


### Building the Server App

Clone the Server Code

### Server API Endpoints


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


### Server API Calls

All of the non-chat client/server messages are passed to the 'apiURL' end point.

The API calls are POSTed to the server as protobuf (or JSON) messages.  The format POSTed determines the format returned:  post JSON in, get JSON back.
