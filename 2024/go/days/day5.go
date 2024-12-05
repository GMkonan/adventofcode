package day

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
		// 75 | 23
		// find left

		fmt.Println(pages)
		fmt.Println(rule.left, rule.right)
		leftIdx := checkRule(rule.left, pages)
		rightIdx := checkRule(rule.right, pages)
		if leftIdx == -1 || rightIdx == -1 {
			continue
		}

		// if idx on the left is bigger

		// check if its out of order
		if leftIdx > rightIdx {
			return false

		}
		// find right
		// check if they are correct

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

	fmt.Println(rules)
	var sum int
	for _, p := range listOfPages {
		rowOfPages := strings.Split(p, ",")
		// [75, 23, 56]
		if check(rules, rowOfPages) {
			mn, _ := strconv.Atoi(rowOfPages[len(rowOfPages)/2])
			fmt.Println(mn)
			sum += mn
		} else {
			// reorder them
			for _, n := range rules {
				checkRule(n., rowOfPages)

			}
		}

	}
	fmt.Println(sum)
	// try the inverrse logic
	// apply rule to update line
	// not line to the rule

	// define sum = 0

	return 0, 0
}
