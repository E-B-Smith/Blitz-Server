

--  Sample Queries


-- insert into friendtable (userid, friendid, friendstatus) values (
--     '801D2014-6113-401E-B6AA-5E2B7D630D43',
--     '1FB481C1-58B1-436F-9A71-26BFF1FC6639',
--     1);


-- insert into friendtable (userid, friendid, friendstatus) values (
--     '1FB481C1-58B1-436F-9A71-26BFF1FC6639',
--     '801D2014-6113-401E-B6AA-5E2B7D630D43',
--     2);


-- insert into friendtable (userid, friendid, friendstatus) values (
--     '801D2014-6113-401E-B6AA-5E2B7D630D43',
--     '1FB481C1-58B1-436F-9A71-26BFF1FC6639',
--     5);
--
--
-- insert into friendtable (userid, friendid, friendstatus) values (
--     '1FB481C1-58B1-436F-9A71-26BFF1FC6639',
--     '801D2014-6113-401E-B6AA-5E2B7D630D43',
--     5);

-- alter table devicetable add column deviceudid text;
-- \copy deviceudidtable (devicename, deviceudid) from 'DeviceUDID.txt'


select usereventtable.userid, usertable.name, count(*) from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    group by usertable.userid, usereventtable.userid;


select usertable.name, count(*) from usertable
    join scoretable on usertable.userid = scoretable.userid
    group by usertable.userid, scoretable.userid;


select usertable.name, count(*) from usertable
    join usereventtable on usertable.userid = usereventtable.userid
    group by usertable.userid, usereventtable.userid;


select usertable.name, usertable.userid, devicetable.appID, devicetable.notificationToken
    from devicetable join usertable on usertable.userid = devicetable.userid
    where devicetable.notificationToken is not null;


select  usertable.name, deviceudidtable.devicename, usertable.lastseen,
        devicetable.modelname, devicetable.appversion
    from usertable
    left join devicetable on devicetable.userid = usertable.userid
    left join deviceudidtable on deviceudidtable.deviceudid = devicetable.deviceudid
    order by lastseen desc limit 40;


select friendtable.userid, u1.name, u2.name, friendstatus from friendtable
    left join usertable as u1 on (u1.userid = friendtable.userid)
    left join usertable as u2 on (u2.userid = friendtable.friendid)
    order by friendtable.userid;


select useridentitytable.userid, usertable.lastseen from useridentitytable
    left join usertable on usertable.userid = useridentitytable.userid
    where identitystring in ('facebookedwardsmith', 'facebook10152986211559513')
    group by usertable.userid, useridentitytable.userid
    order by usertable.lastseen;


select MorphUserIDIntoUserID(
    'A9FD20AE-76D5-4A8C-947F-4CD424B8CF44',
    'E9F6F4A8-CAB8-4399-876E-7A9DDCBF6B36');


select count(*), usereventtable.userid, usertable.name
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    group by usereventtable.userid, usertable.name
    order by count(*) desc;


select usereventtable.userid, usertable.name, usereventtable.timestamp from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    order by usereventtable.timestamp desc
    group by usereventtable.userid;


select distinct usereventtable.userid, max(usereventtable.timestamp), usertable.name, devicetable.modelName
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    join devicetable on usereventtable.userid = devicetable.userid
    group by usereventtable.userid, usertable.name, devicetable.modelName
    order by max(usereventtable.timestamp) desc;


select message, count(*), avg(elapsed) as "Avg Response Sec.",
    sum(bytesin) as "Bytes In", sum(bytesout) as "Bytes Out"
    from MessageStatTable group by message;


--
-- Score stats
--


-- happyscore  | basescore | displayscore | physical |  mental  |  vital   | environment

select current_timestamp - date_trunc('day', age(timestamp)) as day,
    avg(happyscore),
    avg(basescore),
    avg(displayscore),
    avg(physical),
    avg(mental),
    avg(vital),
    avg(environment)
      from scoretable
      where date_trunc('day', age(timestamp)) < interval '7 days' and userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
      group by day order by day;

-- B369F342-5BE4-4C47-9C7A-98A525114059
-- 6CC55B19-976F-4D08-AE69-80CEF296460B

-- Weather


select count(*), (weather).weatherType as "wt", avg(happyscore)
  from scoretable where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
  group by "wt" order by "wt";

--  Hearts

select sum(case when messagetype=8 or messagetype=9 then 1 else 0 end) as sent
    from messagetable where senderid = 'B369F342-5BE4-4C47-9C7A-98A525114059';

