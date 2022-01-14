package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}
	PORT := ":" + arguments[1]
	i,err := net.Listen("tcp",PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer i.Close()
	c,err := i.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	newData,_ := bufio.NewReader(c).ReadString('\n')
	if strings.TrimSpace(string(newData)) == "STOP" {
		fmt.Println("Exiting TCP server!")
		return
	}
	fmt.Print("->",string(newData))
	t := time.Now()
	myTIme := t.Format(time.RFC3339)+"\n"
	c.Write([]byte(myTIme))
}