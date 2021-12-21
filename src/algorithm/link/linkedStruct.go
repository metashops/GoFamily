package main

import "fmt"

type HeroNode struct {
	node     int
	name     string
	nickname string
	next     *HeroNode //指向下一个节点
}

// InsertHeroNode 尾添加
func (this *HeroNode) InsertHeroNode(newHeroNode *HeroNode) {
	//末尾添加，首先遍历到最后一个节点
	temp := this
	for {
		if temp.next == nil {
			break
		} else {
			temp = temp.next
		}
	}
	//添加
	temp.next = newHeroNode
}
func (this *HeroNode) InsertHeroNode1(newHeroNode *HeroNode) {
	tmp := this
	flag := true
	for {
		if tmp.next != nil {
			break
		} else if tmp.next.node > newHeroNode.node {
			break
		} else if tmp.next.node == newHeroNode.node {
			flag = false
			break
		}
		tmp = tmp.next
	}
	if !flag {
		return
	} else {
		newHeroNode.next = tmp.next
		tmp.next = newHeroNode
	}
}

// ListHeroNode 遍历所有节点
func (this *HeroNode) ListHeroNode() {
	if this.next == nil {
		fmt.Println("Linked is empty")
		return
	}
	tmp := this.next
	for {
		fmt.Printf("[no:%v,name:%v,nickname:%v]\n", tmp.node, tmp.name, tmp.nickname)
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
}
func main() {
	//实例化表头
	head := &HeroNode{}

	//实例化第一条数据
	head1 := &HeroNode{
		node:     1,
		name:     "诸葛亮",
		nickname: "神机妙算",
	}
	head2 := &HeroNode{
		node:     2,
		name:     "项羽",
		nickname: "神",
	}
	head3 := &HeroNode{
		node:     3,
		name:     "刘备",
		nickname: "~",
	}
	head.InsertHeroNode(head1)
	head.InsertHeroNode(head2)
	head.InsertHeroNode1(head3)
	head.ListHeroNode()
}
