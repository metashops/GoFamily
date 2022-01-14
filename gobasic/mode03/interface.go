package main

import "fmt"

type Usb interface {
	start()
	stop()
}
type phone struct{}

//让phone实现USB接口方法
func (p phone) start() {
	fmt.Println("手机开始工作！")
}
func (p phone) stop() {
	fmt.Println("手机停止工作！")
}

type camera struct{}

// 相机也实现USB接口方法
func (c camera) start() {
	fmt.Println("相机开始工作！")
}
func (c camera) stop() {
	fmt.Println("相机停止工作！")
}

//competer
type computer struct{}

func (c computer) working(usb Usb) {
	usb.start()
	usb.stop()
}

func main() {
	computer := computer{}
	phone := phone{}
	camera := camera{}
	computer.working(phone)
	computer.working(camera)
}
