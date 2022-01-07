package main

import "fmt"

func bfSearch(source,systax string) int {
	begin := 0
	i,j := 0,0
	n,m := len(source), len(systax)
	for i=0;i < n;begin++ {
		//子串循环
		for j = 0; j < m; j++ {
			if i < n && source[i] == systax[j] {
				i++
			} else {
				//如果没有找到，子串就跳出这个循环，也就是从主串的下一个位置匹配
				break
			}
		}
		if j == m {
			return i-j
		}
		i=begin
		i++
	}
	return -1
}
func bfMatch(source,systax string) int {
	//判断子串长度
	if len(systax) == 0 {
		return 0
	}
	if len(source) == 0 && len(source) < len(systax) {
		return -1
	}
	return bfSearch(source,systax)
}
func main() {
	source := "hello world"
	systax := "world"
	pos := bfMatch(source,systax)
	if pos > 0 {
		fmt.Printf("Find \"%s\"at %d in \"%s\"\n",systax,pos,source)
	} else {
		fmt.Printf("Can not Find \"%s\" | in \"%s\"\n",systax,source)
	}
}