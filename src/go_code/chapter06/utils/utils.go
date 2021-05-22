package utils

import (
	"fmt"
)

var Name string
var Age int

func init() {
	fmt.Println("utils包的 init()...")
	Age = 100
	Name = "TOM"
}
