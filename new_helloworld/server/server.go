package main

import (
	"mooc-grpc/new_helloworld/handler"
	"mooc-grpc/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	//实例化一个 server
	listen, _ := net.Listen("tcp", ":1234")

	//注册逻辑 handler
	_ = server_proxy.RegisterHelloService(&handler.HelloService{})

	for {
		//启动服务
		accept, _ := listen.Accept() //当一个新连接进来的时候
		go rpc.ServeConn(accept)
	}

}
