package main


var ch = make(chan struct{}, 10) // buffered或者unbuffered
var s string
func f() {
	s = "hello, world"
	ch <- struct{}{}
}

func main() {
	go f()
	<-ch
	print(s)
}