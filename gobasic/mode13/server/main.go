package main

import (
	"fmt"
	"io"
	"net"
	"os"
)
func process(conn net.Conn) {
	defer conn.Close()
	//循环接收客户端发送的数据
	for {
		buf := make([]byte,1024)
		s := conn.RemoteAddr().String()
		fmt.Printf("等待客户端{%s}发送消息。。。\n",s)
		n, err := conn.Read(buf)
		if err != io.EOF {
			fmt.Println("Read of server,err:",err)
			return
		}
		//显示client发送消息到server
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	//获取域名
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}
	domain := arguments[1]
	NSs,err := net.LookupMX(domain)
	if err != nil {
		return
	}
	for _,MX := range NSs {
		fmt.Println(MX.Host)
	}
}
