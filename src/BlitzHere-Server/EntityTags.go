//  BlitzHere-Server : EntityTags.go
//
//  E.B.Smith  -  March, 2015


package main


import (
    "strings"
    "violent.blue/GoKit/Log"
    "BlitzMessage"
)


func SetEntityTags(userID, entityID string, entityType BlitzMessage.EntityType, tags []string) {
    Log.LogFunctionName()

    config.DB.Exec(
        `delete from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
        userID, entityID, entityType)

    for _, tag := range tags {
        cleanTag := strings.TrimSpace(tag)
        if len(cleanTag) > 0 {
            _, error := config.DB.Exec(
                `insert into EntityTagTable (userID, entityID, entityType, entityTag)
                    values ($1, $2, $3, $4);`,
                userID, entityID, entityType, cleanTag)
            if error != nil { Log.LogError(error) }
        }
    }
}


func GetEntityTags(userID, entityID string, entityType BlitzMessage.EntityType) []string {
    Log.LogFunctionName()

    tags := make([]string,0, 10)

    rows, error := config.DB.Query(
        `select entityTag from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
            userID, entityID, entityType)
    if error != nil {
        Log.LogError(error)
        return tags
    }
    defer rows.Close()

    for rows.Next() {
        var tag string
        error = rows.Scan(&tag)
        if error == nil {
            s := strings.TrimSpace(tag)
            if len(s) > 0 { tags = append(tags, s) }
        }
    }

    return tags
}

