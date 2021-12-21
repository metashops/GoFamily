package main

import "fmt"

func QuickSort(left int, right int, arr *[6]int) {
	l, r := left, right
	//找到需要分割的点
	pivot := arr[(left+right)>>1]
	//将比pivot小放左边，大就放右边
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		//说明分解完成
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		//优化
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//l==r，再移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, arr)
	}
}

func main() {
	arr := []int{9, -1, 3, 6, 2, 7}
	fmt.Println(arr)
}
