package main

import (
    "fmt"
    "testing"
)


func TestLogRegex(t *testing.T) {
    fmt.Println("Start")
    testString := "107.3.151.67 - - [25/May/2015:11:03:16 -0700] \"GET /beinghappy/ios-labs/HappyLabs.ipa HTTP/1.1\" 200 29874597"

    var test, truth string

    truth = "107.3.151.67"
    test = RegexFind("^([^\\s]*)", testString)
    if  test != truth {
        t.Errorf("Got '%s' but expected '%s'.", test, truth)
    }

    truth = "25/May/2015:11:03:16 -0700"
    test = RegexFindSubstring("\\Q[\\E(.*?)\\Q]\\E", testString)
    if test != truth {
        t.Errorf("Got '%s' but expected '%s'.", test, truth)
    }

    truth = "/beinghappy/ios-labs/HappyLabs.ipa"
    test = RegexFindSubstring("GET (.*?) HTTP", testString)
    if test != truth {
        t.Errorf("Got '%s' but expected '%s'.", test, truth)
    }

    truth = "200"
    test = RegexFindSubstring("\"\\s([[:digit:]].*)\\s", testString)
    if test != truth {
        t.Errorf("Got '%s' but expected '%s'.", test, truth)
    }

    truth = "29874597"
    test = RegexFindSubstring("\"\\s[[:digit:]].*\\s([[:digit:]].*)", testString)
    if test != truth {
        t.Errorf("Got '%s' but expected '%s'.", test, truth)
    }

    fmt.Println("End")
}

