package main

import (
	"context"
	"grpc/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{
		Body: "Hello grpc, are you heare?",
	})
	if err != nil {
		log.Fatal("error when calling SayHello: %v", err)
	}
	log.Printf("Response from the server is equal to: %s", response.Body)

	response, err = c.SayGoodBye(context.Background(), &chat.Message{
		Body: "GoodBye grpc!",
	})
	if err != nil {
		log.Fatal("error when calling SayHello: %v", err)
	}
	log.Printf("Response from the server is equal to: %s", response.Body)
}
