package main

import (
	"fmt"
	"sync"
	"time"
)
var chan1 = make(chan bool,1)
var chan2 = make(chan bool)
var index = make(chan bool,1)
func func1() {
	for i:=1;i<10;i+=2{
		<- chan1
		fmt.Print(i)
		fmt.Print(i+1," ")
		chan2 <- true
	}
}
func func2() {
	for i:='A';i<='I';i+=2{
		<- chan2
		fmt.Print(string(i))
		fmt.Print(string(i+1)," ")
		chan1 <- true
	}
	index <- true
}
var num = make(chan int,1)
var char = make(chan int,1)
var wg sync.WaitGroup
func func3() {
	defer wg.Done()
	for i:=0;i<=12;i++{
		for j:=0;j<2;j++{
			fmt.Print(2*i+j+1," ")
		}
		num <- 1
		<- char
	}
}
func func4(){
	defer wg.Done()
	for i:=0;i<=12;i++{
		<- num
		for j:=0;j<2;j++{
			fmt.Printf("%c ",'A'+(2*i+j))
		}
		char <- 1
	}
}
var ch = make(chan int,1)
func func5(){
	go func() {
		for i:=1;i<=100;i++{
			if i%2==1{
				fmt.Print(i," ")
			}
		}
		fmt.Println()
		ch <- 1
	}()
	go func() {
		<- ch
		for i:=1;i<=100;i++{
			if i%2==0{
				fmt.Print(i," ")
			}
		}
	}()
}

func main() {
	//chan1 <- true
	//go func1()
	//go func2()
	//<- index
	//wg.Add(2)
	//go func3()
	//go func4()
	//wg.Wait()
	//func5()
	func5()
	time.Sleep(1*time.Second)
}