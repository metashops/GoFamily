package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	str := []byte(s)
	reverse1(str,0,n-1)
	reverse1(str,n,len(str)-1)
	reverse1(str,0,len(str)-1)
	return string(str)
}
func reverse1(b []byte,left,right int) {
	for left < right {
		b[left],b[right] = b[right],b[left]
		left++
		right--
	}
}
func main() {
	str := string("helloworld")
	fmt.Println(str)
	fmt.Println(reverseLeftWords(str,5))
}
