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
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
    )


//----------------------------------------------------------------------------------------
//                                                                        UserIsConfirming
//----------------------------------------------------------------------------------------


func UserIsConfirming(session *Session, confirmation *BlitzMessage.ConfirmationRequest,
        ) *BlitzMessage.ServerResponse {

    Log.LogFunctionName()

    var verified bool = false

    if  confirmation.ContactInfo == nil ||
        confirmation.ContactInfo.Contact == nil ||
        confirmation.UserProfile == nil ||
        confirmation.UserProfile.UserID == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, nil)
    }

    profile := ProfileForUserID(nil, *confirmation.UserProfile.UserID)
    //profile := confirmation.UserProfile;
    Log.Debugf("Confirmation contact: %+v.", confirmation.ContactInfo.Contact)
    Log.Debugf("Confirming profile:\n%+v.", profile)

    confirmation.ContactInfo.Contact = Util.CleanStringPtr(confirmation.ContactInfo.Contact)

    for _, contactInfo := range profile.ContactInfo {
        if contactInfo.Contact != nil && *contactInfo.Contact == *confirmation.ContactInfo.Contact {
            verified = true
            contactInfo.IsVerified = &verified
            Log.Debugf("Verified contact detail %s.", *contactInfo.Contact)
            break
        }
    }

    var error error = nil

    if config.TestingEnabled && strings.HasPrefix(*confirmation.ContactInfo.Contact, "555") {

        error = nil
        verified = true

    } else if verified {

        code := "XXXX"
        if  confirmation.ConfirmationCode != nil {
            code = strings.TrimSpace(*confirmation.ConfirmationCode)
        }

        row := config.DB.QueryRow(
            `select code, codeDate from UserContactTable
                where userID = $1
                  and contactType = $2
                  and contact = $3`,
            profile.UserID,
            confirmation.ContactInfo.ContactType,
            confirmation.ContactInfo.Contact,
        )
        var ( dbCode sql.NullString; dbCodeDate pq.NullTime)
        error = row.Scan(&dbCode, &dbCodeDate)
        if error != nil {
            Log.LogError(error)
        }

        Log.Debugf("Code '%s' Secret '%s'.", code, dbCode.String)

        if error != nil || !dbCode.Valid || !dbCodeDate.Valid || code != dbCode.String ||
            time.Since(dbCodeDate.Time) > (time.Hour * 24) {
            Log.Errorf("Confirmation secret wrong. %s != %s.", *confirmation.ConfirmationCode, dbCode.String)
            error = fmt.Errorf("The confirmation code does not match.")
            verified = false
        }

    } else {

        error = errors.New("Error: Didn't verify contact detail with profile.");

    }

    if ! verified {
        Log.LogError(error)
        UpdateProfileStatusForUserID(*profile.UserID, BlitzMessage.UserStatus_USConfirming)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    var result sql.Result
    result, error = config.DB.Exec(
        `update UserContactTable set
            code = NULL, codeDate = current_timestamp, isVerified = true
            where userid = $1
              and contacttype = $2
              and contact = $3`,
        *profile.UserID,
        confirmation.ContactInfo.ContactType,
        confirmation.ContactInfo.Contact,
    )
    error = pgsql.UpdateResultError(result, error)
    if error != nil {
        Log.LogError(error)
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, error)
    }

    UpdateProfileStatusForUserID(*profile.UserID, BlitzMessage.UserStatus_USConfirmed)
    profile = ProfileForUserID(session, *profile.UserID)

    confirmed := BlitzMessage.ConfirmationRequest {
        UserProfile: profile,
    }
    responseCode    := BlitzMessage.ResponseCode_RCSuccess
    responseMessage :=
        config.Localizef(
            "kConfirmConfirmedWelcome", "Hello %s\nWelcome to %s",
            *profile.Name,
            config.AppName,
        )
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseMessage:    &responseMessage,
        ResponseType:       &BlitzMessage.ResponseType { ConfirmationRequest: &confirmed },
    }
    return response
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
            BlitzMessage.UserMessageType_MTNotification,
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

    currentProfile := ProfileForUserID(session, session.UserID)
    if currentProfile == nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No Profile"))
    }

    inviteProfile := ProfileForUserID(session, *invite.UserID)
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

    currentProfile = ProfileForUserID(session, resultProfileID)
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

