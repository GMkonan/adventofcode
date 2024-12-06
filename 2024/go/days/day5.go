package day

// https://www.reddit.com/r/adventofcode/comments/1h7mm3w/2024_day_05_part_2_how_nice_is_the_input_a_binary/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	left  string
	right string
}

// check presence
func checkRule(rule string, pages []string) int {
	for i, page := range pages {
		if rule == page {
			return i
		}
	}
	return -1
}

func check(rules []Rule, pages []string) bool {
	for _, rule := range rules {

		fmt.Println(pages)
		fmt.Println(rule.left, rule.right)
		leftIdx := checkRule(rule.left, pages)
		rightIdx := checkRule(rule.right, pages)
		if leftIdx == -1 || rightIdx == -1 {
			continue
		}

		// check if its out of order
		if leftIdx > rightIdx {
			return false

		}
	}
	return true
}

func Five() (int64, int64) {
	data, err := os.ReadFile("input/05.txt")

	if err != nil {
		fmt.Print(err.Error())
		return 0, 0
	}

	content := data[:len(data)-1]

	textArray := strings.Split(string(content), "\n\n")
	orderingRules := strings.Split(textArray[0], "\n")
	listOfPages := strings.Split(textArray[1], "\n")

	var rules []Rule

	for i := range orderingRules {
		n := strings.Split(string(orderingRules[i]), "|")
		rules = append(rules, Rule{left: n[0], right: n[1]})
	}

	var sum int
	for _, p := range listOfPages {
		rowOfPages := strings.Split(p, ",")
		if check(rules, rowOfPages) {
			mn, _ := strconv.Atoi(rowOfPages[len(rowOfPages)/2])
			fmt.Println(mn)
			sum += mn
		}

	}
	fmt.Println(sum)

	return int64(sum), 0
}
