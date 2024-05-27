package main

// (1) Import console output, network connection, protobuf generated pkg, and grcp modules
import (
	"context"
	"log"
	"net"

	"hello/hello"

	"google.golang.org/grpc"
)

// (2) Create a server class for protoc generated base class
type MyServiceServer struct {
	hello.UnimplementedMyServiceServer
}

// (3) Define the method that was generated from "hello_grpc.proto" for the server class
// Accepts parameter of "request" that is of the message type "MyNumber" generated from "hello_grpc.proto"
// Returns the message - "MyNumber"
func (s MyServiceServer) MyFunction(ctx context.Context, request *hello.MyNumber) (*hello.MyNumber, error) {
	return &hello.MyNumber{Value: request.Value * request.Value}, nil
}

func main() {
	// (4) Open up the port using TCP
	lis, err := net.Listen("tcp", ":50051") // short variable declaration for network connection and error

	// Print possible error while opening the port
	if err != nil { // log if error is initilized, thus error has occured
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// (5) Create grpc server
	server := grpc.NewServer() // grcp server uses goroutines by default -> if needed, instead of limiting goroutines, one should limit num of incoming reqs

	s := MyServiceServer{}
	// (6) Add the handler to the registered server
	hello.RegisterMyServiceServer(server, &s)

	// (7) Start the server until termination
	log.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 50051: %v", err)
		server.GracefulStop()
	}
}
