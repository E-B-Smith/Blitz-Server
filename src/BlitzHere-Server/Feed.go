

//----------------------------------------------------------------------------------------
//
//                                                              BlitzHere-Server : Feed.go
//                                                                       Feed maintainence
//
//                                                                 E.B. Smith, March, 2015
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "errors"
    "strings"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
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
            surveyAnswerSequence,
            amountTotal
        ) = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) where postID = $14;`,
            feedPost.ParentID,
            feedPost.PostType,
            feedPost.PostScope,
            feedPost.UserID,
            feedPost.AnonymousPost,
            feedPost.TimespanActive.NullTimeStart(),
            feedPost.TimespanActive.NullTimeStop(),
            feedPost.HeadlineText,
            feedPost.BodyText,
            feedPost.MayAddReply,
            feedPost.MayChooseMulitpleReplies,
            feedPost.SurveyAnswerSequence,
            feedPost.AmountTotal,
            feedPost.PostID,
        )

    error = pgsql.UpdateResultError(result, error)
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
`   FeedPostTable.postID,
    FeedPostTable.parentID,
    FeedPostTable.postType,
    FeedPostTable.postScope,
    FeedPostTable.userID,
    FeedPostTable.anonymousPost,
    FeedPostTable.timestamp,
    FeedPostTable.timeActiveStart,
    FeedPostTable.timeActiveStop,
    FeedPostTable.headlineText,
    FeedPostTable.bodyText,
    FeedPostTable.mayAddReply,
    FeedPostTable.mayChooseMulitpleReplies,
    FeedPostTable.surveyAnswerSequence,
    FeedPostTable.amountPerReply,
    FeedPostTable.amountTotal
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
        surveyAnswerSequence        sql.NullInt64
        amountPerReply  sql.NullString
        amountTotal     sql.NullString
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
        &surveyAnswerSequence,
        &amountPerReply,
        &amountTotal,
    )
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
        Timestamp:          BlitzMessage.TimestampPtr(timestamp),
        TimespanActive:     BlitzMessage.TimespanFromNullTimes(timeActiveStart, timeActiveStop),
        HeadlineText:       StringPtrFromNullString(headlineText),
        BodyText:           StringPtrFromNullString(bodyText),
        MayAddReply:        BoolPtrFromNullBool(mayAddReply),
        MayChooseMulitpleReplies:   BoolPtrFromNullBool(mayChooseMulitpleReplies),
        SurveyAnswerSequence:       Int32PtrFromNullInt64(surveyAnswerSequence),
