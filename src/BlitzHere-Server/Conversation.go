

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

    result, error := config.DB.Exec(
        `insert into ConversationTable
            (conversationID, status, initiatorUserID, parentFeedPostID, creationDate, closedDate, paymentStatus)
            values ($1, $2, $3, $4, current_timestamp, $5, $6)
         on conflict(conversationID) do
            update set (status, parentFeedPostID, closedDate) = ($2, $4, $5);`,
        conv.ConversationID,
        conv.Status,
        conv.InitiatorUserID,
        conv.ParentFeedPostID,
        conv.ClosedDate,
        conv.PaymentStatus,
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
        conversationID)
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
            initiatorUserID,
            parentFeedPostID,
            creationDate,
            closedDate,
            paymentStatus,
            chargeID
                from ConversationTable
                where conversationID = $1;`, conversationID)

    var (
        status              sql.NullInt64
        initiatorUserID     sql.NullString
        parentFeedPostID    sql.NullString
        creationDate        pq.NullTime
        closedDate          pq.NullTime
        paymentStatus       sql.NullInt64
        chargeID            sql.NullString
    )

    error := row.Scan(
        &conversationID,
        &status,
        &initiatorUserID,
        &parentFeedPostID,
        &creationDate,
        &closedDate,
        &paymentStatus,
        &chargeID,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    var conv = BlitzMessage.Conversation {
        ConversationID:     &conversationID,
        Status:             BlitzMessage.UserMessageStatus(status.Int64).Enum(),
        InitiatorUserID:    &initiatorUserID.String,
        ParentFeedPostID:   StringPtr(parentFeedPostID.String),
        CreationDate:       BlitzMessage.TimestampPtr(creationDate),
        ClosedDate:         BlitzMessage.TimestampPtr(closedDate),
        ConversationType:   BlitzMessage.ConversationType(BlitzMessage.ConversationType_CTConversation).Enum(),
        PaymentStatus:      BlitzMessage.PaymentStatus(paymentStatus.Int64).Enum(),
        ChargeID:           proto.String(chargeID.String),
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
//                                                                       StartConversation
//----------------------------------------------------------------------------------------


func StartConversation(session *Session, req *BlitzMessage.ConversationRequest) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    //  Check the members --

    memberMap := make(map[string]bool)
    memberMap[session.UserID] = true
    for _, memID := range req.UserIDs {
        memberMap[memID] = true
    }

    if len(memberMap) < 2 {
        error := fmt.Errorf("Not enough conversation members")
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    idx := 0
    memberArray := make([]string, len(memberMap))
    for memID, _ := range memberMap {
        memberArray[idx] = memID
        idx++
    }

    //  Check for an existing conversation --

    row := config.DB.QueryRow(
        `select A.conversationID
            from ConversationMemberTable A
            inner join ConversationMemberTable B
               on A.conversationID = B.conversationID
            inner join ConversationTable C
               on A.conversationID = C.conversationID
            where A.memberID = $1
              and B.memberID = $2
              and C.closedDate is null
         order by C.creationDate asc
            limit 1;`,
        memberArray[0],
        memberArray[1],
    )

    var conversationID string
    var conversation *BlitzMessage.Conversation
    error := row.Scan(&conversationID)
    if error != nil {

        var introMessage string

        //  Create a new conversation --

        conversationID = Util.NewUUIDString()
        Log.Debugf("Find existing error was +%v'.", error)
        Log.Debugf("Creating new conversation '%s'.", conversationID)
        conversation = &BlitzMessage.Conversation {
            ConversationID:     &conversationID,
            InitiatorUserID:    &session.UserID,
            Status:             BlitzMessage.UserMessageStatus(BlitzMessage.UserMessageStatus_MSNew).Enum(),
            ParentFeedPostID:   req.ParentFeedPostID,
            MemberIDs:          memberArray,
        }

        if  config.ServiceIsFree {
            conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
            Log.Debugf("Conversation is free: Service is free.")
            introMessage = "Blitz is in free mode.\nEnjoy you chat."
        }

        //  Friends?

        var otherMember string
        for _, member := range conversation.MemberIDs {
            if member != session.UserID {
                otherMember = member
                break
            }
        }

        tags := GetEntityTagMapForUserIDEntityIDType(
            session.UserID,
            otherMember,
            BlitzMessage.EntityType_ETUser,
        )
        if val, ok := tags[kTagFriend]; ok && val {
            Log.Debugf("Conversation is between friends.")
            conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
            if len(introMessage) <= 0 {
                introMessage = "Chat between friends is free.\nEnjoy your chat."
            }
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
            conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
            if len(introMessage) <= 0 {
                introMessage = "This chat is free."
            }
        }

        //  Other user is expert?

        row = config.DB.QueryRow(
            `select isExpert from UserTable where userID = $1;`,
            otherMember,
        )
        var otherIsExpert sql.NullBool
        error = row.Scan(&otherIsExpert)
        if error != nil { Log.LogError(error) }

        if isExpert.Bool {
            if otherIsExpert.Bool {
                conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
                if len(introMessage) <= 0 {
                    introMessage =
                        "As Blitz experts, you have unrestricted access to chat with other experts.\n"+
                        "Please state your objective and provide time for the expert to respond."
                }
            } else {
                conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
                if len(introMessage) <= 0 {
                    introMessage =
                    "A Blitz expert would like to chat with you!\nThis chat session will remain open for the next 5 days."
                }
            }
        } else {
            if otherIsExpert.Bool {
                if len(introMessage) <= 0 {
                    introMessage =
                        "You can send one free message to an expert.\nPlease state your objective and provide "+
                        "time for the expert to respond.\nFor guaranteeing a response, please make a payment "+
                        "– Good luck!"
                }
            } else {
                conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSIsFree).Enum()
                if len(introMessage) <= 0 {
                    introMessage = "Chat with non-experts is free.\nEnjoy your chat."
                }
            }
        }

        if  conversation.PaymentStatus != nil  &&
            *conversation.PaymentStatus == BlitzMessage.PaymentStatus_PSIsFree {
            Log.Debugf("Conversation is free.")
        } else {
            conversation.PaymentStatus = BlitzMessage.PaymentStatus(BlitzMessage.PaymentStatus_PSTrialPeriod).Enum()
            Log.Debugf("Conversation is paid.")
        }

        error = WriteConversation(conversation)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }

        //  Send a system message to the participants --

        if len(introMessage) <= 0 {
            introMessage = fmt.Sprintf("Chat with %s and %s.",
                PrettyNameForUserID(memberArray[0]),
                PrettyNameForUserID(memberArray[1]),
            )
        }

        error = SendUserMessageInternal(
            BlitzMessage.Default_Global_SystemUserID,
            memberArray,
            conversationID,
            introMessage,
            BlitzMessage.UserMessageType_MTConversation,
            "",
            "",
        )
        if error != nil { Log.LogError(error) }
    }

    conversation, error = ReadUserConversation(session.UserID, conversationID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    profiles := make([]*BlitzMessage.UserProfile, 0, 3)
    for _, memberID := range memberArray {
        profile := ProfileForUserID(session.UserID, memberID)
        if profile != nil {
            profiles = append(profiles, profile)
        }
    }

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
        conv.InitiatorUserID    =   &userID.String
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
        conv.InitiatorUserID    =   &senderID.String
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


func UpdateConversationStatus(session *Session, updateStatus *BlitzMessage.UpdateConversationStatus,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if  updateStatus.ConversationID == nil ||
        updateStatus.Status == nil ||
        updateStatus.ConversationType == nil ||
        *updateStatus.Status != BlitzMessage.UserMessageStatus_MSRead {
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
            initiatorUserID,
            closedDate,
            paymentStatus,
            chargeID
                from ConversationTable
                where conversationID = $1;`, purchase.PurchaseTypeID)

    var (
        conversationID      sql.NullString
        initiatorUserID     sql.NullString
        closedDate          pq.NullTime
        paymentStatus       sql.NullInt64
        chargeID            sql.NullString
    )

    error = row.Scan(
        &conversationID,
        &initiatorUserID,
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

    if initiatorUserID.String != session.UserID {
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
            callCharge
            from UserTable where userID = $1;`,
        expertID,
    )
    var name sql.NullString
    var chatCharge, callCharge sql.NullFloat64
    error = row.Scan(&chatCharge, &callCharge)
    if error != nil {
        Log.LogError(error)
        return errors.New("Expert not available")
    }

    purchase.Amount   = proto.String(fmt.Sprintf("%1.02f", chatCharge.Float64))
    purchase.Currency = proto.String("usd")
    purchase.MemoText = proto.String(fmt.Sprintf("Chat with %s to get expert views and opinions.", name.String))

    return nil
}

