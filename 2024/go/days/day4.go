package day

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	x int
	y int
}

func CheckBoundaries(matrix [][]string, row int, col int) bool {
	return row <= 0 || row >= len(matrix)-1 ||
		col <= 0 || col >= len(matrix[row])-1

}

func CheckDirections(row int, col int, matrix [][]string, letters [4]string, lIdx int, count *int, direction Direction) {

	if col >= 0 && col < len(matrix) {
		if row >= 0 && row < len(matrix[col]) {
			if matrix[col][row] == letters[lIdx] {
				if letters[lIdx] == "S" {
					*count += 1
					return
				} else {
					CheckDirections(row+direction.x, col+direction.y, matrix, letters, lIdx+1, count, direction)

				}
			}
		}
	}
}

func SolveTwo(matrix [][]string, row int, col int, count *int) {

	if CheckBoundaries(matrix, row, col) {
		return
	}

	patterns := []string{"MSMS", "SMSM", "MSSM", "SMMS"}
	tr := matrix[row-1][col+1]
	tl := matrix[row-1][col-1]
	br := matrix[row+1][col+1]
	bl := matrix[row+1][col-1]
	r := fmt.Sprintf("%s%s%s%s", tl, br, bl, tr)
	for _, p := range patterns {
		if p == r {
			*count++
		}
	}

}

func Four() (int64, int64) {
	data, err := os.ReadFile("input/04.txt")

	if err != nil {
		fmt.Print(err.Error())
		return 0, 0
	}

	content := data[:len(data)-1]

	textArray := strings.Split(string(content), "\n")

	matrix := make([][]string, len(textArray))
	for i, s := range textArray {
		matrix[i] = strings.Split(s, "")
	}

	directions := []Direction{
		{
			x: 1,
			y: 0,
		},
		{
			x: -1,
			y: 0,
		},
		{
			x: 0,
			y: 1,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: -1,
			y: -1,
		},
		{
			x: -1,
			y: 1,
		},
		{
			x: 1,
			y: -1,
		},
		{
			x: 1,
			y: 1,
		},
	}

	countXmas := 0
	countCrossMAS := 0
	letters := [4]string{"X", "M", "A", "S"}

	for col := range matrix {
		for row := range matrix {
			if matrix[col][row] == "X" {
				for i := range directions {
					cRow := row + directions[i].x
					cCol := col + directions[i].y
					if cCol >= 0 && cCol < len(matrix) {
						if cRow >= 0 && cRow < len(matrix[col]) {
							CheckDirections(cRow, cCol, matrix, letters, 1, &countXmas, directions[i])
						}
					}
				}
			}
		}
	}

	// PART 2
	for row := range matrix {
		for col := range matrix {
			if matrix[row][col] == "A" {
				SolveTwo(matrix, row, col, &countCrossMAS)
			}
		}
	}

	return int64(countXmas), int64(countCrossMAS)
}
