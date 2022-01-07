package main

import (
	"bufio"
	"log"
	"net"
)
//服务端
func handleConn(conn net.Conn) {
	defer conn.Close()
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line,_,err := rd.ReadLine()
		if err != nil {
			log.Printf("read error:%v\n",err)
			return
		}
		wr.WriteString("hello")
		wr.Write(line)
		wr.Flush()//一次性syscall
	}
}
//客户端
func main() {
	listen,err := net.Listen("tcp","127.0.0.1:50000")
	if err != nil {
		log.Fatalf("listen error:%v\n",err)
	}
	for {
		conn,err := listen.Accept()
		if err != nil {
			log.Printf("Accept error:%v\n",err)
			continue
		}
		//begin goroutine listen  concatenate
		go handleConn(conn)

	}
}