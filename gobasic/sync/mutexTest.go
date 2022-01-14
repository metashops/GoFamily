package main

import (
	"fmt"
	"sync"
)

type Counter1 struct {
	mu sync.Mutex
	count uint64
}
func main() {
	var counter =  Counter1{}
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			for j:=0;j<100000;j++{
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Count=",counter)
}
