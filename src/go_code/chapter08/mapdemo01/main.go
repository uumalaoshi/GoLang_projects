package main

import "fmt"

func main() {
	a := make(map[string]map[string]string)
	a["no1"] = make(map[string]string, 4)
	a["no1"]["name"] = "tom"
	a["no1"]["sex"] = "男"
	a["no1"]["age"] = "18"
	a["no1"]["dad"] = "make"

	fmt.Println(a)

}
