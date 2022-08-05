package main

import "fmt"

type User struct {
	name string
	age int
	sex int
}
func ModifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		fmt.Printf("%v用户已存在！",name)
		users[name]["pwd"] = "888888"
	} else {
		//用户不存在，就添加
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称："+name
	}
}
func main() {
	users := make(map[string]map[string]string,10)
	//users["smith"] = make(map[string]string,2)
	ModifyUser(users,"smith")
	ModifyUser(users,"tom")
	ModifyUser(users,"tom")
	fmt.Println(users)
}
