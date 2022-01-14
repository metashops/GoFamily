package main

import (
	"fmt"
	"reflect"
)

// Monster 定一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	score float32
	sex   string
}

//方法，显示s的值
func (s Monster) print() {
	fmt.Println("start...")
	fmt.Println(s)
	fmt.Println("end...")
}
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.score = score
	s.sex = sex
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//如果传入不是struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Println(num)
	for i := 0; i < num; i++ {
		fmt.Println(val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d:tag为=%v\n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	val.Method(1).Call(nil)

}
func main() {
	//创建一个Monster实例
	var a Monster = Monster{
		Name:  "诸葛亮",
		Age:   23,
		score: 59,
	}
	TestStruct(a)
}
