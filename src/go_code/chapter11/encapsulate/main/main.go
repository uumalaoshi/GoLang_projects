package main

import (
	"fmt"
	"github/src/go_code/chapter11/encapsulate/model"
)

func main() {
	p := model.Newperson("bot")
	p.Setage(18)
	p.Setsal(1000)
	fmt.Println(p)
	fmt.Println("name=", p.Name, "age=", p.Getage(), "sal=", p.Getsal())
}
