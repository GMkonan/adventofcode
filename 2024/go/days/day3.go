package day

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// you could for sure make this into a single function that
// receives the info if it should check for do's/don'ts or not
// but I didn't know what the second challenge would be before hand hehe

func FindMulInstructs(text string, multi *int) int {
	rgmul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	rgn := regexp.MustCompile(`\d+`)

	mulmatches := rgmul.FindAllString(text, -1)

	for _, n := range mulmatches {
		numbers := rgn.FindAllString(n, -1)
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])

		*multi += n1 * n2
	}
	return *multi
}

func FindMulInstructsWithDoDont(text string, multi *int, enable *bool) int {

	rgmulAndDo := regexp.MustCompile(`don't|do|mul\((\d+),(\d+)\)`)
	rgn := regexp.MustCompile(`\d+`)

	mulAndDoMatches := rgmulAndDo.FindAllString(text, -1)

	for _, n := range mulAndDoMatches {

		if n == "do" {
			*enable = true
		} else if n == "don't" {
			*enable = false
		} else if *enable {

			numbers := rgn.FindAllString(n, -1)
			n1, _ := strconv.Atoi(numbers[0])
			n2, _ := strconv.Atoi(numbers[1])

			*multi += n1 * n2
		}
	}

	return *multi
}

func Three() (int64, int64) {
	file, err := os.Open("input/03.txt")
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}

	scanner := bufio.NewScanner(file)

	multi := 0
	multi2 := 0
	var enable bool = true

	for scanner.Scan() {

		multi = FindMulInstructs(scanner.Text(), &multi)

		multi2 = FindMulInstructsWithDoDont(scanner.Text(), &multi2, &enable)

	}

	return int64(multi), int64(multi2)
}
