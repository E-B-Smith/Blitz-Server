

//----------------------------------------------------------------------------------------
//
//                                                           BlitzMessage : UserProfile.go
//                                                        The back-end server to BlitzHere
//
//                                                                  E.B. Smith, March 2016
//                        -©- Copyright © 2015-2016 Edward Smith, all rights reserved. -©-
//
//----------------------------------------------------------------------------------------


package BlitzMessage


import (
    "errors"
    "strings"
    "unicode/utf8"
    "violent.blue/GoKit/Util"
)


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


func (profile *UserProfile) ContactInfoOfType(ctType ContactType) *ContactInfo {
    for _, contact := range profile.ContactInfo {
        if contact.ContactType != nil && *contact.ContactType == ctType { return contact }
    }
    return nil
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

