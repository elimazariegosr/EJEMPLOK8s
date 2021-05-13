package main

import (
	"context"

	"log"

	"encoding/json"
	"net/http"

	greetpb "grpcclient/greet.pb"

	"google.golang.org/grpc"
)

func enviarMensaje(name string, location string, gender string, age int64, vaccine_type string) {
	//server := "grpcserver:3000"
	server := "localhost:3000"

	println("CLIENTE: Enviando peticion a ", server)

	//para no utilizar ssl
	cc, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("CLIENTE: error 1 ", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//println("CLIENTE: Conectado a ", server)

	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Name:        name,
			Location:    location,
			Gender:      gender,
			Age:         age,
			Vaccinetype: vaccine_type,
			Path:        "GRPC",
		},
	}

	res, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("CLIENTE:  error 2 ", err)
	}

	println("CLIENTE: resultado ", res.Result)
}

func main() {
	type Data struct {
		Name         string
		Location     string
		Gender       string
		Age          int64
		Vaccine_type string
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var d Data
		if json.NewDecoder(r.Body).Decode(&d) != nil {
			//println("CLIENTE: error 3")
		} else {
			enviarMensaje(d.Name, d.Location, d.Gender, d.Age, d.Vaccine_type)
		}
	})

	println("CLIENTE: servidor de cliente escuchando en 8000")
	http.ListenAndServe(":8000", nil)
}

//Creo que no es necesario con el anterior de greet.proto
//go get github.com/golang/protobuf/proto
//go get google.golang.org/grpc
//go get google.golang.org/protobuf/reflect/protoreflect@v1.25.0
