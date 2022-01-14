package main

import "fmt"

func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length; i++ {
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
			fmt.Println(arr)
		}
	}
	return arr
}
func main() {
	arr := []int{3, 1, 0, 2}
	SelectSort(arr)
}
