package main

import "fmt"

type HeroNode struct {
	node     int
	name     string
	nickname string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向下一个节点
}

// InsertHeroNode 双向链表尾部添加
func (this *HeroNode) InsertHeroNode(newHeroNode *HeroNode) {
	//末尾添加，首先遍历到最后一个节点
	temp := this
	for {
		if temp.next == nil { //表示找到最后
			break
		} else {
			//让temp不断指向下一个结点
			temp = temp.next
		}
	}
	//将新节点加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}
func (this *HeroNode) InsertHeroNode2(newHeroNode *HeroNode) {
	temp := this
	flag := true
	for {
		if temp.next != nil {
			break
		} else if temp.next.node > newHeroNode.node {
			break
		} else if temp.next.node == newHeroNode.node {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，已经存在", newHeroNode)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

// DelHeroNode 双向链表删除
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.node == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag { //找到
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("sorry,删除节点不存在！")
	}
}

// ListHeroNode 遍历所有节点
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Linked is empty")
		return
	}
	tmp := temp.next
	for {
		fmt.Printf("[no:%v,name:%v,nickname:%v]\n", tmp.node, tmp.name, tmp.nickname)
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
}

// ListHeroNode2 遍历所有节点（逆序打印）
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("Linked is empty")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//遍历
	for {
		fmt.Printf("[%d,%s,%s]==>", temp.node, temp.name, temp.nickname)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}
func main() {
	//创建头节点
	head := &HeroNode{}
	//创建新的节点
	head1 := &HeroNode{
		node:     1,
		name:     "诸葛亮",
		nickname: "神机妙算",
	}
	head2 := &HeroNode{
		node:     2,
		name:     "诸葛亮",
		nickname: "奸雄",
	}
	head3 := &HeroNode{
		node:     3,
		name:     "项羽",
		nickname: "义绝",
	}
	head4 := &HeroNode{
		node:     4,
		name:     "周瑜",
		nickname: "谋略家",
	}
	head5 := &HeroNode{
		node:     5,
		name:     "周瑜5",
		nickname: "谋略家5",
	}
	head.InsertHeroNode(head1)
	head.InsertHeroNode(head2)
	head.InsertHeroNode2(head5)
	head.InsertHeroNode(head3)
	head.InsertHeroNode(head4)
	ListHeroNode(head)
	fmt.Println()
	ListHeroNode2(head)

}
