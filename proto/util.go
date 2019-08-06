package proto

import (
	"context"
	fmt "fmt"
)

type Server struct{}

//Function to send the message
func (s *Server) SendString(ctx context.Context, in *RequestString) (*Response, error) {
	a := in.GetMess()
	fmt.Println(a)
	return &Response{Sent: true}, nil
}
