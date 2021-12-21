package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改num int 的值
func reflect01(b interface{}) {
	//1、获取reflect.Value
	rVal := reflect.ValueOf(b)
	//其实它是一个指针
	fmt.Println(rVal.Kind())
	//2、必须通过Elem否则报错，因为它已是指针
	rVal.Elem().SetInt(20)
}
func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println(num)
}
