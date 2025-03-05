package main

import (
	"grpc/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := proto.NewExampleClient(conn)

	stream, err := client.ServerReply(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 5; i++ {
		req := &proto.HelloRequest{
			Somestring: fmt.Sprintf("Hello %d", i),
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatal(i, err)
		}
		fmt.Printf("Sent: %s\n", req.Somestring)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Reply)

}
