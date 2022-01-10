package main

import (
	"container/list"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func InvertTree(root *TreeNode) *TreeNode{
	if root == nil {
		return root
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i:=0;i<length;i++{
			e:=queue.Remove(queue.Front()).(*TreeNode)
			e.Left,e.Right=e.Right,e.Left
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}
func main() {
	root := &TreeNode{Val: 1}
	left2:= &TreeNode{Val: 2}
	left4:= &TreeNode{Val: 4}
	left6:= &TreeNode{Val: 6}
	right3:= &TreeNode{Val: 3}
	right5:= &TreeNode{Val: 5}
	right7:= &TreeNode{Val: 7}
	root.Left = left2
	root.Right = right3
	left2.Left = left4
	left2.Right = right5
	right3.Left = left6
	right3.Right = right7

}
