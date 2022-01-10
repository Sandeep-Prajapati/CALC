package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "example.com/sandeepkp/calc/proto"
	"google.golang.org/grpc"
)

func main() {

	addressPtr := flag.String("address", "127.0.0.1:8080", "Specify server address:port")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect %s", err))
	}
	defer conn.Close()
	fmt.Println("Starting client on", *addressPtr)

	c := pb.NewCalcServiceClient(conn)

	// First call
	request := pb.Request{A: 100, B: 20}
	response, err := c.Add(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Add: %s", err)
	}
	log.Printf("Add %d %d = %d", request.A, request.B, response.Result)

	// Second call
	request = pb.Request{A: 20, B: 10}
	response, err = c.Sub(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Add: %s", err)
	}
	log.Printf("Sub %d %d = %d", request.A, request.B, response.Result)

	// Third call
	request = pb.Request{A: 10, B: 20}
	response, err = c.Mul(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Add: %s", err)
	}
	log.Printf("Mul %d %d = %d", request.A, request.B, response.Result)

	// Fourth call
	request = pb.Request{A: 20, B: 10}
	response, err = c.Div(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Add: %s", err)
	}
	log.Printf("Div %d %d = %d", request.A, request.B, response.Result)

}
