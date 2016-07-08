//  Social.go  -  Update/query social data.
//
//  E.B.Smith  -  May, 2015.


package main


import (
    "strings"
    "database/sql"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)



//----------------------------------------------------------------------------------------
//                                                          UpdateSocialIdentityWithUserID
//----------------------------------------------------------------------------------------


func UpdateSocialIdentityForUserID(userID string, social *BlitzMessage.SocialIdentity) error {
    Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    result, error := config.DB.Exec("update SocialTable set" +
        " (userID, service, socialID, userName, displayName, URI, authToken, authExpire) = " +
        " ($1, $2, $3, $4, $5, $6, $7, $8)" +
        " where userID = $9 and service = $2 and socialID = $3;",
        &userID,
        strings.ToLower(*social.SocialService),
        social.SocialID,
        social.UserName,
        social.DisplayName,
        social.UserURI,
        social.AuthToken,
        social.AuthExpire,
        &userID);

    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }

    if error == nil && rowsUpdated > 0 {
        Log.Debugf("Updated social %s %v.", userID, social.SocialService)
    } else {
        Log.Debugf("Inserting social ID %s %s %s: %v", userID, *social.SocialService, *social.SocialID, error)

        //  Insert instead --

        _, error = config.DB.Exec("insert into SocialTable " +
        " (userID, service, socialID, userName, displayName, URI, authToken, authExpire) values " +
        " ($1, $2, $3, $4, $5, $6, $7, $8);",
            &userID,
            strings.ToLower(*social.SocialService),
            social.SocialID,
            social.UserName,
            social.DisplayName,
            social.UserURI,
            social.AuthToken,
            social.AuthExpire)

        if error != nil {
            Log.Errorf("Error inserting social ID %s %s %s: %v", userID, *social.SocialService, *social.SocialID, error)
        }
    }

    return error;
}



//----------------------------------------------------------------------------------------
//                                                               SocialIdentitiesForUserID
//----------------------------------------------------------------------------------------


func SocialIdentitiesWithUserID(userID string)  []*BlitzMessage.SocialIdentity {
    Log.LogFunctionName()
    defer func() {
        if error := recover(); error != nil { Log.LogStackWithError(error) }
    } ()

    social := make([]*BlitzMessage.SocialIdentity, 0, 10)
    rows, error := config.DB.Query(
        `select service, socialID, userName, displayName, URI
            from SocialTable
           where userID = $1;`,
           userID,
    )
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.Debugf("Error querying social: %v.", error)
        return social
    }

    for rows.Next() {
        var (
            service     sql.NullString
            socialID    sql.NullString
            userName    sql.NullString
            displayName sql.NullString
            URI         sql.NullString
        )
        error = rows.Scan(&service, &socialID, &userName, &displayName, &URI)
        if error == nil {
            sid := BlitzMessage.SocialIdentity {
                SocialService:  proto.String(service.String),
                SocialID:       proto.String(socialID.String),
                UserName:       proto.String(userName.String),
                DisplayName:    proto.String(displayName.String),
                UserURI:        proto.String(URI.String),
            }
            social = append(social, &sid)
        }
    }
    return social
}

