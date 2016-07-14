

//----------------------------------------------------------------------------------------
//
//                                                        BlitzHere-Server : EntityTags.go
//                                                        The server back-end to BlitzHere
//
//                                                                 E.B. Smith, March, 2015
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "strings"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func GetEntityTagMapForUserIDEntityIDType(
        userID,
        entityID string,
        entityType BlitzMessage.EntityType,
    ) map[string]bool {

    result := make(map[string]bool)
    rows, error := config.DB.Query(
        `select entityTag
            from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
        userID,
        entityID,
        entityType,
    )
    if error != nil {
        Log.LogError(error)
        return result
    }
    defer rows.Close()

    for rows.Next() {
        var tag sql.NullString
        rows.Scan(&tag)
        if tag.Valid { result[tag.String] = true }
    }

    return result
}


func CleanEntityTag(tag string) string {
    return strings.ToLower(strings.TrimSpace(tag))
}


func SetEntityTagsForUserIDEntityIDType(
        userID,
        entityID string,
        entityType BlitzMessage.EntityType,
        tags map[string]bool,
    )  {

    _, error := config.DB.Exec(
        `delete from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
        userID,
        entityID,
        entityType,
    )
    if error != nil { Log.LogError(error) }

    for key, val := range tags {
        cleanTag := CleanEntityTag(key)
        if val && len(cleanTag) > 0 {
            _, error = config.DB.Exec(
            `insert into EntityTagTable
                (userID, entityID, entityType, entityTag)
                values ($1, $2, $3, $4);`,
            userID, entityID, entityType, cleanTag)
            if error != nil { Log.LogError(error) }
        }
    }
}


func GetUserIDArrayForEntity(entityType BlitzMessage.EntityType, entityID string, entityTag string) []string {
    Log.LogFunctionName()

    resultArray := make([]string, 0, 10)

    rows, error := config.DB.Query(
        `select distinct userID from EntityTagTable
            where entityType = $1
              and entityTag  = $2
              and entityID   = $3;`,
        entityType,
        entityTag,
        entityID,
    )
    if error != nil {
        Log.LogError(error)
        return resultArray
    }
    defer rows.Close()

    var userID string
    for rows.Next() {
        error = rows.Scan(&userID)
        if error == nil {
            resultArray = append(resultArray, userID)
        }
    }

    return resultArray
}


