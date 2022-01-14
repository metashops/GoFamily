package main

import (
	"fmt"
	"time"
)


func process(timeout time.Duration) bool {
	ch := make(chan bool)

	go func() {
		// 模拟处理耗时的业务
		time.Sleep((timeout + time.Second))
		ch <- true // block
		fmt.Println("exit goroutine")
	}()
	select {
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false
	}
}

func main() {
	ch := make(chan int,1)
	ch <- 1
	close(ch)
	<- ch
	close(ch)
}