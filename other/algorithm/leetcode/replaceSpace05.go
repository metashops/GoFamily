package main

import "fmt"

func replaceSpace(s string) string {
	var str string = ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}
func main() {
	s := string("We are happy.")
	fmt.Println(replaceSpace(s))
}
