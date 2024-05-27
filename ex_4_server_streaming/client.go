package main

import (
	"context"
	"io"
	"log"

	"serverstreaming/serverstreaming"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cound not connect: %s", err)
	}
	defer conn.Close()

	client := serverstreaming.NewServerStreamingClient(conn)

	request := serverstreaming.Number{Value: 5}

	stream, err := client.GetServerResponse(context.Background(), &request)

	if err != nil {
		log.Fatalf("Error with MyFunction: %s", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	go func() {
		for {
			response, err := stream.Recv()

			if response != nil {
				log.Printf("[server to client] %s", response)
			}

			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalf("Failed to receive a message from the server: %v", err)
			}
		}
	}()

	<-ctx.Done()
}
