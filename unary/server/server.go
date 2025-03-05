package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "grpc/proto" // Update this import path to your actual proto package path
)

type server struct {
    proto.UnimplementedExampleServer
}

func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
    fmt.Println("Received request from client:", req.Somestring)
    fmt.Println("Hello from server")
    return &proto.HelloResponse{
        Reply: "Hello from my server",
    }, nil
}

func main() {
    listener, tcpErr := net.Listen("tcp", ":9001")
    if tcpErr != nil {
        log.Fatalf("Failed to listen on port 9001: %v", tcpErr)
    }
    log.Println("Server is listening on port 9001")

    srv := grpc.NewServer()
    proto.RegisterExampleServer(srv, &server{})
    reflection.Register(srv)

    if err := srv.Serve(listener); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}

