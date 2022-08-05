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
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode){
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
func preorderTraversal1(root *TreeNode) (res []int) {
	ret := []int{}
	if root == nil {
		return ret
	}
	st := list.New()
	st.PushBack(root)
	for st.Len()>0 {
		node := st.Remove(st.Back()).((*TreeNode))
		ret = append(ret,node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ret
}
func inorderTraversal(root *TreeNode) (res []int) {
	ret := []int{}
	if root == nil {
		return ret
	}
	//定义队列
	st := list.New()
	cur := root //
	for cur != nil || st.Len()>0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			ret = append(ret, cur.Val)
			cur = cur.Right
		}
	}
	return ret
}
func main() {
	root := &TreeNode{
		Val: 1,
	}
	n2 := &TreeNode{
		Val: 2,
	}
	n3 := &TreeNode{
		Val: 3,
	}
	n4 := &TreeNode{
		Val: 4,
	}
	n5 := &TreeNode{
		Val: 5,
	}
	n6 := &TreeNode{
		Val: 6,
	}
	n7 := &TreeNode{
		Val: 7,
	}
	root.Left = n2
	root.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println("迭代前序：",preorderTraversal(root))
	fmt.Println("迭代前序：",preorderTraversal1(root))
	fmt.Println("迭代中序：",inorderTraversal(root))
}