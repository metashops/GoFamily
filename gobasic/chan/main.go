package main

import (
	"fmt"
	"reflect"
)

func main() {
	var ch1 = make(chan int,10)
	var ch2 = make(chan int,10)
	var cases = createCases(ch1,ch2)
	for i :=0;i<10;i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() {
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		}else {
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}
func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase
	for _, ch := range chs {
		cases = append(cases,reflect.SelectCase{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases,reflect.SelectCase{
			Dir: reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}