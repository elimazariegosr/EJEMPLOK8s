package main

import (
	"context"
	"fmt"
	//"log"
	//"time"

	"github.com/go-redis/redis/v8"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

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
		val, err := client.Do(ctx, "RPUSH", "SOPESP2", msg.Payload).Result()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(val)
	}
}
