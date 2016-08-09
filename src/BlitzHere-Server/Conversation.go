

//----------------------------------------------------------------------------------------
//
//                                                      BlitzHere-Server : Conversation.go
//                                                                           Conversations
//
//                                                                 E.B. Smith, April, 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "time"
    "errors"
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "BlitzMessage"
)


func WriteConversation(conv *BlitzMessage.Conversation) error {
    Log.LogFunctionName()

    suggestedDates := make([]time.Time, 0)
    for _, btime := range conv.SuggestedDates {
        suggestedDates = append(suggestedDates, btime.Time())
    }
    suggestedDatesString := pgsql.NullStringFromTimeArray(suggestedDates)

    result, error := config.DB.Exec(
        `insert into ConversationTable (
            conversationID,
            status,
            conversationType,
            initiatorID,
            parentFeedPostID,
            creationDate,
            closedDate,
            paymentStatus,
            expertID,
            topic,
            callDate,
            suggestedDuration,
            suggestedDates
        ) values (
            $1, $2, $3, $4, $5, current_timestamp, $6, $7, $8, $9, $10, $11, $12
        )
        on conflict(conversationID) do update set (
            status,
            closedDate,
            callDate
        ) = ($2, $6, $10);`,
        conv.ConversationID,
        conv.Status,
        conv.ConversationType,
        conv.InitiatorID,
        conv.ParentFeedPostID,
        conv.ClosedDate,
        conv.PaymentStatus,
        conv.ExpertID,
        conv.Topic,
        conv.CallDate,
        conv.SuggestedDuration,
        suggestedDatesString,
    )
    Log.Debugf("Conversation Create status: %v.", error)

    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return error
    }

    _, error = config.DB.Exec(
        `delete from ConversationMemberTable where conversationID = $1;`,
        conv.ConversationID)
    if error != nil {
        Log.LogError(error)
        return error
    }

    for _, memberID := range conv.MemberIDs {
        Log.Debugf("Conversation %s adding %s", *conv.ConversationID, memberID)
        result, error = config.DB.Exec(
            `insert into ConversationMemberTable
                (conversationID, memberID) values ($1, $2)
                on conflict do nothing;`,
            conv.ConversationID,
            memberID)
        error = pgsql.UpdateResultError(result, error)
        if error != nil { Log.LogError(error) }
    }

    return nil
}


func MembersForConversationID(conversationID string) []string {
    Log.LogFunctionName()

    members := make([]string, 0, 5)

    rows, error := config.DB.Query(
        `select memberID from ConversationMemberTable where conversationID = $1;`,
        conversationID,
    )
    if error != nil {
        Log.LogError(error)
        return members
    }
    defer rows.Close()

    for rows.Next() {
        var member string
        error = rows.Scan(&member)
        if error != nil {
            Log.LogError(error)
        } else {
            members = append(members, member)
        }
    }

    return members
}


