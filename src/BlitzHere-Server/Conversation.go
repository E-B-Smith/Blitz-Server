

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
            (conversationID, status, initiatorUserID, parentFeedPostID, creationDate)
            values ($1, $2, $3, $4, current_timestamp)
         on conflict(conversationID) do
            update set (status, parentFeedPostID) = ($2, $4);`,
        conv.ConversationID,
        conv.Status,
        conv.InitiatorUserID,
        conv.ParentFeedPostID)
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
        _, _ = config.DB.Exec(
            `insert into ConversationMemberTable
                (conversationID, memberID) values ($1, $2);`,
            conv.ConversationID,
            memberID)
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


func ReadConversation(conversationID string) (*BlitzMessage.Conversation, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
            conversationID,
            status,
            initiatorUserID,
            parentFeedPostID,
            creationDate
                from ConversationTable
                where conversationID = $1;`, conversationID)

    var (
        status              sql.NullInt64
        initiatorUserID     sql.NullString
        parentFeedPostID    sql.NullString
        creationDate        pq.NullTime
    )

    error := row.Scan(&conversationID, &status, &initiatorUserID, &parentFeedPostID, &creationDate)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    var conv BlitzMessage.Conversation
    conv.ConversationID     =   &conversationID
    conv.InitiatorUserID    =   &initiatorUserID.String
    conv.Status             =   BlitzMessage.UserMessageStatus(status.Int64).Enum()
    conv.CreationDate       =   BlitzMessage.TimestampPtrFromNullTime(creationDate)
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

    convID := Util.NewUUIDString()
    conv := BlitzMessage.Conversation {
        ConversationID:     &convID,
        InitiatorUserID:    &session.UserID,
        Status:             BlitzMessage.UserMessageStatus(BlitzMessage.UserMessageStatus_MSNew).Enum(),
        ParentFeedPostID:   req.ParentFeedPostID,
        MemberIDs:          memberArray,
    }

    error := WriteConversation(&conv)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    convPtr, error := ReadConversation(convID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    response := BlitzMessage.ConversationResponse { Conversation: convPtr }
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
            convo, error := ReadConversation(convID)
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

