package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]string, 10)
	map1["5"] = "45"
	map1["3"] = "64"
	map1["9"] = "12"
	map1["10"] = "56"
	fmt.Println(map1)

}
