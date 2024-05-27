package main

import (
	"log"
	"net"
	"serverstreaming/serverstreaming"

	"google.golang.org/grpc"
)

func makeMessage(message string) serverstreaming.Message {
	return serverstreaming.Message{Message: message}
}

type ServerStreamingServer struct {
	serverstreaming.UnimplementedServerStreamingServer
}

func (s ServerStreamingServer) GetServerResponse(request *serverstreaming.Number, stream serverstreaming.ServerStreaming_GetServerResponseServer) error {
	messages := [5]serverstreaming.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}

	log.Printf("Server processing gRPC server-streaming %d", request.Value)

	for i := 0; i < len(messages); i++ {
		if err := stream.Send(&messages[i]); err != nil {
			log.Fatalf("Failed to send a message to the client: %v", err)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()

	s := ServerStreamingServer{}

	serverstreaming.RegisterServerStreamingServer(server, &s)

	log.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 50051: %v", err)
		server.GracefulStop()
	}
}
