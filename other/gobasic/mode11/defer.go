package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i," ")
	}
}
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i," ")
		}()
	}
	fmt.Println()
}
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n," ")
		}(i)
	}
	fmt.Println()
}
func main() {
	d1()
	fmt.Println("=====================")
	d2()
	fmt.Println("=====================")
	d3()
	fmt.Println("=====================")
}