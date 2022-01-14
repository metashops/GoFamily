package main

//前缀表
func getNext(next []int,s string) {
	j := -1
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] //回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j //next[i]是i（包括i）之前的最长相等前后缀长度
	}
}
func strStr(haystack,needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int,len(needle))
	getNext(next,needle)
	j := -1
	for i := 0; i < len(haystack); i++ {
		for j >= 0 && haystack[i] != needle[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if haystack[i] == needle[j+1] {
			j++
		}
		if j == len(needle)-1 { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}