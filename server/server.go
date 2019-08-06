package main

import (
	"grpc-go/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	serv := proto.Server{}

	grpcServ := grpc.NewServer()

	proto.RegisterPingServer(grpcServ, &serv)

	if err := grpcServ.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
