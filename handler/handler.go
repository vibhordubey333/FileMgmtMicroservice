package handler

import (
	"FileMgmtMicroservice/restapi/operations/file_mgmt"
	"FileMgmtMicroservice/service"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

type HandlerFunctions interface {
	GetFile(params file_mgmt.GetFileParams) middleware.Responder
	UploadFile(params file_mgmt.UploadFileParams) middleware.Responder
	ListFile() middleware.Responder
}

func init() {
	log.Println("File Management Module starting.")
}

// GetFile (params file_mgmt.GetFileParams) middleware.Responder to return the requested archieve file.
func GetFile(params file_mgmt.GetFileParams) middleware.Responder {
	log.Println("Request recieved to get the file.")
	return service.GetFile(params)
}

func ListFile(params file_mgmt.ListFilesParams) middleware.Responder {
	log.Println("Processing request to list files.")
	return service.ListFile(params)
}

func UploadFile(params file_mgmt.UploadFileParams) middleware.Responder {
	log.Println("Request recieved to upload files.")
	return service.UploadFile(params)
}
