package client_proxy

import (
	"mooc-grpc/new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(address, protoc string) HelloServiceStub {
	dial, err := rpc.Dial(protoc, address)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{dial}
}

func (h *HelloServiceStub) Hello(request string, reply *string) error {
	err := h.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
