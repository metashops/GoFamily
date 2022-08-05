package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan bool,1)
var ch2 = make(chan bool)
func fun1() {
	ch1 <- true
	for i:=1;i<=10;i += 2{
		<- ch1
		fmt.Print(i)
		fmt.Print(i+1," ")
		ch2 <- true
	}
}
func fun2() {
	for i:='A';i<='J';i+=2{
		<- ch2
		fmt.Print(string(i))
		fmt.Print(string(i+1)," ")
		ch1 <- true
	}
}


func main() {
	go fun1()
	go fun2()
	time.Sleep(time.Second)
}