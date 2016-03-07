//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package BlitzMessage

import (
    "math"
    "time"
    "errors"
    "strings"
    "unicode/utf8"
    "github.com/lib/pq"
    "violent.blue/GoKit/Util"
)


func ResponseCodePtr(r ResponseCode) *ResponseCode {
    return &r
}


func (profile *UserProfile) ContactInfoOfType(ctType ContactType) *ContactInfo {
    for _, contact := range profile.ContactInfo {
        if contact.ContactType != nil && *contact.ContactType == ctType { return contact }
    }
    return nil
}

func NullTimeFromTimestamp(timestamp *Timestamp) pq.NullTime {
    if timestamp == nil || timestamp.Epoch == nil || *timestamp.Epoch < -22135596800 {
        return pq.NullTime{};
    }
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    return pq.NullTime{Time:time.Unix(sec, nsec), Valid: true }
}


func TimestampPtrFromNullTime(time pq.NullTime) *Timestamp {
    if ! time.Valid  || time.Time.Unix() == 0 { return nil }
    fepoch := float64(time.Time.Unix()) + float64(time.Time.Nanosecond()) / float64(1000000)
    ts := Timestamp{ Epoch: &fepoch }
    return &ts
}


func TimeFromTimestamp(timestamp *Timestamp) time.Time {
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    return time.Unix(sec, nsec)
}


func (timestamp *Timestamp) Time() time.Time {
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    return time.Unix(sec, nsec)
}


func TimestampFromTime(t time.Time) *Timestamp {
    if t.Unix() == 0 { return nil }
    fepoch := float64(t.Unix()) + float64(t.Nanosecond()) / float64(1000000)
    ts := Timestamp{ Epoch: &fepoch }
    return &ts
}


func ValidateUserID(userID *string) (string, error) {
    //  Format and Validate the userid --

    if userID == nil {
        return "", errors.New("Nil userID")
    } else {
        validUserID := strings.TrimSpace(*userID)
        if utf8.RuneCountInString(validUserID) > 31 {
            return validUserID, nil
        }
    }
    return "", errors.New("UserID too short.")
}


func (profile *UserProfile) AddContactInfo(newInfo *ContactInfo) {
    if newInfo == nil { return; }
    newInfo.Contact = Util.CleanStringPtr(newInfo.Contact)
    for index, info := range(profile.ContactInfo) {
        if  newInfo.ContactType == info.ContactType &&
            newInfo.Contact == info.Contact {
            profile.ContactInfo[index] = newInfo
            return
        }
    }
    if profile.ContactInfo == nil {
        profile.ContactInfo = []*ContactInfo{newInfo}
    } else {
        profile.ContactInfo = append(profile.ContactInfo, newInfo)
    }
}

