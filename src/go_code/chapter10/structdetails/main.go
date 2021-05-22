package main

import (
	"encoding/json"
	"fmt"
)

type mon struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	mon := mon{"牛魔王", 500, "芭蕉扇"}
	monstr, err := json.Marshal(mon)
	if err != nil {
		fmt.Println("err = 处理错误", err)
	}
	fmt.Println("monstr=", string(monstr))

}
