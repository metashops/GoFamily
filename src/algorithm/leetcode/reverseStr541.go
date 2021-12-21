package main

import "fmt"

func reverseStr1(s string, k int) string {
	str := []byte(s)
	fmt.Println("反转前：", string(str))
	var i int = 0
	for i < len(str) {
		left := i
		right := i + k - 1
		if right >= len(str) {
			right = len(str) - 1
		}
		for left < right && left < len(str) {
			str[left], str[right] = str[right], str[left]
			left++
			right--
		}
		i += 2 * k
	}
	return string(str)
}
func reverseStr2(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			reverse(ss[i : i+k])
		} else {
			reverse(ss[i:length])
		}
	}
	return string(ss)
}
func reverse(b []byte) {
	left := 0
	right := len(b)
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseStr3(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i = i + 2*k {
		if i+k <= len(str) {
			reverse(str[i : i+k])
		} else {
			reverse(str[i:len(str)])
		}
	}
	return string(str)
}

func main() {
	//str := "abcdef"
	str := "123456789"
	fmt.Println(reverseStr1(str, 2))
}
