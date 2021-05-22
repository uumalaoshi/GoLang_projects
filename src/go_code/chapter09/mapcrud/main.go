package main

import "fmt"

func main() {
	a := make(map[string]string, 3)
	a["no1"] = "jj"
	a["no2"] = "gg"
	a["no3"] = "ll"
	val, finres := a["no4"]
	if finres {
		fmt.Println("找到了这个val", val)

	} else {
		fmt.Println("没有找到这个val")

	}
	for k1, v1 := range a {
		fmt.Printf("k1=%v v1=%v\n", k1, v1)
		for k2, v2 := range v1 {
			fmt.Printf("k1=%v v1=%v\n", k2, v2)
		}
	}
}
