package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"mooc-grpc/grpc-test/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	md := metadata.Pairs("name", "jokerboozp", "passwd", "imooc")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "joker"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}
