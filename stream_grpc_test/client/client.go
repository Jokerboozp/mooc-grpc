package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mooc-grpc/stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 服务端流模式
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamRequestData{
		Data: "joker",
	})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	// 客户端流模式
	putS, err := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putS.Send(&proto.StreamRequestData{Data: fmt.Sprintf("客户端流模式测试%d", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息：" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamRequestData{Data: "客户端-双向流模式"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
