

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
    "database/sql"
    "github.com/lib/pq"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


func WriteConversation(conv *BlitzMessage.Conversation) error {
    Log.LogFunctionName()

    result, error := config.DB.Exec(
        `insert into ConversationTable
            (conversationID, status, initiatorUserID, parentFeedPostID, creationDate, closedDate)
            values ($1, $2, $3, $4, current_timestamp, $5)
         on conflict(conversationID) do
            update set (status, parentFeedPostID, closedDate) = ($2, $4, $5);`,
        conv.ConversationID,
        conv.Status,
        conv.InitiatorUserID,
        conv.ParentFeedPostID,
        conv.ClosedDate,
    )
    Log.Debugf("Conversation Create status: %v.", error)

    error = pgsql.RowUpdateError(result, error)
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
        error = pgsql.RowUpdateError(result, error)
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
            closedDate
                from ConversationTable
                where conversationID = $1;`, conversationID)

    var (
        status              sql.NullInt64
        initiatorUserID     sql.NullString
        parentFeedPostID    sql.NullString
        creationDate        pq.NullTime
        closedDate          pq.NullTime

        replyCount          sql.NullInt64
        unreadCount         sql.NullInt64

        lastMessage         sql.NullString
        lastActivity        pq.NullTime
    )

    error := row.Scan(
        &conversationID,
        &status,
        &initiatorUserID,
        &parentFeedPostID,
        &creationDate,
        &closedDate,
    )
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    row = config.DB.QueryRow(
        `select
            count(*),
            sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
            from usermessagetable
            where conversationID = $1
              and recipientID = $2
            group by conversationID;`, conversationID, userID)
    error = row.Scan(&replyCount, &unreadCount)
    if error != nil { Log.LogError(error) }

    row = config.DB.QueryRow(
        `select messageText,
            creationDate
            from usermessagetable
            where conversationID = $1
            order by creationDate desc
            limit 1;`, conversationID)
    error = row.Scan(&lastMessage, &lastActivity)
    if error != nil { Log.LogError(error) }

    var conv BlitzMessage.Conversation
    conv.ConversationID     =   &conversationID
    conv.InitiatorUserID    =   &initiatorUserID.String
    conv.Status             =   BlitzMessage.UserMessageStatus(status.Int64).Enum()
    conv.CreationDate       =   BlitzMessage.TimestampPtrFromNullTime(creationDate)
    conv.MessageCount       =   Int32PtrFromNullInt64(replyCount)
    conv.UnreadCount        =   Int32PtrFromNullInt64(unreadCount)
    conv.LastMessage        =   &lastMessage.String
    conv.LastActivityDate   =   BlitzMessage.TimestampPtrFromNullTime(lastActivity)
    conv.ClosedDate         =   BlitzMessage.TimestampPtrFromNullTime(closedDate)

    if parentFeedPostID.Valid {
        conv.ParentFeedPostID = &parentFeedPostID.String
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

        error := WriteConversation(conversation)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
    }

    conversation, error = ReadUserConversation(session.UserID, conversationID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    response := BlitzMessage.ConversationResponse { Conversation: conversation }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { ConversationResponse: &response },
    }

    return serverResponse
}


//----------------------------------------------------------------------------------------
//                                                                      FetchConversations
//----------------------------------------------------------------------------------------


func FetchConversations(session *Session, req *BlitzMessage.FetchConversations) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    rows, error := config.DB.Query(
        `select conversationID from ConversationMemberTable where memberID = $1;`,
        session.UserID)
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

    response := BlitzMessage.FetchConversations { Conversations: convos }
    serverResponse := &BlitzMessage.ServerResponse {
        ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
        ResponseType:       &BlitzMessage.ResponseType { FetchConversations: &response },
    }

    return serverResponse
}

