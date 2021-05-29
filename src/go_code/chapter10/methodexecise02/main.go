package main

import "fmt"

type methodUtils struct {
}

func (k methodUtils) print(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}

}
func (p methodUtils) jisuan(len float64, width float64) {
	res := len * width
	fmt.Println("面积为=", res)
}
func (y methodUtils) panduan(i int) {
	if i%2 == 0 {
		fmt.Println("输入的书为偶数")

	} else {
		fmt.Println("输入的数为奇数")
	}

}
func main() {
	var k methodUtils
	k.print(2, 4)
	var p methodUtils
	p.jisuan(1.3, 2.4)
	var y methodUtils
	y.panduan(5)
}
