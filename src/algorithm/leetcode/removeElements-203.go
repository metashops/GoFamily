package main

type Node struct {
	Val int
	Next *Node
}
func removeElements(head *Node,val int) *Node {
	dummyHead := head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
func reverseList(head *Node) *Node {
	var pre *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}


func main() {

}
