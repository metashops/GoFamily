package main

import "fmt"

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	i, j, pivot := left, right, nums[left]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
func main() {
	arr := []int{9, -1, 3, 6, 2, 7}
	fmt.Println(arr)
	array := sortArray(arr)
	fmt.Println(array)
}
