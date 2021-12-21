package main

import "fmt"

func changeMap(m map[string]string) {
	m["name"] = "曹操"
	fmt.Printf("m=m%p,m=%v", m, m)
}
func main1() {
	m1 := map[string]string{"name": "诸葛亮"}
	fmt.Printf("m1=%p,m1=%v\n", m1, m1)
	changeMap(m1)
}
func main() {
	fmt.Println("hello world")
}
