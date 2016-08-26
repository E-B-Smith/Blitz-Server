//  SendConfirmation  -  Send a user confirmation via email / sms.
//
//  E.B.Smith  -  March, 2014


package main


import (
    "fmt"
    "time"
    "bytes"
    "errors"
    "net/url"
    "strconv"
    "strings"
    "html/template"
    "database/sql"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                                            LoginAsAdmin
//----------------------------------------------------------------------------------------


func LoginAsAdmin(session *Session, login *BlitzMessage.LoginAsAdmin,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select isAdmin from UserTable where userID = $1;`,
        session.UserID,
    )
    var isAdmin sql.NullBool
    error := row.Scan(&isAdmin)
    if error != nil {
        Log.LogError(error)
    } else if isAdmin.Valid && isAdmin.Bool {

        userID := BlitzMessage.Default_Global_BlitzUserID
        login.AdminProfile = ProfileForUserID(userID, userID)
        session.UserID = userID

        response := &BlitzMessage.ServerResponse {
            ResponseCode:       BlitzMessage.ResponseCode(BlitzMessage.ResponseCode_RCSuccess).Enum(),
            ResponseType:       &BlitzMessage.ResponseType { LoginAsAdmin: login },
        }
        return response
    }

    return ServerResponseForError(BlitzMessage.ResponseCode_RCNotAuthorized, nil)
}


//----------------------------------------------------------------------------------------
//                                                                        UserIsConfirming
//----------------------------------------------------------------------------------------


func UserIsConfirming(session *Session, confirmation *BlitzMessage.ConfirmationRequest,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()
    Log.Debugf("Confirmation contact: %+v.", confirmation.ContactInfo.Contact)

    if  confirmation == nil ||
        confirmation.ContactInfo == nil ||
        confirmation.ContactInfo.Contact == nil ||
        confirmation.ContactInfo.ContactType == nil ||
        confirmation.ConfirmationCode == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    var error error = nil
    code := strings.TrimSpace(*confirmation.ConfirmationCode)
    confirmation.ContactInfo.Contact = Util.CleanStringPtr(confirmation.ContactInfo.Contact)

    var dbUserID sql.NullString
    if config.TestingEnabled && strings.HasPrefix(*confirmation.ContactInfo.Contact, "555") {

        //  Figure out the code for testing:

        row := config.DB.QueryRow(
            `select code from usercontacttable
                where contact = $1
                  and contacttype = $2
                  and code is not null
                order by codedate desc
                limit 1;`,
            confirmation.ContactInfo.Contact,
            confirmation.ContactInfo.ContactType,
        )
        var dbCode sql.NullString
        error = row.Scan(&dbCode)
        if error == nil && dbCode.Valid {
            Log.Debugf("Found code '%s'.", dbCode.String)
            code = dbCode.String
        }

    }

    //  Query the confirm code:

    row := config.DB.QueryRow(
        `select userid, codedate from usercontacttable
            where contact = $1
              and contacttype = $2
              and code = $3
            order by codedate
            limit 1;`,
        confirmation.ContactInfo.Contact,
        confirmation.ContactInfo.ContactType,
        code,
    )

    var dbCodeDate pq.NullTime
    error = row.Scan(&dbUserID, &dbCodeDate)
    if error != nil || !dbUserID.Valid || !dbCodeDate.Valid ||
        time.Since(dbCodeDate.Time) > (time.Hour * 24) ||
        len(code) == 0 {
        Log.LogError(error)
        error = fmt.Errorf("The confirmation code does not match or expired.")
    }

    //  Great!  We've confirmed.

    //  Now find the earliest verfied contact.
    //  This is our real profile.

    row = config.DB.QueryRow(
        `select c.userID, c.codeDate from usercontacttable c
             join UserTable u on u.userID = c.userID
            where contact = $1
              and contacttype = $2
              and isVerified = true
              and userStatus >= $3
            order by codedate desc nulls first
            limit 1;`,
        confirmation.ContactInfo.Contact,
        confirmation.ContactInfo.ContactType,
        BlitzMessage.UserStatus_USConfirming,
    )
    var oldestUserID sql.NullString
    var oldestDate pq.NullTime
    error = row.Scan(&oldestUserID, &oldestDate)

    //  Is it a referral?

    confirmation.ReferralCode = Util.CleanStringPtr(confirmation.ReferralCode)
    if confirmation.ReferralCode != nil {
        row = config.DB.QueryRow(
            `select referreeID, createDate from ReferralTable
                where referralCode = $1
                  and codeUseDate is null;`,
            confirmation.ReferralCode,
        )
        var referreeID sql.NullString
        var referralDate pq.NullTime
        error = row.Scan(&referreeID, &referralDate)
        if error == nil && referreeID.Valid {

            if oldestUserID.Valid &&
               oldestUserID.String != referreeID.String {

                row = config.DB.QueryRow(
                    `select MergeUserIDIntoUserID($1, $2);`,
                    referreeID.String,
                    oldestUserID.String,
                )
                var result sql.NullString
                error = row.Scan(&result)
                if error != nil || ! result.Valid || result.String != "User merged" {
                    Log.Errorf("Can't merge user! Error: %v result: %+v.", error, result)
                }
                referreeID = oldestUserID
            }

            _, error = config.DB.Exec(
                `update ReferralTable set
                    referreeID = $1,
                    codeUseDate = $2
                    where referralCode = $3;`,
                referreeID,
                time.Now(),
                confirmation.ReferralCode,
            )
            oldestUserID = referreeID
        }
    }

    //  Now we've mybe got out userID --

    if oldestUserID.Valid && len(oldestUserID.String) > 10 {

        //  An older profile exists.  Merge current profile
        //  into profile.

        row = config.DB.QueryRow(
            `select MergeUserIDIntoUserID($1, $2);`,
            dbUserID.String,
            oldestUserID.String,
        )
        var result sql.NullString
        error = row.Scan(&result)
        if error != nil || ! result.Valid || result.String != "User merged" {
            Log.Errorf("Can't merge user! Error: %v result: %+v.", error, result)
        }
        session.UserID = oldestUserID.String

    } else {

        //  Else this is a new profile.  Mark it verified:

        var result sql.Result
        result, error = config.DB.Exec(
            `update UserContactTable set
                code = NULL, codeDate = current_timestamp, isVerified = true
                where userid = $1
                  and contacttype = $2
                  and contact = $3`,
            dbUserID.String,
            confirmation.ContactInfo.ContactType,
            confirmation.ContactInfo.Contact,
        )
        error = pgsql.UpdateResultError(result, error)
        if error != nil {
            Log.LogError(error)
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
        }
        session.UserID = dbUserID.String
        SetupNewUser(session)
    }

    //  Delete old un-confirmed contact info --

    var result sql.Result
    result, error = config.DB.Exec(
        `delete from UserContactTable
            where contactType = $1
              and contact = $2
              and (isVerified = false or isVerfied is null);`,
        confirmation.ContactInfo.ContactType,
        confirmation.ContactInfo.Contact,
    )
    var r int64
    if error == nil {
        r, _ = result.RowsAffected()
    }
    Log.Debugf("Deleted %d old contacts with error %+v.", r, error)

    UpdateProfileStatusForUserID(session.UserID, BlitzMessage.UserStatus_USConfirmed)
    profile := ProfileForUserID(session.UserID, session.UserID)

    var message *string
    if oldestUserID.Valid && len(*profile.Name) > 0 {
        messageS :=
            config.Localizef(
                "kConfirmConfirmedWelcome", "Hello %s\nWelcome back to %s",
                *profile.Name,
                config.AppName,
            )
        message = &messageS
    }

    confirmed := BlitzMessage.ConfirmationRequest {
        UserProfile: profile,
    }
    responseCode    := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseMessage:    message,
        ResponseType:       &BlitzMessage.ResponseType { ConfirmationRequest: &confirmed },
    }
    return response
}


//----------------------------------------------------------------------------------------
//                                                                            SetupNewUser
//----------------------------------------------------------------------------------------


func SetupNewUser(session *Session) {
    Log.LogFunctionName()

    //  Add Blitz Assistant as a friend and follow user:

    result, error := config.DB.Exec(
        `insert into entitytagtable
            (entityid, entitytype, entitytag, userid) values
            ($1::uuid, 1, '.friend', $2),
            ($2::uuid, 1, '.friend', $1),
            ($1::uuid, 1, '.followed', $2)
            on conflict do nothing;`,
        session.UserID,
        BlitzMessage.Default_Global_BlitzUserID,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
    }

    //  Send a welcome message

    //  Create a conversation:

    convReq := BlitzMessage.ConversationRequest {
        Conversation:           &BlitzMessage.Conversation {
            InitiatorID:        proto.String(BlitzMessage.Default_Global_BlitzUserID),
            ExpertID:           proto.String(session.UserID),
            ConversationType:   BlitzMessage.ConversationType_CTConversation.Enum(),
        },
    }
    resp := StartConversation(session, &convReq)
    message := "Hello, I'm here to help you."

    //  Send an initial message:

    error = SendUserMessageInternal(
        BlitzMessage.Default_Global_BlitzUserID,
        []string { session.UserID },
        *resp.ResponseType.ConversationResponse.Conversation.ConversationID,
        message,
        BlitzMessage.UserMessageType_MTConversation,
        "",
        "",
    )
}


//----------------------------------------------------------------------------------------
//                                                                        UserConfirmation
//----------------------------------------------------------------------------------------


func UserConfirmation(session *Session, confirmation *BlitzMessage.ConfirmationRequest,
    ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var error error
    profile := confirmation.UserProfile;
    if profile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A profile is required."))
    }
/*
    if profile.Name != nil {
        profile.Name = StringPtrFromString(strings.TrimSpace(*profile.Name))
    }
    if profile.Name == nil || len(*profile.Name) <= 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A name is required."))
    }
*/
    contact := confirmation.ContactInfo
    if contact == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Contact info is required."))
    }
    if contact.Contact == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Contact info detail is required."))
    }

    //  See if we're confirming a confirmation request --

    if  confirmation.ConfirmationCode != nil &&
        len(*confirmation.ConfirmationCode) > 0 {
        return UserIsConfirming(session, confirmation)
    }

    longSecret := Util.NewUUIDString()
    i, _ := strconv.ParseInt(longSecret[0:4], 16, 32)
    confirmCode := fmt.Sprintf("%05d", i)
    message := config.Localizef("kConfirmConfirmingContact", "Confirming contact info...")
    message  = url.QueryEscape(message)
    link := config.Localizef("kConfirmAppURL", "%s/?action=confirm&message=%s&code=%s&contact=%s",
                config.AppLinkURL, message, confirmCode, *contact.Contact)
    link, error = LinkShortner_ShortLinkFromLink(link)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    name := "Stranger"
    if profile.Name != nil && len(*profile.Name) > 0 {
        name = *profile.Name
    }

    var templateMap = struct {
        UserName        string
        AppName         string
        AppDeepLink     template.HTML
        AuthCode        string
    } {
        name,
        config.AppName,
        template.HTML(link),
        confirmCode,
    }

    var confirmationMessage bytes.Buffer
    error = config.Template.ExecuteTemplate(&confirmationMessage, "ConfirmAccountEmail.html", templateMap)
    if error != nil { Log.LogError(error) }

    var confirmationSubject bytes.Buffer
    error = config.Template.ExecuteTemplate(&confirmationSubject, "ConfirmAccountEmailSubject.html", templateMap)
    if error != nil { Log.LogError(error) }

    switch *contact.ContactType {
        //case BlitzMessage.ConfirmationType_CTFacebook:

        case BlitzMessage.ContactType_CTEmail:
            email, error := Util.ValidatedEmailAddress(*contact.Contact)
            if error != nil {
                return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A valid email address is required."))
            }
            error = config.SendEmail(email, confirmationSubject.String(), confirmationMessage.String())
            if error != nil {
                return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
            }

        case BlitzMessage.ContactType_CTPhoneSMS:
            phone, error := Util.ValidatedPhoneNumber(*contact.Contact)
            if error != nil {
                return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A valid phone number is required."))
            }
            if config.TestingEnabled && strings.HasPrefix(phone, "555") {
            } else {
                error = Util.SendSMS(phone, confirmationMessage.String())
                if error != nil {
                    return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
                }
            }

        default:
            return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Unknown confirmation type %d.", contact.ContactType))
    }

    AddContactInfoToUserID(*profile.UserID, contact)
    UpdateProfileStatusForUserID(*profile.UserID, BlitzMessage.UserStatus_USConfirming)

    var result sql.Result
    result, error = config.DB.Exec(
        `update UserContactTable set
            code = $1, codeDate = current_timestamp
            where userid = $2
              and contacttype = $3
              and contact = $4`,
        confirmCode,
        *profile.UserID,
        contact.ContactType,
        contact.Contact,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
    }

    //  Send a 'friend accepted' notification --

    if confirmation.InviterUserID != nil {
        message := fmt.Sprintf("%s accepted your connection.", *profile.Name)
        SendUserMessageInternal(
            BlitzMessage.Default_Global_SystemUserID,
            []string{ *confirmation.InviterUserID },
            "",
            message,
            BlitzMessage.UserMessageType_MTActionNotification,
            "AppIcon",
            "",
        )
    }

    return response
}


//----------------------------------------------------------------------------------------
//                                                                            AcceptInvite
//----------------------------------------------------------------------------------------


func CompareTime(a, b time.Time) int {
    switch {
    case a.Before(b):
        return -1
    case a.After(b):
        return 1
    default:
        return 0
    }
}


func AcceptInviteRequest(session *Session, invite *BlitzMessage.UserInvite,
        ) *BlitzMessage.ServerResponse {
    //
    //  AcceptInvite
    //
    //  1. If existing profile is already confirmed => Possible friend request, possible message.
    //  2. If currentID is older than invite ID     =>  ditto
    //  3. If currentID is already confirmed        =>  ditto
    //  4. If no contact info provided              =>  ditto
    //  4. Then:
    //     Merge currentID into inviteID
    //     Confirm profile with contact info
    //     => Possible friend request, possible message.

    Log.LogFunctionName()

    if  invite.ContactInfo == nil || invite.UserID == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No contact info"))
    }

    currentProfile := ProfileForUserID(session.UserID, session.UserID)
    if currentProfile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No Profile"))
    }

    inviteProfile := ProfileForUserID(session.UserID, *invite.UserID)
    if inviteProfile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No Profile"))
    }

    var oldProfileID string
    var resultProfileID string

    if session.UserID == *invite.UserID {
        resultProfileID = session.UserID
    } else {
        if CompareTime(currentProfile.CreationDate.Time(), inviteProfile.CreationDate.Time()) < 0 {
            oldProfileID = *inviteProfile.UserID
            resultProfileID = *currentProfile.UserID
        } else {
            oldProfileID = *currentProfile.UserID
            resultProfileID = *inviteProfile.UserID
        }

        error := MergeProfileIDIntoProfileID(oldProfileID, resultProfileID)
        if error != nil {
            return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
        }
    }

    currentProfile = ProfileForUserID(session.UserID, resultProfileID)
    var status BlitzMessage.UserStatus = BlitzMessage.UserStatus_USConfirmed
    currentProfile.UserStatus = &status
    if currentProfile.ContactInfo == nil { currentProfile.ContactInfo = make([]*BlitzMessage.ContactInfo, 0, 1) }
    currentProfile.ContactInfo = append(currentProfile.ContactInfo, invite.ContactInfo)
    UpdateProfile(currentProfile)

    //  Write the response --

    var profiles []*BlitzMessage.UserProfile = make([]*BlitzMessage.UserProfile, 0, 10)
    profiles = append(profiles, currentProfile)

    inviteResponse := BlitzMessage.UserInvite {
        UserID:     &resultProfileID,
        FriendID:   invite.FriendID,
        Message:    invite.Message,
        Profiles:   profiles,
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        ResponseType:   &BlitzMessage.ResponseType { AcceptInviteResponse: &inviteResponse},
    }
    return response
}

