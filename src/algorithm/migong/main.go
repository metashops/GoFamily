package main

import "fmt"
//编写函数完成找路
func setWay(mymap *[8][7]int,i,j int) bool {
	if mymap[6][5] == 2 {
		return true
	} else {
		if mymap[i][j] == 0 {
			mymap[i][j] = 2
			if setWay(mymap,i+1,j){//下
				return true
			} else if setWay(mymap,i,j+1) { //右
				return true
			} else if setWay(mymap,i-1,j) { //上
				return true
			} else if setWay(mymap,i,j+1) { //左
				return true
			} else { //思路
				mymap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}
func main() {
	//1、如果元素值为1，就是墙
	//2、如果元素值为0，是没有走过的点
	//3、如果元素值为2，是一个通路
	//4、如果元素值为3，是走过的点，但是走不通
	var mymap [8][7]int
	//先把地图最上和最下设置为1
	for i := 0; i < 7; i++ {
		mymap[0][i] = 1
		mymap[7][i] = 1
	}
	for i := 0; i < 7; i++ {
	mymap[i][0] = 1
	mymap[i][6] = 1
	}
	mymap[3][1] = 1
	mymap[3][2] = 1
	setWay(&mymap, 1, 1)
	fmt.Println("探测完毕！")
	//输出地图
	for i:=0;i<8;i++{
		for j:=0;j<7;j++{
			fmt.Print(mymap[i][j]," ")
		}
		fmt.Println()
	}
}
