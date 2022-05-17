package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int{
	m := make(map[int]int) //存储已经遇到的数字
	for _,v := range nums {
		if _, ok :=m[v]; ok{
			fmt.Println(m)
			return v
		} else {
			m[v] = 1
		}
	}
	return -1
}
func findRepeatNumber1(nums []int) int{
	sort.Ints(nums)
	ns := len(nums) - 1
	for i := 0; i < ns; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{7, 5, 1, 0, 2, 5, 3}
	number := findRepeatNumber1(nums)
	fmt.Println(number)
}
