![grpc服务端流程图](https://github.com/metashops/GoFamily/blob/main/assets/image/grpc%E6%9C%8D%E5%8A%A1%E7%AB%AF%E5%90%AF%E5%8A%A8%E6%B5%81%E7%A8%8B.png)

### 1、服务器注册&初始化阶段

**1.1 注册服务**

```go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
  // 实现helloworld.GreeterServer
	pb.UnimplementedGreeterServer
}
// 实现要提供的服务
// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
  // 将服务注册到grpc服务器端
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

```

**1.2 解析器初始化**

解析器有很多种，你也可以自定义解析器，例如 passthrough，dns解析器在grpc服务器启动时会自己注册。而像manual，xds解析器，需要在代码显示注册才能生效。

**1.3 平衡构建器的注册**

**1.4 编解码器初始化**

在grace-go框架中使用 port 作为默认的编解码器,源码位置：`encoding/proto/proto.go`

```go
// 省略部分代码
func init() {
  // 注册编解码器 code
	encoding.RegisterCodec(codec{})
}
```

RegisterCodec 如何实现的，函数内部：

```go
func RegisterCodec(codec Codec) {
	contentSubtype := strings.ToLower(codec.Name())
  // 将 proto 注册到 registeredCodes 容器里（非线程安全）
	registeredCodecs[contentSubtype] = codec
}
```

**1.5 拦截器初始化**

拦截器的初始化主要分为两大步骤：

（1）自定义拦截器

```go
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
  // 客户端调用的服务
	m, err := handler(ctx, req)
	if err != nil {
		logger("RPC failed with error %v", err)
	}
	return m, err
}
```

（2）将拦截器注册到服务器端

```go
1．func main() {
2．   flag.Parse()

3．   lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
4．   if err != nil {
5．      log.Fatalf("failed to listen: %v", err)
6．   }
  // 将拦截器注册到 grpc 服务器
7．   s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))

8．   // Register EchoServer on the server.
9．   pb.RegisterEchoServer(s, &server{})

10．   if err := s.Serve(lis); err != nil {
11．      log.Fatalf("failed to serve: %v", err)
12．   }
13．}

```

### 2、服务器监听工作

**2.1 grpc 服务器是如何监听客户端的请求**

实例源码：

```go
func main() {
	flag.Parse()
  // 创建一个监听目标，只监听tcp协议的请求，端口号是port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// 省略。。。
  // 创建 grpc 服务器
	s := grpc.NewServer()
  // 将服务注册到 grpc 服务器里头
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
  // 启动 grpc 服务器对监听目标开始监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

serve 源码:

```go
func (s *Server) Serve(lis net.Listener) error {
  // 省略次要的代码。。。
  // 表示grpc服务器处于运行状态
	s.serve = true
  // 将监听到进行存储，true表示在监听状态
  ls := &listenSocket{Listener: lis}
	for {
    // 阻塞方式监听客户端请求，如果没有请求时会一直阻塞在这里
		rawConn, err := lis.Accept()
		// 针对新的链接，grpc服务器开启一个协程处理客户端的链接，一个请求对应一个协程
		go func() {
			s.handleRawConn(lis.Addr().String(), rawConn)
			s.serveWG.Done()
		}()
	}
}
```



> 下一篇：grpc 客户端与服务器端如何建立起链接

