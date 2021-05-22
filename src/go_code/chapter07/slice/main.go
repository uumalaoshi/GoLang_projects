package main

import (
	"fmt"
)

func main() {
	var heols [5]int = [...]int{34, 23, 26, 53, 55}
	var slice = heols[2:3]
	fmt.Println("heols的元素=", heols)
	fmt.Println("slice的元素", slice)
	fmt.Println("slice的长度", len(slice))
	fmt.Printf("slice第一个元素的地址=%p", &slice[0])
	fmt.Printf("heols第3个元素的地址=%p ", &heols[2])
}
