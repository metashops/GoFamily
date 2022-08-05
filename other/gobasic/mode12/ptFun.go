package main

import "fmt"

func getPtr(v *float64) float64 {
	return *v + *v
}
func returnPtr(x int) *int {
	y := x + x
	return &y
}
//闭包
func funReturnFun() func() int{
	i := 0
	return func() int {
		i++
		return i*i
	}
}
func main() {
	var myInt interface{} = 123
	i,ok := myInt.(int)
	if ok {
		fmt.Println("success:",i)
	}
	v,ok := myInt.(string)
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("Failed without panicking!")
	}
}
