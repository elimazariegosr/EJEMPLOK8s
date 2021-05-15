package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var host = "104.197.236.53"
var ctx = context.Background()
var collection = Connection()

type Data struct {
	Name         string
	Location     string
	Gender       string
	Age          int64
	Vaccine_type string
	Path         string
}



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

	return client.Database("proyecto2").Collection("vacunacions")
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

	client := redis.NewClient(&redis.Options{
		Addr:     "104.197.236.53:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	sub := client.Subscribe(ctx, "channel1")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Payload)
		
		var bdoc interface{}
		errb := bson.UnmarshalExtJSON([]byte(msg.Payload), true, &bdoc)
		if err != nil {
			fmt.Println(errb)
		}
		var err1 error
		ctx := context.Background()
		_, err1 = collection.InsertOne(ctx, bdoc)
		if err1 != nil {
			log.Fatal(err1)
		}
		ctx = context.Background()
		val, err2 := client.Do(ctx, "RPUSH", "REGISTROP2", msg.Payload).Result()
		if err2 != nil {
			fmt.Println("Error: ", err2)
		}
		fmt.Println(val)
	}
}
