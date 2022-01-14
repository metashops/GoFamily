package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("hello"))
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
