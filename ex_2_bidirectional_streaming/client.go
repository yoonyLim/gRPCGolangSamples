package main

import (
	"bidirectional/bidirectional"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func makeMessage(message string) bidirectional.Message {
	return bidirectional.Message{Message: message}
}

func sendMessages(stream bidirectional.Bidirectional_GetServerResponseClient) {
	messages := [5]bidirectional.Message{
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
		// time.Sleep(time.Millisecond * 100) // if needed to receive and log the response simultaneously
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln(err)
	}
}

func receiveMessages(stream bidirectional.Bidirectional_GetServerResponseClient, done chan bool) {
	for {
		response, err := stream.Recv()

		if err == io.EOF {
			close(done) // close
			return
		}

		if err != nil {
			log.Fatalf("Failed to receive a message from the server: %v", err)
		}

		log.Printf("[server to client] %s", response.Message)
	}
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cound not connect: %s", err)
	}
	defer conn.Close()

	client := bidirectional.NewBidirectionalClient(conn)

	stream, err := client.GetServerResponse(context.Background())

	if err != nil {
		log.Fatalf("Error with GetServerResponses: %s", err)
	}

	ctx := stream.Context()

	go sendMessages(stream)

	done := make(chan bool) // to keep the channel alive to listen to the server

	go receiveMessages(stream, done)

	<-ctx.Done() // finish stream
}