func ReadUserConversation(userID string, conversationID string) (*BlitzMessage.Conversation, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
            conversationID,
            status,
            initiatorID,
            parentFeedPostID,
            creationDate,
            closedDate,
            paymentStatus,
            chargeID,
            acceptDate,
            expertID,
            conversationType,
            topic,
            callDate,
            extract(epoch from suggestedDuration),
            suggestedDates
                from ConversationTable
                where conversationID = $1;`,
        conversationID,
    )

    var (
        conversationIDX     sql.NullString
        status              sql.NullInt64
        initiatorID         sql.NullString
        parentFeedPostID    sql.NullString
        creationDate        pq.NullTime
        closedDate          pq.NullTime
        paymentStatus       sql.NullInt64
        chargeID            sql.NullString
        acceptDate          pq.NullTime
        expertID            sql.NullString
        conversationType    sql.NullInt64
        topic               sql.NullString
        callDate            pq.NullTime
        suggestedDuration   sql.NullFloat64
        suggestedDatesString sql.NullString
    )

    error := row.Scan(
        &conversationIDX,
        &status,
        &initiatorID,
        &parentFeedPostID,
        &creationDate,
        &closedDate,
        &paymentStatus,
        &chargeID,
        &acceptDate,
        &expertID,
        &conversationType,
        &topic,
        &callDate,
        &suggestedDuration,
        &suggestedDatesString,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    var conv = BlitzMessage.Conversation {
        ConversationID:     &conversationID,
        Status:             BlitzMessage.UserMessageStatus(status.Int64).Enum(),
        InitiatorID:        &initiatorID.String,
        ParentFeedPostID:   StringPtr(parentFeedPostID.String),
        CreationDate:       BlitzMessage.TimestampPtr(creationDate),
        ClosedDate:         BlitzMessage.TimestampPtr(closedDate),
        PaymentStatus:      BlitzMessage.PaymentStatus(paymentStatus.Int64).Enum(),
        ChargeID:           proto.String(chargeID.String),
        AcceptDate:         BlitzMessage.TimestampPtr(acceptDate),
        ExpertID:           proto.String(expertID.String),
        ConversationType:   BlitzMessage.ConversationType(conversationType.Int64).Enum(),
        Topic:              proto.String(topic.String),
        CallDate:           BlitzMessage.TimestampPtr(callDate),
        SuggestedDuration:  proto.Float64(suggestedDuration.Float64),
    }
    suggestedDates := pgsql.TimeArrayFromNullString(&suggestedDatesString)
    for _, time := range suggestedDates {
        conv.SuggestedDates = append(conv.SuggestedDates, BlitzMessage.TimestampPtr(time))
    }

    row = config.DB.QueryRow(
        `select
            count(*),
            sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
            from usermessagetable
            where conversationID = $1
              and recipientID = $2
            group by conversationID;`, conversationID, userID)
    var replyCount, unreadCount sql.NullInt64
    error = row.Scan(&replyCount, &unreadCount)
    if error != nil { Log.LogError(error) }

    conv.MessageCount = Int32PtrFromInt64(replyCount.Int64)
    conv.UnreadCount  = Int32PtrFromInt64(unreadCount.Int64)

    var rows *sql.Rows
    rows, error = config.DB.Query(
        `with msgs as (
         select messageText,
            creationDate,
            senderID,
            actionURL
            from usermessagetable
            where conversationID = $1
              and recipientID = $2
            order by creationDate desc
            limit 3
        )
        select * from msgs order by creationDate asc;`,
        conversationID,
        userID,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    defer rows.Close()

    for rows.Next() {
        var (
            lastMessage         sql.NullString
            lastActivity         pq.NullTime
            lastUserID          sql.NullString
            lastActionURL       sql.NullString
        )

        error = rows.Scan(&lastMessage, &lastActivity, &lastUserID, &lastActionURL)
        if error != nil {
            Log.LogError(error)
            continue
        }

        if lastActivity.Valid {
            conv.LastActivityDate = BlitzMessage.TimestampPtr(lastActivity)
        }

        if lastUserID.String == BlitzMessage.Default_Global_SystemUserID {
            conv.HeadlineText = &lastMessage.String
        } else {
            conv.LastMessage = &lastMessage.String
            conv.LastActivityUserID = &lastUserID.String
        }

        if lastActionURL.Valid && len(lastActionURL.String) > 0 {
            conv.LastActionURL = &lastActionURL.String
        }
    }

    if conv.LastActivityDate == nil {
        if conv.ClosedDate != nil {
            conv.LastActivityDate = conv.ClosedDate
        } else {
            conv.LastActivityDate = conv.CreationDate
        }
    }

    conv.MemberIDs = MembersForConversationID(conversationID)
    return &conv, nil
}


//----------------------------------------------------------------------------------------
//
//                                                                       StartConversation
//                                                            Start conversation functions
//
//----------------------------------------------------------------------------------------


//----------------------------------------------------------------------------------------
//                                                        PaymentStatusForChatConversation
//----------------------------------------------------------------------------------------


func PaymentStatusForChatConversation(
        session *Session,
        conversation *BlitzMessage.Conversation,
    ) (message string, paymentStatus *BlitzMessage.PaymentStatus) {
    Log.LogFunctionName()

    var error error

    if  config.ServiceIsFree {
        conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
        Log.Debugf("Conversation is free: Service is free.")
        conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
        return "Blitz is in free mode.\nEnjoy you chat.", conversation.PaymentStatus
    }

    //  Friends?

    tags := GetEntityTagMapForUserIDEntityIDType(
        session.UserID,
        *conversation.ExpertID,
        BlitzMessage.EntityType_ETUser,
    )
    if val, ok := tags[kTagFriend]; ok && val {
        Log.Debugf("Conversation is between friends.")
        conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
        return "Chat between friends is free.\nEnjoy your chat.", conversation.PaymentStatus
    }

    //  Free for user?

    row := config.DB.QueryRow(
        `select isFree, isExpert from UserTable where userID = $1;`,
        session.UserID,
    )
    var isFree, isExpert sql.NullBool
    error = row.Scan(&isFree, &isExpert)
    if error != nil { Log.LogError(error) }
    if isFree.Bool {
        Log.Debugf("Conversation is free for user.")
        conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
        return "This chat is free.", conversation.PaymentStatus
    }

    //  Other user is expert?

    row = config.DB.QueryRow(
        `select isExpert from UserTable where userID = $1;`,
        conversation.ExpertID,
    )
    var otherIsExpert sql.NullBool
    error = row.Scan(&otherIsExpert)
    if error != nil { Log.LogError(error) }

    memberName := PrettyNameForUserID(session.UserID)
    expertName := PrettyNameForUserID(*conversation.ExpertID)

    if isExpert.Bool {
        if otherIsExpert.Bool {
            conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
            return  "As Blitz experts, you have unrestricted access to chat with other experts.\n"+
                    "Please state your objective and provide time for the expert to respond.",
                    conversation.PaymentStatus
        } else {
            conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
            return "A Blitz expert would like to chat with you!\nThis chat session will remain open for the next 24 hours.",
                conversation.PaymentStatus
        }
    } else {
        if otherIsExpert.Bool {
            msg := fmt.Sprintf(
                    "%s – you have one free message\nto connect with %s.\n" +
                    "After this message, you'll be prompted to make\na payment " +
                    "to continue your chat with %s.",
                    memberName,
                    expertName,
                    expertName,
                )
            conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSTrialPeriod.Enum()
            return msg, conversation.PaymentStatus
        } else {
            conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSIsFree.Enum()
            return "Chat with non-experts is free.\nEnjoy your chat.", conversation.PaymentStatus
        }
    }
}


func StartChatConversation(
        session *Session,
        conversation *BlitzMessage.Conversation,
    ) error {

    var error error
    var introMessage string
    introMessage, conversation.PaymentStatus =
        PaymentStatusForChatConversation(session, conversation)

    if  conversation.PaymentStatus == nil {
        return errors.New("Server chat payment error")
    }

    error = WriteConversation(conversation)
    if error != nil {
        Log.LogError(error)
        return error
    }

    error = SendUserMessageInternal(
        BlitzMessage.Default_Global_SystemUserID,
        conversation.MemberIDs,
        *conversation.ConversationID,
        introMessage,
        BlitzMessage.UserMessageType_MTConversation,
        "",
        "",
    )
    if error != nil { Log.LogError(error) }

    return error
}


//----------------------------------------------------------------------------------------
//                                                                   StartCallConversation
//----------------------------------------------------------------------------------------


func StartCallConversation(
        session *Session,
        conversation *BlitzMessage.Conversation,
    ) error {

    var error error
    row := config.DB.QueryRow(
        `select isExpert from UserTable where userID = $1;`,
        conversation.ExpertID,
    )
    var otherIsExpert sql.NullBool
    error = row.Scan(&otherIsExpert)
    if error != nil { Log.LogError(error) }
    if ! (otherIsExpert.Valid && otherIsExpert.Bool) {
        error = errors.New("You can only talk with an expert")
        return error
    }

    conversation.Topic = Util.CleanStringPtr(conversation.Topic)
    if conversation.Topic == nil ||
        conversation.SuggestedDuration == nil ||
        len(conversation.SuggestedDates) == 0 {
        return errors.New("Missing fields")
    }

    conversation.PaymentStatus = BlitzMessage.PaymentStatus_PSTrialPeriod.Enum()
    error = WriteConversation(conversation)
    if error != nil {
        Log.LogError(error)
        return error
    }

    return nil
}


//----------------------------------------------------------------------------------------
//                                                                       StartConversation
//----------------------------------------------------------------------------------------


func StartConversation(
        session *Session,
        req *BlitzMessage.ConversationRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    conversation := req.Conversation
    if conversation == nil ||
       conversation.ConversationType == nil ||
       conversation.InitiatorID == nil ||
       conversation.ExpertID == nil ||
       conversation.ConversationType == nil ||
       *conversation.InitiatorID == *conversation.ExpertID {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Missing fields"))
    }

    //  Check for an existing conversation --

    row := config.DB.QueryRow(
        `select conversationID from conversationTable
            where (initiatorID = $1 or initiatorID = $2)
              and (expertID = $1 or expertID = $2)
              and conversationType = $3
              and closedDate is null
            order by creationDate
            limit 1;`,
        conversation.InitiatorID,
        conversation.ExpertID,
        conversation.ConversationType,
    )
    var conversationID string
    error := row.Scan(&conversationID)
    if error == nil {

        //  Found an existing conversation.  Return it:

        var response BlitzMessage.ConversationResponse
        response.Conversation, error = ReadUserConversation(session.UserID, conversationID)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
        }

        response.Profiles = make([]*BlitzMessage.UserProfile, 2)
        response.Profiles[0] = ProfileForUserID(session.UserID, *conversation.InitiatorID)
        response.Profiles[1] = ProfileForUserID(session.UserID, *conversation.ExpertID)

        serverResponse := &BlitzMessage.ServerResponse {
            ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
            ResponseType:       &BlitzMessage.ResponseType { ConversationResponse: &response },
        }
        return serverResponse
    }

    if session.UserID != *conversation.InitiatorID {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Can't start conversation for another user."))
    }

    //  Create a new conversation --

    conversationID = Util.NewUUIDString()
    Log.Debugf("Find existing error was +%v'.", error)
    Log.Debugf("Creating new conversation '%s'.", conversationID)
    conversation.ConversationID = proto.String(conversationID)
    conversation.Status = BlitzMessage.UserMessageStatus(BlitzMessage.UserMessageStatus_MSNew).Enum()
    conversation.MemberIDs = []string { *conversation.InitiatorID, *conversation.ExpertID }
    conversation.CreationDate = BlitzMessage.TimestampPtr(time.Now())
    conversation.LastActivityDate = conversation.CreationDate

    //  Figure out the payment status --

    switch *conversation.ConversationType {

    case BlitzMessage.ConversationType_CTConversation:
        error = StartChatConversation(session, conversation)

    case BlitzMessage.ConversationType_CTCall:
        error = StartCallConversation(session, conversation)

    default:
        error = fmt.Errorf("Invalid conversation type %d.", *conversation.ConversationType)
    }

    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    //  Send back a conversation --

    conversation, error = ReadUserConversation(session.UserID, conversationID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    profiles := make([]*BlitzMessage.UserProfile, 2)
    profiles[0] = ProfileForUserID(session.UserID, *conversation.InitiatorID)
    profiles[1] = ProfileForUserID(session.UserID, *conversation.ExpertID)

    response := BlitzMessage.ConversationResponse {
        Conversation:   conversation,
        Profiles:       profiles,
    }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { ConversationResponse: &response },
    }

    return serverResponse
}


//----------------------------------------------------------------------------------------
//                                                           FetchFeedPostsAsConversations
//----------------------------------------------------------------------------------------


func FetchFeedPostsAsConversations(userID string) []*BlitzMessage.Conversation {
    Log.LogFunctionName()

    resultArray := make([]*BlitzMessage.Conversation, 0, 10)

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

        var conv BlitzMessage.Conversation
        conv.InitiatorID        =   &userID.String
        conv.ParentFeedPostID   =   &postID.String
        conv.Status             =   BlitzMessage.UserMessageStatus(BlitzMessage.UserMessageStatus_MSRead).Enum()
        conv.CreationDate       =   BlitzMessage.TimestampPtr(createDate)
        conv.LastMessage        =   &replyText.String
        conv.MessageCount       =   Int32PtrFromNullInt64(replyCount)
        conv.HeadlineText       =   &headline.String
        conv.ConversationType   =   &conversationType

        if replyDate.Valid {
            conv.LastActivityDate = BlitzMessage.TimestampPtr(replyDate)
        } else {
            conv.LastActivityDate = BlitzMessage.TimestampPtr(createDate)
        }

        if replyUser.Valid {
            conv.LastActivityUserID = &replyUser.String
        }

        resultArray = append(resultArray, &conv)
    }

    Log.Debugf("Found %d feed posts.", len(resultArray))
    return resultArray
}


//----------------------------------------------------------------------------------------
//                                                       FetchNotificationsAsConversations
//----------------------------------------------------------------------------------------


func FetchNotificationsAsConversations(userID string) []*BlitzMessage.Conversation {
    Log.LogFunctionName()

    ary := make([]*BlitzMessage.Conversation, 0, 20)

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

        var conv BlitzMessage.Conversation
        conv.ConversationID     =   &messageID.String
        conv.InitiatorID        =   &senderID.String
        conv.Status             =   BlitzMessage.UserMessageStatus(status.Int64).Enum()
        conv.CreationDate       =   BlitzMessage.TimestampPtr(creationDate)
        conv.MessageCount       =   proto.Int32(1)
        conv.UnreadCount        =   proto.Int32(unreadCount)
        conv.LastMessage        =   &lastMessage.String
        conv.LastActivityDate   =   BlitzMessage.TimestampPtr(creationDate)
        conv.LastActionURL      =   &lastActionURL.String
        conv.ClosedDate         =   BlitzMessage.TimestampPtr(readDate)
        conv.MemberIDs          =   []string { senderID.String, recipientID.String }
        conv.ConversationType   =   &conversationType

        ary = append(ary, &conv)
    }
    return ary
}


//----------------------------------------------------------------------------------------
//                                                                      FetchConversations
//----------------------------------------------------------------------------------------


func FetchConversations(session *Session, req *BlitzMessage.FetchConversations) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var rows *sql.Rows
    var error error

    if len(req.UserID) > 0 {

        //  Fetch all conversations between a set of users (i.e., by member id)

        queryString :=
            `select conversationID from conversationmembertable `

        for _, uid := range req.UserID {
            queryString += fmt.Sprintf(
                `intersect select conversationID from conversationmembertable where memberID = '%s' `,
                uid,
            )
        }

        queryString += ";"
        Log.Debugf("Query is '%s'.", queryString)
        rows, error = config.DB.Query(queryString)

    } else {

        rows, error = config.DB.Query(
            `select conversationID from ConversationMemberTable where memberID = $1;`,
            session.UserID,
        )
    }

    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }
    defer rows.Close()

    convos := make([]*BlitzMessage.Conversation, 0, 10)
    for rows.Next() {
        var convID string
        error = rows.Scan(&convID)
        if error != nil {
            Log.LogError(error)
        } else {
            convo, error := ReadUserConversation(session.UserID, convID)
            if error == nil {
                convos = append(convos, convo)
            }
        }
    }

    if len(req.UserID) == 0 {
        convos = append(convos, FetchFeedPostsAsConversations(session.UserID)...)
        convos = append(convos, FetchNotificationsAsConversations(session.UserID)...)
    }

    response := BlitzMessage.FetchConversations { Conversations: convos }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FetchConversations: &response },
    }

    return serverResponse
}


//----------------------------------------------------------------------------------------
//                                                                UpdateConversationStatus
//----------------------------------------------------------------------------------------


func UpdateConversationPaymentStatus(
        session *Session,
        updateStatus *BlitzMessage.UpdateConversationStatus,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
            initiatorID,
            paymentStatus,
            closedDate
            from ConversationTable
                where conversationID = $1;`,
        updateStatus.ConversationID,
    )
    var (
        clientID    sql.NullString
        status      sql.NullInt64
        closedDate  pq.NullTime
    )
    error := row.Scan(&clientID, &status, &closedDate)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }

    var expertID string
    for _, mid := range MembersForConversationID(*updateStatus.ConversationID) {
        if mid != clientID.String {
            expertID = mid
            break
        }
    }
    if expertID != session.UserID ||
        BlitzMessage.PaymentStatus(status.Int64) !=
        BlitzMessage.PaymentStatus_PSExpertNeedsAccept {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }
    if closedDate.Valid {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("This conversation is closed"))
    }
    var callDate pq.NullTime
    if updateStatus.CallDate != nil {
        callDate.Time = updateStatus.CallDate.Time()
        callDate.Valid = true
    }
    if *updateStatus.ConversationType == BlitzMessage.ConversationType_CTCall &&
        (updateStatus.CallDate == nil || time.Since(callDate.Time) < 0) {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Must select a date in the future"))
    }
    var message string
    var acceptDate pq.NullTime
    expertName := PrettyNameForUserID(expertID)

    switch *updateStatus.PaymentStatus {

    case BlitzMessage.PaymentStatus_PSExpertAccepted:
        if *updateStatus.ConversationType == BlitzMessage.ConversationType_CTCall {
            message = fmt.Sprintf(
                "Congrats!  %s has accepted your scheduled call.",
                expertName,
            )
        } else {
            message = fmt.Sprintf(
                "Congrats!  %s has accepted your invitation.\nEnjoy your chat.",
                expertName,
            )
        }
        acceptDate.Time = time.Now()
        acceptDate.Valid = true

    case BlitzMessage.PaymentStatus_PSExpertRejected:
        closedDate.Valid = true
        closedDate.Time = time.Now()
        message = fmt.Sprintf(
            "Sorry, %s is unavailable now.",
            expertName,
        )

    default:
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }

    result, error := config.DB.Exec(
        `update ConversationTable set
            paymentStatus = $1,
            closedDate = $2,
            acceptDate = $3,
            callDate = $4
              where conversationID = $5;`,
        *updateStatus.PaymentStatus,
        closedDate,
        acceptDate,
        callDate,
        updateStatus.ConversationID,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    actionURL :=
        fmt.Sprintf("%s?action=showchat&chatid=%s",
            config.AppLinkURL,
            *updateStatus.ConversationID,
        )

    error = SendUserMessageInternal(
        BlitzMessage.Default_Global_SystemUserID,
        MembersForConversationID(*updateStatus.ConversationID),
        *updateStatus.ConversationID,
        message,
        BlitzMessage.UserMessageType_MTConversation,
        "",
        actionURL,
    )
    if error != nil {
        Log.LogError(error)
    }

    conversation, _ := ReadUserConversation(session.UserID, *updateStatus.ConversationID)
    response := BlitzMessage.ConversationResponse {
        Conversation:   conversation,
    }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { ConversationResponse: &response },
    }

    return serverResponse
}


