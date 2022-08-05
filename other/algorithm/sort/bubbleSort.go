package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr)- i -1; j++ {
			if arr[j] > arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
}
func main() {
	arr := []int{3,8,6,5,2,1,4,9,7} //1,2,3,4,5,6,7,8,9
	BubbleSort(arr)
	for i,v := range arr{
		fmt.Printf("下标%d值 = %d\n",i,v)
	}
}
