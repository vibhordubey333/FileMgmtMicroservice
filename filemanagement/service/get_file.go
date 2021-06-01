package service

import (
	"FileMgmtMicroservice/filemanagement/models"
	"FileMgmtMicroservice/filemanagement/restapi/operations/file_mgmt"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	//"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type GetInterface interface {
	GetFile(params file_mgmt.GetFileParams) middleware.Responder
}

func GetFile(params file_mgmt.GetFileParams) middleware.Responder {
	log.Println("Processing request to GetFile.")

	if params.FileName != nil {
		
		fileName := *params.FileName

		//MongoDB connection object.
		connObj := GetMongoObject()

		collection := connObj.Database("Archieves")

		files := collection.Collection("fs.files")

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var results bson.M

		err := files.FindOne(ctx, bson.M{}).Decode(&results)

		if err != nil {

			log.Fatalln(GET_FILE_ERROR, "Error while finding document. ", err)
			errMsg := GET_FILE_ERROR + "Error while finding document. " + err.Error()
			return middleware.ResponderFunc(func(w http.ResponseWriter, r runtime.Producer) {

				_ = r // To avoid code smell issue from being reported by SONAR using variable r.

				w.Header().Set("Content-Type", "application/zip")

				err := models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg}

				b, _ := json.Marshal(err)

				_, _ = w.Write(b)

			})
		}

		bucket, _ := gridfs.NewBucket(
			collection,
		)

		var buf bytes.Buffer

		_, err = bucket.DownloadToStreamByName(fileName, &buf)

		if err != nil {

			log.Fatalln(GET_FILE_ERROR, DOWNLOADING_ERROR, err)

			errMsg := GET_FILE_ERROR + DOWNLOADING_ERROR + err.Error()

			return middleware.ResponderFunc(func(w http.ResponseWriter, r runtime.Producer) {

				_ = r // To avoid code smell issue from being reported by SONAR using variable r.

				w.Header().Set("Content-Type", "application/zip")

				err := models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg}

				b, _ := json.Marshal(err)

				_, _ = w.Write(b)

			})

		}

		ioutil.WriteFile(fileName, buf.Bytes(), 0600)

		var openFileFunc = os.OpenFile

		fileCreate, err := openFileFunc(
			DOWNLOAD_PATH+fileName,
			os.O_RDWR|os.O_TRUNC|os.O_CREATE,
			0600)

		fileCreate.Write(buf.Bytes())

		filePath := DOWNLOAD_PATH + fileName

		var openFileHandler = os.Open

		fileCreate, err = openFileHandler(filePath)

		if err != nil {

			log.Fatalln(GET_FILE_ERROR, "Error while creating directory. ", err)

			errMsg := GET_FILE_ERROR + "Error while creating directory. " + err.Error()

			return middleware.ResponderFunc(func(w http.ResponseWriter, r runtime.Producer) {

				_ = r // To avoid code smell issue from being reported by SONAR using variable r.

				w.Header().Set("Content-Type", "application/json")

				err := models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg}

				b, _ := json.Marshal(err)

				_, _ = w.Write(b)

			})

		}

		return middleware.ResponderFunc(func(w http.ResponseWriter, r runtime.Producer) {
			_ = r //To avoid code smell from being reported by SONAR , in future if we run .
			fn := filepath.Base(filePath)
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fn))
			_, _ = io.Copy(w, fileCreate)
			log.Println("Request to GetFile processed successfully.")
		})
	} else {

		return middleware.ResponderFunc(func(w http.ResponseWriter, r runtime.Producer) {

			errMsg := "FileName is empty."

			_ = r // To avoid code smell issue from being reported by SONAR using variable r.

			w.Header().Set("Content-Type", "application/json")

			err := models.Error{Code: INTERNAL_SERVER_ERROR, Message: &errMsg}

			b, _ := json.Marshal(err)

			_, _ = w.Write(b)

		})

	}
}
