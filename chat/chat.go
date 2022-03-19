package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received Message body from client: %s", message.Body)
	return &Message{Body: "Hello from the server"}, nil
}

func (s *Server) SayGoodBye(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received Message body from client: %s", message.Body)
	return &Message{Body: "Goodbye from the server"}, nil
}
