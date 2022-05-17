package main

var stack1 []int
var stack2 []int

func Push(num int) {
	stack1 = append(stack1,num)
}
func Pop() int{
	if len(stack1) != 0 && len(stack2) == 0 {
		//将stack1压入stack2
		for i:=0;i< len(stack1);i++{
			st := stack1[len(stack1)-1-i]
			stack2 = append(stack2,st)
		}
		//最后将stack1设置为空
		stack1 = []int{}
	}
	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}
func main() {

}
