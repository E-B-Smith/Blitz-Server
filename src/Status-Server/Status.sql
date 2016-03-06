
--  ServerStats.sql
--
--  EB Smith, March 2015

\echo <div><h3>Up Since</h3><pre><code>
select message, StringTimeInterval(Now(), timestamp) as "Time"
    from MessageStatTable
    where message in ( 'Started', 'Terminated' )
    order by timestamp desc limit 1;
\echo </code></pre></div>


\echo <div><h3>Total Users</h3><pre><code>
select count(*) from usertable;
\echo </code></pre></div>


\echo <div><h3>Active Users</h3><pre><code>
select count(*) from usertable where userstatus > 1;
\echo </code></pre></div>


\echo <div><h3>Recent Activity</h3><pre><code>
select distinct usereventtable.userid,
        max(usereventtable.timestamp) as "Last Active",
        usertable.name,
        devicetable.modelName
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    join devicetable on usereventtable.userid = devicetable.userid
    group by usereventtable.userid, usertable.name, devicetable.modelName
    order by max(usereventtable.timestamp) desc;
\echo </code></pre></div>


\echo <div><h3>Most Active Users</h3><pre><code>
select count(*), usereventtable.userid, usertable.name
    from usereventtable
    join usertable on usereventtable.userid = usertable.userid
    group by usereventtable.userid, usertable.name
union
select count(*), 'Total', ' '
    from usereventtable
order by count desc;
\echo </code></pre></div>


\echo <div><h3>Friends</h3><pre><code>
select friendtable.userid, u1.name, u2.name, StringFromFriendStatus(friendstatus) as "Status"
    from friendtable
    left join usertable as u1 on (u1.userid = friendtable.userid)
    left join usertable as u2 on (u2.userid = friendtable.friendid)
    order by friendtable.userid;
\echo </code></pre></div>


\echo <div><h3>Network Messages</h3><pre><code>
select message as "Message", count(*) as "Count",
    avg(elapsed) as "Avg Response Sec.",
    sum(bytesin) as "Bytes In", sum(bytesout) as "Bytes Out"
    from MessageStatTable group by message
union
select ' ', count(*), avg(elapsed),  sum(bytesin), sum(bytesout)
    from MessageStatTable;
\echo </code></pre></div>
