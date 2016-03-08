
Stats
=====

To Do
-----
* Downloads
* Most scoring users
* Happiest users
* Saddest losers
* Device stats

* Use frequency by time of day
* Use frequency by week
* Most connected users
* Longest distance users
* System usage
* Send sms on crash
* Send sms on resource limits


Done
----
* Up/Down time
        select messageType, StringTimeInterval(Now(), timestamp)
            from ServerStatTable
            where message in ( 'Started', 'Terminated' )
            order by timestamp desc limit 1;


* Total users
        select count(*) as "Total Users" from usertable;

* Active users
        select count(*) as "Active Users" from usertable where userstatus > 1;

* Most recent users
        select distinct usereventtable.userid, max(usereventtable.timestamp), usertable.name, devicetable.modelName
            from usereventtable
            join usertable on usereventtable.userid = usertable.userid
            join devicetable on usereventtable.userid = devicetable.userid
            group by usereventtable.userid, usertable.name, devicetable.modelName
            order by max(usereventtable.timestamp) desc;

* Most active users
        select count(*), usereventtable.userid, usertable.name
            from usereventtable
            join usertable on usereventtable.userid = usertable.userid
            group by usereventtable.userid, usertable.name
            order by count(*) desc;

* Friend status
        select friendtable.userid, u1.name, u2.name, StringFromFriendStatus(friendstatus)
            from friendtable
            left join usertable as u1 on (u1.userid = friendtable.userid)
            left join usertable as u2 on (u2.userid = friendtable.friendid)
            order by friendtable.userid;

* Message type, message bytes in/out, response code, start/stop
        select messageType, count(*), avg(elapsed) as "Avg Response Sec.",
            sum(bytesin) as "Bytes In", sum(bytesout) as "Bytes Out"
            from ServerStatTable group by message;
