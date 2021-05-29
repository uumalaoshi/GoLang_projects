package main

import (
	"fmt"
	"github/src/go_code/chapter10/factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 3.8)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "score", stu.Getscore())
}
