package main

import "fmt"

func main() {
	hash := make(map[string]int,2)
	//hash["Ann"] = 1
	//hash["Jon"] = 2
	for k,v := range hash {
		fmt.Printf("key=%s,value=%d\n",k,v)
	}
}
