**首先安装 Protocol Buffers v3**

```shell
$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protobuf-all-3.11.4.tar.gz
$ sudo tar -zxvf protobuf-all-3.11.4.tar.gz -C /usr/local
$ cd /usr/local/protobuf-3.11.4/
$ sudo ./configure
$ sudo make
$ sudo make install
$ sudo ldconfig
```

**新建项目Go后，提前把相关插件安装： protoc-gen-go、protoc-gen-grpc-gateway、protoc-gen-swagger**

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

**编写proto文件，注意google/api/annotations.proto这个目录自己新建，然后去GitHub搜索Googleapis拷贝过来即可。**

```protobuf
syntax = "proto3";

package helloworld;
option go_package = "./;helloworld";

import "google/api/annotations.proto";

// Here is the overall greeting service definition where we define all our endpoints
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
      post: "/v1/example/echo"
      body: "*"
    };
  }
}

// The request message containing the user's name
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}

```

##### 使用protoc命令编译生成go源码及swagger.json

```shell
# 生成proto对应的gateway源码
➤ protoc -I ./ -I $GOPATH/src -I $GOPATH/src/google/api --grpc-gateway_out=logtostderr=true:. ./helloworld.proto 

# 生成proto对应的swagger.json文件
 ➤ protoc -I ./ -I $GOPATH/src -I $GOPATH/src/google/api --swagger_out=logtostderr=true:. ./helloworld.proto 
 
 # 将proto文件编译成go文件，生成golang的服务代码
 ➤ protoc -I ./ -I $GOPATH/src -I $GOPATH/src/google/api --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto 
```

* `--go_out`：生成`golang`源文件
* `-IPATH, --proto_path=PATH`：指定`import`搜索的目录，可指定多个，不指定则默认当前工作目录

目录结构

```shell
➤ tree                                                                                                                                                                                                               
.
├── go.mod
├── go.sum
└── proto
    ├── google
    │   └── api
    │       ├── annotations.proto
    │       └── http.proto
    └── helloworld
        ├── helloworld.pb.go
        ├── helloworld.pb.gw.go
        ├── helloworld.proto
        ├── helloworld.swagger.json
        └── helloworld_grpc.pb.go
```

使用docker启动swagger ui服务，查看生成的API文档

```shell
docker run --rm -d -p 80:8080 -e SWAGGER_JSON=/foo/hello.swagger.json -v /path/to/project/proto:/foo swaggerapi/swagger-ui
```

编写 server

```go
package main

import (
	`context`
	`fmt`
	`net`

	`google.golang.org/grpc`

	pb `examples/proto/helloworld`
)

const (
	address = "172.0.0.1:8080"
)
type Servers struct {}


func (s *Servers) SayHello(ctx context.Context, request *pb.HelloRequest)(*pb.HelloReply,error) {
	return &pb.HelloReply{
		Message: "hello" + request.Name,
	},nil
}

func main() {
	// 1、实例化gRPC的server
	r := grpc.NewServer()

	// 2、注册
	pb.RegisterGreeterServer(r,&Servers{})

	// 3、监听
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to Listening,err:%s",err)
	}
	err = r.Serve(lis)
	if err != nil {
		println("failed...")
	}
}

```

client

```go
package main

import (
	`context`
	`fmt`

	`google.golang.org/grpc`

	pb `examples/proto/helloworld`
)

const (
	address = "127.0.0.1:8080"
)
func main() {
	// 1、连接
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		println("failed...")
	}
	defer conn.Close()

	// 2、实例化一个Client
	c := pb.NewGreeterClient(conn)

	// 3、调用
	r,err := c.SayHello(context.Background(),&pb.HelloRequest{
		Name: " bobby",
	})
	if err != nil {
		println("failed...")
	}
	fmt.Println(r.Message)

}
```

ingress 规则会生效到所有安装 ingress Contreoller 的机器的 nginx 配置
