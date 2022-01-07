package main

import (
	"fmt"
	"time"
)

func function() {
	for i :=0;i<10;i++{
		fmt.Print(i," ")
	}
	fmt.Println()
	time.Sleep(2*time.Second)
}

func main() {
	go function()
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Print(i," ")
		}
	}()
	time.Sleep(1*time.Second)
}