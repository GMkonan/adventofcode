package main

import (
	day "aoc/2024/days"
	"fmt"
)

func main() {
	days := map[int]func() (int64, int64){
		1: day.One,
		2: day.Two,
		// "2A": day.TwoA,
		// "2B": day.TwoB,
	}

	oneA, oneB := days[1]()
	fmt.Println(oneA)
	fmt.Println(oneB)
	twoA, twoB := days[2]()
	fmt.Println(twoA)
	fmt.Println(twoB)
	// val1B := days["1B"]()
	// val2A := days["2A"]()
	// val2B := days["2B"]()
	// fmt.Println(val1A)
	// fmt.Println(val1B)
	// fmt.Println(val2A)
	// fmt.Println(val2B)
}
