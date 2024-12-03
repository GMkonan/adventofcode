package day

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc/2024/utils"
)

func CalculatTotalDistance(list1 []string, list2 []string) int {
	var sumOfDiffs int
	for i := range list1 {
		n1, _ := strconv.Atoi(list1[i])
		n2, _ := strconv.Atoi(list2[i])
		sumOfDiffs += utils.Diff(n1, n2)

	}
	return sumOfDiffs
}

func CalculateSimilarityScore(list1 []string, list2 []string) int {
	var countSimilar int
	var sumAll int

	for i := range list1 {

		n1, _ := strconv.Atoi(list1[i])

		for l := range list2 {
			n2, _ := strconv.Atoi(list2[l])
			if n1 == n2 {
				countSimilar += 1
			}
		}

		sumAll += n1 * countSimilar
		countSimilar = 0

	}
	return sumAll
}

func One() (int64, int64) {
	file, err := os.Open("input/01.txt")
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}

	scanner := bufio.NewScanner(file)

	var list1 []string
	var list2 []string

	for scanner.Scan() {
		list1 = append(list1, strings.Fields(scanner.Text())[0])
		sort.Sort(sort.StringSlice(list1))

		list2 = append(list2, strings.Fields(scanner.Text())[1])
		sort.Sort(sort.StringSlice(list2))
	}

	// part 1 (A)
	sumOfDiffs := CalculatTotalDistance(list1, list2)
	// part 2 (B)
	sumAll := CalculateSimilarityScore(list1, list2)

	return int64(sumOfDiffs), int64(sumAll)
}
