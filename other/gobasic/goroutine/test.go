package main

import (
	"fmt"
	"time"
)
var ch1 = make(chan bool,1)
var ch2 = make(chan bool)
var ch3 = make(chan bool)
var ch4 = make(chan bool)
func f1() {
	for i:=0;i<5;i++{
		<- ch1
		fmt.Print("1")
		ch2 <- true
	}
}
func f2() {
	for i:=0;i<5;i++{
		<- ch2
		fmt.Print("2")
		ch3 <- true
	}
}
func f3() {
	for i:=0;i<5;i++{
		<- ch3
		fmt.Print("3")
		ch4 <- true
	}
}
func f4() {
	for i:=0;i<5;i++{
		<- ch4
		fmt.Print("4"," ")
		ch1 <- true
	}
}
func main() {
	ch1 <- true
	go f1()
	go f2()
	go f3()
	go f4()
	time.Sleep(time.Second)
}
