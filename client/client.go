package main

import (
	"context"
	"fmt"
	"grpc-go/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	cli := proto.NewPingClient(conn)
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for t := range ticker.C {
			_ = t //just to use t to avoid error
			response, err := cli.SendString(context.Background(), &proto.RequestString{Mess: "Hello"})
			if err != nil {
				log.Fatalf("Error in client :%v", err)
			}
			fmt.Println("Message recieved : ", response.Sent)
		}
	}()
	time.Sleep(300000 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
