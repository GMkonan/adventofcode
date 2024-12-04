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
	}

	// oneA, oneB := days[1]()
	// fmt.Println(oneA)
	// fmt.Println(oneB)
	// twoA, twoB := days[2]()
	// fmt.Println(twoA)
	// fmt.Println(twoB)

	threeA, threeB := days[3]()
	fmt.Println(threeA)
	fmt.Println(threeB)
}
