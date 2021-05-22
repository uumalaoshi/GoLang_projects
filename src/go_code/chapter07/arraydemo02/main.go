package main

import (
	"fmt"
)

func main() {
	var sco [5]float64
	for i := 0; i < len(sco); i++ {
		fmt.Printf("请输入第%d的元素的值\n", i+1)
		fmt.Scanln(&sco[i])

	}
	for i := 0; i < len(sco); i++ {
		fmt.Printf("sco[%d]=%v\n", i, sco[i])

	}
}
