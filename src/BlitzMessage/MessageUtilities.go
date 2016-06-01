

//----------------------------------------------------------------------------------------
//
//                                                      BlitzMessage : MessageUtilities.go
//                                                        The back-end server to BlitzHere
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package BlitzMessage


import (
    "fmt"
    "math"
    "time"
    "errors"
    "strings"
    "unicode/utf8"
    "github.com/lib/pq"
    "violent.blue/GoKit/Util"
)



func (profile *UserProfile) ContactInfoOfType(ctType ContactType) *ContactInfo {
    for _, contact := range profile.ContactInfo {
        if contact.ContactType != nil && *contact.ContactType == ctType { return contact }
    }
    return nil
}


func TimestampPtr(inputValue interface{}) *Timestamp {
    var t *time.Time = nil

    switch val := inputValue.(type) {
    case pq.NullTime:
        if val.Valid && val.Time.Unix() != 0 {
            t = &val.Time
        }

    case *pq.NullTime:
        if val.Valid && val.Time.Unix() != 0 {
            t = &val.Time
        }

    case time.Time:
        t = &val

    case *time.Time:
        t = val

    default:
        panic(fmt.Errorf("Unhandled type: %+v.", val))
    }

    if t == nil { return nil }
    var fepoch float64 = float64(t.UnixNano()) / float64(1000000000)
    ts := Timestamp{ Epoch: &fepoch }
    return &ts
}


func (timestamp *Timestamp) Time() time.Time {
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000.0)
    return time.Unix(sec, nsec)
}


func (timestamp *Timestamp) TimePtr() *time.Time {
    t := timestamp.Time()
    return &t
}

/*
func (timestamp *Timestamp) NullTime() pq.NullTime {
    if timestamp == nil || timestamp.Epoch == nil || *timestamp.Epoch < -22135596800 {
        return pq.NullTime{};
    }
    i, f := math.Modf(*timestamp.Epoch)
    var  sec int64 = int64(math.Floor(i))
    var nsec int64 = int64(f * 1000000)
    return pq.NullTime{Time:time.Unix(sec, nsec), Valid: true }
}
*/

func (timestamp *Timestamp) NullTime() pq.NullTime {
    var nt pq.NullTime
    if timestamp != nil && timestamp.Epoch != nil && *timestamp.Epoch > -22135596800 {
        nt.Valid = true
        nt.Time = timestamp.Time()
    }
    return nt
}


func (timestamp *Timestamp) NullTimePtr() *pq.NullTime {
    nt := timestamp.NullTime()
    return &nt
}


func NullTimeFromTimespanStart(timespan *Timespan) pq.NullTime {
    if timespan == nil || timespan.StartTimestamp == nil { return pq.NullTime{} }
    return timespan.StartTimestamp.NullTime()
}


func NullTimeFromTimespanStop(timespan *Timespan) pq.NullTime {
    if timespan == nil || timespan.StopTimestamp == nil { return pq.NullTime{} }
    return timespan.StopTimestamp.NullTime()
}


func TimespanFromNullTimes(startDate, stopDate pq.NullTime) *Timespan {
    if !startDate.Valid && !stopDate.Valid { return nil }

    var t Timespan
    if startDate.Valid {
        t.StartTimestamp = TimestampPtr(startDate)
    }
    if stopDate.Valid {
        t.StopTimestamp = TimestampPtr(stopDate)
    }

    return &t
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

