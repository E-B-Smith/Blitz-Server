

//----------------------------------------------------------------------------------------
//
//                                                            BlitzHere-Server : Search.go
//                                                                         Search facility
//
//                                                                 E.B. Smith, April, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "strings"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                     AutocompleteRequest
//
//----------------------------------------------------------------------------------------


func AutocompleteRequest(session *Session, query *BlitzMessage.AutocompleteRequest,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if query == nil || query.Query == nil || len(*query.Query) < 2 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    var rows *sql.Rows
    var error error

    if query.SearchType != nil && *query.SearchType == BlitzMessage.SearchType_STTopics {

        rows, error = config.DB.Query(
            `select distinct entityTag, similarity(entityTag, $1) as similarity
                from EntityTagTable
                where substring(entityTag from 1 for 1) <> '.'
                order by similarity desc, entityTag
                limit 5;`,
            *query.Query,
        )

    } else {

        rows, error = config.DB.Query(
            `select word, similarity(word, $1) as similarity
                from autocompletetable
                order by similarity desc, word
                limit 5;`,
            *query.Query,
        )

    }
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    suggestions := make([]string, 0, 5)
    for rows.Next() {
        var (
            word    sql.NullString
            rank    sql.NullFloat64
        )
        error = rows.Scan(&word, &rank)
        if error != nil {
            Log.LogError(error)
        } else {
            suggestions = append(suggestions, word.String)
        }
    }

    results := BlitzMessage.AutocompleteResponse {
        Query:          query.Query,
        Suggestions:    suggestions,
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCodePtr(BlitzMessage.ResponseCode_RCSuccess),
        ResponseType:       &BlitzMessage.ResponseType { AutocompleteResponse: &results },
    }
    return response
}


//----------------------------------------------------------------------------------------
//
//                                                                           SearchRequest
//
//----------------------------------------------------------------------------------------


func UserSearchRequest(session *Session, query *BlitzMessage.UserSearchRequest,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if query == nil || query.Query == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    parts := strings.Split(*query.Query, " ")
    var queryString string
    for _, part := range parts {
        if len(part) == 0 { continue }

        if len(queryString) > 0 {
            queryString += " & " + part
        } else {
            queryString = part
        }
    }

    rows, error := config.DB.Query(
        `select userid from usertable
            where search @@ to_tsquery('english', $1)
            order by ts_rank(search, to_tsquery('english', $1)) desc;`,
        queryString,
    )
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    profiles := make([]*BlitzMessage.UserProfile, 0, 20)
    for rows.Next() {

        var userID  sql.NullString
        error = rows.Scan(&userID)
        if error != nil || ! userID.Valid {
            Log.LogError(error)
        } else {
            userprofile := ProfileForUserID(session, userID.String)
            if userprofile != nil { profiles = append(profiles, userprofile) }
        }
    }

    results := BlitzMessage.UserSearchResponse {
        Query:       query.Query,
        Profiles:    profiles,
    }

    response := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCodePtr(BlitzMessage.ResponseCode_RCSuccess),
        ResponseType:       &BlitzMessage.ResponseType { UserSearchResponse: &results },
    }
    return response
}


