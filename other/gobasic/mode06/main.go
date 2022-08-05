package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取取传入的变量type，kind，值
	//1、获取reflect.Type
	rTyp := reflect.TypeOf(b)
	//2、获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("Type=%T,Value=%v\n", rTyp, rVal) //Type=*reflect.rtype,Value=10
	//获取值后通过运算
	n := rVal.Int() + 1
	fmt.Println(n) //11

	//需求：将rVal转成interface{}
	iV := rVal.Interface()
	//将interface{}通过断言转成需要的类型
	num1 := iV.(int)
	fmt.Println("num1=", num1) //10
}

type Student struct {
	Name string
	Sex  string
	Age  int
}

func reflectTest02(b interface{}) {
	//1、获取reflect.Type
	rTyp := reflect.TypeOf(b)
	//2、获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("Type=%T,Value=%v\n", rTyp, rVal)

	//3、将rVal转成 Interface{}
	n := rVal.Interface()
	fmt.Println(n)

	//将interface{}通过断言转成需要的类型
	stu, ok := n.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}
func main() {
	var num int = 1
	reflectTest01(num)
	stu := Student{
		Name: "诸葛亮",
		Sex:  "男",
		Age:  23,
	}
	reflectTest02(stu)
}
