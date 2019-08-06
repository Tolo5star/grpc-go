package proto

import (
	"context"
	"log"
)

type Server struct{}

// Function to send the message
func (s *Server) SendString(ctx context.Context, in *RequestString) (*Response, error) {
	a := in.GetMess()
	log.Println("%s", a)
	return &Response{Sent: true}, nil
}
