package main

import (
	"fmt"
)

func binaryFind(arr *[6]int, leftIndex int, rightIndex int, findVar int) {
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return

	}
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVar {
		binaryFind(arr, leftIndex, middle-1, findVar)

	} else if (*arr)[middle] < findVar {
		binaryFind(arr, middle+1, rightIndex, findVar)

	} else {
		fmt.Printf("找到了 下标为%v\n", middle)

	}

}
func main() {
	arr := [6]int{2, 7, 24, 55, 67, 86}
	binaryFind(&arr, 0, len(arr)-1, 67)

}
