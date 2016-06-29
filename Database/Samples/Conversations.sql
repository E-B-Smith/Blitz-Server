select creationDate, conversationID, senderID, messageText
    from UserMessageTable
    where recipientID = '24f1daba-3555-45d9-b19a-97ccc648fd0e'
    order by creationDate desc;


select conversationID, senderID
    from UserMessageTable
    where recipientID = '24f1daba-3555-45d9-b19a-97ccc648fd0e';


select conversationID, senderID, messageText, creationDate
from (
    select conversationID, senderID, messageText, creationDate,
        rank() over (partition by conversationID, senderID order by creationDate desc) as r
    from UserMessageTable
    where recipientID = '24f1daba-3555-45d9-b19a-97ccc648fd0e'
) as conv
where r = 1
order by conversationID, creationDate desc;


select conversationID, senderID, u.name, messageText, conv.creationDate
from (
    select conversationID, senderID, messageText, creationDate,
        rank() over (partition by senderID order by creationDate desc) as r
    from UserMessageTable
    where recipientID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
) as conv
left join usertable u on u.userid = senderID
where r = 1 and conversationID is not null
order by conversationID, creationDate desc;

with conv as (
select a.conversationID as cid, a.memberID as mid, b.memberID
	from conversationMemberTable a
	join conversationMemberTable b on a.conversationID = b.conversationID and b.memberID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
	where a.memberID = '4a952764-779d-4ea4-b402-6cec9ddcb099'
)
select count(*),
	sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
	from conv, usermessagetable where conversationID = conv.cid and recipientID = conv.mid;

select
	count(*),
	sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
	from usermessagetable
	where conversationID = '83803dbd-12b8-4fa8-be0e-646bf3e2ec8c'
	  and recipientID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0';


select count(*),
	sum(case when messageStatus <= 2 or messageStatus is null then 1 else 0 end)
	from (
		select a.conversationID as cid, a.memberID as mid, b.memberID
			from conversationMemberTable a
			join conversationMemberTable b on a.conversationID = b.conversationID and b.memberID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
			where a.memberID = '4a952764-779d-4ea4-b402-6cec9ddcb099'
	) as conv,
	usermessagetable where conversationID = conv.cid and recipientID = conv.mid;

select * from UserMessageTable
	where recipientID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
	  and messageType = 4;

select * from UserMessageTable
	where messageType = 4;

select * from userMessageTable
	where messageText ilike ;

select a.conversationID as cid, a.memberID as mid, b.memberID
	from conversationMemberTable a
	join conversationMemberTable b on a.conversationID = b.conversationID and b.memberID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
	where a.memberID = '4a952764-779d-4ea4-b402-6cec9ddcb099';


select * from feedposttable
	where postid = 'c55d4fd1-50a7-468f-9cb1-ae5e7514d402'
	   or parentid = 'c55d4fd1-50a7-468f-9cb1-ae5e7514d402';

with c as (
	select conversationID from conversationmembertable
		intersect select conversationID from conversationmembertable where memberID = '24f1daba-3555-45d9-b19a-97ccc648fd0e'
		intersect select conversationID from conversationmembertable where memberID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
)
select * from conversationtable a, c
	where a.conversationID = c.conversationID;


select conversationID from conversationmembertable
	intersect select conversationID from conversationmembertable where memberID = '24f1daba-3555-45d9-b19a-97ccc648fd0e';

