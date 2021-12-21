package main

import "fmt"

type B interface {
	testB()
}
type C interface {
	testC()
}
type A interface {
	testB()
	testC()
	testA()
}

// Stu 如何要实现A接口，那么需要将B、C接口方法都实现
type Stu struct{}

func (stu Stu) testB() {
	fmt.Println("testB")
}
func (stu Stu) testC() {
	fmt.Println("testC")
}
func (stu Stu) testA() {
	fmt.Println("testA")
}

type T interface{}

func main() {
	var stu Stu
	var t T = stu
	fmt.Println(t)

}
