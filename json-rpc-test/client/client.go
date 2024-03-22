package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:1234")

	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	_ = client.Call("HelloService.Hello", "jokerboozp", &reply)

	fmt.Println(reply)
}
