package main

import (
	"context"
	"fmt"
	proto "grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewExampleClient(conn)
	req := &proto.HelloRequest{Somestring: "hello from client"}
	res,err := client.ServerReply(context.TODO(), req)
	if err != nil{
		log.Fatal(err)
		
	}
	fmt.Println(res.Reply)

}
