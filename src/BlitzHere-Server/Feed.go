//  BlitzHere-Server : Feed.go  -  Feed maintainence.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "errors"
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

    SetEntityTags(*feedPost.UserID, feedPost.PostTags)

    return error
}


//----------------------------------------------------------------------------------------
//                                                                       FeedPostForPostID
//----------------------------------------------------------------------------------------

var kScanFeedRowString =
`           postID,
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
`


type RowScanner interface {
    Scan(dest ...interface{}) error
}


func ScanFeedPostRow(row RowScanner) (*BlitzMessage.FeedPost, error) {
    Log.LogFunctionName()

    var (
        postID          sql.NullString
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
        &postID,
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
        PostID:             StringPtrFromNullString(postID),
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

    feedPost.PostTags = GetEntityTags(*feedPost.UserID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost)

    return &feedPost, nil
}


func FeedPostForPostID(postID string) (*BlitzMessage.FeedPost, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select ` + kScanFeedRowString +
        `   where postID = $1`, postID)

    return ScanFeedPostRow(row)
}


//----------------------------------------------------------------------------------------
//                                                                          UpdateFeedPost
//----------------------------------------------------------------------------------------


func UpdateFeedPost(session *Session, feedPostUpdate *BlitzMessage.FeedPostUpdateRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if feedPostUpdate.FeedPost.UserID == nil ||
        session.UserID != *feedPostUpdate.FeedPost.UserID {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCNotAuthorized, errors.New("Not authorized"))
    }

    if feedPostUpdate.UpdateVerb == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    if *feedPostUpdate.UpdateVerb == BlitzMessage.UpdateVerb_UVCreate ||
       *feedPostUpdate.UpdateVerb == BlitzMessage.UpdateVerb_UVUpdate {
        error := WriteFeedPost(feedPostUpdate.FeedPost)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
        }
        return ServerResponseForCode(BlitzMessage.ResponseCode_RCSuccess, nil)
    }

    if *feedPostUpdate.UpdateVerb == BlitzMessage.UpdateVerb_UVDelete {
        result, error := config.DB.Exec(
            `update FeedPostTable set postStatus = $1 where postID = $2;`,
                BlitzMessage.FeedPostStatus_FPSDeleted, feedPostUpdate.FeedPost.PostID)
        error = pgsql.RowUpdateError(result, error)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
        return ServerResponseForCode(BlitzMessage.ResponseCode_RCSuccess, nil)
    }

    return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid,
            fmt.Errorf("Unknown verb '%d'", feedPostUpdate.UpdateVerb))
}


//----------------------------------------------------------------------------------------
//                                                                          FetchFeedPosts
//----------------------------------------------------------------------------------------


func FetchFeedPosts(session *Session, fetchRequest *BlitzMessage.FeedPostFetchRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select ` + kScanFeedRowString +
        `   from FeedPostTable
            where postStatus = $1
              and parentID is null
              and timeActiveStart <= current_timestamp
              and timeActiveStop   > current_timestamp
            order by timestamp desc;`, BlitzMessage.FeedPostStatus_FPSActive)

    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    feedPosts := make([]*BlitzMessage.FeedPost, 0, 10)
    for rows.Next() {
        feedPost, error := ScanFeedPostRow(rows)
        if error != nil {
            Log.LogError(error)
        } else {
            feedPosts = append(feedPosts, feedPost)
        }
    }

    feedResponse := BlitzMessage.FeedPostFetchResponse {
        FeedPosts:      feedPosts,
    }
    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseType:       &BlitzMessage.ResponseType { FeedPostFetchResponse: &feedResponse },
    }

    return response
}

