package main

import "fmt"

func QuickSort(arr []int,begin,end int) {
	if begin < end {
		loc := partition(arr,begin,end) //切分
		QuickSort(arr,begin,loc-1) //对左部分进行快排
		QuickSort(arr,loc+1,end) //对右部分进行快排
	}
}
func partition(arr []int,begin,end int) int{
	i := begin+1
	j := end
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
func main() {
	arr := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
