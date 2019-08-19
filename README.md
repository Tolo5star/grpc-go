# grpc-go
<h2>This is a simple demonstration of gRPC based server client interaction using Protocol buffers in containers(Docker)</h2>
<h1>Implementation</h1>

The message and service format in the proto is:
```
message RequestString {
    string mess = 1;
}

message Response{
    bool sent = 2;
}

service  Ping {
    rpc SendString(RequestString) returns (Response);
}
```
Where message RequestString contains the string that we want to send , message Response will indicate if the message is sent to the server or not 
service Ping will send the ResquestString from the client to the server and return Response from the server to the client

Compiling the service.proto file will provide us with generated code (format : filename<b>.pb.go</b>) which has Classes and Interfaces based on the ```.proto``` file

In util.go we use struct <b>Server</b> to implement the interface that was generated in the <b>.pb.go</b> file.Then we implement the SendString function wrt to the generated code.

In the server implementation we create an instance of the gRPC server and register our service implementation with it (```RegisterPingServer() ```).

In the client we create a grpc channel to communicate with the server pass this channel to ```NewPingClient() ``` to be able to use the service methods. Ticker is used to send the message every 5 seconds. 

Setup.sh will compile the protobuf and create binaries for both client and server. It will then load these binaries into a Docker Scratch Image and build it. Scratch container is used here to keep the size of the containers minimum (compared to Golang or alpine bases)


Demo.sh will run both the client and server containers and their interaction can be seen. Client runs with "net=host" otherwise it will require the server container to be inside the client container.

To set up a 2 Node Kubernetes cluster I have used Vagrant to setup 3 machines (2 workers and 1 master on <b>ubuntu 16.04 box</b>) and then performed the following on them :
1. Installing prerequisites
2. Install docker
3. Install Kubeadm ,Kubectl,kubelet and Calico in master , then copy the 4 yml files in it.
4. Create Join command shell script and store it in the master.
5. Make the worker nodes join the kubernetes cluster using the script stored in master.

Then create a server pod and NodePort service on the kubernetes cluster and make the client interact with the service it from the host system.

<h1>Running the project</h1>

```
./setup.sh
./demo.sh
```
For the Kubernetes Service (server)
Use ```vagrant up``` to start the cluster
then ssh into master using ```vagrant ssh master``` and start the pod and service with ->
```
cd Yamls
kubectl apply -f server_pod.yml
kubectl apply -f server_npservice.yml
```
this will start the server container on the cluster in the pod and creates a Nodeport service to expose it

run the testClient.go file to interact with the Service
```
go run testClient.go
```
