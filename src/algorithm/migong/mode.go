package main

import "fmt"

func test(n uint64) (res uint64){
	if n > 1 {
		res = n * test(n-1)
		return res
	}
	return 1
}
func main() {
	res := test(10)
	fmt.Println(res)
}
