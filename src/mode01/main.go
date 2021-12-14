package main

import "fmt"
func main() {
	var a map[string]string
	a = make(map[string]string,10)
	a["no1"] = "诸葛亮"
	a["no2"] = "曹操"
	a["no3"] = "项羽"
	a["no4"] = "关羽"
	fmt.Println(a)
	cities := make(map[string]string)
	cities["no1"] = "广州"
	cities["no2"] = "深圳"
	cities["no3"] = "上海"
	fmt.Println(cities)
	data := map[string]string{
		"name":"lisa",
		"age" :"23",
		"sex" :"female",
	}
	data["id"] = "1"
	for k1,v1 := range data {
		fmt.Printf("k1=%v\tv1=%v",k1,v1)
	}

}
