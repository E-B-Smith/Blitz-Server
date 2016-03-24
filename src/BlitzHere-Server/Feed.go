//  BlitzHere-Server : Feed.go  -  Feed maintainence.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "errors"
    "strings"
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
            mayChooseMulitpleReplies,
            surveyAnswerSequence
        ) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) where postID = $13;`,
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
            feedPost.SurveyAnswerSequence,
            feedPost.PostID,
        )

    error = pgsql.RowUpdateError(result, error)
    if error != nil {
        Log.LogError(error)
        return error
    }

    SetEntityTagsWithUserID(*feedPost.UserID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost, feedPost.PostTags)

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
            mayChooseMulitpleReplies,
            surveyAnswerSequence
`


type RowScanner interface {
    Scan(dest ...interface{}) error
}


func ScanFeedPostRowForUserID(queryUserID string, row RowScanner) (*BlitzMessage.FeedPost, error) {
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
        surveyAnswerSequence sql.NullInt64
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
        &mayChooseMulitpleReplies,
        &surveyAnswerSequence)
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
        SurveyAnswerSequence:       Int32PtrFromNullInt64(surveyAnswerSequence),
    }

    feedPost.PostTags = GetEntityTagsWithUserID(queryUserID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost)

    return &feedPost, nil
}


func FeedPostForPostID(userID string, postID string) *BlitzMessage.FeedPost {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select ` + kScanFeedRowString +
        `   from FeedPostTable
            where postID = $1`, postID)

    feedPost, error := ScanFeedPostRowForUserID(userID, row)
    if error != nil {
        Log.LogError(error)
    }
    return feedPost
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

        //  If it's a survey question, write the responses --

        for _, reply := range feedPostUpdate.FeedPost.Replies {
            error := WriteFeedPost(reply)
            if error != nil { Log.LogError(error) }
        }

        //  Send a notification if it's a response --

        if  feedPostUpdate.FeedPost.ParentID != nil {
            Log.Debugf("Try to send a notification to the original poster:")
            parentPost := FeedPostForPostID(session.UserID, *feedPostUpdate.FeedPost.ParentID)
            if  parentPost != nil {
                replyPoster := ProfileForUserID(session.UserID)
                name := "Someone"
                if  replyPoster != nil &&
                    replyPoster.Name != nil &&
                    len(*replyPoster.Name) > 0 {
                    name = *replyPoster.Name
                }
                message := fmt.Sprintf("%s responded to your question.", name)
                SendUserMessage(BlitzMessage.Default_Globals_SystemUserID,
                    [] string { *parentPost.UserID },
                    message,
                    BlitzMessage.UserMessageType_MTNotification,
                    "AppIcon",
                    "",
                )
            }
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
//                                                          FetchTopOpenRepliesForFeedPost
//----------------------------------------------------------------------------------------


func FetchTopOpenRepliesForFeedPost(queryUserID string, parentPostID string, limit int) []*BlitzMessage.FeedPost {
    Log.LogFunctionName()

    feedPosts := make([]*BlitzMessage.FeedPost, 0, 10)
    rows, error := config.DB.Query(
        `select ` + kScanFeedRowString +
        `   from FeedPostTable
            where postStatus = $1
              and parentID = $2
              order by timestamp limit $3 ;`,
        BlitzMessage.FeedPostStatus_FPSActive,
        parentPostID,
        limit,
    )

    if error != nil {
        Log.LogError(error)
        return feedPosts
    }

    for rows.Next() {
        feedPost, error := ScanFeedPostRowForUserID(queryUserID, rows)
        if error != nil {
            Log.LogError(error)
        } else {
            feedPosts = append(feedPosts, feedPost)
        }
    }

    Log.Debugf("Found %d top posts.", len(feedPosts))
    return feedPosts
}


//----------------------------------------------------------------------------------------
//                                                        FetchTopSurveyRepliesForFeedPost
//----------------------------------------------------------------------------------------


func FetchTopSurveyRepliesForFeedPost(queryUserID string, parentPostID string, limit int) []*BlitzMessage.FeedPost {
    Log.LogFunctionName()

    feedPosts := make([]*BlitzMessage.FeedPost, 0, 10)
    rows, error := config.DB.Query(
        `select ` + kScanFeedRowString +
        `   from FeedPostTable
            where postStatus = $1
              and parentID = $2
              order by surveyAnswerSequence nulls last, timestamp
              limit $3;`,
        BlitzMessage.FeedPostStatus_FPSActive,
        parentPostID,
        limit,
    )

    if error != nil {
        Log.LogError(error)
        return feedPosts
    }

    for rows.Next() {
        feedPost, error := ScanFeedPostRowForUserID(queryUserID, rows)
        if error != nil {
            Log.LogError(error)
        } else {
            feedPosts = append(feedPosts, feedPost)
        }
    }

    Log.Debugf("Found %d top posts.", len(feedPosts))
    return feedPosts
}


//----------------------------------------------------------------------------------------
//                                                                          FetchFeedPosts
//----------------------------------------------------------------------------------------


func FetchFeedPosts(session *Session, fetchRequest *BlitzMessage.FeedPostFetchRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var ( rows *sql.Rows; error error )

    var parentID *string = nil
    if fetchRequest.ParentID != nil {
        p := strings.ToLower(strings.TrimSpace(*fetchRequest.ParentID))
        if len(p) > 0 {
            parentID = &p
        }
    }

    if parentID == nil {

        rows, error = config.DB.Query(
            `select ` + kScanFeedRowString +
            `   from FeedPostTable
                where postStatus = $1
                  and parentID is null
                  and timeActiveStart <= current_timestamp
                  and timeActiveStop   > current_timestamp
                order by timestamp desc;`,
            BlitzMessage.FeedPostStatus_FPSActive)

    } else {

        rows, error = config.DB.Query(
            `select ` + kScanFeedRowString +
            `   from FeedPostTable
                where postStatus = $1
                  and parentID = $2
                order by timestamp desc;`,
            BlitzMessage.FeedPostStatus_FPSActive,
            parentID)

    }

    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    feedPosts := make([]*BlitzMessage.FeedPost, 0, 10)
    for rows.Next() {
        feedPost, error := ScanFeedPostRowForUserID(session.UserID, rows)
        if error != nil {
            Log.LogError(error)
        } else {
            feedPosts = append(feedPosts, feedPost)
        }
    }

    //  Now go back through the feed posts to update their responses:

    for _, feedPost := range feedPosts {
        var limit int = 0

        switch *feedPost.PostType {

        case BlitzMessage.FeedPostType_FPOpenEndedQuestion:
            limit = 6
            feedPost.Replies = FetchTopOpenRepliesForFeedPost(session.UserID, *feedPost.PostID, limit)

        case BlitzMessage.FeedPostType_FPSurveyQuestion:
            limit = 10
            feedPost.Replies = FetchTopSurveyRepliesForFeedPost(session.UserID, *feedPost.PostID, limit)

        }

        if len(feedPost.Replies) >= limit && len(feedPost.Replies) > 0 {
            feedPost.AreMoreReplies = BoolPtr(true)
            feedPost.Replies = feedPost.Replies[:len(feedPost.Replies)-1]
        } else {
            feedPost.AreMoreReplies = BoolPtr(false)
        }

    }

    Log.Debugf("Found %d feed posts.", len(feedPosts))
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

