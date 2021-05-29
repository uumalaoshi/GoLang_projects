package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (p Phone) Call() {
	fmt.Println("手机呼叫开始工作")
}

type Camera struct {
	Name string
}

func (p Camera) Start() {
	fmt.Println("相机开始工作")
}
func (p Camera) Stop() {
	fmt.Println("相机停止工作")
}

type computer struct {
}

func (computer computer) Working(usb Usb) {
	usb.Start()
	if Phone, OK := usb.(Phone); OK {
		Phone.Call()
	}
	usb.Stop()
}
func main() {
	var Usbarr [3]Usb
	Usbarr[0] = Phone{"vivo"}
	Usbarr[1] = Phone{"小米"}
	Usbarr[2] = Camera{"尼康"}
	fmt.Println(Usbarr)
	var computer computer
	for _, v := range Usbarr {
		computer.Working(v)
	}

}
