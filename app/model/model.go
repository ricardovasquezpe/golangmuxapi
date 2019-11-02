package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseUser struct {
	Users  []User `json:"users"`
	Status int    `json:"status"`
}

func GetUsers(client *mongo.Client, filter bson.M) []User {
	var users []User
	collection := client.Database("contactlist").Collection("users")
	usr, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for usr.Next(context.TODO()) {
		var user User
		err = usr.Decode(&user)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		users = append(users, user)
	}
	return users
}

func InsertUser(client *mongo.Client, user User) {
	collection := client.Database("contactlist").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	//return insertResult.InsertedID
}
