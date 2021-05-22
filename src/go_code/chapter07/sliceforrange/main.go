package main

import (
	"fmt"
)

func main() {
	var ran [5]int = [...]int{47, 243, 535, 34, 53}
	for i := 0; i < len(ran); i++ {
		fmt.Printf("ran[%v]=%v", i, ran[i])
	}

}
