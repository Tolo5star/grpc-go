#!/bin/bash

#Assign the current directory to a variable
CDIR=$(pwd)

#compile proto
protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto

#Create binary for the client and server 
cd server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
cd ..
cd client
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .
cd ..

#Build docker images
sudo docker build -t server-scratch -f server/Dockerfile.scratch .
sudo docker build -t client-scratch -f client/Dockerfile .