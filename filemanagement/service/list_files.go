package service

import (
	"FileMgmtMicroservice/filemanagement/models"
	"FileMgmtMicroservice/filemanagement/restapi/operations/file_mgmt"
	"context"
	"log"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
)

type ListInterface interface {
	ListFile() middleware.Responder
}

func ListFile(params file_mgmt.ListFilesParams) middleware.Responder {
	log.Println("Processing request to ListFiles.")

	//MongoDB connection object.
	connObj := GetMongoObject()

	collection := connObj.Database("Archieves")

	files := collection.Collection("fs.files")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := files.Find(ctx, bson.D{})

	if err != nil {
		log.Println("Error while finding all the documents. \n : ", err)
		msg := LIST_FAIL_MSG + err.Error()
		return file_mgmt.NewListFilesInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &msg})
	}

	//Storing archieve present in DB in slice.
	fileNames := make([]interface{}, 0)

	//Iterating over the cursor returned by DB to list the archieve names.
	for result.Next(ctx) {
		var output bson.M
		err := result.Decode(&output)

		if err != nil {
			log.Println("Error while decoding the documents. \n : ", err)
			msg := LIST_FAIL_MSG + err.Error()
			return file_mgmt.NewListFilesInternalServerError().WithPayload(&models.Error{Code: INTERNAL_SERVER_ERROR, Message: &msg})
		}

		//Storing in slice.
		fileNames = append(fileNames, output["filename"])
	}

	log.Println("Request to list files processed successfully.")
	//Success return .
	return file_mgmt.NewListFilesOK().WithPayload(&file_mgmt.ListFilesOKBody{Message: "Archieve Present In DB List:", FileName: fileNames})
}
