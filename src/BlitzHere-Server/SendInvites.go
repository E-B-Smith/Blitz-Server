

//----------------------------------------------------------------------------------------
//
//                                                       BlitzHere-Server : SendInvites.go
//                                                                            Send invites
//
//                                                                 E.B. Smith, August 2016
//                        -©- Copyright © 2014-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package main


import (
    "fmt"
    "errors"
    "net/url"
    "database/sql"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "BlitzMessage"
)


func SendInvite(inviterUserID string, invite *BlitzMessage.UserInvite) error {
    //  No: If already a friend on Blitz, do nothing.  Done.
    //  No: If already on Blitz, send a friend request.  Done.
    //  If not on Blitz, create a user profile.
    //      - Generate an invite link.
    //      - Send friend request.

    if invite == nil {
        return errors.New("No invite")
    }

    var error error
    error = CleanContactInfo(invite.ContactInfo)
    if error != nil {
        return error
    }

    row := config.DB.QueryRow(
        `select userID from UserContactTable
            where contactType = $1
              and contact = $2;`,
        invite.ContactInfo.ContactType,
        invite.ContactInfo.Contact,
    )
    var userID sql.NullString
    error = row.Scan(&userID)
    if error != nil {
        Log.LogError(error)
    }

    var friendProfile *BlitzMessage.UserProfile
    if userID.Valid {
        friendProfile = ProfileForUserID("", userID.String)
    }

    if friendProfile == nil {
        friendProfile = &BlitzMessage.UserProfile {
            UserID:         proto.String(Util.NewUUIDString()),
            Name:           invite.Name,
            ContactInfo:    []*BlitzMessage.ContactInfo { invite.ContactInfo },
            UserStatus:     BlitzMessage.UserStatus(BlitzMessage.UserStatus_USInvited).Enum(),
        }
        UpdateProfile(friendProfile)
        friendProfile = ProfileForUserID("", *friendProfile.UserID)
    }
    if friendProfile.UserStatus == nil {
        friendProfile.UserStatus = BlitzMessage.UserStatus(BlitzMessage.UserStatus_USInvited).Enum()
    }
    // if *friendProfile.UserStatus >= BlitzMessage.UserStatus_USConfirming {
    //     return
    // }

    name := PrettyNameForUserID(inviterUserID)
    message := fmt.Sprintf("%s sent you an invitation to Blitz", name)
    if invite.Message != nil && len(*invite.Message) > 0 {
        message += ":\n\n" + *invite.Message
    }

    Log.Debugf("%v %v %v %v",
        friendProfile.UserID,
        invite.ContactInfo.ContactType,
        invite.ContactInfo.Contact,
        message,
    )

    inviteURL := fmt.Sprintf(
        "%s?action=invited&inviteeid=%s&contacttype=%d&contact=%s&message=%s",
        config.AppLinkURL,
        *friendProfile.UserID,
        *invite.ContactInfo.ContactType,
        url.QueryEscape(*invite.ContactInfo.Contact),
        url.QueryEscape(message),
    )
    shortLink, _ := LinkShortner_ShortLinkFromLink(inviteURL)
    message += "\nGet Blitz here: " + shortLink

    Log.Debugf("Invite is: %s.", message)

    switch *invite.ContactInfo.ContactType {

    case BlitzMessage.ContactType_CTPhoneSMS:
        Util.SendSMS(*invite.ContactInfo.Contact, message)

    case BlitzMessage.ContactType_CTEmail:
        config.SendEmail(*invite.ContactInfo.Contact, "Join Blitz, a network of vetted experts", message)

    default:
        Log.Errorf("Unkown contactType %d.", *invite.ContactInfo.ContactType)
    }
    invite.Profiles = []*BlitzMessage.UserProfile { friendProfile }
    return nil
}


func SendUserInvites(session *Session, invites *BlitzMessage.UserInvites,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    var firstError, error error
    for _, userInvite := range invites.UserInvites {
        error = SendInvite(session.UserID, userInvite)
        if error != nil && firstError ==nil {
            firstError = error
        }
    }

    var message *string
    var code BlitzMessage.ResponseCode = BlitzMessage.ResponseCode_RCSuccess
    if firstError != nil {
        code = BlitzMessage.ResponseCode_RCInputInvalid
        message = proto.String("Some invites not sent. (Is the invite address correct?)")
    }

    return &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    message,
        ResponseType:       &BlitzMessage.ResponseType {
            UserInvitesResponse:    invites,
        },
    }
}

