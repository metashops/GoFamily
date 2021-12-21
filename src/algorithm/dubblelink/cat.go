package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func AddCat(head *CatNode, newHead *CatNode) {
	if head.next == nil {
		head.no = newHead.no
		head.name = newHead.name
		head.next = head //形成环状
		fmt.Printf("欢迎[%v]第一个加入环形", newHead)
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newHead
	newHead.next = head
}
func ShowLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty")
		return
	}
	for {
		fmt.Printf("猫信息为：[id=%d,name=%s]->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}
func main() {
	//初始化头节点
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tmo",
	}
	cat2 := &CatNode{
		no:   2,
		name: "lisa",
	}
	cat3 := &CatNode{
		no:   3,
		name: "smi",
	}
	cat4 := &CatNode{
		no:   4,
		name: "mali",
	}
	AddCat(head, cat1)
	AddCat(head, cat2)
	AddCat(head, cat3)
	AddCat(head, cat4)
	fmt.Println()
	ShowLink(head)
}
