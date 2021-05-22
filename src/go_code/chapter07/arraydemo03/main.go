package main

import "fmt"

func main() {
	var heros [3]string = [3]string{"松江", "无用", "林冲"}
	for i, v := range heros {
		fmt.Printf("i=%v v=%v\n", i, v)

	}
}
