//  BlitzHere-Server  -  Track the user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "sync"
    "time"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "violent.blue/GoKit/Util"
    "BlitzMessage"
)



//----------------------------------------------------------------------------------------
//
//                                                                                 Session
//
//----------------------------------------------------------------------------------------


type Session struct {
    UserID              string
    DeviceID            string
    SessionToken        string
    Timestamp           time.Time   //  Time session was created.
    Secret              string      //  For confirming a user's ID, etc.
    Device              BlitzMessage.DeviceInfo
    AppOptions          *BlitzMessage.AppOptions
}


type SessionLock struct {
    Map         *map[string](*Session)
    DoneChannel chan bool
}


const kSessionDuration time.Duration  = 120*time.Minute
var   SessionLockRequest              = make(chan bool)
var   SessionLockChannel              = make(chan SessionLock)


//------------------------------------------------------------------ Session_CreateSession


func Session_CreateSession(userID, deviceID string) *Session {
    Log.LogFunctionName()

    session := new(Session)
    session.UserID = userID
    session.DeviceID = deviceID
    session.SessionToken = Util.NewUUIDString()
    session.Timestamp = time.Now()
    session.Secret = Util.NewUUIDString()

    rowResult, error := config.DB.Exec("delete from SessionTable where userid = $1 and deviceID = $2;", userID, deviceID);
    if error != nil { Log.LogError(error); }
    if rowResult != nil {
        var count int64
        count, error = rowResult.RowsAffected()
        Log.Debugf("Removed %d old sessions for %s.", count, userID)
    }

    _, error = config.DB.Exec(
        `insert into SessionTable (
            userID,
            sessionToken,
            timestamp,
            deviceID,
            secret) values ($1, $2, $3, $4, $5);`,
            session.UserID,
            session.SessionToken,
            session.Timestamp,
            session.DeviceID,
            session.Secret)
    if error != nil {
        Log.LogError(error);
        return nil
    }

    SessionLockRequest <- true
    sessionLock := <- SessionLockChannel
    defer func() { sessionLock.DoneChannel <- true; } ()

    deleteArray := make([](*Session), 0, 5)
    for _, oldSession := range (*sessionLock.Map) {
        if oldSession.UserID == userID && oldSession.DeviceID == deviceID {
           deleteArray = append(deleteArray, oldSession)
        }
    }
    for _, oldSession := range deleteArray {
        Log.Debugf("Expiring %s.", oldSession.SessionToken)
        delete(*sessionLock.Map, oldSession.SessionToken)
    }

    (*sessionLock.Map)[session.SessionToken] = session
    return session
}


//-------------------------------------------------------- Session_DeleteSessionsForUserID


func Session_DeleteSessionsForUserID(userID string) {
    Log.LogFunctionName()

    rowResult, error := config.DB.Exec("delete from SessionTable where userid = $1;", userID);
    if error != nil { Log.LogError(error); }
    if rowResult != nil {
        var count int64
        count, error = rowResult.RowsAffected()
        Log.Debugf("Removed %d old sessions for %s.", count, userID)
    }

    SessionLockRequest <- true
    sessionLock := <- SessionLockChannel
    defer func() { sessionLock.DoneChannel <- true; } ()

    deleteArray := make([](*Session), 0, 5)
    for _, oldSession := range (*sessionLock.Map) {
        if oldSession.UserID == userID {
           deleteArray = append(deleteArray, oldSession)
        }
    }
    for _, oldSession := range deleteArray {
        Log.Debugf("Expiring %s.", oldSession.SessionToken)
        delete(*sessionLock.Map, oldSession.SessionToken)
    }
}


//--------------------------------------------------------------- Session_SessionFromToken


func Session_SessionFromToken(sessionToken string) *Session {
    Log.LogFunctionName()

    SessionLockRequest <- true
    sessionLock := <- SessionLockChannel
    defer func() { sessionLock.DoneChannel <- true; } ()

    session := (*sessionLock.Map)[sessionToken]
    if session != nil && time.Since(session.Timestamp) < kSessionDuration {
        return session
    }
    return nil
}


//------------------------------------------------------------- Session_InitializeSessions


var SessionExpirationTimer *time.Timer


