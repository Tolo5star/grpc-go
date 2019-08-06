package main

import (
	"context"
	"grpc-go/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	//defer conn.Close()

	cli := proto.NewPingClient(conn)

	response, err := cli.SendString(context.Background(), &proto.RequestString{Mess: "Hello"})
	if err != nil {
		log.Fatalf("Error in client :%v", err)
	}
	log.Printf("Message recieved : %t", response.Sent)

}
