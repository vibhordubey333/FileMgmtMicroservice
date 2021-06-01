package service

import (
	"FileMgmtMicroservice/filemanagement/models"
	"FileMgmtMicroservice/filemanagement/restapi/operations/file_mgmt"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type UploadInterface interface {
	UploadFile(params file_mgmt.UploadFileParams) middleware.Responder
}

func UploadFile(params file_mgmt.UploadFileParams) middleware.Responder {
	log.Println("Processing request to upload file.")

	_, handler, err := params.HTTPRequest.FormFile("inputFile")
	if err != nil {
		errMsg := "Request to upload file failed."
		log.Fatalln(errMsg)
		errMsg += err.Error()
		return file_mgmt.NewUploadFileInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg})
	}

	TMP_UPLOAD_DIR := os.Getenv("TMP_UPLOAD_DIR")

	if TMP_UPLOAD_DIR == "" {
		TMP_UPLOAD_DIR = DEFAULT_UPLOAD_DIR
	}

	//MongoDB connection object.
	connObj := GetMongoObject()
	bs, err := ioutil.ReadAll(params.InputFile)
	if err != nil {
		errMsg := "Request to upload file failed.Error while Reading File ."
		log.Fatalln(errMsg)
		errMsg += err.Error()
		return file_mgmt.NewUploadFileInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg})
	}
	buck, err := gridfs.NewBucket(
		connObj.Database("Archieves"),
	)

	if err != nil {
		errMsg := "Request to upload file failed."
		log.Fatalln(errMsg)
		errMsg += err.Error()
		return file_mgmt.NewUploadFileInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg})
	}

	uploadStream, err := buck.OpenUploadStream(
		handler.Filename,
	)
	defer uploadStream.Close()
	log.Println("Opening stream error : ",err)
	if err != nil {
		errMsg := "Request to upload file failed.Error while opening stream."
		log.Fatalln(errMsg)
		errMsg += err.Error()
		return file_mgmt.NewUploadFileInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg})
	}
	fileSize, err := uploadStream.Write(bs)
	if err != nil {
		errMsg := "Request to upload file failed.Error while writing to DB."
		log.Fatalln(errMsg)
		errMsg += err.Error()
		return file_mgmt.NewUploadFileInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg})
	}
	fmt.Println("Write is successful , filesize is : ", fileSize)
	return file_mgmt.NewUploadFileOK().WithPayload(&file_mgmt.UploadFileOKBody{Code: "200", Message: "File " + handler.Filename + " uploaded successfully."})
}
