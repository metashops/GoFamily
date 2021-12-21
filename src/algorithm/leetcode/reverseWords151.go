package main

import (
	"fmt"
	"strings"
)

// split是Go内置分割函数，返回是一个切片
// strings.Join 用于元素类型为string的切片使用分割符号冰姐成一个字符串
func reverseWords(s string) string {
	list := strings.Split(s, " ")
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		if len(list[i]) > 0 {
			res = append(res, list[i])
		}
	}
	s = strings.Join(res, " ")
	return s
}

func main() {
	str := string("a good example   this")
	fmt.Println(reverseWords(str))
	fmt.Println()
}
