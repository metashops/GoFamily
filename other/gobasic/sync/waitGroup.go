package main

import (
	"fmt"
	"sync"
	"time"
)
type Counters struct {
	mu sync.Mutex
	count uint64
}
func (c *Counters) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
func (c *Counters) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func Worker(c *Counters,wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}
func main() {
	var counter Counters
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go Worker(&counter,&wg)
	}
	wg.Wait()
	fmt.Println("count:",counter.Count())

}