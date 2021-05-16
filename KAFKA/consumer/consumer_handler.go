package consumer

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"ftm"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConsumerGroupHandler represents the sarama consumer group
type ConsumerGroupHandler struct{}

// Setup is run before consumer start consuming, is normally used to setup things such as database connections
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages(), here is supposed to be what you want to
// do with the message. In this example the message will be logged with the topic name, partition and message value.
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		m := string(msg.Value)

		if m != "prueba" {
			var d Data
			err := json.Unmarshal([]byte(m), &d)
			if err != nil {
				println("Error occured during unmarshaling. Error: " + err.Error())
			}

			creacion(d)
		}

		//message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder(out)}

	}
	return nil
}

var host = "104.197.236.53"

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

	return client.Database("proyecto2").Collection("vacunacions")
}

func creacion(data Data) {
	client := redis.NewClient(&redis.Options{
		Addr:     "104.197.236.53:6379",
		Password: "",
		DB:       0,
	})
	s := ""
	concatenated := fmt.Sprintf("%d%s", req.GetGreeting().GetAge(), s)
	
	msg := `{ "name": "` + d.Name + `",
	"location": "` + d.Location + `",
	"gender": "` + d.Gender + `",
	"age": `+ concatenated +`,
	"vaccine_type": "` + d.Vaccine_type + `",
	"path": "Kafka"  
	 }`
	defer client.Close()
	
	var ctx = context.Background()
	val, err := client.Do(ctx, "RPUSH", "REGISTROP2", msg).Result()
	if err != nil {
		fmt.Println("Error: ", err)
	}	
	fmt.Println(val)

	var err error

	ctx := context.Background()

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
}
