package main

import (
	"fmt"
	"sync"
	"time"
)
var (
	c1 = make(chan int,1)
	c2 = make(chan int,1)
	c3 = make(chan int,1)
	c4 = make(chan int,1)
)
func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	c4 <- 1
	go func() {
		for {//2
			time.Sleep(time.Second)
			<- c4
			fmt.Println(1," ")
			c1 <- 1
			wg.Done()
		}
	}()
}