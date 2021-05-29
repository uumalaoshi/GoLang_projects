package main

import (
	"fmt"
)

type Acccount struct {
	AcccountNo string
	Pwd        string
	Balance    float64
}

//存款
func (account *Acccount) Depositr(money float64, Pwd string) {
	//检查输入的密码是否正确
	if Pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return

	}
	if money <= 0 {
		fmt.Println("您输入的金额不正确")
		return

	}
	account.Balance += money
	fmt.Println("存款成功")

}

//取款
func (account *Acccount) WithDraw(money float64, Pwd string) {
	//检查输入的密码是否正确
	if Pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return

	}
	if money <= 0 || money >= account.Balance {
		fmt.Println("您输入的金额不正确")
		return

	}
	account.Balance -= money
	fmt.Println("取款成功")

}
func (account *Acccount) Query(Pwd string) {
	//检查输入的密码是否正确
	if Pwd != account.Pwd {
		fmt.Println("您输入的密码不正确")
		return

	}
	fmt.Printf("您的账号为%v 余额为%v \n", account.AcccountNo, account.Balance)

}
func main() {
	account := Acccount{
		AcccountNo: "23245",
		Pwd:        "33333",
		Balance:    100.0,
	}
	account.Query("33333")
	account.Depositr(100.0, "33333")
	account.Query("33333")
	account.WithDraw(120.0, "33333")
	account.Query("33333")

}
