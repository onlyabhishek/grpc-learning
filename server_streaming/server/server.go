package main

import (
	"fmt"
	"log"
	"net"

	"grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedExampleServer
}

func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
	reply := "Hello, " + req.GetSomestring()
	for i := 0; i < 5; i++ {
		if err := stream.Send(&proto.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterExampleServer(s, &server{})

	fmt.Println("Server is running on port 9000...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
