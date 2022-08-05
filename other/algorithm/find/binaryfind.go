package main
import "fmt"

func BinaryFind(arr *[6]int,left,right,findVal int) {
	if left > right {
		return
	}
	middle := (left+right) / 2
	if (*arr)[middle] > findVal {
		BinaryFind(arr,left,middle-1,findVal)
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle+1,right,findVal)
	} else {
		fmt.Println(middle)
	}
}
func main() {
	arr := [6]int{1,3,5,6,7,8}
	BinaryFind(&arr,0, len(arr)-1,8)
}
