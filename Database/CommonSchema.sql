
--  Happiness Common Database Schema
--
--  E.B.Smith  -  November 2014


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
    ,imageURL           text[]
    );


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
     userID             UserID not null
    ,service            text not null
    ,socialID           text not null
    ,userName           text
    ,displayName        text
    ,URI                text
    ,authToken          text
    ,authExpire         timestamptz
    );
create unique index SocialTableUniqueIndex on SocialTable(userID, service, socialID);


create domain WeatherType as smallint;
create domain PrecipitationType as smallint;


create type Weather as
    (
     weatherType        WeatherType
    ,temperature        real
    ,cloudCover         real
    ,precipitation      real
    ,precipitationType  PrecipitationType
    ,pressure           real
    ,windSpeed          real
    ,windBearing        real
    );


create type UserResponse as
    (
     emotionID          int
    ,emotionCount       int
    ,emotionValue       real
    );


create type ScoreComponent as
    (
     label              text
    ,score              real
    );


create table ScoreTable
    (
     userID             UserID          not null
    ,timestamp          timestamptz     not null
    ,previousTimestamp  timestamptz
    ,previousBaseScore  real
    ,happyScore         real
    ,baseScore          real
    ,displayScore       real
    ,physical           real
    ,mental             real
    ,vital              real
    ,environmental      real
    ,components         ScoreComponent[]
    ,location           Location
    ,weather            Weather
    ,testID             text
    ,userResponse       UserResponse[]
    ,userTestAssessment real
    ,unique(userID, timestamp)
    );
create unique index ScoreTableUniqueIndex on ScoreTable(userID, timestamp);


-- create table UserDeviceTable
--     (
--      userID             UserID      not null primary key
--     ,deviceID           DeviceID
--     );
-- create unique index UserDeviceTableUniqueIndex on UserDeviceTable(userID, deviceID);


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


create domain StoryType as smallint;


create table StoryTable
    (
     storyID            text unique not null primary key
    ,storyType          StoryType
    ,creationDate       timestamptz
    ,happyScore         real
    ,storyText          text
    ,storyAttribution   text
    );


create domain FriendStatus as smallint;

