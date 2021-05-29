package main

import (
	"fmt"
	"github/src/go_code/chapter11/encapexercise/model"
)

func main() {
	account := model.Newaccount("jd134344", "444444", 90)
	if account != nil {
		fmt.Println("创建成功", account)
	} else {
		fmt.Println("创建失败")
	}

}
