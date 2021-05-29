package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

func Newperson(name string) *person {
	return &person{
		Name: name,
	}
}
func (p *person) Setage(age int) {
	if age > 0 && age < 150 {
		p.age = age

	} else {
		fmt.Println("年龄范围不正常")
	}

}
func (p *person) Getage() int {
	return p.age

}
func (p *person) Setsal(sal float64) {
	if sal > 3000.0 && sal < 15000.0 {
		p.sal = sal

	} else {
		fmt.Println("薪资范围不正常")
	}

}
func (p *person) Getsal() float64 {
	return p.sal
}
