package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"grpc/proto" // import the generated Go code
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewExampleClient(conn)

	stream, err := client.ServerReply(context.Background())
	if err != nil {
		log.Fatalf("could not initiate stream: %v", err)
	}

	
	go func() {
		for i := 1; i <= 3; i++ {
			req := &proto.HelloRequest{Somestring: fmt.Sprintf("Message %d", i)}
			if err := stream.Send(req); err != nil {
				log.Fatalf("could not send message: %v", err)
			}
		}
		stream.CloseSend()
	}()


	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving response: %v", err)
		}
		fmt.Printf("Received response: %s\n", resp.Reply)
	}
}