func CloseConversationID(conversationID string) error {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `update ConversationTable set
            closedDate = current_timestamp
            where conversationID = $1;`,
        conversationID,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return error
    }

    //  Add a system to the participants --

    message := fmt.Sprintf("This conversation has been closed.")
    error = SendUserMessageInternal(
        BlitzMessage.Default_Global_SystemUserID,
        MembersForConversationID(conversationID),
        conversationID,
        message,
        BlitzMessage.UserMessageType_MTConversation,
        "",
        "",
    )
    if error != nil {
        Log.LogError(error)
    }
    return error
}


//----------------------------------------------------------------------------------------
//                                                                UpdateConversationStatus
//----------------------------------------------------------------------------------------


func UpdateConversationStatus(session *Session, updateStatus *BlitzMessage.UpdateConversationStatus,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  updateStatus.ConversationID == nil ||
        updateStatus.ConversationType == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }
    if  (*updateStatus.ConversationType == BlitzMessage.ConversationType_CTConversation ||
         *updateStatus.ConversationType == BlitzMessage.ConversationType_CTCall) &&
        updateStatus.PaymentStatus != nil {
        return UpdateConversationPaymentStatus(session, updateStatus)
    }
    if updateStatus.Status == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }
    if *updateStatus.Status == BlitzMessage.UserMessageStatus_MSClosed {
        error := CloseConversationID(*updateStatus.ConversationID)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCSuccess, error)
    }
    if *updateStatus.Status != BlitzMessage.UserMessageStatus_MSRead {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Invalid input"))
    }


    var error error
    switch (*updateStatus.ConversationType) {
    case BlitzMessage.ConversationType_CTConversation:
        _, error = config.DB.Exec(
            `update UserMessageTable set
                messageStatus = $3,
                readDate = transaction_timestamp()
                    where conversationID = $1
                      and recipientID = $2
                      and (messageStatus is null or messageStatus < $3);`,
            updateStatus.ConversationID,
            session.UserID,
            updateStatus.Status,
        )

    case BlitzMessage.ConversationType_CTNotification:
        _, error = config.DB.Exec(
            `update UserMessageTable set
                messageStatus = $3,
                readDate = transaction_timestamp()
                    where messageID = $1
                      and recipientID = $2;`,
            updateStatus.ConversationID,
            session.UserID,
            updateStatus.Status,
        )
    }

    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    conversation, _ := ReadUserConversation(session.UserID, *updateStatus.ConversationID)
    response := BlitzMessage.ConversationResponse {
        Conversation:   conversation,
    }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { ConversationResponse: &response },
    }

    return serverResponse
}


