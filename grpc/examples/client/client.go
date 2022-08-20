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