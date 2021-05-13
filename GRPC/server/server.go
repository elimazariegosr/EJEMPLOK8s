package main

import (
	//"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	//"net/http"
	//"encoding/json"

	greetpb "grpcserver/greet.pb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

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
	var err error

	ctx := context.Background()

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("SERVER: recibiendo data ", req.GetGreeting().GetPath())
	//host := "http://34.121.110.42/"
	//host := "http://localhost:5000/"
	//host := "http://0.0.0.0:5000/"

	result := "todo bien SERVER"

	//println(req.GetGreeting().GetName())

	data := Data{
		Name:         req.GetGreeting().GetName(),
		Location:     req.GetGreeting().GetLocation(),
		Gender:       req.GetGreeting().GetGender(),
		Age:          req.GetGreeting().GetAge(),
		Vaccine_type: req.GetGreeting().GetVaccinetype(),
		Path:         req.GetGreeting().GetPath(),
	}

	creacion(data)
	/*data, _ := json.Marshal(req.GetGreeting())
	http.Post(host, "application/json", bytes.NewBuffer(data))*/

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	host := ":3000"
	//host := "0.0.0.0:3000"

	fmt.Println("SERVER: server iniciado en ", host)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("SERVER: error 1 ", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err == nil {
		log.Fatalf("SERVER: error ", err)
	}
}

//instalar libreria mongo
//go get -u go.mongodb.org/mongo-driver/mongo