//----------------------------------------------------------------------------------------
//                                                UpdatePurchaseDescriptionForConversation
//----------------------------------------------------------------------------------------


func UpdatePurchaseDescriptionForConversation(session *Session, purchase *BlitzMessage.PurchaseDescription) error {
    Log.LogFunctionName()

    purchase.MemoText = nil
    purchase.Amount = nil
    purchase.Currency = nil

    if *purchase.PurchaseType != BlitzMessage.PurchaseType_PTChatConversation {
        return errors.New("Invalid Input")
    }

    var error error
    row := config.DB.QueryRow(
        `select
            conversationID,
            initiatorID,
            closedDate,
            paymentStatus,
            chargeID
                from ConversationTable
                where conversationID = $1;`, purchase.PurchaseTypeID)

    var (
        conversationID      sql.NullString
        initiatorID         sql.NullString
        closedDate          pq.NullTime
        paymentStatus       sql.NullInt64
        chargeID            sql.NullString
    )

    error = row.Scan(
        &conversationID,
        &initiatorID,
        &closedDate,
        &paymentStatus,
        &chargeID,
    )
    if error != nil {
        Log.LogError(error)
        return errors.New("No such conversation")
    }
    if closedDate.Valid {
        return errors.New("Conversation already closed")
    }

    switch BlitzMessage.PaymentStatus(paymentStatus.Int64) {

    case BlitzMessage.PaymentStatus_PSIsFree:
        return errors.New("Conversation is free")

    case BlitzMessage.PaymentStatus_PSUnknown,
         BlitzMessage.PaymentStatus_PSTrialPeriod,
         BlitzMessage.PaymentStatus_PSPaymentRequired:
        {}

    case BlitzMessage.PaymentStatus_PSExpertNeedsAccept:
        return errors.New("Already purchased. Waiting for expert")

    case BlitzMessage.PaymentStatus_PSExpertRejected:
        return errors.New("Expert unavailable. Purchase refunded")

    case BlitzMessage.PaymentStatus_PSExpertAccepted:
        return errors.New("Expert chat purchased and expert is available")

    }

    if initiatorID.String != session.UserID {
        return errors.New("Not buyer")
    }

    var expertID string
    members := MembersForConversationID(conversationID.String)
    for _, mid := range members {
        if mid != session.UserID {
            expertID = mid
            break
        }
    }
    if len(expertID) <= 0 {
        return errors.New("Expert not available")
    }

    row = config.DB.QueryRow(
        `select
            name,
            chatCharge,
            callCharge,
            isExpert
            from UserTable where userID = $1;`,
        expertID,
    )
    var name, chatCharge, callCharge sql.NullString
    var isExpert sql.NullBool
    error = row.Scan(
        &name,
        &chatCharge,
        &callCharge,
        &isExpert,
    )
    if error != nil || ! isExpert.Bool {
        Log.LogError(error)
        return errors.New("Expert not available")
    }

    purchase.Amount   = proto.String(chatCharge.String)
    purchase.Currency = proto.String("usd")
    purchase.MemoText = proto.String(fmt.Sprintf("Chat with %s to get expert views and opinions.", name.String))

    return nil
}



