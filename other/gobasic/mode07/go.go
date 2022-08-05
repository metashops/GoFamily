package main

import "fmt"

func newInt() *int {
	a := 11
	return &a
}
func main() {
	a := newInt()
	//打印值
	fmt.Println(*a)
	//打印出来的是地址
	fmt.Println(a)
}
