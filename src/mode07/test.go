package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person = Person{
		Name: "诸葛亮",
		Age:  11,
	}
	fmt.Println(p)
	p1 := Person{
		Name: "关羽",
	}
	fmt.Println(p1)
	var p2 *Person = new(Person)
	(*p2).Name = "smith" //等价=p.Name="john"
	p2.Age = 2
	fmt.Println(*p2)
	var person *Person = &Person{}
	person.Name = "李四"
	(*person).Name = "张三"
	(*person).Age = 26
	fmt.Println(&person) //获取地址
	fmt.Println(person)
	fmt.Printf("p=%T", person)
}