select sum(case when messagetype=8 or messagetype=9 then 1 else 0 end) as received
    from messagetable where recipientid = 'B369F342-5BE4-4C47-9C7A-98A525114059';


-- Circle score


select friendid from friendtable where friendstatus = 5 and userid =  'B369F342-5BE4-4C47-9C7A-98A525114059';


select current_timestamp - date_trunc('day', age(timestamp)) as day,
    count(*),
    avg(happyscore),
    avg(basescore),
    avg(displayscore),
    avg(physical),
    avg(mental),
    avg(vital),
    avg(environment)
      from scoretable
      where date_trunc('day', age(timestamp)) < interval '7 days'
        and (userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
        or userid in
        (select friendid from friendtable where friendstatus = 5 and userid =  'B369F342-5BE4-4C47-9C7A-98A525114059'))
      group by day order by day;

-- Global

select current_timestamp - date_trunc('day', age(timestamp)) as day,
    avg(happyscore),
    avg(basescore),
    avg(displayscore),
    avg(physical),
    avg(mental),
    avg(vital),
    avg(environment)
      from scoretable
      where date_trunc('day', age(timestamp)) < interval '7 days' and userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
      group by day order by day;


-- Emo-circle:


update scoretable set userResponse = array[(1, 1.0)::userresponse, (3, 1.0)::userresponse, (5, 1.0)::userresponse ]
    where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
    and timestamp = '2015-05-15 19:27:03.000275-07';


update scoretable set userResponse = array[(2, 1.0)::userresponse, (4, 1.0)::userresponse, (6, 1.0)::userresponse ]
    where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
     and timestamp = '2015-05-15 19:21:58.000663-07';


update scoretable set userResponse = array[(2, 1.0)::userresponse, (4, 1.0)::userresponse, (6, 1.0)::userresponse ]
    where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
     and timestamp = '2015-05-15 19:28:25.000016-07';


select count(*), (r).emotionid, sum((r).emotionvalue) from
    ( select unnest(userresponse)::userresponse as r from scoretable ) s
    group by (r).emotionid order by (r).emotionid;


-- 8F93B6BE-E1DA-4761-91C6-EE9D56925568
-- B369F342-5BE4-4C47-9C7A-98A525114059


select date_trunc('day', age(timestamp)) as day,
    happyscore,
    basescore,
    displayscore,
    physical,
    mental,
    vital,
    environment
      from scoretable
      where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
       or userid in
        (select friendid from friendtable where friendstatus = 5 and userid =  'B369F342-5BE4-4C47-9C7A-98A525114059')
    order by day;

select userid from scoretable
       where userid = 'B369F342-5BE4-4C47-9C7A-98A525114059'
       or userid in
        (select friendid from friendtable where friendstatus = 5 and userid =  'B369F342-5BE4-4C47-9C7A-98A525114059');


select userid, timestamp from scoretable
    where userid in (
        '51951A93-FA1F-487B-802F-75B07C58F837',
        '5FD32EE0-46AB-4756-A010-D6E1EC0412F8',
        '429F899D-F5A9-4818-A755-4CB580794E05',
        'B369F342-5BE4-4C47-9C7A-98A525114059');


        (select friendid as userid from friendtable where userid =  'B369F342-5BE4-4C47-9C7A-98A525114059' and friendstatus = 5);


select scoretable.userid, timestamp, name from scoretable join usertable on scoretable.userid = usertable.userid;


insert into friendtable (userid, friendid, friendstatus) values
    ('B369F342-5BE4-4C47-9C7A-98A525114059'
    ,'6CC55B19-976F-4D08-AE69-80CEF296460B'
    ,5);



select MergeUserIDIntoUserID('22D25384-DC83-4444-81C3-1252B17C25AA', '40C6DA10-1B67-4B4B-B826-698BBA8BAB38')


* User profile updates at:
  - Update session
  - Update profile score
  - Update profile

select * from pg_stat_activity;


--  New users by month


select
    to_char(date_trunc('month', CreationDate), 'Mon YYYY') as "Month",
    count(*)::int as "New",
    lpad('', (count(*))::int, '#') as "New Users"
    from UserTable
    where CreationDate is not null
    group by date_trunc('month', CreationDate)
    order by date_trunc('month', CreationDate)
;


select count(*) from usertable;
select count(*) from usertable where lastseen is not null;
select count(*) from usertable where CreationDate is null;


--  User activity by month


