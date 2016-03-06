//  ResourceData.go  -  Generated resource file.
//  Resources - E.B.Smith  -  Sun Jun 21 23:36:28 PDT 2015


package Resource


import (
    "bytes"
    "strings"
    "reflect"
    "io/ioutil"
    "compress/gzip"
    "encoding/base64"
    "violent.blue/GoKit/ServerUtil"
    "violent.blue/GoKit/Log"
)


type ResourceData struct {
    name    string
    data    string
    header  gzip.Header
}


var Resource ResourceData = ResourceData{}


func (resource *ResourceData) Name() string {
    return resource.name
}


func (resource *ResourceData) Header() gzip.Header {
    return resource.header
}


func (resource *ResourceData) Bytes() *[]byte {
    compressedBytes, error := base64.StdEncoding.DecodeString(resource.data)
    if error != nil { return nil; }

    buffer := bytes.NewBuffer(compressedBytes)
    gz, error := gzip.NewReader(buffer)
    if error != nil { Log.LogError(error); return nil; }
    defer gz.Close()

    rawbytes, error := ioutil.ReadAll(gz)
    if error != nil { Log.LogError(error); return nil; }
    resource.header = gz.Header

    return &rawbytes;
}


func (resource *ResourceData) ResourceBytesNamed(name string) *[]byte {
    name = ServerUtil.ReplaceCharactersNotInSetWithRune(name, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", '_')
    name = strings.Title(name)
    Log.Debugf("Loading resource '%s'.", name)
    methodV := reflect.ValueOf(resource).MethodByName(name)
    if methodV.IsValid() {
        values := methodV.Call([]reflect.Value{ })
        //Log.Debugf("Values: %+v.", values)
        if len(values) == 1 && values[0].CanInterface() {
            data, found := values[0].Interface().(*ResourceData)
            if (found) { return data.Bytes() }
        }
    }
    return nil
}


func (resource *ResourceData) TestResource_txt() *ResourceData {
    return &ResourceData{
    name: "TestResource_txt",
    data:
`
H4sIAOysh1UCAwvJyCxWAKLcSoWS1OIShZTEkkQ9LgDDwtLUFgAAAA==
`   }
}

