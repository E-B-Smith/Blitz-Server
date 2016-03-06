//  ServerInfo  -  ServerInfo for each monitored server.
//
//  E.B.Smith  -  March, 2016


package main


import (
    "fmt"
    "net"
    "time"
    "path/filepath"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/ServerUtil"
)


type ServerInfo struct {
    ConfigFilename  string
    Config          ServerUtil.Configuration
}


func (serverInfo *ServerInfo) LoadConfiguration() error {
    filename := filepath.Dir(flagInputFilename)
    filename  = filepath.Join(filename, serverInfo.ConfigFilename)
    Log.Debugf("Loading config from file '%s'.", filename)
    error := serverInfo.Config.ParseFilename(filename)
    if error != nil {
        Log.Warningf("Can't open '%s': %v.", filename, error)
    }
    Log.Debugf("Loaded server info for %s:%d", serverInfo.Config.ServiceName, serverInfo.Config.ServicePort)
    return error
}


func (serverInfo *ServerInfo) GetStatus() string {
    portnum := serverInfo.Config.ServicePort + 1

    host := fmt.Sprintf("localhost:%d", portnum)
    connection, error := net.Dial("tcp", host)
    if error != nil {
        return error.Error()+"\n"
    }
    defer connection.Close()

    timeout := time.Now().Add(time.Duration(1.5 * float64(time.Second)))
    connection.SetDeadline(timeout)
    _, error = connection.Write([]byte("status\n"))
    if error != nil {
        return error.Error()+"\n"
    }

    reply := make([]byte, 1024)
    n, error := connection.Read(reply)
    if error != nil {
        return error.Error()+"\n"
    }
    return string(reply[:n])
}

