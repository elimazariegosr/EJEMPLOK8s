package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type Data struct {
	Name         string
	Location     string
	Gender       string
	Age          json.Number
	Vaccine_type string
}


func send(w http.ResponseWriter, r *http.Request) {
	client := redis.NewClient(&redis.Options{
		Addr:     "104.197.236.53:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()
	w.Header().Set("Content-Type", "application/json")

	var d Data

	_ = json.NewDecoder(r.Body).Decode(&d)
	json.NewEncoder(w).Encode(&d)
	err := client.Publish(r.Context(), "channel1",
		`{ "name": "`+d.Name+`",
	"location": "`+d.Location+`",
	"gender": "`+d.Gender+`",
	"age": "`+string(d.Age)+`",
	"vaccine_type": "`+d.Vaccine_type+`",
	"path: REDIS"  
	 }`).Err()
	if err != nil {
		panic(err)
	}
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", send).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
