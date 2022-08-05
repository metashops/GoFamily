package main

import "fmt"

// 闭包
func GetTaskFunc() func(name string, hours int) (progress int) {
	var progress int = 0
	f := func(name string, hours int) int {
		fmt.Printf("%s头带领队行军%d个小时\n", name, hours)
		progress += hours
		return progress
	}
	return f
}

func main() {
	var p1, p2 int
	f1 := GetTaskFunc()
	f2 := GetTaskFunc()
	p2 = f2("鲁达", 13)
	p1 = f1("武松", 24)
	p2 = f2("鲁达", 12)
	p1 = f1("武松", 1)
	fmt.Println("二哥进度：", p1)
	fmt.Println("大哥进度：", p2)

}
