package main

import (
	"errors"
	"fmt"
	_"strconv"
	_ "time"
)

/*func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)

	}
}

func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行demo03耗费的时间为%v秒\n", end-start)
}*/
func read(name string) (err error) {
	if name == "yan.ip" {

		return nil

	} else {
		return errors.New("文件提取错误")
	}
}

func test02() {
	err := read("yan.ipg")
	if err != nil {
		panic(err)

	}
	fmt.Println("test02的代码")

}
func main() {
	test02()
	fmt.Println("main()//")

}
