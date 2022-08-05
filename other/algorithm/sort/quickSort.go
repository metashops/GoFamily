package main

import "fmt"

func QuickSort(arr []int,begin,end int) {
	if begin < end {
		loc := Partition(arr,begin,end) //切分
		QuickSort(arr,begin,loc-1) //对左部分进行快排
		QuickSort(arr,loc+1,end) //对右部分进行快排
	}
}
func Partition(arr []int,begin,end int) int{
	i,j := begin+1,end
	for i < j {
		if arr[i] > arr[begin] {
			arr[i],arr[j] = arr[j],arr[i]
			j--
		} else {
			i++
		}
	}
	if arr[i] >= arr[begin] {
		i--
	}
	arr[begin],arr[i] = arr[i],arr[begin]
	return i
}
func quickSort1(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j, base := left, right, nums[left]               //优化点，随机选基准
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort1(nums, left, i - 1)
	quickSort1(nums, i + 1, right)
}

func main() {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(arr, 0, len(arr)-1)
	//fmt.Println(arr)
	quickSort1(arr,0,len(arr)-1)
	fmt.Println(arr)
}
