package main

import (
	"fmt"
)

func bubb(arr *[5]int) {
	fmt.Println("排序前arrs", arr)
	temp := 0
	for j := 0; j < 4; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}

	}
	fmt.Println("第一次排序后的arr", arr)
	for j := 0; j < 3; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}

	}
	fmt.Println("第二次排序后的arr", arr)
	for j := 0; j < 2; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}

	}
	fmt.Println("第三次排序后的arr", arr)
	for j := 0; j < 1; j++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}

	}
	fmt.Println("第四次排序后的arr", arr)
}
func main() {
	var arr [5]int = [5]int{147, 34, 123, 42, 5834}
	bubb(&arr)

}
