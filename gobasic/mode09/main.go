package main

import (
	"fmt"
)

func main() {
	var flag = []string{"zero", "one", "two"}
	{
		flag := map[int]string{0: "zero", 1: "one", 2: "two"}
		fmt.Printf("%q\n", flag[2])
		//自动转换为[]string
		a, _ := interface{}(flag).([]string)
		fmt.Printf("aa=%T\n", a)
	}
	a := interface{}(flag).([]string)
	fmt.Printf("a=%T\n", a)

}