//      AmountPerReply:     StringPtrFromNullString(amountPerReply),
        AmountTotal:        StringPtrFromNullString(amountTotal),
    }

    feedPost.PostTags = GetEntityTagsWithUserID(queryUserID, *feedPost.PostID, BlitzMessage.EntityType_ETFeedPost)

    var panelRows *sql.Rows
    panelRows, error = config.DB.Query(
        `select
             memberID
            ,bountyAmount
            ,dateAnswered
            from FeedPanelTable
            where postID = $1;`,
        postID,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    for panelRows.Next() {
        var (
            memberID        string
            bountyAmount    sql.NullString
            dateAnswered    pq.NullTime
        )
        error = panelRows.Scan(
            &memberID,
            &bountyAmount,
            &dateAnswered,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }
        member := BlitzMessage.FeedPanelMember {
            UserID:         proto.String(memberID),
            BountyAmount:   proto.String(bountyAmount.String),
            DateAnswered:   BlitzMessage.TimestampPtr(dateAnswered.Time),
        }
        feedPost.Panel = append(feedPost.Panel, &member)
    }
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
//
//                                                                          UpdateFeedPost
//
//----------------------------------------------------------------------------------------


//------------------------------------------------------------------------- DeleteFeedPost


func DeleteFeedPost(session *Session, postID string) error {
    Log.LogFunctionName()

    _, error := config.DB.Exec(
        `update FeedPostTable set PostStatus = $1
            where postID = $2
              and userID = $3;`,
        BlitzMessage.FeedPostStatus_FPSDeleted,
        postID,
        session.UserID,
    )
    return error
}


//------------------------------------------------------------------------- UpdateFeedPost


func UpdateFeedPost(session *Session, feedPost *BlitzMessage.FeedPost) error {
    Log.LogFunctionName()

    if  feedPost == nil ||
        feedPost.UserID == nil ||
        session.UserID != *feedPost.UserID {
        return errors.New("Not authorized")
    }

    _, error := config.DB.Exec(
        `update FeedPostTable set
             postScope                  = $1
            ,anonymousPost              = $2
            ,timeActiveStart            = $3
            ,timeActiveStop             = $4
            ,headlineText               = $5
            ,bodyText                   = $6
            ,mayAddReply                = $7
            ,mayChooseMulitpleReplies   = $8
        where postID = $9
          and userID = $10;`,
        feedPost.PostScope,
        feedPost.AnonymousPost,
        feedPost.TimespanActive.NullTimeStart(),
        feedPost.TimespanActive.NullTimeStop(),
        feedPost.HeadlineText,
        feedPost.BodyText,
        feedPost.MayAddReply,
        feedPost.MayChooseMulitpleReplies,
        feedPost.PostID,
        session.UserID,
    )
    return error
}


//------------------------------------------------------------------------- CreateFeedPost


func CreateFeedPost(session *Session, feedPost *BlitzMessage.FeedPost) error {
    Log.LogFunctionName()

    if  feedPost.UserID == nil ||
        session.UserID != *feedPost.UserID ||
        feedPost.HeadlineText == nil {
        return errors.New("Not authorized")
    }

    error := WriteFeedPost(feedPost)
    if error != nil {
        Log.LogError(error)
        return error
    }

    //  If there's a panel, update the tags --

    for _, panelMember := range feedPost.Panel {
        _, error = config.DB.Exec(
            `insert into EntityTagTable (
                userID,
                entityID,
                entityType,
                entityTag
            ) values ($1, $2, $3, $4);`,
            panelMember.UserID,
            feedPost.PostID,
            BlitzMessage.EntityType_ETFeedPost,
            ".panel",
        )
        if error != nil  {
            Log.LogError(error)
        }

        _, error = config.DB.Exec(
            `insert into FeedPanelTable (
                postID,
                memberID,
                bountyAmount
            ) values ($1, $2, $3);`,
            feedPost.PostID,
            panelMember.UserID,
            panelMember.BountyAmount,
        )
        if error != nil {
            Log.LogError(error)
        }
    }
    Log.Debugf("Added %d panel members.", len(feedPost.Panel))

    //  Save the panel --

    //  Send a notification if it's a response --

    if  feedPost.ParentID != nil {
        Log.Debugf("Try to send a notification to the original poster:")
        actionURL := fmt.Sprintf("%s?action=showpost&postid=%s",
            config.AppLinkURL, *feedPost.ParentID)
        parentPost := FeedPostForPostID(session.UserID, *feedPost.ParentID)
        if  parentPost != nil {
            name, _ := NameForUserID(session.UserID)
            if len(name) == 0 { name = "Someone" }
            message := fmt.Sprintf("%s responded to your post.", name)
            SendUserMessageInternal(*feedPost.UserID,
                [] string { *parentPost.UserID },
                "",
                message,
                BlitzMessage.UserMessageType_MTNotification,
                "AppIcon",
                actionURL,
            )
        }
    }

    //  Send a notification to the user's followers --

    followingUsers := GetUserIDArrayForEntity(
        BlitzMessage.EntityType_ETUser,
        *feedPost.UserID,
        ".followed",
    )
    postID := *feedPost.PostID
    if  feedPost.ParentID != nil {
        postID = *feedPost.ParentID
    }
    actionURL := fmt.Sprintf("%s?action=showpost&postid=%s",
        config.AppLinkURL, postID)

    name, _ := NameForUserID(*feedPost.UserID)
    if len(followingUsers) > 0 && len(name) > 0 {
        message := *feedPost.HeadlineText
        SendUserMessageInternal(*feedPost.UserID,
            followingUsers,
            "",
            message,
            BlitzMessage.UserMessageType_MTNotification,
            "AppIcon",
            actionURL,
        )
    }

    return nil
}


//-------------------------------------------------------------------- UpdateFeedPostBatch


func UpdateFeedPostBatch(session *Session, feedPostUpdate *BlitzMessage.FeedPostUpdateRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  feedPostUpdate.UpdateVerb == nil ||
        len(feedPostUpdate.FeedPosts) == 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    var error error
    switch *feedPostUpdate.UpdateVerb  {

    case BlitzMessage.UpdateVerb_UVCreate:

        for _, feedPost := range feedPostUpdate.FeedPosts {
            error = CreateFeedPost(session, feedPost)
            if error != nil {
                Log.LogError(error)
                break
            }
        }

    case BlitzMessage.UpdateVerb_UVUpdate:

        for _, feedPost := range feedPostUpdate.FeedPosts {
            error = UpdateFeedPost(session, feedPost)
            if error != nil {
                Log.LogError(error)
                break
            }
        }

    case BlitzMessage.UpdateVerb_UVDelete:

        for _, feedPost := range feedPostUpdate.FeedPosts {
            error = DeleteFeedPost(session, *feedPost.PostID)
            if error != nil {
                Log.LogError(error)
                break
            }
        }

    default:
        Log.Errorf("Invalid case: %d.", *feedPostUpdate.UpdateVerb)
        error = errors.New("Invalid verb")
    }

    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    //  Read & return the updated posts --

    Log.Debugf("Reading updated feed posts.")
    updateArray := make([]*BlitzMessage.FeedPost, 0, len(feedPostUpdate.FeedPosts))
    for _, feedPost := range feedPostUpdate.FeedPosts {
        updatedPost := FeedPostForPostID(session.UserID, *feedPost.PostID)
        if updatedPost != nil {
            updateArray = append(updateArray, updatedPost)
        }
    }

    feedResponse := BlitzMessage.FeedPostResponse {
        FeedPosts:      updateArray,
    }
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FeedPostResponse: &feedResponse },
    }
    return response
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
              and surveyAnswerSequence is not null
              and surveyAnswerSequence <> 0
              order by surveyAnswerSequence, timestamp
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

    //  Get the global feed posts --

    if parentID == nil {

        // rows, error = config.DB.Query(
        //     `select ` + kScanFeedRowString +
        //     `   from FeedPostTable
        //         where postStatus = $1
        //           and (parentID is null or parentID = postID)
        //           and timeActiveStart <= current_timestamp
        //           and timeActiveStop   > current_timestamp
        //           and amountTotal is null
        //         order by timestamp desc;`,
        //     BlitzMessage.FeedPostStatus_FPSActive)

        rows, error = config.DB.Query(
            `select `  + kScanFeedRowString +
            `   from FeedPostTable
                left join EntityTagTable  ett
                    on (ett.userID = $1
                      and ett.entityTag = '.panel'
                      and ett.entityType = $2
                      and ett.entityID = FeedPostTable.postID)
                where postStatus = $3
                  and (parentID is null or parentID = postID)
                  and timeActiveStart <= current_timestamp
                  and timeActiveStop   > current_timestamp
                  and (postType = $4
                        or entityTag is not null
                        or FeedPostTable.userID = $1)
                order by timestamp desc;`,
            session.UserID,
            BlitzMessage.EntityType_ETFeedPost,
            BlitzMessage.FeedPostStatus_FPSActive,
            BlitzMessage.FeedPostType_FPWantedQuestion,
        )

    } else {

        rows, error = config.DB.Query(
            `select ` + kScanFeedRowString +
            `   from FeedPostTable
                where (postID = $1
                   or parentID = $1)
                  and postStatus <> $2
                order by timestamp desc;`,
            parentID,
            BlitzMessage.FeedPostStatus_FPSDeleted,
        )

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

        var limit int = 20
        var replies []*BlitzMessage.FeedPost

        switch *feedPost.PostType {

        case BlitzMessage.FeedPostType_FPOpenEndedQuestion:
            replies = FetchTopOpenRepliesForFeedPost(session.UserID, *feedPost.PostID, limit)

        case BlitzMessage.FeedPostType_FPSurveyQuestion:
            replies = FetchTopSurveyRepliesForFeedPost(session.UserID, *feedPost.PostID, limit)

        }
        feedPosts = append(feedPosts, replies...)

    }

    Log.Debugf("Found %d feed posts.", len(feedPosts))
    feedResponse := BlitzMessage.FeedPostResponse {
        FeedPosts:      feedPosts,
    }
    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseType:       &BlitzMessage.ResponseType { FeedPostResponse: &feedResponse },
    }

    return response
}

