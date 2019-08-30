package stringx

import "testing"

func TestStringReverse(t *testing.T) {
	// We are given the expected reversed words as an array
	expectedArr := []string{"tac", "esuom", "god", "esroh"}
	lnExpectedArr := len(expectedArr)

	arr := ReverseStrings()
	lnArr := len(arr) // do len() only once

	// Optimize search with map of words expected
	// We could have done this the other way putting
	// the values of the resulting array (arr) in the map
	wordsMap := make(map[string]bool, lnArr)

	found := 0

	for i := 0; i < lnArr; i++ {
		// On the first result check for a match in the expected
		// while saving values scanned in expected into a map for faster lookup O(1)
		if i == 0 {
			for j := 0; j < lnExpectedArr; j++ {
				if arr[i] == expectedArr[j] {
					found++
					continue
				}
				// We are given the expected strings as an array, but search for subsequent items against a map
				wordsMap[expectedArr[j]] = true
			}
		} else {
			if _, ok := wordsMap[arr[i]]; ok {
				found++
			}
		}
	}

	t.Log("Reversed strings found and valid:", found)
	if found != lnExpectedArr {
		t.Errorf("All reversed strings not found")
	}
}