select
    to_char(date_trunc('month', timestamp), 'Mon YYYY') as "Month",
    count(*)::int as "Events",
    lpad('', (count(*)/150)::int, '#') as "User Activity"
    from usereventtable
    where timestamp is not null
    group by date_trunc('month', timestamp)
    order by date_trunc('month', timestamp)
;


--
--  Unique users and activty by month
--

with monthlyusers as (
select
    date_trunc('month', timestamp) as timestamp,
    userid,
    count(*) as events
        from usereventtable
        group by 1,2
        order by 1,2
)
select
    to_char(date_trunc('month', timestamp), 'Mon YYYY') as "Month",
    count(*) as "Unique Users",
    sum(monthlyusers.events) as "User Events",
    rpad(lpad('', (count(*)/10)::int, '#'),
        (sum(monthlyusers.events)/100 - (count(*)/10)::int)::int,
        '+') as "Users / Activity by Month"
    from monthlyusers
        group by timestamp
        order by timestamp
;



--
--  User stats
--

-- New Users:
select
    date_trunc('month', CreationDate) as timestamp,
    count(*) as newusers
    from UserTable
    where CreationDate is not null
    group by date_trunc('month', CreationDate)
    order by date_trunc('month', CreationDate)
;

-- Unique visitors:
select date_trunc('month', timestamp) as timestamp,
    count(distinct userid)
    from usereventtable
    group by 1 order by 1
;


with newusers as (
select
    date_trunc('month', CreationDate) as timestamp,
    count(*) as usercount
    from UserTable
    where CreationDate is not null
    group by date_trunc('month', CreationDate)
    order by date_trunc('month', CreationDate)
),
uniqueusers as (
select date_trunc('month', timestamp) as timestamp,
    count(distinct userid) as usercount
    from usereventtable
    group by 1 order by 1
)
select
    to_char(date_trunc('month', newusers.timestamp), 'Mon YYYY') as "Month",
    newusers.usercount as "New",
    uniqueusers.usercount as "Unique",
    rpad(lpad('', (newusers.usercount/10)::int, '#'),
        ((uniqueusers.usercount - newusers.usercount)/10)::int,
        '-') as "New / Unique by Month"
        from newusers
        full outer join uniqueusers on newusers.timestamp = uniqueusers.timestamp
;

--
--  Repeat visits by month:
--

with userVisitDays as (
select
    date_trunc('day', timestamp) as timestamp,
    userid
        from usereventtable
        group by 1, 2
        order by 2, 1
),
userVisitsPerMonth as (
select
    userid,
    date_trunc('month', timestamp) as timestamp,
    count(*) as visitsPerMonth
        from userVisitDays
        group by 1, 2
        order by 1, 2
)
select
    to_char(date_trunc('month', timestamp), 'Mon YYYY') as "Month",
    to_char(avg(visitsPerMonth), 'FM99.00') as "avg",
    max(visitsPerMonth),
    rpad(
        lpad('', (avg(visitsPerMonth))::int, '#'),
        (max(visitsPerMonth) - avg(visitsPerMonth))::int,
        '-') as "Returning user visit days per month"
        from userVisitsPerMonth
        group by timestamp
        order by timestamp
;


with recursive tree(parent, item, isLeaf, parentlist) as (
	select parent, item, isLeaf, array[ parent ]
		from CategoryTable
		where parent = 'root'
union
	select ct.parent, ct.item, ct.isLeaf, array_append(parentlist, ct.parent)
		from CategoryTable as ct
	inner join tree t
	   on (ct.parent = t.item)
)
select parentlist, item, isLeaf, UserTable.name from tree
	left join EntityTagTable on
		(entityTag = item and entityType = 1 and entityID::text = userID::text)
	join UserTable on UserTable.userID = EntityTagTable.userID
	where isLeaf;

select * from EntityTagTable
	where userID = '24f1daba-3555-45d9-b19a-97ccc648fd0e'
	  and EntityTagTable.entityID::text = EntityTagTable.userID::text
	  and entityType = 1;

with recursive tree(parent, item, isLeaf, parentlist) as (
	select parent, item, isLeaf, array[ parent ]
		from CategoryTable
		where parent = 'root'
union
	select ct.parent, ct.item, ct.isLeaf, array_append(parentlist, ct.parent)
		from CategoryTable as ct
	inner join tree t
	   on (ct.parent = t.item)
)
select parentlist, item, isLeaf from tree where isLeaf;

select distinct et.entityTag, ct.item as "Category" from EntityTagTable et
	left join CategoryTable ct on et.entityTag = ct.item
	where ct.item is null
	  and et.entityTag not like '.%'
;


