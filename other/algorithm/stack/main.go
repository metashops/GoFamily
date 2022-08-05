package main

import (
	"errors"
	"fmt"
)

//使用数组模拟栈
type stack struct {
	MaxTop int    //表示栈顶最大
	Top    int    //表示栈顶
	arr    [5]int //数组模拟栈
}

//入栈方法
func (this *stack) push(val int) (err error) {
	if this.Top == this.MaxTop-1 {
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//出栈
func (this *stack) pop() (val int, err error) {
	if this.Top == -1 {
		return 0, errors.New("stack empty")
	}
	data := this.arr[this.Top]
	this.Top--
	return data, nil
}

//遍历栈
func (this *stack) list() {
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack pull")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {
	stack := &stack{
		MaxTop: 5,
		Top:    -1, //表示栈为空
	}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.list()
}
