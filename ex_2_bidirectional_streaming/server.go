package main

import (
	"bidirectional/bidirectional"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BidirectionalServer struct {
	bidirectional.UnimplementedBidirectionalServer
}

func (s BidirectionalServer) GetServerResponse(stream bidirectional.Bidirectional_GetServerResponseServer) error {
	log.Println("Server processing gRPC bidirectional streaming.")

	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Failed to receive a message from the client: %v", err)
			continue
		}

		response := bidirectional.Message{Message: request.Message}

		if err := stream.Send(&response); err != nil {
			log.Printf("Failed to send a message to the client: %v", err)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()

	s := BidirectionalServer{}

	bidirectional.RegisterBidirectionalServer(server, &s)

	log.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 50051: %v", err)
		server.GracefulStop()
	}
}
