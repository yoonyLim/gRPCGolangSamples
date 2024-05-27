package main

// (1) Import console output, protobuf generated pkg, and grcp modules
import (
	"context"
	"log"

	"hello/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// (3) Create gRPC channel with insecure credentials
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cound not connect: %s", err)
	}
	defer conn.Close()

	// (4) Create a client stub as generated from "hello_grpc.proto" for the client class
	client := hello.NewMyServiceClient(conn)

	// (5) Create a message that is of the type "MyNumber" generated from "hello_grpc.proto"
	request := hello.MyNumber{Value: 4}

	// (6) Call the remote function with stub
	response, err := client.MyFunction(context.Background(), &request)

	// Print possible error while getting the response
	if err != nil {
		log.Fatalf("Error with MyFunction: %s", err)
	}

	// (7) Print the response
	log.Printf("gRCP result: %d", response.Value)
}
