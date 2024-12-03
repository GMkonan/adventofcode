package day

import (
	"aoc/2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SafeCheck(ns []int) bool {
	var isIncrease bool
	var isDesc bool
	for i := range ns {

		if i+1 == len(ns) {
			break
		}

		n := ns[i]
		adjN := ns[i+1]

		if utils.Diff(n, adjN) > 3 || utils.Diff(n, adjN) < 1 {
			return false
		}

		if adjN > n {
			if i == 0 {
				isIncrease = true
				continue
			} else if isIncrease != true {
				return false
			}
		}

		if adjN < n {
			if i == 0 {
				isDesc = true
				continue
			} else if isDesc != true {
				return false
			}
		}

	}

	return true
}

func SafeCheckDamped(levelInts []int) bool {

	for i := range levelInts {
		fmt.Println(levelInts)
		if SafeCheck(utils.RemoveElementFromSlice(levelInts, i)) {
			return true
		}
	}

	return false

}

func Two() (int64, int64) {
	file, err := os.Open("input/02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	safeNums := 0
	safeNumsDamped := 0
	for scanner.Scan() {
		level := strings.Fields(scanner.Text())

		levelInts := make([]int, 0)

		for i := range level {
			n, _ := strconv.Atoi(level[i])
			levelInts = append(levelInts, n)
		}

		if SafeCheck(levelInts) {
			safeNums += 1
		}

		if SafeCheckDamped(levelInts) {
			safeNumsDamped += 1
		}

	}
	return int64(safeNums), int64(safeNumsDamped)
}
