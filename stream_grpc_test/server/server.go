package main

import (
	"fmt"
	"google.golang.org/grpc"
	"mooc-grpc/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
}

func (s *server) GetStream(req *proto.StreamRequestData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResponseData{Data: fmt.Sprintf("服务端流模式测试%d", i)})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allStr.Send(&proto.StreamResponseData{Data: "服务端-双向流模式"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		panic(err)
	}

}
