package main

import "fmt"

func main() {
	str := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(str)
	//写法1
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Println(str)
	//写法2
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	fmt.Println(str)
}
