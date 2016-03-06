//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package main


import (
    "fmt"
    "time"
    "errors"
    "strings"
    "net/url"
    "net/http"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "github.com/golang/protobuf/proto"
    "happiness"
    )


//----------------------------------------------------------------------------------------
//                                                             SendMessageFromUserToFriend
//----------------------------------------------------------------------------------------



func SendMessageTypeFromUserToFriend(messageType happiness.MessageType, userID string, friendID string, userName *string) {


    message    := ""
    name       := ""
    if userName != nil {
        name = strings.TrimSpace(*userName)
    }

    switch messageType {

    case happiness.MessageType_MTFriendRequest:
        message = config.Localizef("kFriendAskJoin", "%s wants to join you.", name)

    case happiness.MessageType_MTFriendAccept:
        message = config.Localizef("kFriendAccepted", "%s accepted your friend request.", name)

    default:
        panic(fmt.Errorf("Unknown message type %v.", messageType))
    }

    SendAppMessage(userID,
        []string { friendID },
        message,
        messageType,
        "",
        "",
    )
}


//----------------------------------------------------------------------------------------
//                                                                      UpdateFriendStatus
//----------------------------------------------------------------------------------------


/*
enum FriendStatus {
  FSUnknown     = 0;
  FSInviter     = 1;
  FSInvitee     = 2;
  FSIgnored     = 3;
  FSAccepted    = 4;
  FSCircleDeprecated = 5;
}
*/


func FriendStatusForFriend(userID *string, friendID *string) happiness.FriendStatus {
    friendStatus := happiness.FriendStatus_FSUnknown
    if userID == nil || friendID == nil { return friendStatus }


    rows, error := config.DB.Query(
        "select friendStatus from FriendTable where userID = $1 and friendID = $2;", userID, friendID)
    if  rows != nil {
        defer rows.Close()
        if rows.Next() { rows.Scan(&friendStatus) }
    }

    Log.Debugf("Found friend status %d error %v.", friendStatus, error)
    return friendStatus
}


func UpdateDatabaseFriendStatus(
        userID *string,
        friendID *string,
        friendStatus happiness.FriendStatus) {
    if (userID == nil || friendID == nil) { return; }

    _, error := config.DB.Exec(
        "insert into FriendTable (userid, friendid, friendstatus)" +
        "  values ($1, $2, $3);",
            userID, friendID, friendStatus)

    if error != nil {
        update, error := config.DB.Exec(
            `update FriendTable set (userid, friendid, friendstatus)
              = ($1::userid, $2::userid, $3) where userid = $1::userid and friendid::userid = $2;`,
            userID, friendID, friendStatus)

        var rowsUpdated int64 = 0
        if update != nil { rowsUpdated, _ = update.RowsAffected() }

        if error != nil || rowsUpdated == 0 {
            Log.Errorf("Error updating friend status for '%s' and '%s': %v",
                *userID, *friendID, error)
        }
    }
}


func UpdateDatabaseFriendCircleStatus(userID string, friendID string, isInCircle bool) {
    _, error := config.DB.Exec(
        "update FriendTable set (isInCircle) = ($3) where userid = $1 and friendid = $2;",
        userID, friendID, isInCircle)
    if error != nil { Log.LogError(error); }
}



//----------------------------------------------------------- UpdateFriendStatusWithStatus

/*
enum FriendStatus {
  FSUnknown     = 0;
  FSInviter     = 1;
  FSInvitee     = 2;
  FSIgnored     = 3;
  FSAccepted    = 4;
  FSCircleDeprecated = 5;
}
*/

