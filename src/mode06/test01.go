package main

import (
	"fmt"
	"reflect"
)

func reflectTest03(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Println("reflect.Value=", rVal)
	rTyp := reflect.TypeOf(b)
	//rKin := reflect.Kind()
	fmt.Println("rType=", rTyp)

	iVal := rVal.Interface()
	fmt.Println("转成interface=", iVal)
	//断言
	f := iVal.(float64)
	fmt.Println("断言=", f)
}
func main() {
	var v float64 = 1.2
	reflectTest03(v)
}
