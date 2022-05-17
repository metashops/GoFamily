package main

import "fmt"

func main() {
	var ages map[string]int
	if age,ok := ages["a"];!ok{
		fmt.Println(age)
	}
}