func expire_sessions_body() {
    Log.LogFunctionName()

    SessionLockRequest <- true
    sessionLock := <- SessionLockChannel
    defer func() { sessionLock.DoneChannel <- true; } ()

    deleteArray := make([](*Session), 0, 20)
    Log.Debugf("Checking %d sessions. Max duration: %d.", len(*sessionLock.Map), kSessionDuration)
    for _, session := range *sessionLock.Map {
        Log.Debugf("Session %s since %s (%d).", session.SessionToken, session.Timestamp, time.Since(session.Timestamp))
        if time.Since(session.Timestamp) > kSessionDuration {
            deleteArray = append(deleteArray, session)
        }
    }
    for _, session := range deleteArray {
        Log.Debugf("Expiring %s.", session.SessionToken)
        delete(*sessionLock.Map, session.SessionToken)
        result, error := config.DB.Exec(
            "delete from SessionTable where SessionToken = $1;", session.SessionToken)
        rowCount := int64(0)
        if result != nil { rowCount, _ = result.RowsAffected() }
        if error != nil || rowCount != 1 {
            Log.Errorf("Deleting session error: %v Rows: %d.", error, rowCount)
        }
    }

    SessionExpirationTimer = time.AfterFunc(1 * time.Minute, expire_sessions_body)
}


func initialize_sessions_body() {
    Log.LogFunctionName()


    go func() {
        var sessionMap = make(map[string](*Session))
        for {
            <- SessionLockRequest
            sessionLock := SessionLock {
                Map:            &sessionMap,
                DoneChannel:    make(chan bool),
            }
            SessionLockChannel <- sessionLock
            <- sessionLock.DoneChannel
        }
    } ()


    expire_sessions_body()

    SessionLockRequest <- true
    sessionLock := <- SessionLockChannel
    defer func() { sessionLock.DoneChannel <- true; } ()


    count := 0
    rows, error := config.DB.Query(
        `select
            userID,
            sessionToken,
            timestamp,
            deviceID,
            secret
            from SessionTable;`)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        return
    }
    for rows.Next() {
        var session Session
        error = rows.Scan(
            &session.UserID,
            &session.SessionToken,
            &session.Timestamp,
            &session.DeviceID,
            &session.Secret,
        )
        if error != nil {
            Log.LogError(error)
        } else {
            count++
            (*sessionLock.Map)[session.SessionToken] = &session
        }
    }
    Log.Debugf("Added %d previous sessions.", count)
}


var SessionInitializedOnce sync.Once


func Session_InitializeSessions() {
    SessionInitializedOnce.Do(initialize_sessions_body)
}



//----------------------------------------------------------------------------------------
//
//                                                                           UpdateSession
//
//----------------------------------------------------------------------------------------


