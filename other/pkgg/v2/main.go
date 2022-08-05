package main

var (
	B int
	A int
)
func f1() {
	A = 1
	println(B)
}
func f2() {
	B = 1
	println(A)
}
func main() {
	f1()
	f2()
}
