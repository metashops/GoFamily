package main

import (
	"fmt"
	"sync"
)
type Counter struct {
	mu sync.Mutex
	Count uint64
}
func main() {
	var counter Counter
	//确保所有goroutine执行完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			for j:=0;j<100000;j++{
				counter.mu.Lock()
				counter.Count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count=",counter.Count)
}
