package main

import (
	"container/list"
	"fmt"
)
type TreeNode1 struct {
	Val int
	Left *TreeNode1
	Right *TreeNode1
}
func levelOrderBottom(root *TreeNode1) [][]int {
	queue := list.New()
	res := [][]int{}
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode1)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp =append(tmp,node.Val)
		}
		res=append(res, tmp)
	}
	//反转结果集
	for i:=0;i< len(res)/2;i++ {
		res[i],res[len(res)-i-1] = res[len(res)-i-1],res[i]
	}
	return res
}
func main() {
	root :=  &TreeNode1{
		Val: 1,
	}
	left2 := &TreeNode1{
		Val: 2,
	}
	left4 := &TreeNode1{
		Val: 4,
	}
	left6 := &TreeNode1{
		Val: 6,
	}

	right3 := &TreeNode1{
		Val: 3,
	}
	right5 := &TreeNode1{
		Val: 5,
	}
	right7 := &TreeNode1{
		Val: 7,
	}
	root.Left = left2
	root.Right = right3
	left2.Left = left4
	left2.Right = right5

	right3.Left = left6
	right3.Right = right7

	fmt.Println(levelOrderBottom(root))
}