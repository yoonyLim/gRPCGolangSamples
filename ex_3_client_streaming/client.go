package main

import (
	"clientstreaming/clientstreaming"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeMessage(message string) clientstreaming.Message {
	return clientstreaming.Message{Message: message}
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cound not connect: %s", err)
	}
	defer conn.Close()

	client := clientstreaming.NewClientStreamingClient(conn)

	stream, err := client.GetServerResponse(context.Background())

	if err != nil {
		log.Fatalf("Error with GetServerResponses: %s", err)
	}

	messages := [5]clientstreaming.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}

	for i := 0; i < len(messages); i++ {
		log.Printf("[client to server] %s", messages[i].Message)
		if err := stream.Send(&messages[i]); err != nil {
			log.Fatalf("Failed to send a message to the server: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("[server to client] %v", response.Value)
}
