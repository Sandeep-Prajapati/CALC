package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	rt "runtime"

	pb "example.com/sandeepkp/calc/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedCalcServiceServer
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server")
	addressPtr := flag.String("address", "127.0.0.1:8080", "Specify server address:port")
	flag.Parse()

	listner, err := net.Listen("tcp", *addressPtr)
	if err != nil {
		panic(fmt.Sprintf("Failed to listen %s", err))
	}

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})
	// reflection.Register(s)

	fmt.Println("Server listening at", listner.Addr())

	/////
	// Serve gRPC server
	log.Println("Serving gRPC on", *addressPtr)
	go func() {
		// log.Fatalln(s.Serve(listner))
		s.Serve(listner)
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		*addressPtr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterCalcServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	// log.Fatalln(gwServer.ListenAndServe())
	gwServer.ListenAndServe()

}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a + b
	return &pb.Response{Result: result}, nil
}

func (s *server) Sub(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a - b
	return &pb.Response{Result: result}, nil
}

func (s *server) Mul(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a, b := in.GetA(), in.GetB()
	result := a * b
	return &pb.Response{Result: result}, nil
}

func (s *server) Div(ctx context.Context, in *pb.Request) (resp *pb.Response, e error) {
	a, b := in.GetA(), in.GetB()
	defer func(error) {
		if err := recover(); err != nil {
			log.Println(err)
			e = err.(rt.Error)
		}
	}(e)
	result := a / b
	return &pb.Response{Result: result}, e
}
