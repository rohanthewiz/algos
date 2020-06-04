package sliding_win

import "fmt"

func LongestNonRepeatingSubstring(in string) (out string) {
	runes := []rune(in)
	lnRunes := len(runes)

	memo := make(map[rune]int, lnRunes)

	var currStart, longestStart, longestEnd int

	i := 0
	for ; i < lnRunes; i++ {
		if _, ok := memo[runes[i]]; ok {
			// Prev occur found, so check if preceding span is longest and reset
			fmt.Println("prev span", string(runes[currStart:i]))

			// Check longest
			if i-1-currStart > longestEnd-longestStart {
				longestStart = currStart
				longestEnd = i - 1
				fmt.Println("new longest", string(runes[longestStart:i]))
			}

			// Reset
			currStart = i
			memo = make(map[rune]int, lnRunes-i)
		}

		memo[runes[i]] = i // save the current
	}

	// Check if last span is longest
	if i-1-currStart > longestEnd-longestStart {
		longestStart = currStart
		longestEnd = i - 1
	}

	if lnRunes > 0 {
		out = string(runes[longestStart : longestEnd+1])
	}

	return
}
