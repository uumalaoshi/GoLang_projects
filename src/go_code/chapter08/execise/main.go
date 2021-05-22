package main

import (
	"fmt"
)

func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d个班级第%d个的学生的成绩\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])

		}

	}
	tdad := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]

		}
		tdad += sum
		fmt.Printf("第%d的班级的总分为%v,平均分为%v\n", i+1, sum, sum/float64(len(scores[i])))

	}
	fmt.Printf("所有班级的总分为%v 所以班级的平均分为%v\n", tdad, tdad/15)
}
