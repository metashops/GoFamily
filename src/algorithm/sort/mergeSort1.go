package main

import "fmt"

// MergeSort 自顶向下归并排序，排序范围在 [begin,end) 的数组
func MergeSort(array []int, begin int, end int) {
	// 元素数量大于1时才进入递归
	if end-begin > 1 {

		// 将数组一分为二，分为 array[begin,mid) 和 array[mid,high)
		mid := begin + (end-begin+1) >> 1

		// 先将左边排序好
		MergeSort(array, begin, mid)

		// 再将右边排序好
		MergeSort(array, mid, end)

		// 两个有序数组进行合并
		merge(array, begin, mid, end)
	}
}
// 归并操作
func merge(array []int, begin int, mid int, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin         // 左边数组的长度
	rightSize := end - mid          // 右边数组的长度
	newSize := leftSize + rightSize // 辅助数组的长度
	result := make([]int, 0, newSize)

	left, right := 0, 0
	for left < leftSize && right < rightSize {
		lValue := array[begin+left] // 左边数组的元素
		rValue := array[mid+right]  // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			left++
		} else {
			result = append(result, rValue)
			right++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+left:mid]...)
	result = append(result, array[mid+right:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}

func main() {
	arr := []int{4,7,6,2,8,3,5,9,1}
	MergeSort(arr,0,len(arr))
	fmt.Println(arr)
}
