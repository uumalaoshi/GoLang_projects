package main

import "fmt"

type cat1 struct {
	Name  string
	Age   int
	Color string
}

var cat2 cat1

func main() {
	fmt.Println("Cat1= ", cat2)
}
