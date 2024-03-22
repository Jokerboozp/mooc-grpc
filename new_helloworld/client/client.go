package main

import (
	"fmt"
	"mooc-grpc/new_helloworld/client_proxy"
)

func main() {
	//建立连接
	client := client_proxy.NewHelloServiceClient("localhost:1234", "tcp")

	var reply string
	client.Hello("joker", &reply)
	fmt.Println(reply)
}
