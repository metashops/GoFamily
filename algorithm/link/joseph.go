package main

import "fmt"

type Boy struct {
	No   int  //编号
	Next *Boy //指向下一个指针
}

// AddBoy 编写函数，西形成单向链表
// num 代表小孩子个数
// *Boy 返回该环形的链表的第一个小孩子的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		return first
	}
	//构建该环形链表
	for i := 1; i < num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = boy
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}
	return first
}
func ShowBoy(first *Boy) {
	//处理环形为空情况
	if first.Next == nil {
		fmt.Println("nil...")
		return
	}
	curBoy := first
	for {
		fmt.Printf("编号=%d ->", curBoy.No)
		if curBoy.Next == first {
			break
		}
		//移动下一个指针
		curBoy = curBoy.Next
	}
}
func main() {
	first := AddBoy(5)
	ShowBoy(first)
}
