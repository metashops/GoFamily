package main

import (
	"log"
	"net"
)

func main() {
	listen,err := net.ListenUDP("UDP",&net.UDPAddr{Port: 20000})
	if err != nil {
		log.Fatalf("listen error:%v\n",err)
	}
	defer listen.Close()
	for {
		var buf[1024]byte
		//直接读
		n,addr,err := listen.ReadFromUDP(buf[:])
		if err !=  nil {
			log.Printf("read udp error:%v\n",err)
			continue
		}
		data := append([]byte("hello"),buf[:n]...)
		listen.WriteToUDP(data,addr)
	}

}