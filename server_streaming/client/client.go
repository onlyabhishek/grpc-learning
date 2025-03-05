package main

import (
	"grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return
	}
	defer conn.Close()

	client := proto.NewExampleClient(conn)
	req := &proto.HelloRequest{Somestring: "Hello"}

	stream, err := client.ServerReply(context.Background(), req)
	if err != nil {
		fmt.Printf("Error when calling ServerStream: %v\n", err)
		return
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			fmt.Printf("Error when receiving response: %v\n", err)
			break
		}
		fmt.Printf("Received: %s\n", msg.Reply)
	}

	stream.CloseSend()
}
