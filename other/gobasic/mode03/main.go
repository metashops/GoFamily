package main

import "fmt"

type A interface {
	say()
}
type B struct {
	Name string
}

func (b B) say() {
	fmt.Println("ok")
}

type Person interface {
	hello()
}
type student interface {
	good()
}
type Monster struct{}

func (m Monster) hello() {
	fmt.Println("实现Person接口方法")
}
func (m Monster) good() {
	fmt.Println("实现student接口方法")
}
func main() {
	var monster Monster
	var p Person = monster
	var s student = monster
	p.hello()
	s.good()
}
