package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// client
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println("连接中。。。")
	if err != nil {
		fmt.Println("Client dial err:",err)
		return
	}
	//fmt.Println("conn 成功=",conn)
	// os.Stdin 是代表终端输入
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ReadString err",err)
			return
		}
		if line == "exit" {
			fmt.Println("Goodbye。。。")
			break
		}
		//将line发给服务器
		_,err = conn.Write([]byte(line+"\n"))
		if err != nil {
			fmt.Println("err",err)
			return
		}
	}

}