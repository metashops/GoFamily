package main

import (
	"fmt"
)

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：把每一个节点的左右孩子翻转一下，即可达到整体翻转的效果
// 递归法：根据Carl哥递归三部曲
//        1、确定递归函数的参数和返回值
//		  2、确定终止条件
//		  3、确定单层递归的逻辑
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree1(root.Left)
	invertTree1(root.Right)
	return root
}

// 前序遍历
func PreorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Key)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res

}
func main() {
	root1 := TreeNode{Key: 1}
	root2 := TreeNode{Key: 3}
	root := TreeNode{Key: 2}
	root.Left = &root1
	root.Right = &root2
	tree1 := invertTree1(&root)
	traversal := PreorderTraversal(tree1)
	fmt.Println(traversal)
}
