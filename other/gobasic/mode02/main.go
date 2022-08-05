package main

import "fmt"

func myFunction(i int, arr [2]int) {
	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}

func main() {
	i := 30
	arr := [2]int{66, 77}
	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
	myFunction(i, arr)
	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
}
