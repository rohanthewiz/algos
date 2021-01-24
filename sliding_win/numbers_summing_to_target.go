package sliding_win

import (
	"strconv"
)

// Return index of numbers summing to a target
// Return empty array if not found
func TwoNumberSum(array []int, target int) []int {
	// So we don't have to do multiple passes, save the index
	// of each number whose complement is not found
	numToIdx := make(map[int]int, len(array))

	for i, n := range array {
		comp := target - n
		if idx, ok := numToIdx[comp]; ok {
			println("Index of complement is:", strconv.Itoa(idx))
			return []int{idx, i}
		}
		numToIdx[n] = i
	}

	return []int{}
}
