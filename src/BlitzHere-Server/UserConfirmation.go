//  SendConfirmation  -  Send a user confirmation via email / sms.
//
//  E.B.Smith  -  March, 2014


package main


import (
    "fmt"
    "time"
    "bytes"
    "strings"
    "errors"
    "net/url"
    "net/http"
    "html/template"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "BlitzMessage"
    )


//----------------------------------------------------------------------------------------
//                                                                        UserIsConfirming
//----------------------------------------------------------------------------------------


func UserIsConfirming(writer http.ResponseWriter, session *Session, confirmation *BlitzMessage.ConfirmationRequest) {
    Log.LogFunctionName()

    var verified bool = false
    var userStatus BlitzMessage.UserStatus = BlitzMessage.UserStatus_USConfirmed
    profile := confirmation.Profile;
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

    if ! verified {
        error = errors.New("Error: Didn't verify contact detail with profile.");
    }

    if *confirmation.ConfirmationCode != session.Secret {
        error = fmt.Errorf("Confirmation secret wrong. %s != %s.", *confirmation.ConfirmationCode, session.Secret)
    }

    if ! verified {
        userStatus = BlitzMessage.UserStatus_USActive
        profile.UserStatus = &userStatus
        Log.LogError(error)
        //error = errors.New("Sorry, the confirmation has expired."))
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, error)
        return
    }

    session.Secret = Util.NewUUIDString()
    profile.UserStatus = &userStatus
    UpdateProfile(profile)

    confirmed := BlitzMessage.ConfirmationRequest {
        Profile: profile,
    }
    responseCode    := BlitzMessage.ResponseCode_RCSuccess
    responseMessage := config.Localizef("kConfirmConfirmedWelcome", "Confirmed. Welcome to BeingHappy.")
    response := &BlitzMessage.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseMessage:    &responseMessage,
        Response:           &BlitzMessage.ServerResponse_ConfirmationRequest { ConfirmationRequest: &confirmed },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}



//----------------------------------------------------------------------------------------
//                                                                        UserConfirmation
//----------------------------------------------------------------------------------------


func UserConfirmation(writer http.ResponseWriter, session *Session, confirmation *BlitzMessage.ConfirmationRequest) {
    Log.LogFunctionName()

    var error error
    profile := confirmation.Profile;
    if profile == nil {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A profile is required."))
        return
    }
    if profile.Name != nil {
        profile.Name = StringPtrFromString(strings.TrimSpace(*profile.Name))
    }
    if profile.Name == nil || len(*profile.Name) <= 0 {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A name is required."))
        return
    }
    contact := confirmation.ContactInfo
    if contact == nil {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Contact info is required."))
        return
    }
    if contact.Contact == nil {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Contact info detail is required."))
        return
    }

    //  See if we're confirming a confirmation request --

    if confirmation.ConfirmationCode != nil {
        UserIsConfirming(writer, session, confirmation)
        return
    }

    var userStatus BlitzMessage.UserStatus = BlitzMessage.UserStatus_USConfirming
    profile.UserStatus = &userStatus
    confirmCode := session.Secret
    message := config.Localizef("kConfirmConfirmingContact", "Confirming contact info...")
    message  = url.QueryEscape(message)
    link := config.Localizef("kConfirmAppURL", "%s/?action=confirm&message=%s&code=%s&contact=%s",
                config.AppLinkURL, message, confirmCode, *contact.Contact)
    link, error = LinkShortner_ShortLinkFromLink(link)
    if error != nil {
        Log.LogError(error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    var templateMap = struct {
        UserName        string
        AppName         string
        AppDeepLink     template.HTML
    } {
        *profile.Name,
        config.AppName,
        template.HTML(link),
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
                SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A valid email address is required."))
                return
            }
            error = config.SendEmail(email, confirmationSubject.String(), confirmationMessage.String())
            if error != nil {
                SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
                return
            }

        case BlitzMessage.ContactType_CTPhoneSMS:
            phone, error := Util.ValidatedPhoneNumber(*contact.Contact)
            if error != nil {
                SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("A valid phone number is required."))
                return
            }
            error = Util.SendSMS(phone, confirmationMessage.String())
            if error != nil {
                SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
                return
            }

        default:
            SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, fmt.Errorf("Unknown confirmation type %d.", contact.ContactType))
            return
    }

    UpdateProfile(profile)

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)

    //  Send a 'friend accepted' notification --

    if confirmation.InviterUserID != nil {
        message := fmt.Sprintf("%s accepted your Pulse invite.", *profile.Name)
        SendAppMessage(BlitzMessage.Default_Globals_SystemUserID,
                []string{ *confirmation.InviterUserID },
                message,
                BlitzMessage.MessageType_MTPulse,
                "Pulse", "")
    }

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


func AcceptInviteRequest(writer http.ResponseWriter, session *Session, invite *BlitzMessage.AcceptInviteRequest) {
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
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No contact info"))
        return
    }

    currentProfile := ProfileForUserID(session.UserID)
    if currentProfile == nil {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No Profile"))
        return
    }

    inviteProfile := ProfileForUserID(*invite.UserID)
    if inviteProfile == nil {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No Profile"))
        return
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
            SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
            return
        }
    }

    currentProfile = ProfileForUserID(resultProfileID)
    var status BlitzMessage.UserStatus = BlitzMessage.UserStatus_USConfirmed
    currentProfile.UserStatus = &status
    if currentProfile.ContactInfo == nil { currentProfile.ContactInfo = make([]*BlitzMessage.ContactInfo, 0, 1) }
    currentProfile.ContactInfo = append(currentProfile.ContactInfo, invite.ContactInfo)
    UpdateProfile(currentProfile)

    //  Write the response --

    var profiles []*BlitzMessage.Profile = make([]*BlitzMessage.Profile, 0, 10)
    profiles = append(profiles, currentProfile)

    friends := FriendsForUserID(resultProfileID)
    for _, friend := range friends {
        profile := ProfileForUserID(*friend.FriendID)
        if profile != nil {
            profiles = append(profiles, profile)
        }
    }

    inviteResponse := BlitzMessage.AcceptInviteResponse {
        UserID:     &resultProfileID,
        FriendID:   invite.FriendID,
        Message:    invite.Message,
        Friends:    friends,
        Profiles:   profiles,
    }

    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode:   &code,
        Response:       &BlitzMessage.ServerResponse_AcceptInviteResponse { AcceptInviteResponse: &inviteResponse},
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}

