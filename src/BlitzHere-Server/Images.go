//  Images.go  -  Image handling.
//
//  E.B.Smith  -  June, 2015


package main


import (
    "fmt"
    "time"
    "errors"
    "strconv"
    "net/http"
    "hash/crc32"
    "database/sql"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                             UploadImage
//
//----------------------------------------------------------------------------------------


func ImageURLForImageData(userID string, imageData *BlitzMessage.ImageData) string {
    if imageData.Crc32 == nil { return "" }
    return fmt.Sprintf("%s%s/image?uid=%s&h=%x",
        config.ServerURL,
        config.ServicePrefix,
        userID,
        *imageData.Crc32,
    )
}


func UploadImage(session *Session, imageUpload *BlitzMessage.ImageUpload,
        ) *BlitzMessage.ServerResponse {
    Log.LogFunctionName()

    if imageUpload.ImageData == nil || len(imageUpload.ImageData) == 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No image in message"))
    }
    imageData := imageUpload.ImageData[0]
    if len(imageData.ImageBytes) == 0 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No image in message"))
    }
    if len(imageData.ImageBytes) > 1000000 {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Image > 1 megabyte"))
    }
    kDefaultType := "image/jpeg"
    if imageData.ContentType == nil || *imageData.ContentType == "" {
        imageData.ContentType = &kDefaultType
    }

    var crc int64
    crc = int64(crc32.ChecksumIEEE(imageData.ImageBytes))
    imageData.Crc32 = &crc
    if imageData.DateAdded == nil {
        imageData.DateAdded = BlitzMessage.TimestampFromTime(time.Now())
    }

    var error error
    var result sql.Result
    result, error = config.DB.Exec(
        `insert into ImageTable (
           userID,
           imageContent,
           contentType,
           crc32,
           imageData,
           dateAdded) values ($1, $2, $3, $4, $5, $6);`,
             session.UserID,
             imageData.ImageContent,
             imageData.ContentType,
             imageData.Crc32,
             imageData.ImageBytes,
             BlitzMessage.TimeFromTimestamp(imageData.DateAdded))
    if error != nil || pgsql.RowsUpdated(result) != 1 {
        //Log.LogError(error)
        result, error = config.DB.Exec(
            `update ImageTable set (
               imageContent,
               contentType,
               imageData,
               dateAdded) = ($1, $2, $3, $4)
               where userID = $5 and crc32 = $6;`,
                 imageData.ImageContent,
                 imageData.ContentType,
                 imageData.ImageBytes,
                 BlitzMessage.TimeFromTimestamp(imageData.DateAdded),
                 session.UserID,
                 imageData.Crc32)
    }
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    imageData.ImageURL = StringPtrFromString(ImageURLForImageData(session.UserID, imageData))
    Log.Debugf("ImageURL: %s Updated: %d Error: %v.", *imageData.ImageURL, pgsql.RowsUpdated(result), error)
    if error != nil {
        return ServerResponseForError(BlitzMessage.ResponseCode_RCServerError, error)
    }

    imageData.ImageBytes = nil
    replyImageUpload := BlitzMessage.ImageUpload {
        ImageData:  []*BlitzMessage.ImageData{ imageData },
    }
    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
        ResponseType: &BlitzMessage.ResponseType { ImageUploadReply:  &replyImageUpload },
    }
    return response
}



//----------------------------------------------------------------------------------------
//
//                                                                                GetImage
//
//----------------------------------------------------------------------------------------


func GetImage(writer http.ResponseWriter, httpRequest *http.Request) {
    Log.LogFunctionName()

    if httpRequest.URL == nil {
        http.Error(writer, "Not Found", 404)
        return
    }
    var error error
    uid := httpRequest.URL.Query().Get("uid")
    crc := httpRequest.URL.Query().Get("h")

    uid, error = BlitzMessage.ValidateUserID(&uid)
    if error != nil {
        http.Error(writer, "Not Found", 404)
        return
    }

    var crc32 int64
    crc32, error = strconv.ParseInt(crc, 16, 64)
    if error != nil {
        http.Error(writer, "Not Found", 404)
        return
    }

    Log.Debugf("Getting image for '%s' '%d'...", uid, crc32)
    rows, error := config.DB.Query(
        "select contentType, imageData from ImageTable where userID = $1 and crc32 = $2;",
            uid, crc32)
    defer pgsql.CloseRows(rows)
    if error != nil {
        Log.LogError(error)
        http.Error(writer, "Not Found", 404)
        return
    }
    if rows.Next() {
        var (contentType string; imageData []byte)
        error = rows.Scan(&contentType, &imageData)
        if error == nil {
            writer.Header().Add("Content-Type", contentType)
            bytes, error := writer.Write(imageData)
            Log.Debugf("Wrote %d of %d bytes. Error: %v.", bytes, len(imageData), error)
            return
        } else {
            Log.LogError(error)
            http.Error(writer, "Internal Server Error", 500)
            return
        }
    }
    http.Error(writer, "Not Found", 404)
}


