
--  BlitzHere Common Database Schema
--
--  E.B.Smith  -  March 2016


create domain UUID as varchar(36);
create domain UserID as varchar(36);
create domain SessionToken as varchar(36);
create domain DeviceID as varchar(36);
create domain ShortHash as varchar(8);
create domain UserStatus as smallint;
create domain PlatformType as smallint;
create domain Gender as smallint;
create domain Event as varchar(48);


create type KeyValue as
    (
     key        varchar(32)
    ,value      varchar(64)
    );


create type Size as
    (
     width    real
    ,height   real
    );


create type Location as
    (
     latitude    real
    ,longitude   real
    ,placename   text
    );


create type RGBColor as
    (
     red         smallint
    ,green       smallint
    ,blue        smallint
    );

create table UserContactTable
    (
     userID         UserID      not null
    ,contactType    smallint
    ,contact        text
    ,isverified     boolean
    ,code           text
    ,codeDate       timestamptz
    );
create unique index UserContactTableUniqueIndex
    on UserContactTable(userID, contactType, contact);


create table UserTable
    (
     userID             UserID  unique not null primary key
    ,userStatus         UserStatus
    ,creationDate       timestamptz
    ,lastSeen           timestamptz

    ,name               text
    ,gender             Gender
    ,birthday           timestamptz

    ,backgroundSummary  text
    ,interestTags       text[]

    ,isExpert           boolean
    ,isFree             boolean
    ,isAdmin            boolean
    ,stripeAccount      text
    ,processorid        text
    ,defaultCard        text    -- deprecated

    ,search             tsvector
    ,editProfileID      text

    ,chatCharge         text
    ,callCharge         text
    ,shortQACharge      text
    ,longQACharge       text
    ,charityPercent     real
    );
create index UserSearchIndex on UserTable using gin(search);


create table EmploymentTable
    (
     userID             UserID  not null
    ,isHeadLineItem     boolean
    ,jobTitle           text
    ,companyName        text
    ,location           text
    ,industry           text
    ,startDate          timestamptz
    ,stopDate           timestamptz
    ,summary            text
    );
create index EmploymentTableIndex on EmploymentTable(userID);


create table EducationTable
    (
     userID             UserID not null
    ,schoolName         text
    ,degree             text
    ,emphasis           text
    ,startDate          timestamptz
    ,stopDate           timestamptz
    ,summary            text
    );
create index EducationTableIndex on EducationTable(userID);


create table SessionTable
    (
     userID             UserID
    ,deviceID           DeviceID
    ,sessionToken       SessionToken    unique not null primary key
    ,timestamp          timestamptz
    ,secret             text

    ,unique(userID, deviceID)
    );


create table SocialTable
    (
     userID             UserID  not null
    ,service            text    not null
    ,socialID           text    not null
    ,userName           text
    ,displayName        text
    ,URI                text
    ,authToken          text
    ,authExpire         timestamptz
    );
create unique index SocialTableUniqueIndex on SocialTable(userID, service, socialID);


create table ReferralTable
    (
     referreeID         UserID          not null
    ,referrerID         UserID          not null
    ,creationDate       timestamptz     not null
    ,referralType       smallint
    ,referenceID        text
    ,validFromDate      timestamptz
    ,validToDate        timestamptz
    ,claimDate          timestamptz
    ,redemptionDate     timestamptz
    ,referralCode       text
    );
create unique index ReferralTableIndex on
    ReferralTable(referreeID, referrerID, creationDate);
create unique index ReferralCodeIndex on
    ReferralTable(referralCode);


create table DeviceTable
    (
     userID             UserID       unique not null primary key
--  ,deviceID           DeviceID     not null
    ,timestamp          timestamptz
    ,lastIPAddress      inet
    ,localIPAddress     inet
    ,platformType       PlatformType
    ,modelName          text
    ,systemVersion      text
    ,language           text
    ,timezone           text
    ,phoneCountryCode   text
    ,screenSize         Size
    ,screenScale        float
    ,appID              text
    ,appVersion         text
    ,appIsReleaseVersion boolean
    ,notificationToken  text
    ,vendorID           text
    ,advertisingID      text
    ,deviceUDID         text
    ,systemBuildVersion text
    );


