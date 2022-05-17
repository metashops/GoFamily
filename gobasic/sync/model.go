package main

import (
	"fmt"
	"unsafe"
)

// LKQueue lock-free的queue
type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}
// 通过链表实现，这个数据结构代表链表中的节点
type node struct {
	value interface{}
	next unsafe.Pointer
}
func NewLKQueue() *LKQueue{
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}
//入队
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	fmt.Println(n)
}