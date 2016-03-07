//  Images.go  -  Image handling.
//
//  E.B.Smith  -  June, 2015


package main


import (
    "fmt"
    "errors"
    "net/http"
    "hash/crc32"
    "database/sql"
    "github.com/golang/protobuf/proto"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/pgsql"
    "BlitzMessage"
)


//----------------------------------------------------------------------------------------
//
//                                                                             UploadImage
//
//----------------------------------------------------------------------------------------


func UploadImage(writer http.ResponseWriter, userID string, imageUpload *BlitzMessage.ImageUpload) {
    Log.LogFunctionName()

    if imageUpload.ImageData == nil || len(imageUpload.ImageData) == 0 {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No image in message"))
        return
    }
    imageData := imageUpload.ImageData[0]
    if len(imageData.ImageBytes) == 0 {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("No image in message"))
        return
    }
    if len(imageData.ImageBytes) > 1000000 {
        SendError(writer, BlitzMessage.ResponseCode_RCInputInvalid, errors.New("Image > 1 megabyte"))
        return
    }
    kDefaultType := "image/jpeg"
    if imageData.ContentType == nil || *imageData.ContentType == "" {
        imageData.ContentType = &kDefaultType
    }

    crc := crc32.ChecksumIEEE(imageData.ImageBytes)

    var error error
    var result sql.Result
    result, error = config.DB.Exec(
        `insert into ImageTable (
           userID,
           imageContent,
           contentType,
           crc32,
           imageData) values ($1, $2, $3, $4, $5);`,
             userID,
             imageData.ImageContent,
             imageData.ContentType,
             crc,
             imageData.ImageBytes)
    if error != nil || pgsql.RowsUpdated(result) != 1 {
        //Log.LogError(error)
        result, error = config.DB.Exec(
            `update ImageTable set (
               imageContent,
               contentType,
               crc32,
               imageData) = ($1, $2, $3, $4)
               where userID = $5;`,
                 imageData.ImageContent,
                 imageData.ContentType,
                 crc,
                 imageData.ImageBytes,
                 userID)
    }
    if error != nil {
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }

    imageURL := fmt.Sprintf("%s%s/image?uid=%s&h=%x",
        config.ServerURL,
        config.ServicePrefix,
        userID,
        crc,
    )

    result, error = config.DB.Exec(
        "update UserTable set imageURL = array[ $1 ] where userid = $2;",
            imageURL, userID)
    Log.Debugf("ImageURL: %s Result: %+v Error: %v.", imageURL, result, error)
    if error != nil {
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, error)
        return
    }
    var rowsUpdated int64 = 0
    if result != nil { rowsUpdated, _ = result.RowsAffected() }
    if rowsUpdated != 1 {
        Log.Errorf("Didn't update image URL in UserTable.")
        SendError(writer, BlitzMessage.ResponseCode_RCServerError, errors.New("UserTable error"))
        return
    }
    Log.Debugf("Updated user '%s' image to '%s'.", userID, imageURL)

    replyImageData := BlitzMessage.ImageData {
        ImageContent: imageData.ImageContent,
        ContentType:  imageData.ContentType,
        ImageURL:     &imageURL,
    }
    replyImageUpload := BlitzMessage.ImageUpload {
        ImageData:  []*BlitzMessage.ImageData{ &replyImageData },
    }
    code := BlitzMessage.ResponseCode_RCSuccess
    response := &BlitzMessage.ServerResponse {
        ResponseCode: &code,
        Response:     &BlitzMessage.ServerResponse_ImageUploadReply { ImageUploadReply:  &replyImageUpload },
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
    uid, error = BlitzMessage.ValidateUserID(&uid)
    if error != nil {
        http.Error(writer, "Not Found", 404)
        return
    }

    Log.Debugf("Getting image for '%s'...", uid)
    rows, error := config.DB.Query(
        "select contentType, imageData from ImageTable where userID = $1;", uid)
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