create table UserEventTable
    (
     userID             UserID      not null
    ,timestamp          timestamptz not null
    ,location           Location
    ,event              Event
    ,eventData          text[]

    ,unique(userID, timestamp)
    );
create unique index UserEventUniqueIndex on UserEventTable(userID, timestamp);


create table ShortnerTable
    (
     dataHash      ShortHash    unique not null primary key
    ,dataBlob      text         unique not null
    );
create index ShortnerTableDataIndex on ShortnerTable(dataBlob);


create function ShortHashFromDataBlob(dataBlob text) returns ShortHash as
    $$
    begin
    if dataBlob is null then return null; end if;
    return substring(md5(dataBlob), 17, 6);
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create or replace function InsertStringGivingShortHash(inputstring text) returns ShortHash as
    $$
    declare
        hash ShortHash;
    begin
    if inputstring is null then return null; end if;

    hash = ShortHashFromDataBlob(inputstring);

        begin
        insert into ShortnerTable (dataHash, dataBlob) values (hash, inputstring);
        return hash;
        exception when unique_violation then
            select dataHash into hash from ShortnerTable where dataBlob = inputstring;
            return hash;
        end;

    end;
    $$
    language plpgsql
    returns null on null input;


create function NullIfInvalidTimestamp(ts timestamptz) returns timestamptz as
    $$
    begin
    if ts is null or extract(epoch from ts) < 1 then
        return null;
    else
        return ts;
    end if;
    end;
    $$
    language plpgsql immutable
    returns null on null input;

