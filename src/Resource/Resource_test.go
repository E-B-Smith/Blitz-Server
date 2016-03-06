//  Resource  -  Compile time resource data embeded in the executable
//
//  E.B.Smith  -  June, 2015


package Resource


import (
    "testing"
    "violent.blue/GoKit/Log"
)


func TestResources1(t *testing.T) {
    Log.LogLevel = Log.LevelAll

    resdata := Resource.TestResource_txt().Bytes()
    if resdata == nil {
        t.Errorf("Result was nil!")
        return
    }

    truth := "This is my test data.\n"
    resstring := string(*resdata)
    if truth != resstring {
        t.Errorf("Expected '%s' but got '%s'.", truth, resstring)
    } else {
        Log.Debugf("TestResources1 success.")
    }
}


func TestResources2(t *testing.T) {
    Log.LogLevel = Log.LevelAll

    resdata := Resource.ResourceBytesNamed("TestResource_txt")
    if resdata == nil {
        t.Errorf("Result was nil!")
        return
    }

    truth := "This is my test data.\n"
    resstring := string(*resdata)
    if truth != resstring {
        t.Errorf("Expected '%s' but got '%s'.", truth, resstring)
    } else {
        Log.Debugf("TestResources2 success.")
    }
}



