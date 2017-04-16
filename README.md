# README.md

## Blitz-Server

* [Git and Github](#git-and-github)
* [Accessing the BlitzHere Server](#accessing-the-blitzhere-server)
* [Server Users and Files](#server-users-and-files)
* [Building the Server App](#building-the-server-app)
* [Server API Endpoints](#server-api-endpoints)
* [Server API Calls](#server-api-calls)

## Git and Github

### Git

Git is a version control system: It tracks and keeps a history of the changes made to a directory of
files on your computer. The changes can be reverted or fast forwarded as needed. Changes can be
shared with collaborators, and collaborators can browse and modify files in the project and share
back their changes.

The git version control system is freely available, open source, and was developed to help manage
the source code of the hundreds of developers that contribute to the Linux operating system.

A more detailed introduction can can be found at the git website:
[Git: About Version Control.](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control)

### Github

Github is a company that hosts git repositories. The advantage of using Github to host our
repositories is that it reduces the amount of administration time we have to spend hosting the
service ourself, plus they have some nice value added tools for managing repositories.

[About Github](https://github.com/business)

[Github Help](https://help.github.com/)

### BlitzHere at Github

All the source files for building the apps, the server app, the web site, and the design and
documentation are in private repositories on Github at the
[BlitzHere Github site.](https://github.com/BlitzHere/)

The project repositories are split by function and their names are pretty self explanatory.

[BlitzHere at Github](https://github.com/BlitzHere)

| Project                                                       |                               |
|---------------------------------------------------------------|-------------------------------|
| [Blitz-Android](https://github.com/BlitzHere/Blitz-Android)   | The Android App               |
| [Blitz-Web](https://github.com/BlitzHere/Blitz-Web)           | The Website                   |
| [Blitz-iOS](https://github.com/BlitzHere/Blitz-iOS)           | The iOS App                   |
| [Blitz-Server](https://github.com/BlitzHere/Blitz-Server)     | The Server App                |
| [Blitz-Design](https://github.com/BlitzHere/Blitz-Design)     | Documentation and design documents |
| [Blitz-Legacy](https://github.com/BlitzHere/Blitz-Legacy)     | The original Java server code |


### Inviting new developers to the Github project

Invite a new user to the Blitz Github project [here, at the Blitz Github project page](https://github.com/BlitzHere/)

### Setting up Github access on your local computer

Instructions for setting up `git` for Github can be found
[here, at the Github documentation.](https://help.github.com/articles/set-up-git/)

## Accessing the BlitzHere Server

The BlitzHere server contains all the public facing files for BlitzHere.  It hosts the website, the
database, and the server apps that communicate with the iOS and Android mobile apps.

The server is hosted by the Amazon AWS service: [Amazon AWS.](https://aws.amazon.com)

### Configuring your local computer to access the server

Access the server via ssh.  The `~/.ssh/` directory on your local computer will contains these files
among others, possibly:

| File Name             | Purpose                                                               |
|-----------------------|-----------------------------------------------------------------------|
|`authorized_keys`      | Contains keys of users allowed to log in to the computer.             |
|`config`               | Configuration file that tells `ssh` how to log into remote hosts.     |
|`known_hosts`          | Signatures of known computers.                                        |
|`<something_rsa>`      | Your private key file.  Keep this secret and don't distribute it.     |
|`<something_rsa>.pub`  | The corresponding public key file. This is the file you can share.    |

Edit your `~/.ssh/config` file to make it easier to log in to the Blitz server. Add the lines
below, replacing `~/.ssh/smith.ed.b@gmail.com@github.com` with the name of your own private key file.

```
ServerAliveInterval 60

Host *blitzhere*
    HostName blitzhere.com
    User blitzhere
    IdentityFile ~/.ssh/smith.ed.b@gmail.com@github.com
    Port 22
```

Now you should be able to access the server by typing `ssh blitzhere` from the command line, like this:

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

Log out with the `exit` command:

```
blitzhere@blitzhere:~$ exit
logout
Connection to blitzhere.com closed.
```

## Server Users and Files

There are three users on the server:

* `ubuntu`
  - The super user account. In general, don't use this account. It's the master admin account.
* `sysadmin`
  - The sysadmin account. This account is `sudo` capable and is mostly used for upgrading system
    software and configuring ssh, nginx, and the postgres database.
* `blitzhere`
  - The is a basic, secure, regular user account that contains all the blitzhere files and the
    server apps run under.

### Server Files

* backups - Database backups
* bin - The BlitzHere executable files
* database - Database maintenance files and scripts
* log - Server log files
* www - The BlitzHere website files

### Controlling the Server

#### The `sc` server control script

The `sc` command controls the server apps.

```
    sc  [ -f | --force ]  [ start | stop | restart | status ]  [ <server-app-name> | all ]
```

where `<server-app-name>` is one of `BlitzLabs-Server`, `BlitzHere-Server`, `Status-Server`.  The
`--force` option will force a server to start, stop, or restart if it isn't responding to normal
server commands.  When the `--force` option is used, the server isn't quit gracefully and some small
amount of data may be lost.

The `all` option applies the command to all the server apps, so

```
   sc -f restart all
```

will force restart all the server apps.

#### Server Log Files

The server log files are found on the server in the `log` directory. You can use the command

```
    less -Rqni --follow-name +F log/BlitzLabs-Server.log
```

to watch the real-time status of the BlitzLabs-Server.


## Building the Server App

### Prerequisites

You'll need to download and install some software on your local Mac or linux computer:

* The Go compiler: [The Go Programming Language](https://golang.org/dl/)
* (Mac) Install home brew: [Homebrew](https://brew.sh)
* (Mac) Install `automake`. From the command line:
```
    brew install automake
    brew install autoconf
    brew install libtool
```
* (All) Install the protocol buffer 2.6.1 compiler: [Protocol Buffers v2.6.1.](https://github.com/google/protobuf/releases/tag/v2.6.1)
  - Download and unzip the [protobuf-2.6.1.zip file.](https://github.com/google/protobuf/releases/download/v2.6.1/protobuf-2.6.1.zip)
  - Follow the instructions in the protobuf-2.6.1 `README.md` file to install the files.
    * That is, `cd` to your `protobuf-2.6.1` directory and `./autogen.sh`, etc. etc.
* (All) Build the objc protobuf compiler:
```
    git clone git@github.com:E-B-Smith/protobuf-objc.git
    cd protobuf-objc
    ./scripts/build.sh
```

### Building the Server

The server app is buit on your local computer and deployed to the server to run.
The `make` utility is used to keep track of what needs to be built, how to build it,
and how to deploy to the server.

In brief, clone the server project from GitHub to your local machine, `cd` into the directory,
and make all the apps:

```
    git clone git@github.com:BlitzHere/Blitz-Server.git
    cd Blitz-Server
    make all
```

### `make` options

| Make Command      |                                               |
|-------------------|-----------------------------------------------|
| `make clean`      | Remove all old intermediate build files.      |
| `make compile`    | Compile all the server apps.                  |
| `make deploy`     | Deploy the server apps on the server.         |
| `make restart`    | Restart the server apps.                      |
| `make status`     | Report the server status.                     |
| `make all`        | Do all of the above.                          |

The simplest way to build and deploy the server is:
```
    make all
```
on the command line.

Multiple options can be included on the make command execution. So the command:
```
    make compile deploy restart
```
will compile the source, deploy it to the server, and restart the server app.

### Server Project Files

The project contains these files:

| File / Directory  |                                                       |
|-------------------|-------------------------------------------------------|
| Database          | Files for generating and controlling the database.    |
| Protobuf          | Files for generating the protobuf scheme files.       |
| README.md         | This readme file.                                     |
| Server-Config     | Files for configuring ssh, nginx, etc on the server.  |
| Staging           | Files staged for deployment on the server.            |
| bin               | Intermediate server build files.                      |
| git-subtrees      | A script to pull / push git project subtrees.         |
| makefile          | The file that controls the `make` utility.            |
| pkg               | Intermediate server build files.                      |
| src               | The server app source files.                          |


## Server API Endpoints

The first API call to the Blitz service is the "Get Servers" call, which returns a list of server
end points.  Do a simple GET to `https://blitzhere.com/Servers.json` and the Server.json file is
returned.

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

The `"format": "Servers-JSON-1",` key-value is the format, which you should check to make sure
you're getting an expected format.

After the format, a dictionary of app identifiers is listed, each with a list of server end points.
The app identifier distinguishes between various production and development environments.
The app identifier `com.blitzhere.blitzhere-lab2` contains our current development server end points.

The `apiURL` key value is the API server to which to POST protobuf (preferred) or JSON formed
requests.

The `pushURL` key value is the web socket connection URL for chat messages.

The `statusMessageURL` key value is the URL of a web page that contains our current server state.
The web page is suitable for showing to the user.  This is not really used at the moment.

The `appUpdateURL` key value is the URL for iOS app to check for automatic iOS app updates.

We can add a URL for automatic Android updates too.


## Server API Calls

All of the non-chat client/server messages are passed to the 'apiURL' end point.

The API calls are POSTed to the server as protobuf (or JSON) messages.  The format POSTed determines
the format returned:  post JSON in, get JSON back.
