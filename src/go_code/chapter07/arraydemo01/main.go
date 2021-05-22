package main

import "fmt"

func main() {
	var hens [6]float64
	hens[0] = 2.2
	hens[1] = 3.3
	hens[2] = 4.4
	hens[3] = 5.5
	hens[4] = 6.6
	hens[5] = 7.7
	totlweight := 0.0
	for i := 0; i < len(hens); i++ {
		totlweight += hens[i]

	}
	totlweight2 := fmt.Sprintf("%.2f", totlweight/float64(len(hens)))

	fmt.Printf("totlweight=%v totlweight2=%v", totlweight, totlweight2)

}
