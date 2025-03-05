package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc/proto" // import the generated Go code
)

type server struct {
	proto.UnimplementedExampleServer
}

func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error receiving request: %v", err)
			return fmt.Errorf("error receiving request: %v", err)
		}

		log.Printf("Received request: %s", req.GetSomestring())

		reply := &proto.HelloResponse{Reply: "Hello, " + req.GetSomestring()}
		if err := stream.Send(reply); err != nil {
			log.Printf("error sending response: %v", err)
			return fmt.Errorf("error sending response: %v", err)
		}
		log.Printf("Sent response: %s", reply.GetReply())
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