func UpdateFriendStatusWithStatus(userID *string, friendID *string, friendStatus happiness.FriendStatus) {
    if userID == nil || friendID == nil { return; }

    if friendStatus < happiness.FriendStatus_FSIgnored {
        return;
    }

    user2FriendStatus := FriendStatusForFriend(userID, friendID)
    if user2FriendStatus < happiness.FriendStatus_FSInvitee ||
        user2FriendStatus == friendStatus {
        return;
    }

    if friendStatus == happiness.FriendStatus_FSIgnored {
        UpdateDatabaseFriendStatus(userID, friendID, friendStatus)
        UpdateDatabaseFriendStatus(friendID, userID, happiness.FriendStatus_FSAccepted)
        return
    }

    //  Else accepted:

    UpdateDatabaseFriendStatus(userID, friendID, happiness.FriendStatus_FSAccepted)

    friend2UserStatus := FriendStatusForFriend(friendID, userID)

    if friend2UserStatus == happiness.FriendStatus_FSIgnored ||
       friend2UserStatus == happiness.FriendStatus_FSAccepted {
       return
    }

    UpdateDatabaseFriendStatus(friendID, userID, happiness.FriendStatus_FSAccepted)

    //  Send friend accept --

    var name *string = nil
    userProfile := ProfileForUserID(*userID)
    if userProfile != nil { name = userProfile.Name; }
    SendMessageTypeFromUserToFriend(happiness.MessageType_MTFriendAccept, *userID, *friendID, name)
}



//----------------------------------------------------------------------------------------
//
//                                                                           FriendRequest
//
//----------------------------------------------------------------------------------------


