package main

import "fmt"

func QuickSort(arr []int, begin, end int) {
	if begin < end {
		loc := partition(arr,begin,end)
		QuickSort(arr,begin,loc-1)
		QuickSort(arr,loc,end-1)
	}
}

//切分函数
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
	QuickSort(arr, len(arr)-1, 0)
	fmt.Println(arr)
}
