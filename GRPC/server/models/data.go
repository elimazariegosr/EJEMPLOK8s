//instalar libreria mongo
//go get go.mongodb.org/mongo-driver/mongo

package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name         string
	Location     string
	Gender       string
	Age          int64
	Vaccine_type string
	Path         string
}

func Connection() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/proyecto2"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("proyecto2").Collection("vacunacion")
}

func main() {

	data := Data{
		Name:         "miguel",
		Location:     "guate",
		Gender:       "m",
		Age:          15,
		Vaccine_type: "sv",
		Path:         "grpc",
	}

	collection := Connection()
	if collection != nil {
		print("jalo")
	}

	var err error

	ctx := context.Background()

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
}