func FriendRequest(writer http.ResponseWriter, session *Session, friendRequest *happiness.FriendRequest) {
    //  Makes a friend request:
    //  * Tries to match the friend to an existing user.
    //  * If existing, update profile and return 'invited'.
    //  * If not, create a new user and return 'invitation pending'.

    Log.LogFunctionName()
    if friendRequest.FriendProfile == nil {
        GetShareLink(writer, session, friendRequest)
        return
    }

    responseCode := happiness.ResponseCode_RCSuccess
    responseMessage := ""

    userID := session.UserID
    var profile *happiness.Profile = nil
    var userProfile *happiness.Profile = ProfileForUserID(userID)
    if userProfile == nil {
        panic(fmt.Errorf("No user profile for %s.", userID))
    }

    //  Save the contact info --

    var contact *happiness.ContactInfo = nil
    if friendRequest.FriendProfile != nil &&
       friendRequest.FriendProfile.ContactInfo != nil {
       for _, contactInfo := range(friendRequest.FriendProfile.ContactInfo) {
           if  contactInfo.IsVerified != nil && *contactInfo.IsVerified {
               contact = contactInfo
               break
           }
        }
    }

    if (friendRequest.FriendProfile.UserID != nil) {
        profile = ProfileForUserID(*friendRequest.FriendProfile.UserID)
        if profile != nil { Log.Debugf("Found existing profile with same userID.") }
    }
    if profile == nil {
        //  No profile.  See if we can find a related profile.
        identities := IdentityStringsFromProfile(friendRequest.FriendProfile)
        profile = ExistingProfileFromIdentities(identities)
        if profile != nil {
            Log.Debugf("Found existing profile from profile identities.")
            UpdateProfile(friendRequest.FriendProfile)
            MergeProfileIDIntoProfileID(*friendRequest.FriendProfile.UserID, *profile.UserID)
            profile = ProfileForUserID(*profile.UserID)
            if (contact != nil) {
                profile.AddContactInfo(contact)
                UpdateProfile(profile)
            }
        }
    }
    if profile == nil {
        //  Still no profile.  Create a new profile for this user.
        Log.Debugf("No existing profile found for friend.  Creating a new profile.")
        profile = friendRequest.FriendProfile
        if profile.CreationDate == nil {
            profile.CreationDate = happiness.TimestampFromTime(time.Now())
        }
        UpdateProfile(profile)
    }

    if userID == *profile.UserID {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("You are your own best friend!"))
        return
    }

    Log.Debugf("Profile is %v.", profile)
    friendID := *profile.UserID

    //  Update friend status --

    user2FriendStatus := FriendStatusForFriend(&userID, &friendID)
    friend2UserStatus := FriendStatusForFriend(&friendID, &userID)


    switch {

    case user2FriendStatus == happiness.FriendStatus_FSAccepted &&
         friend2UserStatus == happiness.FriendStatus_FSAccepted:
         responseMessage = config.Localizef("kFriendsAlreadyFriends", "You are already friends!")

    case friend2UserStatus == happiness.FriendStatus_FSIgnored ||
         friend2UserStatus == happiness.FriendStatus_FSAccepted:

         user2FriendStatus = happiness.FriendStatus_FSAccepted
         UpdateDatabaseFriendStatus(&userID, &friendID, user2FriendStatus)
         responseMessage = config.Localizef("kFriendsNewFriend", "You are now friends with %s.", *profile.Name)

    case friend2UserStatus == happiness.FriendStatus_FSInviter:

         user2FriendStatus = happiness.FriendStatus_FSAccepted
         UpdateDatabaseFriendStatus(&userID, &friendID, happiness.FriendStatus_FSAccepted)
         UpdateDatabaseFriendStatus(&friendID, &userID, happiness.FriendStatus_FSAccepted)

         responseMessage = config.Localizef("kFriendsFriend", "You are a friend with %s.", *profile.Name)
         SendMessageTypeFromUserToFriend(happiness.MessageType_MTFriendAccept, userID, friendID, userProfile.Name)

    default:
         user2FriendStatus = happiness.FriendStatus_FSInviter
         UpdateDatabaseFriendStatus(&userID, &friendID, user2FriendStatus)
         UpdateDatabaseFriendStatus(&friendID, &userID, happiness.FriendStatus_FSInvitee)
         SendMessageTypeFromUserToFriend(happiness.MessageType_MTFriendRequest, userID, friendID, userProfile.Name)
    }

    //  Make the response --

    shareLink := ""
    if *profile.UserStatus < happiness.UserStatus_USActive && contact != nil {
        name := ""
        if userProfile.Name != nil {
            name = *userProfile.Name
        }
        message := config.Localizef("kFriendAskConnectLinkMessage", "%s wants to connect with you on BeingHappy.", name)
        message  = url.QueryEscape(message)
        contactValue := url.QueryEscape(*Util.CleanStringPtr(contact.Contact))
        shareLink = fmt.Sprintf("%s/?action=invite&friendid=%s&userid=%s&contacttype=%s&contact=%s&message=%s",
            config.AppLinkURL, userID, friendID, contact.ContactType.String(), contactValue, message)

        shareLink, _ = LinkShortner_ShortLinkFromLink(shareLink)
    }

    UpdateDatabaseFriendCircleStatus(userID, *profile.UserID, *friendRequest.IsInCircle)

    friend := happiness.Friend {
        FriendID:       profile.UserID,
        FriendStatus:   &user2FriendStatus,
        IsInCircle:     friendRequest.IsInCircle,
    }
    if len(shareLink) > 0 {
        friend.InviteLink = &shareLink
    }

    responseType := happiness.FriendResponseType_FRUpdate
    friendResponse := &happiness.FriendResponse {
        Friends:            []*happiness.Friend  { &friend },
        FriendProfiles:     []*happiness.Profile { profile },
        ResponseType:       &responseType,
    }

    response := &happiness.ServerResponse {
        ResponseCode:       &responseCode,
        ResponseMessage:    &responseMessage,
        Response:           &happiness.ServerResponse_FriendResponse{ FriendResponse: friendResponse },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}



//----------------------------------------------------------------------------------------
//
//                                                                       GenerateShareLink
//
//----------------------------------------------------------------------------------------


func GenerateShareLink(userID string) (shareLink string, error error) {
    shareLink = ""
    var userProfile *happiness.Profile = ProfileForUserID(userID)

    name := "me"
    if userProfile != nil && userProfile.Name != nil {
        name = *userProfile.Name
    }
    message := config.Localizef("kFriendShareLinkMessage", "Connect with %s on BeingHappy.", name)
    message  = url.QueryEscape(message)

    if userProfile == nil {
        shareLink = fmt.Sprintf("%s/?action=message&message=%s",
            config.AppLinkURL,
            message)
    } else {
        shareLink = fmt.Sprintf("%s/?action=showprofile&profileid=%s&message=%s",
            config.AppLinkURL,
            *userProfile.UserID,
            message)
    }

    shareLink, error = LinkShortner_ShortLinkFromLink(shareLink)
    return shareLink, error
}


