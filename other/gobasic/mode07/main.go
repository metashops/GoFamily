package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func test(b interface{}) {
	rTyp := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	fmt.Printf("\nType:%T\nValue:%v\n", rTyp, rVal)
	//reflect.ValueOf->Interface
	iVal := rVal.Interface()
	student := iVal.(Student)
	fmt.Printf("\nInterface{}:%T\n断言:%v\n", iVal, student.Name)

}
func main() {
	stu := Student{
		Name: "诸葛亮",
		Age:  12,
	}
	test(stu)

}
