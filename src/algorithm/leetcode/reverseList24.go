package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertHeroNode 尾添加
func (this *ListNode) InsertHeroNode(newHeroNode *ListNode) {
	//末尾添加，首先遍历到最后一个节点
	temp := this
	for {
		if temp.Next == nil {
			break
		} else {
			temp = temp.Next
		}
	}
	//添加
	temp.Next = newHeroNode
}

// 递归反正
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 迭代反正
func ReverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func (this *ListNode) ListHeroNode() {
	if this.Next == nil {
		fmt.Println("Linked is empty")
		return
	}
	tmp := this.Next
	for {
		fmt.Printf("[%v]->", tmp.Val)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
}
func main() {
	head := &ListNode{}
	head1 := &ListNode{
		Val: 1,
	}
	head2 := &ListNode{
		Val: 2,
	}
	head3 := &ListNode{
		Val: 3,
	}
	head.InsertHeroNode(head1)
	head.InsertHeroNode(head2)
	head.InsertHeroNode(head3)
	head.ListHeroNode()
	fmt.Println()
	list := ReverseList(head)
	list.ListHeroNode()

}
