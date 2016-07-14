

//----------------------------------------------------------------------------------------
//
//                                                 BlitzHere-Server : ConversationGroup.go
//                                                                           Conversations
//
//                                                                  E.B. Smith, June, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                       FetchFeedPostsAsConversationGroup
//----------------------------------------------------------------------------------------


func FetchFeedPostsAsConversationGroup(userID string) []*BlitzMessage.ConversationGroup {
    Log.LogFunctionName()

    resultArray := make([]*BlitzMessage.ConversationGroup, 0, 10)

    rows, error := config.DB.Query(
        `with FeedPostIDTable as (
            select entityid as postID from entitytagtable
                where entitytype = 2
                  and userid = $1
                  and entityTag = '.followed'
            union
            select coalesce(parentID, postID) as postID
                from feedposttable
                where userid = $1
        )
        , FeedPost as (
            select  FeedPostTable.postID,
                    FeedPostTable.parentID,
                    FeedPostTable.userID,
                    FeedPostTable.timestamp,
                    FeedPostTable.headlineText
                 from FeedPostIDTable
                inner join FeedPostTable
                   on FeedPostTable.postID = FeedPostIDTable.postID
        )
        select
            FeedPost.*,
            (select count(*) from feedposttable where parentID = FeedPost.postID),
            Latest.timestamp,
            Latest.headlineText,
            Latest.userID
            from FeedPost
            left join FeedPostTable as Latest
             on Latest.parentID = FeedPost.postID
            and latest.timestamp =
                (select max(timestamp) from FeedPostTable as ff
                    where ff.parentID = FeedPost.postID);`,
        userID,
    )
    if error != nil {
        Log.LogError(error)
        return resultArray
    }
    defer rows.Close()

    conversationType := BlitzMessage.ConversationType_CTFeedPost

    for rows.Next() {
        var (
            postID      sql.NullString
            parentID    sql.NullString
            userID      sql.NullString
            createDate  pq.NullTime
            headline    sql.NullString
            replyCount  sql.NullInt64
            replyDate   pq.NullTime
            replyText   sql.NullString
            replyUser   sql.NullString
        )

        error = rows.Scan(
            &postID,
            &parentID,
            &userID,
            &createDate,
            &headline,
            &replyCount,
            &replyDate,
            &replyText,
            &replyUser,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }

        var conv BlitzMessage.ConversationGroup
        conv.GroupID            =   &postID.String
        conv.GroupType          =   &conversationType
        conv.UserID             =   &userID.String
        if replyDate.Valid {
            conv.ActivityDate   =   BlitzMessage.TimestampPtr(replyDate)
        } else {
            conv.ActivityDate   =   BlitzMessage.TimestampPtr(createDate)
        }
        conv.HeadlineText       =   &headline.String
        conv.LastMessage        =   &replyText.String
        if replyUser.Valid {
            conv.LastUserID     =   &replyUser.String
        }
        conv.TotalCount         =   Int32PtrFromNullInt64(replyCount)

        resultArray = append(resultArray, &conv)
    }

    Log.Debugf("Found %d feed posts.", len(resultArray))
    return resultArray
}


//----------------------------------------------------------------------------------------
//                                                       FetchNotificationsAsConversations
//----------------------------------------------------------------------------------------


func FetchNotificationsAsConversationGroup(userID string) []*BlitzMessage.ConversationGroup {
    Log.LogFunctionName()

    ary := make([]*BlitzMessage.ConversationGroup, 0, 20)

    rows, error := config.DB.Query(
        `select
            messageID,
            senderID,
            recipientID,
            messageStatus,
            creationDate,
            readDate,
            messageText,
            actionURL
        from UserMessageTable
        where recipientID = $1
          and recipientID <> senderID
          and messageType = $2;`,
        userID,
        BlitzMessage.UserMessageType_MTActionNotification,
    )
    if error != nil {
        Log.LogError(error)
        return ary
    }
    defer rows.Close()

    conversationType := BlitzMessage.ConversationType_CTNotification

    for rows.Next() {
        var (
            messageID           sql.NullString
            senderID            sql.NullString
            recipientID         sql.NullString
            status              sql.NullInt64
            creationDate        pq.NullTime
            readDate            pq.NullTime
            lastMessage         sql.NullString
            lastActionURL       sql.NullString
        )

        error = rows.Scan(
            &messageID,
            &senderID,
            &recipientID,
            &status,
            &creationDate,
            &readDate,
            &lastMessage,
            &lastActionURL,
        )

        var unreadCount int32 = 1
        if readDate.Valid {
            unreadCount = 0
        }

        var conv BlitzMessage.ConversationGroup
        conv.GroupID            =   &messageID.String
        conv.GroupType          =   &conversationType
        conv.UserID             =   &senderID.String
        conv.ActivityDate       =   BlitzMessage.TimestampPtr(creationDate)
        conv.LastMessage        =   &lastMessage.String
        conv.LastUserID         =   &senderID.String
        conv.TotalCount         =   proto.Int32(1)
        conv.UnreadCount        =   proto.Int32(unreadCount)
        conv.ActionURL      =   &lastActionURL.String

        ary = append(ary, &conv)
    }
    return ary
}


