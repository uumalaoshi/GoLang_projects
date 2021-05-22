package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	fbn4 := make([]uint64, n)
	fbn4[0] = 1
	fbn4[1] = 1
	for i := 2; i < n; i++ {
		fbn4[i] = fbn4[i-1] + fbn4[i-2]
	}
	return fbn4
}
func main() {
	res := fbn(30)
	fmt.Println("res", res)

}
