

//----------------------------------------------------------------------------------------
//
//                                                            BlitzHere-Server : Review.go
//                                                                     Review maintainence
//
//                                                                 E.B. Smith, April, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "time"
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
//                                                                            Write Review
//
//----------------------------------------------------------------------------------------

/*
message UserReview {
  optional string       userID          = 1;
  optional string       reviewerID      = 2;
  optional Timestamp    timestamp       = 3;
  optional string       conversationID  = 4;
  optional double       responsiveness  = 5;
  optional double       satisfaction    = 6;
  optional double       recommended     = 7;
  optional string       reviewText      = 8;
  repeated string       tags            = 9;
}
*/

func WriteReview(session *Session, review *BlitzMessage.UserReview,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  review.UserID == nil ||
        review.ReviewerID == nil ||
        review.ConversationID == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Missing fields"))
    }

    if *review.ReviewerID != session.UserID {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCNotAuthorized, fmt.Errorf("Not authorized"))
    }

    if review.Timestamp == nil {
        review.Timestamp = BlitzMessage.TimestampPtr(time.Now())
    }

    var reviewText *string
    if review.ReviewText != nil {
        s := strings.TrimSpace(*review.ReviewText)
        if len(s) > 0 { reviewText = &s }
    }

    result, error := config.DB.Exec(
        `insert into ReviewTable(
            userID,
            reviewerID,
            timestamp,
            conversationID,
            responsetime,
            responsive,
            outgoing,
            recommended,
            reviewText,
            tags
        ) values (
            $1::UserID, $2, $3, $4::uuid,
            ResponseTimeForConversationUser($4::text, $1::text),
            $5, $6, $7, $8, $9
        );`,
        review.UserID,
        review.ReviewerID,
        review.Timestamp.NullTime(),
        review.ConversationID,
        review.Responsive,
        review.Outgoing,
        review.Recommended,
        reviewText,
        pgsql.NullStringFromStringArray(review.Tags),
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    error = CloseConversationID(*review.ConversationID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    return ServerResponseForError(BlitzMessage.ResponseCode_RCSuccess, nil)
}


//----------------------------------------------------------------------------------------
//
//                                                                     AddReviewsToProfile
//
//----------------------------------------------------------------------------------------


func AddReviewsToProfile(profile *BlitzMessage.UserProfile) error {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select
            reviewerID,
            timestamp,
            responseTime,
            responsive,
            outgoing,
            recommended,
            reviewText
                from ReviewTable
                where userID = $1;`,
        profile.UserID,
    )
    if error != nil {
        Log.LogError(error)
        return error
    }
    defer rows.Close()

    var ratingCount int32
    var responsiveT, outgoingT, recommendedT, responseT float64
    var responsiveC, outgoingC, recommendedC, responseC float64
    reviews := make([]*BlitzMessage.UserReview, 0)

    for rows.Next() {
        ratingCount++
        var (
            reviewerID      string
            timestamp       pq.NullTime
            responseTime    sql.NullFloat64
            responsive      sql.NullFloat64
            outgoing        sql.NullFloat64
            recommended     sql.NullFloat64
            reviewText      sql.NullString
        )

        error = rows.Scan(
            &reviewerID,
            &timestamp,
            &responseTime,
            &responsive,
            &outgoing,
            &recommended,
            &reviewText,
        )
        if error != nil {
            Log.LogError(error)
            continue
        }

        if responsive.Valid && responsive.Float64 > 0 {
            responsiveT += responsive.Float64
            responsiveC += 1.0
        }

        if outgoing.Valid && outgoing.Float64 > 0 {
            outgoingT += outgoing.Float64
            outgoingC += 1.0
        }

        if recommended.Valid && recommended.Float64 > 0 {
            recommendedT += recommended.Float64
            recommendedC += 1.0
        }

        if responseTime.Valid && responseTime.Float64 > 0 {
            responseT += responseTime.Float64
            responseC += 1.0
        }

        s := strings.TrimSpace(reviewText.String)
        if reviewText.Valid && len(s) > 0 {
            r := &BlitzMessage.UserReview{
                ReviewerID:     &reviewerID,
                Timestamp:      BlitzMessage.TimestampPtr(timestamp),
                ReviewText:     &s,
            }
            reviews = append(reviews, r)
        }
    }

    profile.RatingCount         = proto.Int32(ratingCount)
    profile.RatingResponsive    = proto.Float64(0.0)
    profile.RatingOutgoing      = proto.Float64(0.0)
    profile.RatingRecommended   = proto.Float64(0.0)
    profile.ResponseSeconds     = proto.Float64(0.0)

    if responsiveC > 0.0  { profile.RatingResponsive    = proto.Float64(responsiveT / responsiveC) }
    if outgoingC > 0.0    { profile.RatingOutgoing      = proto.Float64(outgoingT / outgoingC) }
    if recommendedC > 0.0 { profile.RatingRecommended   = proto.Float64(recommendedT / recommendedC) }
    if responseC > 0.0    { profile.ResponseSeconds     = proto.Float64(responseT / responseC) }

    profile.Reviews = reviews
    return nil
}