create or replace
function StringFromUserStatus(userStatus UserStatus) returns text as
    $$
    declare

    statusString text[] := array
        [ 'UserStatusUnknown',
          'UserStatusBlocked',
          'UserStatusInvited',
          'UserStatusActive',
          'UserStatusConfirming',
          'UserStatusConfirmed' ];

    begin
    if userStatus is null then return null; end if;
    return statusString[userStatus+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table UserIdentityTable
    (
     userID         UserID not null
    ,identityString varchar not null
    );
create index UserIdentityIndex on UserIdentityTable(identityString);
create unique index UserIdentityUniqueIndex on UserIdentityTable(userID, identityString);


create table DeviceUDIDTable
    (
     userID             UserID
    ,IPAddress          inet
    ,deviceUDID         text
    ,deviceName         text
    ,name               text
    ,email              text
    ,phone              text
    ,modificationDate   timestamptz
    ,tempID             text
    ,imei               text
    ,iccid              text
    ,meid               text
    ,macAddress         text
    ,product            text
    ,version            text
    ,serial             text
    ,notes              text
    );


create table AppDownloadTable
    (
     timestamp   timestamptz    unique not null primary key
    ,IPAddress          inet
    ,filename           text
    ,httpCode       smallint
    ,totalBytes       bigint
    );


create table ServerStatTable
    (
     timestamp          timestamptz unique not null primary key
    ,elapsed            real
    ,messageType        text
    ,bytesIn            int
    ,bytesOut           int
    ,statusCode         int
    ,responseCode       text
    ,responseMessage    text
    );


create domain MessageType as smallint;


create function StringFromUserMessageType(messageType MessageType) returns text as
    $$
    declare

    statusString text[] := array
        [ 'MessageTypeUnknown',
          'MessageTypeSystem'
          'MessageTypeNotification'];

    begin
    if messageType is null then return null; end if;
    return statusString[messageType+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table UserMessageTable
    (
     messageID          UUID            not null
    ,senderID           UserID          not null
    ,conversationID     UUID
    ,recipientID        UserID          not null
    ,creationDate       timestamptz     not null
    ,notificationDate   timestamptz
    ,readDate           timestamptz
    ,messageType        MessageType     not null
    ,messageStatus      smallint
    ,messageText        text
    ,actionIcon         text
    ,actionURL          text
    );
create unique index UserMessageUniqueIndex on UserMessageTable(messageID, senderID, recipientID);
create index UserMessageDeliveryIndex on UserMessageTable(recipientID, creationDate);


create domain ImageContent as smallint;


create function StringFromImageContent(imageContent ImageContent) returns text as
    $$
    declare

    labels text[] := array
        [ 'ImageContentUnknown',
          'ImageContentUserProfile'
          'ImageContentUserBackground'
        ];

    begin
    if imageContent is null then return null; end if;
    return labels[imageContent+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table ImageTable
    (
     userID             UserID          not null
    ,dateAdded          timestamptz     not null
    ,imageContent       ImageContent
    ,contentType        text
    ,crc32              bigint          not null
    ,deleted            boolean
    ,imageData          bytea
    );
create unique index ImageTableUniqueIndex on ImageTable(UserID, crc32);



------------------------------------------------------------------------------------------
--
--                                                                              Feed Posts
--
------------------------------------------------------------------------------------------


create domain FeedPostType   as smallint;
create domain FeedPostScope  as smallint;
create domain FeedPostStatus as smallint;


create table FeedPostTable
    (
     postID                     UUID            unique not null primary key
    ,parentID                   UUID
    ,postType                   FeedPostType
    ,postScope                  FeedPostScope
    ,postStatus                 FeedPostStatus
    ,userID                     UserID
    ,anonymousPost              boolean
    ,timestamp                  timestamptz
    ,timeActiveStart            timestamptz
    ,timeActiveStop             timestamptz
    ,headlineText               text
    ,bodyText                   text
    ,mayAddReply                bool
    ,mayChooseMulitpleReplies   bool
    ,surveyAnswerSequence       int
    ,amountPerReply             text
    ,amountTotal                text
    );
create index FeedPostTimestampIndex on FeedPostTable(timestamp desc);
create index FeedReplyTable         on FeedPostTable(parentID);


create table FeedPanelTable
    (
     postID                     UUID        not null
    ,memberID                   UserID      not null
    ,bountyAmount               text
    ,dateAnswered               timestamptz
    );
create unique index FeedPanelTableIndex on FeedPanelTable(postID, memberID);


------------------------------------------------------------------------------------------
--
--                                                                          EntityTagTable
--
------------------------------------------------------------------------------------------


create domain EntityType as smallint;


create table EntityTagTable
    (
     entityID           UUID        not null
    ,entityType         EntityType  not null
    ,userID             UserID      not null
    ,entityTag          text        not null
    );
create unique index EntityTagTableIndex
    on  EntityTagTable(entityID, entityType, userID, entityTag);

create index EntityTagTableAutocompleteIndex
    on EntityTagTable (entityTag text_pattern_ops);


------------------------------------------------------------------------------------------
--
--                                                                           Conversations
--
------------------------------------------------------------------------------------------


create table ConversationTable
    (

     conversationID             UUID        unique not null primary key
    ,status                     smallint    not null
    ,initiatorID                UserID      not null
    ,parentFeedPostID           UUID
    ,creationDate               timestamptz not null
    ,acceptDate                 timestamptz
    ,closedDate                 timestamptz
    ,paymentStatus              smallint
    ,chargeID                   uuid

    --  Added to support calls:

    ,expertID                   userID
    ,conversationType           smallint
    ,topic                      text
    ,callDate                   timestamptz
    ,suggestedDuration          interval
    ,suggestedDates             timestamptz[]
    ,callPhoneNumber            text
    ,callStatus                 smallint
    );


create table ConversationMemberTable
    (
     conversationID             UUID        not null
    ,memberID                   UserID      not null
    );
create unique index ConversationMemberTableIndex
    on ConversationMemberTable(conversationID, memberID);


------------------------------------------------------------------------------------------
--
--                                                                                 Reviews
--
------------------------------------------------------------------------------------------


create table ReviewTable
    (
     userID         UserID      not null
    ,reviewerID     UserID      not null
    ,timestamp      timestamptz not null
    ,conversationID UUID
    ,responseTime   real        --  In seconds
    ,responsive     real
    ,outgoing       real
    ,recommended    real
    ,reviewText     text
    ,tags           text[]
    );
create unique index ReviewTableIndex
    on ReviewTable(userID, reviewerID, timestamp);


create or replace
function ResponseTimeForConversationUser(conversationIDIn text, userIDIn text) returns real as
    $$
    declare
        result real;
        start timestamptz;
        stop timestamptz;
    begin

    select
        ConversationTable.creationDate,
        UserMessageTable.creationDate
            into start, stop
        from UserMessageTable
        inner join ConversationTable
        on ConversationTable.conversationID = UserMessageTable.conversationID
        where UserMessageTable.conversationID = conversationIDIn::uuid
          and UserMessageTable.senderID = userIDIn
          order by UserMessageTable.creationDate limit 1;

    result = extract(epoch from (stop - start));
    return result;
    end;
    $$
    language plpgsql immutable
    returns null on null input;


------------------------------------------------------------------------------------------
--
--                                                                        PhoneNumberTable
--
------------------------------------------------------------------------------------------


create table PhoneNumberTable
    (
     phoneNumber        text    unique not null primary key
    ,conversationID     text
    ,clientPhoneNumber  text
    ,expertPhoneNumber  text
    ,startDate          timestamptz
    ,stopDate           timestamptz
    );


------------------------------------------------------------------------------------------
--
--
--                                                               User Profile Manipulation
--
--
------------------------------------------------------------------------------------------


------------------------------------------------------------------------------------------
--                                                                             EraseUserID
------------------------------------------------------------------------------------------


create or replace function EraseUserID(eraseID UserID) returns text as
    $$
    begin
    if eraseID is null then
        return null;
        end if;
    delete from DeviceTable where userID = eraseID;
    delete from ImageTable where userID = eraseID;
    delete from UserMessageTable where senderID = eraseID;
    delete from UserMessageTable where recipientID = eraseID;
    delete from SocialTable where userID = eraseID;
    delete from UserContactTable where userID = eraseID;
    delete from UserEventTable where userID = eraseID;
    delete from UserIdentityTable where userID = eraseID;
    delete from UserTable where userID = eraseID;
    delete from SessionTable where userID = eraseID;
    delete from EntityTagTable where userID = eraseID;
    delete from EntityTagTable where entityID = eraseID::uuid and entityType = 1;
    return 'User erased';
    end;
    $$
    language plpgsql
    returns null on null input;


------------------------------------------------------------------------------------------
--                                                                   MergeUserIDIntoUserID
------------------------------------------------------------------------------------------


create or replace
function MergeUserIDIntoUserID(oldID UserID, newID UserID) returns text as
    $$
    declare
        result text;
        oldidentity text;
    begin
    if oldID is null or newID is null then
        return null;
    end if;
    if oldID = newID then
        return 'User merged';
    end if;

    --  UserContactTable

    with recursive merge as (
        select userid, contacttype, contact from usercontacttable where userid = oldID
    )
    update usercontacttable set (userid) = (newID)
        from merge where usercontacttable.userid = merge.userid
                     and usercontacttable.contacttype = merge.contacttype
                     and usercontacttable.contact = merge.contact
        and (select 1 from usercontacttable
                where usercontacttable.userid = newID
                  and usercontacttable.contacttype = merge.contacttype
                  and usercontacttable.contact = merge.contact)
                  is null;

    --  SocialTable

    with recursive merge as (
        select userid, service, socialid from socialtable where userid = oldID
    )
    update socialtable set (userid) = (newID)
        from merge where socialtable.userid = merge.userid
                     and socialtable.service = merge.service
                     and socialtable.socialid = merge.socialid
        and (select 1 from socialtable
                where socialtable.userid = newID
                  and socialtable.service = merge.service
                  and socialtable.socialid = merge.socialid)
                  is null;

    --  UserDeviceTable

    update DeviceTable set (userid) = (newID)
        where userid = oldID
        and (select 1 from DeviceTable
            where DeviceTable.userid = newID) is null;

    --  UserEventTable

    with recursive merge as (
        select userid, timestamp from usereventtable where userid = oldID
    )
    update usereventtable set (userid) = (newID)
        from merge where usereventtable.userid = merge.userid
                     and usereventtable.timestamp = merge.timestamp
        and (select 1 from usereventtable
            where usereventtable.userid = newID
              and usereventtable.timestamp = merge.timestamp) is null;

    --  UserIdentityTable

    with recursive merge as (
        select userid, identitystring from UserIdentityTable where userid = oldID
    )
    update UserIdentityTable set (userid) = (newID)
        from merge where UserIdentityTable.userid = merge.userid
                     and UserIdentityTable.identitystring = merge.identitystring
        and (select 1 from UserIdentityTable
            where UserIdentityTable.userid = newID
              and UserIdentityTable.identitystring = merge.identitystring) is null;

    --  UserMessageTable

    with recursive merge as (
        select messageid, senderid, recipientid from UserMessageTable where senderid = oldID
    )
    update UserMessageTable set (senderid) = (newID)
        from merge where UserMessageTable.messageid = merge.messageid
                     and UserMessageTable.senderid = merge.senderid
                     and UserMessageTable.recipientid = merge.recipientid
        and (select 1 from UserMessageTable
            where UserMessageTable.messageid = merge.messageid
              and UserMessageTable.senderid = newID
              and UserMessageTable.recipientid = merge.recipientid) is null;

    --

    with recursive merge as (
        select messageid, senderid, recipientid from UserMessageTable where recipientid = oldID
    )
    update UserMessageTable set (recipientid) = (newID)
        from merge where UserMessageTable.messageid = merge.messageid
                     and UserMessageTable.senderid = merge.senderid
                     and UserMessageTable.recipientid = merge.recipientid
        and (select 1 from UserMessageTable
            where UserMessageTable.messageid = merge.messageid
              and UserMessageTable.senderid = merge.senderid
              and UserMessageTable.recipientid = newID) is null;

    --  ImageTable

    update imagetable set userid = newID
        where userid = oldID;

    --  Merge EntityTags

    create temporary table TempEntityTagTable
        (like EntityTagTable) on commit drop;

    insert into TempEntityTagTable
        select * from EntityTagTable
        where userID = oldID;

    insert into TempEntityTagTable
        select * from EntityTagTable
        where entityID = oldID::uuid
          and entityType = 1;

    update TempEntityTagTable set
        userID = newID
        where userID = oldID;

    update TempEntityTagTable set
        entityID = newID::uuid
        where entityID = oldID::uuid
          and entityType = 1;

    insert into EntityTagTable
        select * from TempEntityTagTable
        on conflict (entityID, entityType, userID, entityTag)
            do nothing;

    --  Done

    select EraseUserID(oldID) into result;
    if result = 'User erased'::text then
        result = 'User merged';
    else
        result = 'Merge failed';
    end if;

    return result;
    end;
    $$
    language plpgsql
    returns null on null input;


------------------------------------------------------------------------------------------
--                                                                      CopyUserIDToUserID
------------------------------------------------------------------------------------------


create or replace
function CopyUserIDToUserID(fromID text, newID text) returns text as
    $$
    -- declare
    --     fromID text;
    --     newID  text;
    begin

    -- fromID := '24f1daba-3555-45d9-b19a-97ccc648fd0e';
    --  newID := '24f1daba-3555-45d9-b19a-97ccc648fd0f';

    --  Copy user tables:
    --      UserTable
    --      EducationTable
    --      EmploymentTable
    --      UserContactTable
    --      EntityTagTable
    --      ImageTable

    -- UserTable

    create temporary table TempUserTable
        (like UserTable) on commit drop;

    insert into TempUserTable
        select * from UserTable
        where userID = fromID;

    update TempUserTable set userid = newID;

    delete from UserTable where userID = newID;
    insert into UserTable
        select * from TempUserTable;

    -- EducationTable

    create temporary table TempEducationTable
        (like EducationTable) on commit drop;

    insert into TempEducationTable
        select * from EducationTable
        where userID = fromID;

    update TempEducationTable set userid = newID;

    delete from EducationTable where userID = newID;
    insert into EducationTable
        select * from TempEducationTable;

    -- EmploymentTable

    create temporary table TempEmploymentTable
        (like EmploymentTable) on commit drop;

    insert into TempEmploymentTable
        select * from EmploymentTable
        where userID = fromID;

    update TempEmploymentTable set userid = newID;

    delete from EmploymentTable where userID = newID;
    insert into EmploymentTable
        select * from TempEmploymentTable;

    -- UserContactTable

    create temporary table TempUserContactTable
        (like UserContactTable) on commit drop;

    insert into TempUserContactTable
        select * from UserContactTable
        where userID = fromID;

    update TempUserContactTable set userid = newID;

    delete from UserContactTable where userID = newID;
    insert into UserContactTable
        select * from TempUserContactTable;

    --      EntityTagTable

    create temporary table TempEntityTagTable
        (like EntityTagTable) on commit drop;

    insert into TempEntityTagTable
        select * from EntityTagTable
        where userID = fromID
          and entityID = fromID::uuid
          and entityType = 1;

    update TempEntityTagTable set
        userid = newID,
        entityID = newID::uuid;

    delete from EntityTagTable
        where userID = newID
          and entityID = newID::uuid
          and entityType = 1;
    insert into EntityTagTable
        select * from TempEntityTagTable;

    -- Images

    create temporary table TempImageTable
        (like ImageTable) on commit drop;

    insert into TempImageTable
        select * from ImageTable
        where userID = fromID;

    update TempImageTable set userid = newID;

    delete from ImageTable where userID = newID;
    insert into ImageTable
        select * from TempImageTable;

    return 'User copied';
    end;
    $$
    language plpgsql;


------------------------------------------------------------------------------------------
--                                                                       CreateEditProfile
------------------------------------------------------------------------------------------


create or replace
function CreateEditProfile(fromID text, newID text) returns text as
    $$
    declare
        result text;
    begin

    select CopyUserIDToUserID(fromID, newID) into result;
    if result <> 'User copied' then
        return result;
    end if;

    update usertable set editprofileid = newID
        where userID = fromID;

    update userTable set
        editprofileid = fromID,
        userStatus = null
        where userID = newID;

    return 'Profile created';
    end;
    $$
    language plpgsql;


------------------------------------------------------------------------------------------
--                                                                      ApproveEditProfile
------------------------------------------------------------------------------------------


create or replace
function ApproveEditProfile(profileID text) returns text as
    $$
    declare
        result text;
        editID text;
        editStatus smallint;
        fromID text;
        toID   text;
    begin

    select editprofileid, userStatus from usertable
        where userID = profileID
        into editID, editStatus;

    if character_length(editID) < 16 then
        return 'No change to approve';
    end if;

    -- Swap the editID and profileID if they're reversed:

    if editStatus = 5 then
        fromID := editID;
          toID := profileID;
    else
        fromID := profileID;
          toID := editID;
    end if;

    select CopyUserIDToUserID(fromID, toID) into result;
    if result <> 'User copied' then
        return result;
    end if;

    update usertable set
        editprofileid = null,
        userStatus = 5
        where userID = toID;

    select EraseUserID(fromID) into result;
    if result <> 'User erased' then
       return result;
    end if;

    return 'Approved';
    end;
    $$
    language plpgsql;


------------------------------------------------------------------------------------------
--                                                                      DiscardEditProfile
------------------------------------------------------------------------------------------


create or replace
function DiscardEditProfile(profileID text) returns text as
    $$
    declare
        result text;
        editID text;
    begin

    select editprofileid from usertable
        where userID = profileID
        into editID;

    if character_length(editID) < 16 then
        return 'Not discarded';
    end if;

    update usertable set
        editprofileid = null,
        userStatus = 5
        where userID = profileID;

    select EraseUserID(editID) into result;
    if result <> 'User erased' then
       return result;
    end if;

    return 'Discarded';
    end;
    $$
    language plpgsql;


------------------------------------------------------------------------------------------
--
--
--                                                                        Helper Functions
--
--
------------------------------------------------------------------------------------------


create function StringFromTimeInterval(timestamp1 timestamptz, timestamp2 timestamptz) returns text as
    $$
    declare
        s text;
        dys int; hrs int; m int;
        sec real;
        ti interval;
    begin
    if timestamp1 is null or timestamp2 is null then
       return null;
       end if;

    if timestamp1 <= '-infinity'::timestamptz or timestamp1 >= 'infinity'::timestamptz or
       timestamp2 <= '-infinity'::timestamptz or timestamp2 >= 'infinity'::timestamptz then
        return 'infinity';
        end if;

    ti  := timestamp1 - timestamp2;
    dys := extract(day from ti);
    hrs := extract(hour from ti);
    m   := extract(minutes from ti);
    sec := extract(seconds from ti);

    case
    when dys > 0 then s := format('%s days, %s:%s:%s hours',
        to_char(dys, 'FM999'), to_char(hrs, 'FM99'), to_char(m, 'FM09'), to_char(sec, 'FM09D0'));
    when hrs > 0 then s := format('%s:%s:%s hours',
        to_char(hrs, 'FM99'), to_char(m, 'FM09'), to_char(sec, 'FM09D0'));
    when m > 0 then
        s := format('%s:%s minutes', to_char(m, 'FM99'), to_char(sec, 'FM09D0'));
    else
        s := format('%s seconds', to_char(sec, 'FM09D0'));
    end case;

    return s;
    end;
    $$
    language plpgsql immutable
    returns null on null input;


-- Tests for UserNameFromIPAddress:
--
-- select UserNameFromIPAddress('107.3.151.67');
-- select UserNameFromIPAddress('73.162.198.203');
-- select UserNameFromIPAddress('76.14.58.112');
-- select UserNameFromIPAddress('166.171.250.116');
-- select UserNameFromIPAddress('66.87.119.207');
-- select UserNameFromIPAddress('172.56.38.238');


create or replace
function UserNameFromIPAddress(ipaddress inet) returns text as
    $$
    declare
        uid UserID;
        uname text;
    begin

    if ipaddress is null then
       return null;
       end if;

    select usertable.name into uname from devicetable
        left join usertable on usertable.userid = devicetable.userid
        where lastipaddress = ipaddress
        and timestamp is not null
        and character_length(usertable.name) > 0
          order by devicetable.timestamp desc
          limit 1;
    if found then
        return uname;
        end if;

    select userid into uid from devicetable
      where lastipaddress = ipaddress and timestamp is not null
       order by timestamp desc;
    if found then
       return uid::text;
       end if;

    return ipaddress::text;

    end;
    $$
    language plpgsql immutable
    returns null on null input;


------------------------------------------------------------------------------------------
--                                                                                AppTable
------------------------------------------------------------------------------------------


create table AppTable
    (
     appID              text    unique not null primary key
    ,appName            text
    ,minAppVersion      text
    ,minAppDataDate     timestamptz
    );


------------------------------------------------------------------------------------------
--
--                                                                         HTTPDeepLinking
--
------------------------------------------------------------------------------------------


create table HTTPDeepLinkTable
    (
     deviceSignature    text        not null
    ,deviceRPM          real        not null
    ,creationDate       timestamptz not null
    ,claimDate          timestamptz
    ,deviceType         text
    ,inviteData         bytea       not null
    ,referrer           text
    );
create unique index HTTPDeepLinkIndex on HTTPDeepLinkTable(deviceSignature, creationDate);


------------------------------------------------------------------------------------------
--
--                                                                   Search & Autocomplete
--
------------------------------------------------------------------------------------------


create table AutocompleteTable
    (
     word       text    not null
    ,rank       int     not null
    );
create unique index AutocompleteIndex on autocompletetable (word text_pattern_ops);


create or replace
function UpdateAutocompleteTable(words text) returns void as
    $$
    declare
        wordarray text[];
        newword   text;
    begin

    words := lower(words);
    wordarray := regexp_split_to_array(words, E'\\s+');

    foreach newword in array wordarray
        loop
        if char_length(newword) > 1 then
            insert into autocompletetable
                (rank, word) values (1, newword)
            on conflict(word) do
                update set rank = autocompletetable.rank + 1;
            end if;
        end loop;

    end;
    $$
    language plpgsql;


create or replace
function UpdateSearchIndexForUserID(indexID text) returns text as
    $$
    declare
        --indexID   text = 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0';
        searchtext  text[];
        result      text;
        job         employmenttable%rowtype;
        a text; b text; c text; d text;
        edu         educationtable%rowtype;
    begin

    searchtext[1] = '';
    searchtext[2] = '';
    searchtext[3] = '';
    searchtext[4] = '';

    --  User Table

    select name, backgroundsummary into a, b from usertable where userID = indexID;
    searchtext[1] = concat(searchtext[1], ' ', a);
    searchtext[2] = concat(searchtext[2], ' ', b);

    --  Entity Table

    for a in select entitytag from entitytagtable where entityid = indexid::uuid
        loop
        searchtext[1] = concat(searchtext[1], ' ', a);
        end loop;

    --  Employment

    for job in select * from employmenttable where userid = indexid
        loop

        searchtext[1] = concat(searchtext[1], ' ', job.jobtitle);
        searchtext[1] = concat(searchtext[1], ' ', job.companyname);

        searchtext[2] = concat(searchtext[2], ' ', job.industry);

        searchtext[3] = concat(searchtext[3], ' ', job.location);
        searchtext[3] = concat(searchtext[3], ' ', job.summary);

        end loop;

    --  Education

    for edu in select * from educationtable where userid = indexid
        loop

        searchtext[1] = concat(searchtext[1], ' ', edu.schoolname);

        searchtext[2] = concat(searchtext[2], ' ', edu.degree);
        searchtext[2] = concat(searchtext[2], ' ', edu.emphasis);

        searchtext[3] = concat(searchtext[3], ' ', edu.summary);

        end loop;

    --  Done.  Update search column:

    update usertable set search =
        setweight(to_tsvector('english', searchtext[1]), 'A') ||
        setweight(to_tsvector('english', searchtext[2]), 'B') ||
        setweight(to_tsvector('english', searchtext[3]), 'C') ||
        setweight(to_tsvector('english', searchtext[4]), 'D')
            where userid = indexid;

    --  Update autocomplete:

    perform UpdateAutocompleteTable(searchtext[1]);
    perform UpdateAutocompleteTable(searchtext[2]);
    perform UpdateAutocompleteTable(searchtext[3]);
    perform UpdateAutocompleteTable(searchtext[4]);

    result =
        '1: ' || searchtext[1] || E'\n'
        '2: ' || searchtext[2] || E'\n'
        '3: ' || searchtext[3] || E'\n'
        '4: ' || searchtext[4] || E'\n'
        ;

    return result;
    end;
    $$
    language plpgsql;


------------------------------------------------------------------------------------------
--
--                                                                         Cards & Charges
--
------------------------------------------------------------------------------------------


create table ChargeTable
    (
     chargeID           uuid        unique not null primary key
    ,timestamp          timestamptz not null
    ,chargeStatus       smallint    not null

    ,payerID            UserID      not null
    ,purchaseType       smallint    not null
    ,purchaseTypeID     uuid        not null

    ,memoText           text        not null
    ,amount             numeric(12, 4) not null
    ,currency           text        not null
    ,chargeToken        text

    ,processorReason    text
    ,processorChargeID  text

    ,refundDate         timestamptz
    ,refundProcessorID  text
    ,refundMemo         text
    );


------------------------------------------------------------------------------------------
--
--                                                                              Categories
--
------------------------------------------------------------------------------------------


create table CategoryTable
    (
     item               text        not null
    ,parent             text        not null
    ,descriptiontext    text
    ,isLeaf             boolean     not null
    );
create unique index CategoryIndex on
    CategoryTable(item, parent);


------------------------------------------------------------------------------------------
--
--                                                                        Pretty Functions
--
------------------------------------------------------------------------------------------


create or replace
function pretty_size(sz bigint) returns text as
    $$
    begin
    return pg_size_pretty(sz);
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create or replace
function pretty_int(sz bigint) returns text as
    $$
    begin
    return to_char(sz, 'FM999,999,999,999,999');
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create or replace
function pretty_float(fl double precision) returns text as
    $$
    begin
    return to_char(fl, 'FM999,999,999,999,999.000');
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create or replace
function trimwhitespace(s text) returns text as
    $$
    begin
    return regexp_replace(regexp_replace(s, '^\s+', ''), '\s+$', '');
    end;
    $$
    language plpgsql
    immutable
    returns null on null input;


------------------------------------------------------------------------------------------
--
--                                                                          Default Values
--
------------------------------------------------------------------------------------------


insert into usertable (userid, name, userstatus, creationdate, isexpert, isFree) values
    ('a8277a5e-b461-476b-9f4a-922a50b97f26', 'Team Blitz', 5, 'Jun  1 12:00:00 GMT 2016', false, true);

