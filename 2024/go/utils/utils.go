package utils

func Diff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func RemoveElementFromSlice(slice []int, s int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice[:s]...)
	return append(newSlice, slice[s+1:]...)
}
