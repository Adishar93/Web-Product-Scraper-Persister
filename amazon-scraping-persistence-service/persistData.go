package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func persistData(Request *productDocument) {
	uri := "mongodb://mongoDB:mongoDB@db:27017"
	var doc productDocument

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	collection := client.Database("amascr").Collection("products")

	//Check whether url already exists
	err = collection.FindOne(context.TODO(), bson.M{"url": Request.URL}).Decode(&doc)
	if err != nil {
		fmt.Println("Data not found on database")
		collection.InsertOne(context.TODO(), Request)
	} else {
		fmt.Println("Data Found on database")
		collection.ReplaceOne(context.TODO(), bson.M{"url": Request.URL}, Request)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}

}