//----------------------------------------------------------------------------------------
//                                                                 FetchConversationGroups
//----------------------------------------------------------------------------------------


/*
message ConversationGroup {
  optional string             groupID       = 1;    //  For feed items, feed.postID, conversations then 'other' memberID.
  optional ConversationType   groupType     = 2;
  optional string             userID        = 3;    //  Feed: Initiator | Message: other userID.
  optional Timestamp          activityDate  = 4;
  optional string             headlineText  = 5;
  optional string             statusText    = 6;
  optional string             lastMessage   = 7;
  optional string             lastUserID    = 8;    //  UserID from last message.
  optional int32              totalCount    = 9;
  optional int32              unreadCount   = 10;
  optional string             actionURL     = 11;
}
*/

func FetchConversationGroups(
        session *Session, req *BlitzMessage.FetchConversationGroups,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    userForConversationID := func(conversationID, sessionUserID string) *string {
        Log.Debugf("cid: %s sid: %s.", conversationID, sessionUserID)
        row := config.DB.QueryRow(
            `select memberID from ConversationMemberTable
                where conversationID = $1
                  and memberID <> $2
                limit 1;`,
            conversationID,
            sessionUserID,
        )
        var memberID sql.NullString
        error := row.Scan(&memberID)
        if error != nil || ! memberID.Valid {
            return nil
        }
        Log.Debugf("Found: %s.", memberID.String)
        return &memberID.String
    }

    rows, error := config.DB.Query(
        `select distinct on (conversationID, senderID)
            conversationID, senderID, creationDate, messageText
        from (
            select conversationID, senderID, messageText, creationDate,
                rank() over (partition by senderID order by creationDate desc) as r
            from UserMessageTable
            where recipientID = $1
              and conversationID is not null
              order by conversationID, creationDate desc
        ) as conv
        order by conversationID, senderID, creationDate desc;`,
        session.UserID,
    )
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }
    defer rows.Close()

    groups := make([]*BlitzMessage.ConversationGroup, 0, 10)
    var groupConversationID string
    var group *BlitzMessage.ConversationGroup
    for rows.Next() {
        var (
            convID          string
            senderID        string
            creationDate    pq.NullTime
            messageText     string
        )
        error = rows.Scan(&convID, &senderID, &creationDate, &messageText)
        if error != nil {
            Log.LogError(error)
            continue
        }
        if convID != groupConversationID && group != nil {
            if group.GroupID == nil {
                group.GroupID = userForConversationID(groupConversationID, session.UserID)
                group.UserID = group.GroupID
            }
            groups = append(groups, group)
            group = nil
        }
        groupConversationID = convID
        if group == nil {
            group = &BlitzMessage.ConversationGroup {
                GroupType:  BlitzMessage.ConversationType(BlitzMessage.ConversationType_CTConversation).Enum(),
            }
        }
        group.ActivityDate = BlitzMessage.TimestampPtr(creationDate)
        if senderID == BlitzMessage.Default_Global_SystemUserID {
            group.StatusText = &messageText
        } else {

            if senderID != session.UserID {
                group.GroupID = &senderID
                group.UserID  = &senderID
            }

            group.LastMessage = &messageText
            group.LastUserID  = &senderID
        }
    }
    if group != nil {
        if group.GroupID == nil {
            group.GroupID = userForConversationID(groupConversationID, session.UserID)
            group.UserID = group.GroupID
        }
        groups = append(groups, group)
    }

    //  Get un-read message counts --

    for _, group = range groups {
        row := config.DB.QueryRow(
            `select count(*),
                sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
                from (
                    select a.conversationID as cid, a.memberID as mid, b.memberID
                        from conversationMemberTable a
                        join conversationMemberTable b
                            on a.conversationID = b.conversationID
                           and b.memberID = $1
                         where a.memberID = $2
                ) as conv,
                usermessagetable
                    where conversationID = conv.cid and recipientID = conv.mid;`,
            group.GroupID,
            session.UserID,
        )
        var total, unread sql.NullInt64
        error = row.Scan(&total, &unread)
        if error != nil {
            Log.LogError(error)
        }
        group.TotalCount  = Int32PtrFromInt64(total.Int64)
        group.UnreadCount = Int32PtrFromInt64(unread.Int64)
    }

    groups = append(groups, FetchFeedPostsAsConversationGroup(session.UserID)...)
    groups = append(groups, FetchNotificationsAsConversationGroup(session.UserID)...)

    response := BlitzMessage.FetchConversationGroups { Conversations: groups }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FetchConversationGroups: &response },
    }

    return serverResponse
}
