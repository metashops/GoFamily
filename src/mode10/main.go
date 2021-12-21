package main

import "fmt"

type Dog struct {
	Name string
}

// SetName 接收者是*Dog类型，那么该方法就是基本类型Dog的指针方法
func (dog *Dog) SetName(name string) {
	dog.Name = name
}
func main() {
	chan1 := make(chan int, 3)
	//向通道写数据
	chan1 <- 3
	chan1 <- 1
	chan1 <- 5
	//从通道读取数据
	elem1 := <-chan1
	elem2 := <-chan1
	elem3 := <-chan1
	elem4 := <-chan1
	fmt.Println(elem1)
	fmt.Println(elem2)
	fmt.Println(elem3)
	fmt.Println(elem4)
	//fmt.Println(elem4)
}
