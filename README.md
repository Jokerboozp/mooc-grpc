### protobuf官方文档

```text
https://protobuf.dev/programming-guides/proto3/
```

### GRPC 和 protobuf 安装更新

- Mac 环境

```shell
brew install protobuf

go get -u google.golang.org/protobuf/cmd/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### protobuf生成 grpc 文件

```shell
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. helloworld.proto
```

- 命令解释

```text
protoc: 这是Protocol Buffers的编译器，用于编译.proto文件。.proto文件是一种接口描述语言文件，用于定义数据结构和服务接口。

--go_out=.: 这个参数指定了生成的Go语言代码的输出目录。在这个例子中，.表示当前目录。这部分是用于生成普通的Go语言代码，通常包括结构体（对应.proto文件中的消息）和序列化/反序列化函数。

--go-grpc_out=require_unimplemented_servers=false:.: 这个参数也是指定输出目录，同样是当前目录。require_unimplemented_servers=false是一个选项，告诉编译器在生成的gRPC服务代码中，如果有接口没有实现，不要生成报错代码。这个选项有时候在开发阶段很有用，因为你可能还没实现所有的服务接口。这部分是专门用于生成gRPC相关的Go代码，包括服务接口和客户端存根。

helloworld.proto: 这是要编译的protobuf文件名。它包含了要通过protobuf定义的数据结构和gRPC服务接口。
```

### grpc的四种模式

- 简单模式
```text
这种模式最为传统，即客户端发起一次请求，服务端响应一次数据
```

- 服务端数据流模式
```text
这种模式是客户端发起一次请求，服务端返回一段连续的数据流。典型的例子就是客户端向服务端发送一段股票代码，服务端就把该股票代码的实时数据源不断的返回给客户端
```

- 客户端数据流模式
```text
与客户端数据流模式相反，这次是客户端不断向服务端发送数据，而在发送结束后，由服务端返回一个响应。典型的例子就是物联网终端向服务器报送数据
```

- 双向数据流模式
```text
这种就是客户端和服务端都可以向对方发送数据流，这个时候双方的数据可以同时互相发送，也就是可以实现实时交互。典型的例子就是聊天机器人
```
