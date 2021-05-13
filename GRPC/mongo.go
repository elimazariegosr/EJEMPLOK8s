//instalar libreria mongo
//go get -u go.mongodb.org/mongo-driver/mongo

//crear una variable con el struct Data
//enviar Data de parametro en la funcion creacion

package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var host = "localhost"

type Data struct {
	Name         string
	Location     string
	Gender       string
	Age          int64
	Vaccine_type string
	Path         string
}

var collection = Connection()

func Connection() *mongo.Collection {
	dir := "mongodb://" + host + ":27017/proyecto2"

	client, err := mongo.NewClient(options.Client().ApplyURI(dir))
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

func creacion(data Data) {
	var err error

	ctx := context.Background()

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data := Data{
		Name:         "Miguel",
		Location:     "Mixco",
		Gender:       "Masculino",
		Age:          15,
		Vaccine_type: "Sputnic V",
		Path:         "GRPC",
	}

	creacion(data)
}
