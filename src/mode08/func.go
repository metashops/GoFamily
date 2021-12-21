package main

import "fmt"

func main() {
	//len=5,cap=5
	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("%p\n", slice)
	//追加后，len=7,cap=10
	nSlice := append(slice, 100, 200)
	fmt.Printf("%p\n", nSlice)
	fmt.Println(len(nSlice), cap(nSlice))
}
