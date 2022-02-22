package main

import "fmt"

func sum(x int) (int,int,bool) {
	return x,x*x,true
}
func main() {
	fmt.Println(sum(3))
}
