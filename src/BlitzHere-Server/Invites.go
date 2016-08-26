

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
    "math"
    "time"
    "errors"
    "strings"
    "net/url"
    "database/sql"
    "github.com/lib/pq"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//                                                                          Referral Codes
//----------------------------------------------------------------------------------------


var encodeSymbols string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"



func referralNumberFromRowCount(rc int64) int64 {
    return (rc * 100313) % 1679609
}


func rowCountFromReferralNumber(ref int64) int64 {
    return (ref * 4638) % 1679609
}


func encodeReferralNumber(ref int64) string {
    var result string
    for i := 0; i < 4; i++ {
        r := math.Remainder(float64(ref), 36)
        if r < 0 { r += 36 }
        result += string(encodeSymbols[int(r)])
        ref = ref / 36
    }
    return result
}


func decodeReferralString(ref string) int64 {
    var result int64
    var m int64 = 1
    for i := 0; i < 4; i++ {
        idx := strings.Index(encodeSymbols, string(ref[i]))
        if idx < 0 { return -1 }
        result += m*int64(idx)
        m *= 36
    }
    return result
}


func referralStringFromRowCount(rowCount int64) string {
    return encodeReferralNumber(referralNumberFromRowCount(rowCount))
}


/*
    Test

func main() {
    var i int64
    for i = 1; i < 1679609; i++ {
        if (i * 100313) % 1679609 == 1 {
            fmt.Println("Inverse is ", i)
            break
        }
    }

    var rc1 int64 = 20
    ref := referralNumberFromRowCount(rc1)
    rc2 := rowCountFromReferralNumber(ref)
    enc := encodeReferralNumber(ref)
    dec := decodeReferralString(enc)
    fmt.Println("Rc1:", rc1, "Ref: ", ref, "RC2:", rc2, "Enc:", enc, "Dec:", dec)
}
*/


//----------------------------------------------------------------------------------------
//                                                                                Referral
//----------------------------------------------------------------------------------------


type Referral struct {
    referreeID         string
    referrerID         string
    creationDate       time.Time
    referralType       *BlitzMessage.InviteType
    referenceID        *string
    validFromDate      pq.NullTime
    validToDate        pq.NullTime
    redemptionDate     pq.NullTime
    referralCode       string
    claimDate          pq.NullTime
}


func (ref *Referral) InsertNew() (errRet error) {
    Log.LogFunctionName()

    var err error
    var tx *sql.Tx
    tx, err = config.DB.Begin()
    if err != nil {
        Log.LogError(err)
        return err
    }
    defer func() {
        Log.Debugf("tx: %+v error: %+v.", tx, err)
        if err != nil {
            Log.Errorf("Error was %v.", err)
            Log.Debugf("Rolling back.")
            tx.Rollback()
        } else {
            Log.Debugf("Commit.")
            err = tx.Commit()
            if err != nil {
                Log.LogError(err)
            }
        }
        errRet = err
    } ()

    _, err = tx.Exec(
        `lock table ReferralTable in access exclusive mode;`,
    )
    if err != nil { return err }

    row := tx.QueryRow(
        `select count(*) from ReferralTable;`,
    )
    var rowCount int64
    err = row.Scan(&rowCount)
    if err != nil { return err }

    ref.referralCode = referralStringFromRowCount(rowCount+1)
    Log.Debugf("Row count: %d Code: %s.", rowCount, ref.referralCode)

    var result sql.Result
    result, err = tx.Exec(
        `insert into ReferralTable (
            referreeID,
            referrerID,
            creationDate,
            referralType,
            referenceID,
            validFromDate,
            validToDate,
            redemptionDate,
            referralCode,
            claimDate
        ) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`,
            ref.referreeID,
            ref.referrerID,
            ref.creationDate,
            ref.referralType,
            ref.referenceID,
            ref.validFromDate,
            ref.validToDate,
            ref.redemptionDate,
            ref.referralCode,
            ref.claimDate,
    )
    err = pgsql.UpdateResultError(result, err)
    if err != nil {
        Log.LogError(err)
        return err
    }

    return nil
}


//----------------------------------------------------------------------------------------
//                                                                              SendInvite
//----------------------------------------------------------------------------------------


func SendInvite(inviterUserID string, invite *BlitzMessage.UserInvite) error {
    //  No: If already a friend on Blitz, send message.  Done.
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
              and contact = $2
            order by isVerified desc nulls last;`,
        invite.ContactInfo.ContactType,
        invite.ContactInfo.Contact,
    )
    var userID sql.NullString
    error = row.Scan(&userID)
    if error != nil {
        Log.LogError(error)
    }

    var claimDate pq.NullTime
    var friendProfile *BlitzMessage.UserProfile
    if userID.Valid {
        friendProfile = ProfileForUserID("", userID.String)
        claimDate.Time = time.Now()
        claimDate.Valid = true
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
        friendProfile.UserStatus = BlitzMessage.UserStatus_USInvited.Enum()
    }

    var fromDate, toDate pq.NullTime
    if invite.InviteType != nil &&
        *invite.InviteType == BlitzMessage.InviteType_ITFeedPost {
        fromDate.Time = time.Now()
        fromDate.Valid = true
        toDate.Time = time.Now().Add(time.Hour * 24 * 30)
        toDate.Valid = true
    }
    ref := Referral {
        referreeID:     inviterUserID,
        referrerID:     *friendProfile.UserID,
        creationDate:   time.Now(),
        referralType:   invite.InviteType,
        referenceID:    invite.ReferenceID,
        validFromDate:  fromDate,
        validToDate:    toDate,
        claimDate:      claimDate,
    }
    error = ref.InsertNew()
    if error != nil {
        return error
    }

    name := PrettyNameForUserID(inviterUserID)
    message := fmt.Sprintf("%s invited you to Blitz", name)
    if invite.Message != nil && len(*invite.Message) > 0 {
        message += ":\n" + *invite.Message
    }

    Log.Debugf("%v %v %v %v",
        friendProfile.UserID,
        invite.ContactInfo.ContactType,
        invite.ContactInfo.Contact,
        message,
    )

    inviteURL := fmt.Sprintf(
        "%s?action=invited&inviteeid=%s&contacttype=%d&contact=%s&message=%s&ref=%s",
        config.AppLinkURL,
        *friendProfile.UserID,
        *invite.ContactInfo.ContactType,
        url.QueryEscape(*invite.ContactInfo.Contact),
        url.QueryEscape(message),
        ref.referralCode,
    )
    shortLink, _ := LinkShortner_ShortLinkFromLink(inviteURL)

    message += "\n\nReferral Code: " + ref.referralCode
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
        message = proto.String("Some invites not sent. (Are all the invitee addresses correct?)")
    }

    return &BlitzMessage.ServerResponse {
        ResponseCode:       &code,
        ResponseMessage:    message,
        ResponseType:       &BlitzMessage.ResponseType {
            UserInvitesResponse:    invites,
        },
    }
}

