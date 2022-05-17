package main
import (
	"fmt"
)
func main(){
	s := "a b c"
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(len(s))
}