//----------------------------------------------------------------------------------------
//
//                                                                      ConversationCloser
//                                           Closes expired conversations.  Refunds Money.
//
//----------------------------------------------------------------------------------------


func CloseConversationIDTestMode(conversationID string) {
    //Log.Debugf("Would close conversation %s.", conversationID)
    CloseConversationID(conversationID)
}


func RefundChargeIDTestMode(chargeID string, memoText string) {
    //Log.Debugf("Would refund chargeID %s.", chargeID)
    RefundChargeID(chargeID, memoText)
}


func ConversationCloser() {
    Log.LogFunctionName()

    //  Close & refund non-accepted expert conversaations --
    Log.Debugf("Close non-accepted conversations...")

    rows, error := config.DB.Query(
        `select conversationID, chargeID from ConversationTable
            where paymentStatus > $1
              and paymentStatus < $2
              and closedDate is null
              and (now() - creationDate) >= (to_char($3::real, '999D999') || ' hours')::interval;`,
        BlitzMessage.PaymentStatus_PSIsFree,
        BlitzMessage.PaymentStatus_PSExpertAccepted,
        config.ChatLimitHours,
    )
    if error != nil {
        Log.LogError(error)
    }
    defer pgsql.CloseRows(rows)

    for rows != nil  && rows.Next() {
        var conversationID, chargeID sql.NullString
        error = rows.Scan(&conversationID, &chargeID)
        if error != nil {
            Log.LogError(error)
            continue
        }
        CloseConversationIDTestMode(conversationID.String)
        if chargeID.Valid {
            RefundChargeIDTestMode(chargeID.String, "Expert not available.")
        }
    }

    //  Close old open expert conversations --
    Log.Debugf("Close old paid expert conversations...")

    rows, error = config.DB.Query(
        `select conversationID from ConversationTable
            where paymentStatus = $1
              and closedDate is null
              and (now() - acceptDate) >= (to_char($2::real, '999D999') || ' hours')::interval;`,
        BlitzMessage.PaymentStatus_PSExpertAccepted,
        config.ChatLimitHours,
    )
    if error != nil {
        Log.LogError(error)
    }
    defer pgsql.CloseRows(rows)

    for rows != nil  && rows.Next() {
        var conversationID string
        error = rows.Scan(&conversationID)
        if error != nil {
            Log.LogError(error)
            continue
        }
        CloseConversationIDTestMode(conversationID)
    }

    //  Close expert to non-expert conversations --
    Log.Debugf("Close old expert to non-expert conversations...")

    rows, error = config.DB.Query(
        `select ct.conversationID from ConversationTable ct
            join UserTable ut on ut.userID = ct.initiatorID
            join ConversationMemberTable cmt on
                (cmt.conversationID = ct.conversationID and cmt.memberID <> ut.userID)
            join UserTable utx on utx.userID = cmt.memberID
            where closedDate is null
              and paymentStatus <= $1
              and ut.isExpert = true
              and (utx.isExpert = false or utx.isExpert is null)
              and (now() - ct.creationDate) >= (to_char($2::real, '999D999') || ' hours')::interval;`,
        BlitzMessage.PaymentStatus_PSIsFree,
        config.ChatLimitHours,
    )
    if error != nil {
        Log.LogError(error)
    }
    defer pgsql.CloseRows(rows)

    for rows != nil  && rows.Next() {
        var conversationID string
        error = rows.Scan(&conversationID)
        if error != nil {
            Log.LogError(error)
            continue
        }
        CloseConversationIDTestMode(conversationID)
    }
}

