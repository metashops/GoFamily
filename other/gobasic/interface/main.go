package main

import "fmt"
type Cat struct{Name string}
type Duck interface {
	Quack()
}
func (c *Cat) Quack() {//实现Duck接口
	fmt.Println("meow")
}
func main() {
	var c Duck = &Cat{Name: "draven"}
	c.Quack()
}
