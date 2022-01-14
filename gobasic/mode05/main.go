package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name string
	Age  int
}
type HeroSlice []Hero

func (hs HeroSlice) Len(i, j int) int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlice) swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	//先定义一个数组/切片
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
}
