package main

import (
	"clientstreaming/clientstreaming"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ClientStreamingServer struct {
	clientstreaming.UnimplementedClientStreamingServer
}

func (s ClientStreamingServer) GetServerResponse(stream clientstreaming.ClientStreaming_GetServerResponseServer) error {
	log.Println("Server processing gRPC client-streaming.")

	var count int

	for {
		msg, err := stream.Recv()

		// just to prevent msg not being used error
		if msg == nil {
		}

		if err == io.EOF {
			return stream.SendAndClose(&clientstreaming.Number{Value: int32(count)})
		}

		if err != nil {
			log.Fatalf("Failed to receive a message from the client: %v", err)
		}

		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()

	s := ClientStreamingServer{}

	clientstreaming.RegisterClientStreamingServer(server, &s)

	log.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 50051: %v", err)
		server.GracefulStop()
	}
}
