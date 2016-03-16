//  BlitzHere-Server : EntityTags.go
//
//  E.B.Smith  -  March, 2015


package main


import (
    "strings"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func SetEntityTags(userID string, tags []*BlitzMessage.EntityTag) {
    Log.LogFunctionName()

    for _, tag := range tags {
        var error error
        var result sql.Result

        cleanTag := strings.ToLower(strings.TrimSpace(*tag.EntityTagName))
        if len(cleanTag) <= 0 { continue }

        if *tag.EntityIsTagged {

            result, error = config.DB.Exec(
                `insert into EntityTagTable
                    (userID, entityID, entityType, entityTag)
                    values ($1, $2, $3, $4);`,
                userID, tag.EntityID, tag.EntityType, cleanTag)

        } else {

            result, error = config.DB.Exec(
                `delete from EntityTagTable
                    where userID = $1
                      and entityID = $2
                      and entityType = $3
                      and entityTag  = $4;`,
                userID, tag.EntityID, tag.EntityType, cleanTag)

        }

        error = pgsql.RowUpdateError(result, error)
        if error != nil { Log.LogError(error) }
    }
}


func GetEntityTags(userID, entityID string, entityType BlitzMessage.EntityType) []*BlitzMessage.EntityTag {
    Log.LogFunctionName()

    tagArray := make([]*BlitzMessage.EntityTag, 0, 10)

    rows, error := config.DB.Query(
        `select entityTag,
            (select count(*) from EntityTagTable where entityID = $2 and entityType = $3)
            from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
            userID, entityID, entityType)
    if error != nil {
        Log.LogError(error)
        return tagArray
    }
    defer rows.Close()

    for rows.Next() {
        var (
            tag     string;
            count   int64;
        )
        error = rows.Scan(&tag, &count)
        if error != nil { continue }

        cleanTag := strings.ToLower(strings.TrimSpace(tag))
        if len(cleanTag) <= 0 { continue }

        entityTag := BlitzMessage.EntityTag {
            EntityID:       &entityID,
            EntityType:     &entityType,
            EntityTagName:  &cleanTag,
            EntityIsTagged: BoolPtr(true),
            EntityTagCount: Int32Ptr(int32(count)),
        }

        tagArray = append(tagArray, &entityTag)
    }

    return tagArray
}


func UpdateEntityTags(session *Session, tagUpdate *BlitzMessage.EntityTagList,
    ) *BlitzMessage.ServerResponse {

    SetEntityTags(session.UserID, tagUpdate.Tags)
    return ServerResponseForCode(BlitzMessage.ResponseCode_RCSuccess, nil)
}

