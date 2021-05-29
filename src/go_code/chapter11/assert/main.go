package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var x interface{}
	var b float32 = 4.5
	x = b
	y, ok := x.(float32)
	if ok {
		fmt.Println("代码正确")
		fmt.Printf("y的数据类型是 %T y的值的%v\n", y, y)
	} else {
		fmt.Println("代码错误")
	}
	fmt.Println("继续执行")

}
