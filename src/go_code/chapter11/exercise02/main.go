package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

func (M *Monkey) climbing() {
	fmt.Println(M.Name, "生来会爬树")
}

type littleMonkey struct {
	Monkey
}

func main() {
	monkey := littleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
}
