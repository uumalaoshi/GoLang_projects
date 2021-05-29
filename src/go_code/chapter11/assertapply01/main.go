package main

import "fmt"

func TypeJuder(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool 值是%v\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64 值是%v\n", i, x)
		case int, int64:
			fmt.Printf("param #%d is a int , int 64 值是%v\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil 值是%v\n", i, x)
		case string:
			fmt.Printf("param #%d is a string 值是%v\n", i, x)
		default:
			fmt.Printf("param #%d  type is a unknow 值是%v\n", i, x)
		}
	}
}

func main() {
	var n1 int = 65
	var n2 bool = true
	var n3 string = "tom"
	n4 := "北京"
	var n5 float32 = 6.5
	TypeJuder(n1, n2, n3, n4, n5)

}
