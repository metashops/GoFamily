package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
  }

func levelOrder(root *TreeNode) [][]int {
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
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tempArr=append(tempArr,node.Val)
		}
		res=append(res,tempArr)
		tempArr=[]int{}
	}
	return res
}
func main() {
	root := &TreeNode{
		Val: 1,
	}
	left2 := &TreeNode{
		Val: 2,
	}
	left4 := &TreeNode{
		Val: 4,
	}
	left6 := &TreeNode{
		Val: 6,
	}

	right3 := &TreeNode{
		Val: 3,
	}
	right5 := &TreeNode{
		Val: 5,
	}
	right7 := &TreeNode{
		Val: 7,
	}
	root.Left = left2
	root.Right = right3
	left2.Left = left4
	left2.Right = right5

	right3.Left = left6
	right3.Right = right7

	fmt.Println(levelOrder(root))
}