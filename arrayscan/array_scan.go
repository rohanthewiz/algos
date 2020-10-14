package arrayscan

import "math"

// Given two arrays of differing length
// find the sum of pairs that is closest to target
// Algo: Normally this would be a quadratic operation (n*n)
// however for exact matches, by storing the values to position (index) of the second array on the first pass
// we will likely not scan that array again.
// If there are no exact matches we can only know that by scanning n * n then return the closest found
func PairsSumNearToTarget(arr1, arr2 []int, target int) (arr1Idx, arr2Idx int) {
	valueToPosition := make(map[int]int, len(arr2))

	closest, ci, cj := 999, 0, 0

	for i := 0; i < len(arr1); i++ {
		// Speed up: Is the current element's exact complement wrt target in the map?
		comp := target - arr1[i]
		// If comp is in map, return it
		if idx, ok := valueToPosition[comp]; ok {
			return i, idx
		}

		// Scan the second array
		for j := 0; j < len(arr2); j++ {
			// Core check
			if arr1[i]+arr2[j] == target {
				return i, j
			}
			// Is it the closest?
			if cl := math.Abs(float64(target - arr1[i] - arr2[j])); cl < float64(closest) {
				closest, ci, cj = int(cl), i, j
			}

			// On first scan store arr2 values -> arr2 index in the map
			if i == 0 {
				valueToPosition[arr2[j]] = j
			}
		}
	}

	// We didn't find an exact match so just return the closest
	println("closest indices", ci, cj, "diff:", closest)
	return ci, cj
}
