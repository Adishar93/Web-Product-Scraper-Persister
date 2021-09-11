package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func fetchProducts() []productDocument {
	uri := "mongodb://mongoDB:mongoDB@db:27017"
	var Products []productDocument

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	collection := client.Database("amascr").Collection("products")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	if err = cursor.All(context.TODO(), &Products); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Products)
	return Products

}
