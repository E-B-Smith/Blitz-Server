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

/*
func SetEntityTags(userID, entityID string, entityType BlitzMessage.EntityType tags []*BlitzMessage.EntityTag) {
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
                userID, entityID, entityType, cleanTag)

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
*/


func SetEntityTagsWithUserID(userID, entityID string, entityType BlitzMessage.EntityType, tags []*BlitzMessage.EntityTag) {
    Log.LogFunctionName()

    var error error
    _, error = config.DB.Exec(
        `delete from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
        userID, entityID, entityType)
    if error != nil { Log.LogError(error) }

    for _, tag := range tags {
        if tag.UserHasTagged == nil ||
           tag.TagName == nil ||
           ! *tag.UserHasTagged {
            continue
        }

        cleanTag := strings.ToLower(strings.TrimSpace(*tag.TagName))
        if len(cleanTag) <= 0 { continue }

        var result sql.Result
        result, error = config.DB.Exec(
            `insert into EntityTagTable
                (userID, entityID, entityType, entityTag)
                values ($1, $2, $3, $4);`,
            userID, entityID, entityType, cleanTag)

        error = pgsql.RowUpdateError(result, error)
        if error != nil { Log.LogError(error) }
    }
}


func GetEntityTagsWithUserID(userID, entityID string, entityType BlitzMessage.EntityType) []*BlitzMessage.EntityTag {
    Log.LogFunctionName()

    tagArray := make([]*BlitzMessage.EntityTag, 0, 10)

    rows, error := config.DB.Query(
        `select
            entityTag,
            count(*),
            sum(case when userid = $1 then 1 else 0 end)
        from EntityTagTable
        where entityID = $2
          and entityType = $3
        group by entityTag;`,
        userID, entityID, entityType,
    )
    if error != nil {
        Log.LogError(error)
        return tagArray
    }
    defer rows.Close()

    for rows.Next() {
        var (
            tag             string;
            count           int64;
            userSelected    sql.NullBool;
        )
        error = rows.Scan(&tag, &count, &userSelected)
        if error != nil {
            Log.LogError(error)
            continue
        }

        cleanTag := strings.ToLower(strings.TrimSpace(tag))
        if len(cleanTag) <= 0 { continue }

        entityTag := BlitzMessage.EntityTag {
            TagName:        &cleanTag,
            TagCount:       Int32Ptr(int32(count)),
            UserHasTagged:  BoolPtrFromNullBool(userSelected),
        }

        tagArray = append(tagArray, &entityTag)
    }

    return tagArray
}


func UpdateEntityTags(session *Session, tagList *BlitzMessage.EntityTagList,
    ) *BlitzMessage.ServerResponse {

    if  tagList.EntityID == nil ||
        tagList.EntityType == nil ||
        *tagList.EntityType == BlitzMessage.EntityType_ETUnknown {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    SetEntityTagsWithUserID(session.UserID, *tagList.EntityID, *tagList.EntityType, tagList.EntityTags)
    return ServerResponseForCode(BlitzMessage.ResponseCode_RCSuccess, nil)
}

