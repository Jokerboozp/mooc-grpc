package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//建立连接
	client, _ := rpc.Dial("tcp", "localhost:1234")

	var reply string
	err := client.Call("HelloService.Hello", "jokerboozp", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
