package main

import (
	day "aoc/2024/days"
	"fmt"
)

func main() {
	days := map[int]func() (int64, int64){
		1: day.One,
		2: day.Two,
		3: day.Three,
		4: day.Four,
		5: day.Five,
	}

	// make this loop later
	a, b := days[5]()
	fmt.Println(a)
	fmt.Println(b)
}
