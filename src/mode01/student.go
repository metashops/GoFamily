package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}

func main() {
	students := make(map[string]Student,10)
	stu1 := Student{
		1,"tom",23,
	}
	stu2 := Student{
		2,"mary",22,
	}
	students["k1"] = stu1
	students["k2"] = stu2
	fmt.Println(students)
	fmt.Printf("K1=%v",students["k1"])
	for k,v := range students {
		fmt.Printf("\t学生编号%v,值%v",k,v)
	}
}
