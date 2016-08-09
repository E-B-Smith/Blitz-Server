
-- SQL to update database for talk conversations.

alter table conversationTable rename initiatoruserid to initiatorID;

alter table conversationTable
    add column expertID userID,
    add column conversationType smallint,
    add column topic text,
    add column callTime timestamptz,
    add column suggestedDuration interval,
    add column suggestedTimes timestamptz[]
    ;

update conversationTable ct set conversationType = 1
    where ct.parentfeedpostid is null;

select * from conversationTable where conversationType is null;

update conversationTable ct1 set expertID = cmt.memberID
    from conversationTable ct
    join conversationmembertable cmt on
        (cmt.conversationID = ct.conversationID
        and cmt.memberID <> ct.initiatorID)
    where ct1.conversationID = cmt.conversationID
        ;

select ct.expertID, cmt.memberID
    from conversationTable ct
    join conversationmembertable cmt on
        (cmt.conversationID = ct.conversationID
        and cmt.memberID <> ct.initiatorID)
        ;
