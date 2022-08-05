package main

import "fmt"

// 单调递减栈
func dailyTemperatures(num []int) []int{
	ans := make([]int,len(num))
	var stack []int
	fmt.Println(len(stack))
	for i,v := range num {
		// 栈不为空 && 当前天温度大于栈顶天温度
		for len(stack) !=0 && v > num[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			// 设置栈顶天的结果值
			ans[top] = i - top
		}
		// 入栈
		stack = append(stack,i)
	}
	return ans
}
func main() {
	num := []int{73, 74, 75, 71, 71, 72, 76, 73}
	temp := dailyTemperatures(num)
	fmt.Println(temp)
}
