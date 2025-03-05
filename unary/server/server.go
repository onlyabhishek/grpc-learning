package main

import (
	"context"
	//"errors"
	"fmt"
	proto "grpc/proto"
	"net"
	"log"
	"google.golang.org/grpc"
	//got"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}


func (s *server) ServerReply(c context.Context,req *proto.HelloRequest)(*proto.HelloResponse,error){
	fmt.Println("recieve request from client",req.Somestring)
	fmt.Println("hello from server")
	return &proto.HelloResponse{
		Reply: "hello from my server",
	},nil
}



func main(){
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		log.Fatal(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	//reflection.Register(srv)
	
    
	// it gives all the clients to discover the available services and method run time 

	fmt.Println("Server is running")
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)

	}



}


