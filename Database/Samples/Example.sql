select distinct on (1)
    UserMessageTable.recipientID,
    UserMessageTable.messageText,
    UserMessageTable.actionURL,
    DeviceTable.appID,
    DeviceTable.notificationToken,
    DeviceTable.appIsReleaseVersion
      from UserMessageTable
      join DeviceTable on DeviceTable.userID = UserMessageTable.recipientID
        where UserMessageTable.notificationDate is null
          and DeviceTable.notificationToken is not null
          and DeviceTable.appID is not null
        order by UserMessageTable.recipientID, UserMessageTable.creationDate;




select dateAdded, imageContent, contentType, crc32
    from ImageTable
    where userID = '4a952764-779d-4ea4-b402-6cec9ddcb099'
    order by dateAdded desc;


select entityTag,
    (select count(*)
        from EntityTagTable
        where entityID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
        and entityType = 1
        and entityTag = EntityTagTable.entityTag)
    from EntityTagTable
    where userID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
      and entityID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
      and entityType = 1;

select entityTag, count(*)
    from EntityTagTable
    where entityID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
    and entityType = 1
    and entityTag = 'ios';


select entityTag as tagName,
    (select count(*)
        from EntityTagTable
        where entityID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
        and entityType = 1
        and entityTag = t1.entityTag)
    from EntityTagTable t1
    where userID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
      and entityID = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0'
      and entityType = 1;



        `select entityTag as tagName,
            (select count(*)
                from EntityTagTable
                where entityID = $2
                and entityType = $3
                and entityTag = t1.entityTag)
            from EntityTagTable t1
            where userID = $1
              and entityID = $2
              and entityType = $3;`,
            userID, entityID, entityType,

select
    entityTag,
    count(*),
    (select 1 if )
    from EntityTagTable
    where entityID = $1
      and entityType = $2
    group by entityTag;


select
    entityTag,
    count(*),
    count(userid = '4a952764-779d-4ea4-b402-6cec9ddcb099')
    from EntityTagTable
    where entityID = '1559e965-3c89-4ddb-9384-264309c36e26'
      and entityType = 2
    group by entityTag;

select
    entityTag,
    count(*),
    sum(case when userid = '4a952764-779d-4ea4-b402-6cec9ddcb099' then 1 else 0 end)
from EntityTagTable
group by entityTag
where entityID = '1559e965-3c89-4ddb-9384-264309c36e26'
  and entityType = 2;


-- Make Bobby Blitz a friend to all!

insert into entitytagtable
    (entityid, entitytype, entitytag, userid)
    select 'a8277a5e-b461-476b-9f4a-922a50b97f26', 1, '.friend', userid
        from usertable
        on conflict do nothing;

insert into entitytagtable
    (userid, entitytype, entitytag, entityid)
    select 'a8277a5e-b461-476b-9f4a-922a50b97f26', 1, '.followed', userid::uuid
        from usertable
        on conflict do nothing;


