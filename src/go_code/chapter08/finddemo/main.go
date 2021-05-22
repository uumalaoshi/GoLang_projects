package main

import (
	"fmt"
)

func main() {
	names := [4]string{"严青峰", "王三", "王菲菲", "余杰"}
	var herName string = ""
	fmt.Println("请输入需要查找的名字")
	fmt.Scanln(&herName)
	for i := 0; i < len(names); i++ {
		if herName == names[i] {
			fmt.Printf("找到%v , 下标%v \n", herName, i)
			break

		} else if i == (len(names) - 1) {
			fmt.Println("没有找到")

		}

	}
}
