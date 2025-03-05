package main

import (
	
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	 proto "grpc/proto"
)

type server struct {
	proto.UnimplementedExampleServer
}

func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	fmt.Println("Server Reply")
	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				res:= &proto.HelloResponse{Reply: "HELLO From my server"}
				stream.SendAndClose(res)
			}
			log.Print(err)
			break
		}
		log.Print(req.Somestring)
	}
	
	return nil
	
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterExampleServer(s, &server{})
	fmt.Println("Server listening ")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
