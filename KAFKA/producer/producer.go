package producer

import (
	"log"
	"os"
	"os/signal"

	"encoding/json"
	"net/http"

	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers = []string{":9093"}
	KafkaTopic   = "sarama_topic"
	enqueued     int
)

// StartProducer runs the AsyncProducer
func StartProducer() {

	producer, err := setupProducer()
	if err != nil {
		panic(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	produceMessages(producer, signals)

	log.Printf("Kafka AsyncProducer finished with %d messages produced.", enqueued)
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	return sarama.NewAsyncProducer(kafkaBrokers, config)
}

// produceMessages will send 'testing 123' to KafkaTopic each second, until receive a os signal to stop e.g. control + c
// by the user in terminal
func produceMessages(producer sarama.AsyncProducer, signals chan os.Signal) {
	message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("prueba")}
	select {
	case producer.Input() <- message:
		enqueued++
		//log.Println("New Message produced")
	case <-signals:
		producer.AsyncClose() // Trigger a shutdown of the producer.
		return
	}

	type Data struct {
		Name         string
		Location     string
		Gender       string
		Age          int64
		Vaccine_type string
		Path         string
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var d Data
		if json.NewDecoder(r.Body).Decode(&d) != nil {
			//println("CLIENTE: error 3")
		} else {
			d.Path = "Kafka"
			out, err := json.Marshal(d)
			if err != nil {
				println(err)
			}
			println(out)
			message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder(out)}
			select {
			case producer.Input() <- message:
				enqueued++
				//log.Println("New Message produced")
			case <-signals:
				producer.AsyncClose() // Trigger a shutdown of the producer.
				return
			}
		}
	})

	println("CLIENTE: servidor de cliente escuchando en 8000")
	http.ListenAndServe(":8000", nil)

}
