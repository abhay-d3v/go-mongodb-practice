package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoDbUrl = "mongodb+srv://abhay:OCitKVzD2qa4tQsb@test.ubl4da9.mongodb.net/?retryWrites=true&w=majority&appName=Test"

var Client *mongo.Client
var UsersCollection *mongo.Collection

func init() {
	var err error

	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		fmt.Println("[Error]: Couldn't connect to URL")
		panic(err)
	}

	err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()

	if err != nil {
		fmt.Println("[Error]: Couldn't ping to db")
		panic(err)
	}

	UsersCollection = Client.Database("Test").Collection("users")
}