func GetShareLink(writer http.ResponseWriter, session *Session, friendRequest *happiness.FriendRequest) {
    Log.LogFunctionName()

    if friendRequest.FriendProfile != nil {
        FriendRequest(writer, session, friendRequest)
        return
    }

    shareLink, error := GenerateShareLink(session.UserID)
    if len(shareLink) <= 0 {
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }


    responseType := happiness.FriendResponseType_FRUpdate
    friendResponse := &happiness.FriendResponse {
        InviteLink:     &shareLink,
        ResponseType:   &responseType,
    }

    code := happiness.ResponseCode_RCSuccess
    response := &happiness.ServerResponse {
        ResponseCode:   &code,
        Response:       &happiness.ServerResponse_FriendResponse{ FriendResponse: friendResponse },
    }

    data, error := proto.Marshal(response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}



//----------------------------------------------------------------------------------------
//                                                                        FriendsForUserID
//----------------------------------------------------------------------------------------


func FriendsForUserID(userID string) []*happiness.Friend {
    Log.LogFunctionName()

    errorCount := 0
    var firstError error = nil
    var friends []*happiness.Friend

    rows, error := config.DB.Query("select friendID, friendStatus, isInCircle from FriendTable where userID = $1;", userID)
    defer pgsql.CloseRows(rows)
    if error != nil {
        if firstError == nil { firstError = error }
        Log.Errorf("User %s can't query friends: %v.", userID, error)
        errorCount++
    } else {
        for rows.Next() {
            var ( friendID string; friendStatus happiness.FriendStatus; isInCircle bool)
            error = rows.Scan(&friendID, &friendStatus, &isInCircle)
            if error != nil {
                if firstError == nil { firstError = error }
                Log.Errorf("Can't read friend row: %v", error)
                errorCount++
                continue
            } else {
                friend := happiness.Friend {
                    FriendID: &friendID,
                    FriendStatus: &friendStatus,
                    IsInCircle: &isInCircle,
                }
                friends = append(friends, &friend)
            }
        }
    }

    Log.Debugf("Found %d friends, %d errors.", len(friends), errorCount)

    if len(friends) > 0 {
        return friends;
    }
    return nil
}



//----------------------------------------------------------------------------------------
//
//                                                                           UpdateFriends
//
//----------------------------------------------------------------------------------------


func UpdateFriends(writer http.ResponseWriter, session *Session, friendUpdate *happiness.FriendUpdate) {
    Log.LogFunctionName()

    if len(friendUpdate.Friends) > 1 {
        SendError(writer, happiness.ResponseCode_RCInputInvalid, errors.New("Too many friends"))
        return
    }

    var response happiness.ServerResponse
    code := happiness.ResponseCode_RCSuccess
    response.ResponseCode = &code

    if len(friendUpdate.Friends) == 1 {

        friend := friendUpdate.Friends[0]
        UpdateFriendStatusWithStatus(&session.UserID, friend.FriendID, *friend.FriendStatus)
        UpdateDatabaseFriendCircleStatus(session.UserID, *friend.FriendID, *friend.IsInCircle)

    } else {

        friends := FriendsForUserID(session.UserID)

        var profiles []*happiness.Profile = make([]*happiness.Profile, 0, 10)
        for _, friend := range friends {
            profile := ProfileForUserID(*friend.FriendID)
            if profile != nil {
                profiles = append(profiles, profile)
            }
        }

        shareLink, _ := GenerateShareLink(session.UserID)

        responseType := happiness.FriendResponseType_FRList
        friendResponse := &happiness.FriendResponse {
            Friends:            friends,
            FriendProfiles:     profiles,
            InviteLink:         &shareLink,
            ResponseType:       &responseType,
        }

        response.Response = &happiness.ServerResponse_FriendResponse{ FriendResponse: friendResponse }
    }

    data, error := proto.Marshal(&response)
    if error != nil {
        Log.Errorf("Error marshaling data: %v.", error)
        SendError(writer, happiness.ResponseCode_RCServerError, error)
        return
    }

    writer.Write(data)
}

