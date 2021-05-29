package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("Name=%v Age=%v Score=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试。。。")
}

type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试。。。")
}
func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "TOM"
	pupil.Student.Age = 6
	pupil.testing()
	pupil.Student.SetScore(80)
	pupil.Student.ShowInfo()
	graduate := &Graduate{}
	graduate.Student.Name = "TOM"
	graduate.Student.Age = 26
	graduate.testing()
	graduate.Student.SetScore(70)
	graduate.Student.ShowInfo()

}
