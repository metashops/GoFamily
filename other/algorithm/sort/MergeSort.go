package main

// arr 排序原始数组
// left,mid,right 左中右索引
// temp 辅助数组
func merge1(arr []int,left,mid,right int,temp []int) {
	i := left //初始化i，左边有序序列的初始索引
	j := mid + 1 //初始化j，右边有序序列的初始索引
	t := 0 //指向temp数组当前索引
	//先把左右两边的数据按照拷贝到temp数组，直到左右两边有序序列有一边处理完毕
	for i <= mid && j <= right {
		//左边有序序列当前元素 <= 右边有序序列当前元素，拷贝到temp
		if arr[i] < arr[j] {
			temp[t] = arr[i]
			t += 1
			i += 1
		} else {
			temp[t] = arr[j]
			t += 1
			j += 1
		}
	}
	//
	for i <= mid {
		temp[t] = arr[i]
		t += 1
		i += 1
	}
	for j <= right {
		temp[t] = arr[i]
		t += 1
		j += 1
	}
	//将temp拷贝到arr
	t = 0
}
