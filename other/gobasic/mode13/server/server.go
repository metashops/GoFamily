package main

import (
	"fmt"
	"io"
	"net"
)
func Process(conn net.Conn) {
	conn.Close()
	buf := make([]byte,1024)
	fmt.Printf("等待客户端{%s}发送消息。。。\n",conn.RemoteAddr().String())
	n,err := conn.Read(buf)
	if err != io.EOF {
		fmt.Println("Read Error:",err)
		return
	}
	//显示client发送消息到server
	fmt.Print(string(buf[:n]))
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Listen Error:",err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("Waiting for client to connect 。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept Error:",err)
			return
		} else {
			addr := conn.RemoteAddr().String()
			fmt.Printf("suc Client:%v\n",addr)
		}
		go Process(conn)
	}
}
