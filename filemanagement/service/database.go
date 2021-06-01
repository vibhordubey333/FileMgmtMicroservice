package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbObject *mongo.Client

func GetMongoObject() *mongo.Client {

	var client *mongo.Client
	var err error

	opts := options.Client()
	opts.ApplyURI(URI)
	opts.SetMaxPoolSize(0) // Default connection size is 100 when it is 0.

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Fatalln("Error while connecting with DB. Error: ", err)
	}

	return client
}

func init() {
	log.Println("Initializing database connection.")
	dbObject = GetMongoObject()
	log.Println("Succesfully established connection with database.")
}