func SetEntityTagsWithUserID(userID, entityID string, entityType BlitzMessage.EntityType, tags []*BlitzMessage.EntityTag) {
    Log.LogFunctionName()

    var error error
    var result sql.Result
    result, error = config.DB.Exec(
        `delete from EntityTagTable
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
        userID, entityID, entityType)
    if error != nil { Log.LogError(error) }
    Log.Debugf("Deleted %d tags.", pgsql.RowsUpdated(result))

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

        error = pgsql.UpdateResultError(result, error)
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


const (
    kTagFriendDidAsk     = ".frienddidask"
    kTagFriendWasAsked   = ".friendwasasked"
    kTagFriendAccepted   = ".friendaccepted"
    kTagFriend           = ".friend"
    kTagFriendIgnored    = ".friendignored"
)


func SendFriendRequest(session *Session, request *BlitzMessage.FriendUpdate) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  request.FriendID == nil ||
        request.FriendStatus == nil {
        return ServerResponseForCode(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    userTagMap   := GetEntityTagMapForUserIDEntityIDType(
        session.UserID,
        *request.FriendID,
        BlitzMessage.EntityType_ETUser,
    )
    friendTagMap := GetEntityTagMapForUserIDEntityIDType(
        *request.FriendID,
        session.UserID,
        BlitzMessage.EntityType_ETUser,
    )

    /*
    State Map:

    FSDidAsk:
        friendStatus == FSIgnored   =>      userStatus.DidAsk,  friendStatus.WasAsked, No notification.
        friendStatus == FSFiends    =>      userStatus.Friends, friendStatus.Friends,  No notification.
        friendStatus == FSDidAsk    =>      userStatus.Friends, friendStatus.Friends,  Notification.
        friendStatus == FSWasAsked  =>      userStatus.DidAsk,  friendStatus.WasAsked, No Notification.
        default                     =>      userStatus.DidAsk,  friendStatus.WasAsked, Notification.

    FSIgnore:
        => -userStatus.Friend, +friendStatus.Ignored -friendStatus.Friend, No Notification.

    FSAccept:
        friendStatus == FSIgnored   =>      userStatus.DidAsk  -userStatus.FSFriend, friendStatus.WasAsked   No notification.
        friendStatus == FSAccepted  =>      userStatus.Friend, friendStatus.Friend,  No notification.
        friendStatus == FSFriends   =>      userStatus.Friend, friendStatus.Friend,  No notification.
        friendStatus == FSDidAsk    =>      userStatus.Friend, friendStatus.Friend,   Notification.
        friendStatus == FSWasAsked  =>      userStatus.DidAsk, friendStatus.WasAsked, No notification.
        default                     =>      userStatus.DidAsk,  friendStatus.WasAsked, Notification.

    FSUnknown:
        Remove friend tags from both.
    */

    var ok bool
    message := ""
    actionURL := ""
    friendStatus := BlitzMessage.FriendStatus_FSUnknown

    if _, ok = friendTagMap[kTagFriendIgnored]; ok {

        friendStatus = BlitzMessage.FriendStatus_FSIgnored

    } else if _, ok = friendTagMap[kTagFriend]; ok {

        friendStatus = BlitzMessage.FriendStatus_FSFriends

    } else if _, ok = friendTagMap[kTagFriendAccepted]; ok {

        friendStatus = BlitzMessage.FriendStatus_FSAccepted

    } else if _, ok = friendTagMap[kTagFriendDidAsk]; ok {

        friendStatus = BlitzMessage.FriendStatus_FSDidAsk

    } else if _, ok = friendTagMap[kTagFriendWasAsked]; ok {

        friendStatus = BlitzMessage.FriendStatus_FSWasAsked

    }

    Log.Debugf("Friend update request: %s FriendStatus: %s.",
        request.FriendStatus.String(),
        friendStatus.String(),
    )

    if *request.FriendStatus == BlitzMessage.FriendStatus_FSUnknown {

        userTagMap[kTagFriendDidAsk] = false;
        userTagMap[kTagFriendWasAsked] = false;
        userTagMap[kTagFriendAccepted] = false;
        userTagMap[kTagFriend] = false;
        userTagMap[kTagFriendIgnored] = false;

        friendTagMap[kTagFriendDidAsk] = false;
        friendTagMap[kTagFriendWasAsked] = false;
        friendTagMap[kTagFriendAccepted] = false;
        friendTagMap[kTagFriend] = false;
        friendTagMap[kTagFriendIgnored] = false;

    } else
    if *request.FriendStatus == BlitzMessage.FriendStatus_FSIgnored {

        userTagMap[kTagFriendIgnored] = true
        userTagMap[kTagFriend] = false

        friendTagMap[kTagFriendDidAsk] = false;
        friendTagMap[kTagFriendWasAsked] = false;
        friendTagMap[kTagFriend] = false;

        Log.Debugf("EndState: Ignored.")

    } else {

        switch friendStatus {

        case BlitzMessage.FriendStatus_FSIgnored:
            userTagMap[kTagFriend] = false
            userTagMap[kTagFriendDidAsk] = true;
            userTagMap[kTagFriendWasAsked] = false;
            userTagMap[kTagFriendAccepted] = false;

            friendTagMap[kTagFriend] = false
            Log.Debugf("EndState: Ignoring.")

        case BlitzMessage.FriendStatus_FSFriends:
            userTagMap[kTagFriend] = true
            userTagMap[kTagFriendAccepted] = true;
            userTagMap[kTagFriendIgnored] = false
            friendTagMap[kTagFriend] = true
            friendTagMap[kTagFriendAccepted] = true;
            friendTagMap[kTagFriendIgnored] = false
            Log.Debugf("EndState: Made friends.")

        case BlitzMessage.FriendStatus_FSAccepted:
            userTagMap[kTagFriend] = true
            userTagMap[kTagFriendAccepted] = true;
            userTagMap[kTagFriendIgnored] = false
            friendTagMap[kTagFriend] = true
            friendTagMap[kTagFriendAccepted] = true;
            friendTagMap[kTagFriendIgnored] = false
            Log.Debugf("EndState: Send friend re-accept.")

        case BlitzMessage.FriendStatus_FSDidAsk:
            userTagMap[kTagFriend] = true
            userTagMap[kTagFriendAccepted] = true;
            userTagMap[kTagFriendIgnored] = false
            friendTagMap[kTagFriend] = true
            friendTagMap[kTagFriendAccepted] = true;
            friendTagMap[kTagFriendIgnored] = false
            message =
                fmt.Sprintf("%s accepted your friend request.",
                    PrettyNameForUserID(session.UserID))
            actionURL =
                fmt.Sprintf("%s?action=showuser&userid=%s",
                    config.AppLinkURL,
                        session.UserID)
            Log.Debugf("EndState: Send friend accept.")

        case BlitzMessage.FriendStatus_FSWasAsked:
            userTagMap[kTagFriendDidAsk] = true
            friendTagMap[kTagFriendWasAsked] = true
            Log.Debugf("EndState: Friend request already sent.")

        default:
            userTagMap[kTagFriendDidAsk] = true
            friendTagMap[kTagFriendWasAsked] = true
            message =
                fmt.Sprintf("%s sent you a friend request.",
                    PrettyNameForUserID(session.UserID))
            actionURL =
                fmt.Sprintf("%s?action=friend&userid=%s",
                    config.AppLinkURL,
                        session.UserID)
            Log.Debugf("EndState: Send friend request.")

        }
    }

    SetEntityTagsForUserIDEntityIDType(
        session.UserID,
        *request.FriendID,
        BlitzMessage.EntityType_ETUser,
        userTagMap,
    )
    SetEntityTagsForUserIDEntityIDType(
        *request.FriendID,
        session.UserID,
        BlitzMessage.EntityType_ETUser,
        friendTagMap,
    )

    if len(message) > 0 {
        SendUserMessageInternal(
            session.UserID,
            []string { *request.FriendID },
            "",
            message,
            BlitzMessage.UserMessageType_MTActionNotification,
            "",
            actionURL,
        )
    }

    profiles := make([]*BlitzMessage.UserProfile, 2)
    profiles[0] = ProfileForUserID(session.UserID, session.UserID)
    profiles[1] = ProfileForUserID(session.UserID, *request.FriendID)

    response := BlitzMessage.FriendUpdate {
        FriendStatus:       request.FriendStatus,
        FriendID:           request.FriendID,
        Profiles:           profiles,
    }

    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:        BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FriendResponse: &response },
    }
    return serverResponse
}



