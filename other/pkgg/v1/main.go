package main

import (
	"fmt"
)
func f(x *float64) float64 {
	return *x * *x
}
func ff(x int) *int {
	y := x + x
	return &y
}
func fc() func() int {
	i := 1
	return func() int {
		i++
		return i * i
	}
}
func f1(a int) int {
	return a + a
}
func f2(a int) int {
	return a * a
}
func f3(f func(int) int,v int) int{
	return f(v)
}


func main() {
	var myInt interface{} = 425
	k,ok := myInt.(int)
	if ok {
		fmt.Println("SUCCESS:",k)
	}
	v,ok := myInt.(float64)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Failed without panicking!")
	}
}