create function StringFromFriendStatus(friendStatus FriendStatus) returns text as
    $$
    declare
        statusString text[] := array
          [ 'FriendStatusUnknown',
            'FriendStatusInviter',
            'FriendStatusInvitee',
            'FriendStatusIgnored',
            'FriendStatusAccepted',
            'FriendStatusCircleDeprecated' ];
    begin
    if friendStatus is null then return null; end if;
    return statusString[friendStatus+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table FriendTable
    (
     userID         UserID not null
    ,friendID       UserID not null
    ,friendStatus   FriendStatus not null check (friendStatus > 0)
    ,isInCircle     boolean not null default false
    );
create unique index FriendUniqueIndex on FriendTable(userID, friendID);
create index FriendIndex on FriendTable(friendID);


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


create table MessageStatTable
    (
     timestamp          timestamptz unique not null primary key
    ,elapsed            real
    ,message            text
    ,bytesIn            int
    ,bytesOut           int
    ,statusCode         int
    ,responseCode       int
    ,responseMessage    text
    );


create domain MessageType as smallint;


create function StringFromMessageType(messageType MessageType) returns text as
    $$
    declare

    statusString text[] := array
        [ 'MessageTypeUnknown',
          'MessageTypeJoined',
          'MessageTypeFriendRequest',
          'MessageTypeFriendAccept',
          'MessageTypeFriendCircle',
          'MessageTypeScored',
          'MessageTypeScoreRequest',
          'MessageTypeHearted',
          'MessageTypeHeartedBack',
          'MessageTypeSystem' ];

    begin
    if messageType is null then return null; end if;
    return statusString[messageType+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table MessageTable
    (
     messageID          UUID            not null
    ,senderID           UserID          not null
    ,recipientID        UserID          not null
    ,creationDate       timestamptz     not null
    ,notificationDate   timestamptz
    ,readDate           timestamptz
    ,messageType        MessageType     not null
    ,messageText        text
    ,actionIcon         text
    ,actionURL          text
    );
create unique index MessageUniqueIndex on MessageTable(messageID, senderID, recipientID);
create index MessageDeliveryIndex on MessageTable(recipientID, creationDate);


create domain ImageContent as smallint;


create function StringFromImageContent(imageContent ImageContent) returns text as
    $$
    declare

    labels text[] := array
        [ 'ImageContentUnknown',
          'ImageContentProfile' ];

    begin
    if messageType is null then return null; end if;
    return labels[messageType+1];
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create table ImageTable
    (
     userID             UserID      unique not null primary key
    ,imageContent       ImageContent
    ,contentType        text
    ,crc32              int8
    ,imageData          bytea
    );


create table TestTable
    (
     testID             UUID        unique not null primary key
    ,testName           text
    ,testNote           text
    ,testItems          UUID[]
    );


create table TestItemTable
    (
     testItemID         UUID        unique not null primary key
    ,itemNote           text
    ,itemCaption        text
    ,averageScore       real
    ,maleScore          real
    ,femaleScore        real
    ,mental             real
    ,physical           real
    ,vital              real
    ,color              RGBColor
    ,imageOnURL         text
    ,imageOffURL        text
    );


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


create or replace function EraseUserID(eraseID UserID) returns text as
    $$
    begin
    if eraseID is null then
        return null;
        end if;
    delete from DeviceTable where userID = eraseID;
    delete from FriendTable where userID = eraseID;
    delete from FriendTable where friendID = eraseID;
    delete from ImageTable where userID = eraseID;
    delete from MessageTable where senderID = eraseID;
    delete from MessageTable where recipientID = eraseID;
    delete from ScoreTable where userID = eraseID;
    delete from SocialTable where userID = eraseID;
    delete from UserContactTable where userID = eraseID;
    delete from UserEventTable where userID = eraseID;
    delete from UserIdentityTable where userID = eraseID;
    delete from UserTable where userID = eraseID;
    delete from SessionTable where userID = eraseID;
    return 'User erased';
    end;
    $$
    language plpgsql
    returns null on null input;


create or replace function MergeUserIDIntoUserID(oldID UserID, newID UserID) returns text as
    $$
    declare
        result text;
        oldidentity text;
    begin
    if oldID is null or newID is null then return null; end if;

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

    --  ScoreTable

    with recursive merge as (
        select userid, timestamp from scoretable where userid = oldID
    )
    update scoretable set (userid) = (newID)
        from merge where scoretable.userid = merge.userid
                     and scoretable.timestamp = merge.timestamp
        and (select 1 from scoretable
            where scoretable.userid = newID
              and scoretable.timestamp = merge.timestamp) is null;

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

    --  FriendTable

    with recursive merge as (
        select userid, friendid from FriendTable where userid = oldID
    )
    update FriendTable set (userid) = (newID)
        from merge where FriendTable.userid = merge.userid
                     and FriendTable.friendid = merge.friendid
        and (select 1 from FriendTable
            where FriendTable.userid = newID
              and FriendTable.friendid = merge.friendid) is null;

    --

    with recursive merge as (
        select userid, friendid from FriendTable where friendid = oldID
    )
    update FriendTable set (friendid) = (newID)
        from merge where FriendTable.userid = merge.userid
                     and FriendTable.friendid = merge.friendid
        and (select 1 from FriendTable
            where FriendTable.userid = merge.userid
              and FriendTable.friendid = newID) is null;

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

    --  MessageTable

    with recursive merge as (
        select messageid, senderid, recipientid from MessageTable where senderid = oldID
    )
    update MessageTable set (senderid) = (newID)
        from merge where MessageTable.messageid = merge.messageid
                     and MessageTable.senderid = merge.senderid
                     and MessageTable.recipientid = merge.recipientid
        and (select 1 from MessageTable
            where MessageTable.messageid = merge.messageid
              and MessageTable.senderid = newID
              and MessageTable.recipientid = merge.recipientid) is null;

    --

    with recursive merge as (
        select messageid, senderid, recipientid from MessageTable where recipientid = oldID
    )
    update MessageTable set (recipientid) = (newID)
        from merge where MessageTable.messageid = merge.messageid
                     and MessageTable.senderid = merge.senderid
                     and MessageTable.recipientid = merge.recipientid
        and (select 1 from MessageTable
            where MessageTable.messageid = merge.messageid
              and MessageTable.senderid = merge.senderid
              and MessageTable.recipientid = newID) is null;

    --  ImageTable

    update imagetable set (userid) = (newID)
        where userid = oldID
        and (select 1 from imagetable
            where imagetable.userid = newID) is null;

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


-- Tests for UserNameFromIPAddress:
--
-- select UserNameFromIPAddress('107.3.151.67');
-- select UserNameFromIPAddress('73.162.198.203');
-- select UserNameFromIPAddress('76.14.58.112');
-- select UserNameFromIPAddress('166.171.250.116');
-- select UserNameFromIPAddress('66.87.119.207');
-- select UserNameFromIPAddress('172.56.38.238');


create or replace function UserNameFromIPAddress(ipaddress inet) returns text as
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
--                                                                                   Pulse
--
------------------------------------------------------------------------------------------


create type ColorRGB256 as
    (
     red        smallint
    ,green      smallint
    ,blue       smallint
    );


create domain PulseStatus as smallint;
create domain PulseBeatState as smallint;


create table PulseTable
    (
     pulseID            UUID            unique not null primary key
    ,senderID           UserID          not null
    ,pulseStatus        PulseStatus
    ,title              text
    ,body               text
    ,color              ColorRGB256
    ,teamIsVisible      boolean
    ,creationDate       timestamptz
    ,updateDate         timestamptz
    ,testID             UUID
    );
create index PulseSenderIndex on PulseTable(senderID);


create table PulseBeatTable
    (
     pulseID            UUID            not null
    ,beatDate           timestamptz     not null
    ,expirationDate     timestamptz     not null
    ,updateDate         timestamptz
    ,responseRate       real    --  These are all 0.0 - 1.0
    ,happyScore         real
    ,components         ScoreComponent[]
    );
create unique index PulseBeatDateIndex on PulseBeatTable(pulseID, beatDate);


create table PulseBeatMemberTable
    (
     pulseID            UUID            not null
    ,beatDate           timestamptz     not null
    ,memberID           UserID          not null
    ,memberPulseStatus  PulseStatus
    ,beatState          PulseBeatState
    ,scoreDate          timestamptz
    );
create unique index PulseBeatMemberTableUniqueIndex on PulseBeatMemberTable(pulseID, beatDate, memberID);
create index PulseBeatMemberTableIndex on PulseBeatMemberTable(memberID);


create table StoreTransactionTable
    (
     transactionID      UUID            not null primary key
    ,storeID            text            not null
    ,storeTransactionID text            not null
    ,userID             UserID          not null
    ,quantity           int             not null
    ,purchase           text            not null
    ,purchaseDate       timestamptz     not null
    ,locale             text
    ,localizedPrice     text
    );
create unique index StoreTransactionIndex on StoreTransactionTable(storeID, storeTransactionID);


create function pretty_size(sz bigint) returns text as
    $$
    begin
    return pg_size_pretty(sz);
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create function pretty_int(sz bigint) returns text as
    $$
    begin
    return to_char(sz, 'FM999,999,999,999,999');
    end;
    $$
    language plpgsql immutable
    returns null on null input;


create or replace function pretty_float(fl double precision) returns text as
    $$
    begin
    return to_char(fl, 'FM999,999,999,999,999.000');
    end;
    $$
    language plpgsql immutable
    returns null on null input;