func UpdateSession(ipAddress string,
                   sessionToken string,
                   request *BlitzMessage.SessionRequest,
                   ) *BlitzMessage.ServerResponse {

    //  * If the User doesn't exist:
    //      - Check the vendorID, otherwise
    //      - Check the advertsingID, otherwise
    //      - Check the userTags, otherwise:
    //      => Create a record.
    //  * Update the device info.
    //  * Update the user session state.

    Log.LogFunctionName()

    //  Validate the userID --

    if request.Profile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }
    userID, error := BlitzMessage.ValidateUserID(request.Profile.UserID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    if request.DeviceInfo == nil || request.DeviceInfo.VendorUID == nil || request.DeviceInfo.AppID == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    //  Check the app version --

    var (appID string; minAppVersion string; minAppDataDate time.Time;)
    row := config.DB.QueryRow(
        "select appID, minAppVersion, minAppDataDate from AppTable where AppID = $1;",
            request.DeviceInfo.AppID)
    error = row.Scan(&appID, &minAppVersion, &minAppDataDate)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    if  request.DeviceInfo.AppVersion == nil ||
        Util.CompareVersionStrings(*request.DeviceInfo.AppVersion, minAppVersion) < 0 {
        version := "Uknown"
        if request.DeviceInfo.AppVersion != nil { version = *request.DeviceInfo.AppVersion }
        error = fmt.Errorf("Client too old.  %s < %s.", version, minAppVersion)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCClientTooOld, error)
    }

    //  Check the session --

    session := Session_SessionFromToken(sessionToken)
    if session == nil {
        session  = Session_CreateSession(userID, *request.DeviceInfo.VendorUID)
        if session == nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
    }

    request.DeviceInfo.IPAddress = &ipAddress

    //  Check for a new user --

    var invite *BlitzMessage.AcceptInviteRequest = nil
    if request.Profile.UserStatus == nil || *request.Profile.UserStatus < BlitzMessage.UserStatus_USActive {
        //  Check to see if we have an invite saved --
        Log.Debugf("Checking new user for invite...")
        invite = InviteRequestForDevice(request.DeviceInfo)
        if invite != nil {
            profile := ProfileForUserID(*invite.FriendID)
            if profile == nil {
                invite = nil
            } else {
                request.Profile = profile
                senderProfile := ProfileForUserID(*invite.UserID)
                if senderProfile != nil {
                    invite.Profiles = [] *BlitzMessage.UserProfile { senderProfile }
                }
            }
        }
    }

    //  Update user & device --

    Log.Debugf("Updating user %s.", userID)

    identities := IdentityStringsFromProfile(request.Profile)
    identities  = IdentityStringsForDevice(identities,
        request.DeviceInfo.VendorUID,
        request.DeviceInfo.AdvertisingUID,
        request.DeviceInfo.DeviceUDID)

    profile := ExistingProfileFromIdentities(identities)
    if profile == nil { profile = request.Profile }

    if userID != *profile.UserID {
        Log.Debugf("Morphing user %s into existing user %s.", userID, *profile.UserID)
        UpdateProfile(request.Profile) // Save data.
        MergeProfileIDIntoProfileID(userID, *profile.UserID) // Morph data in database.
        userID = *profile.UserID
    }

    Log.Debugf("Insert user.")
    _, error = config.DB.Exec("insert into usertable (userid, userstatus) "+
      "  values ($1, $2);", userID, BlitzMessage.UserStatus_USActive)
    Log.Debugf("Result: %v.", error)

    var status BlitzMessage.UserStatus = BlitzMessage.UserStatus_USUnknown
    if profile.UserStatus == nil {
        profile.UserStatus = &status
    }

    Log.Debugf("Update user.")
    if *profile.UserStatus == BlitzMessage.UserStatus_USBlocked {
        Log.Warningf("Blocking user %s.", *profile.UserID)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    } else if *profile.UserStatus < BlitzMessage.UserStatus_USActive {
        _, error = config.DB.Exec(
            "update usertable set (lastseen, userstatus) = ($1, $2) where userID = $3;",
             time.Now(), BlitzMessage.UserStatus_USActive, userID)
        Log.Debugf("Result: %v.", error)
    } else {
        _, error  = config.DB.Exec(
            "update usertable set (lastseen) = ($1) where userID = $2;",
             time.Now(), userID)
        Log.Debugf("Result: %v.", error)
    }

    Log.Debugf("Insert device: %+v", request.DeviceInfo)
    _, error  = config.DB.Exec("insert into devicetable (userid) values ($1);", userID)
    Log.Debugf("Result: %v.", error)

    var d = request.DeviceInfo
    _, error = config.DB.Exec(
        `update DeviceTable set (
            platformType,
            modelName,
            systemVersion,
            appID,
            appVersion,
            notificationToken,
            vendorID,
            advertisingID,
            deviceUDID,
            language,
            timezone,
            phoneCountryCode,
            screenSize,
            screenScale,
            appIsReleaseVersion,
            lastIPAddress,
            timestamp,
            systemBuildVersion,
            localIPAddress)
            = ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, row($13, $14), $15, $16, $17, $18, $19, $20)
                where userID = $21;`,
            d.PlatformType,
            d.ModelName,
            d.SystemVersion,
            d.AppID,
            d.AppVersion,
            d.NotificationToken,
            d.VendorUID,
            d.AdvertisingUID,
            d.DeviceUDID,
            d.Language,
            d.Timezone,
            d.PhoneCountryCode,
            d.ScreenSize.Width,
            d.ScreenSize.Height,
            d.ScreenScale,
            d.AppIsReleaseVersion,
            d.IPAddress,
            time.Now(),
            d.SystemBuildVersion,
            d.LocalIPAddress,
            userID)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    session.Device = *request.DeviceInfo
    //session.AppOptions = AppOptionsForSession(session)

    //UpdateProfile(profile)
    profile = ProfileForUserID(userID)

    sessionResponse := &BlitzMessage.SessionResponse {
        UserID:             &userID,
        SessionToken:       &session.SessionToken,
        UserProfile:        profile,
        InviteRequest:      invite,
        ResetAllAppData:    BoolPtrFromBool(false),
        AppOptions:         session.AppOptions,
    }


    Log.Debugf("Last app data reset date: %+v", request.LastAppDataResetDate)
    if request.LastAppDataResetDate == nil {
       sessionResponse.ResetAllAppData = BoolPtrFromBool(true)
    } else {
        resetTime := request.LastAppDataResetDate.Time()
        if resetTime.Before(minAppDataDate) {
           sessionResponse.ResetAllAppData = BoolPtrFromBool(true)
        }
    }
    Log.Debugf("ResetAllAppData: %+v", sessionResponse.ResetAllAppData)

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        Response:       &BlitzMessage.ResponseType { SessionResponse: sessionResponse },
    }
    return response
}


