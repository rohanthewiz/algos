package arrayscan


// Given two arrays of differing length
// find the sum of pairs that is closest to target

func PairsSumNearToTarget(arr1, arr2 []int, target int) (arr1Idx, arr2Idx int) {
	valueToPositions := make(map[int]int, len(arr2))

	for i := 0; i < len(arr1); i++ {
		// Is the current elements complement in the map?
		comp := target - arr1[i] // TODO - closest means we want some tolerance here
		// If comp is in map, return it
		if idx, ok := valueToPositions[comp]; ok {
			return i, idx
		}

		// Scan the second array
		for j := 0; j < len(arr2); j++ {
			// Core check
			if arr1[i] + arr2[j] == target {
				return i, j
			}

			// On first scan store arr2 values -> arr2 index in the map
			if i == 0 {
				valueToPositions[arr2[j]] = j
			}
		}
	}

	return
}
