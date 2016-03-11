//  BlitzHere-Server : Feed.go  -  Feed maintainence.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                   FeedPost Read / Write
//
//----------------------------------------------------------------------------------------


func WriteFeedPost(feedPost *BlitzMessage.FeedPost) error {
    Log.LogFunctionName()

    _, error := config.DB.Exec(
        `insert into FeedPostTable (postID, postStatus, timestamp)
            values ($1, $2, current_timestamp);`, feedPost.PostID, BlitzMessage.FeedPostStatus_FPSActive)
    Log.Debugf("FeedPost Create status: %v.", error)

    result, error := config.DB.Exec(
        `update FeedPostTable set (
            parentID,
            postType,
            postScope,
            userID,
            anonymousPost,
            timeActiveStart,
            timeActiveStop,
            headlineText,
            bodyText,
            mayAddReply,
            mayChooseMulitpleReplies
        ) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) where postID = $12;`,
            feedPost.ParentID,
            feedPost.PostType,
            feedPost.PostScope,
            feedPost.UserID,
            feedPost.AnonymousPost,
            BlitzMessage.NullTimeFromTimespanStart(feedPost.TimespanActive),
            BlitzMessage.NullTimeFromTimespanStop(feedPost.TimespanActive),
            feedPost.HeadlineText,
            feedPost.BodyText,
            feedPost.MayAddReply,
            feedPost.MayChooseMulitpleReplies,
            feedPost.PostID)

    error = pgsql.RowUpdateError(result, error)
    if error != nil {
        Log.LogError(error)
        return error
    }

    SetEntityTags(*feedPost.PostID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost, feedPost.TopicTags)

    return error
}



func FeedPostForPostID(postID string) (*BlitzMessage.FeedPost, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
            parentID,
            postType,
            postScope,
            userID,
            anonymousPost,
            timestamp,
            timeActiveStart,
            timeActiveStop,
            headlineText,
            bodyText,
            mayAddReply,
            mayChooseMulitpleReplies
                where postID = $1`, postID)
    var (
        parentID        sql.NullString
        postType        sql.NullInt64
        postScope       sql.NullInt64
        userID          sql.NullString
        anonymousPost   sql.NullBool
        timestamp       pq.NullTime
        timeActiveStart pq.NullTime
        timeActiveStop  pq.NullTime
        headlineText    sql.NullString
        bodyText        sql.NullString
        mayAddReply     sql.NullBool
        mayChooseMulitpleReplies    sql.NullBool
    )
    error := row.Scan(
        &parentID,
        &postType,
        &postScope,
        &userID,
        &anonymousPost,
        &timestamp,
        &timeActiveStart,
        &timeActiveStop,
        &headlineText,
        &bodyText,
        &mayAddReply,
        &mayChooseMulitpleReplies)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    feedPost := BlitzMessage.FeedPost {
        ParentID:           StringPtrFromNullString(parentID),
        PostType:           BlitzMessage.FeedPostType(postType.Int64).Enum(),
        PostScope:          BlitzMessage.FeedPostScope(postScope.Int64).Enum(),
        UserID:             StringPtrFromNullString(userID),
        AnonymousPost:      BoolPtrFromNullBool(anonymousPost),
        Timestamp:          BlitzMessage.TimestampPtrFromNullTime(timestamp),
        TimespanActive:     BlitzMessage.TimespanFromNullTimes(timeActiveStart, timeActiveStop),
        HeadlineText:       StringPtrFromNullString(headlineText),
        BodyText:           StringPtrFromNullString(bodyText),
        MayAddReply:        BoolPtrFromNullBool(mayAddReply),
        MayChooseMulitpleReplies:   BoolPtrFromNullBool(mayChooseMulitpleReplies),
    }

    feedPost.TopicTags = GetEntityTags(*feedPost.PostID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost)

    return &feedPost, nil
}




