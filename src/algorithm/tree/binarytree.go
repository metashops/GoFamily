package main

import (
	"container/list"
	"fmt"
)

type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

// PreOrder 前序遍历：先root，再左子树，最后右子树
func PreOrder(root *Hero) {
	if root == nil {
		return
	}
	//前序遍历：先root，再左子树，最后右子树
	fmt.Printf("no:%d,name:%s\n",root.No,root.Name)
	PreOrder(root.Left)
	//fmt.Println(root.No) //中序：左、中、右
	PreOrder(root.Right)
	//fmt.Println(root.No) //后序：左、右、中
}
func levelOrder(root *Hero) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tempArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Hero)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr=append(tempArr,node.No)
		}
		res=append(res,tempArr)
		tempArr=[]int{}
	}
	return res
}
func main() {
	//构建二叉树
	root := &Hero{
		No: 1,
		Name: "诸葛亮",
	}
	left1 := &Hero{
		No: 2,
		Name: "曹操",
	}
	right1 := &Hero{
		No: 3,
		Name: "刘邦",
	}
	right2 := &Hero{
		No: 4,
		Name: "关羽",
	}
	root.Left = left1
	root.Right = right1
	right1.Right = right2
	PreOrder(root)
}
