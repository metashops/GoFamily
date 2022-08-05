package main

import "fmt"

type Queue struct {
	maxSize int    //大小
	array   [4]int //数组
	front   int    //指向队列首
	rear    int    //指向队列尾
}

func (this *Queue) AddQueue(val int) (err error) {
	if this.rear == this.maxSize-1 {
		return err
	}
	this.rear++
	this.array[this.rear] = val
	return
}
func (this *Queue) showQueue() {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Println(i)
	}
}
func main() {
	var queue *Queue = &Queue{}
	queue.maxSize = 5
	queue.front = -1
	queue.rear = -1
	queue.AddQueue(2)
	queue.AddQueue(1)
	queue.AddQueue(3)
	queue.showQueue()

}
