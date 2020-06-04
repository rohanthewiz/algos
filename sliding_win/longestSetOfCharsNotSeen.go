package sliding_win

import (
	"fmt"
)

func longestSetOfCharsNotAlreadySeen(in string) (out string) {
	runes := []rune(in)
	lnRunes := len(runes)

	var start, longestStart, longestEnd int

	memo := make(map[rune]int, lnRunes)

	i := 0
	for ; i < lnRunes; i++ {
		// fmt.Println("i", i)

		// Letter already seen?
		if val, ok := memo[runes[i]]; ok {
			if i-start > longestEnd-longestStart {
				fmt.Println("Prev pos:", val, "i:", i)
				longestStart = start
				longestEnd = i - 1
			}
			start = i + 1 // slide forward
			// memo = make(map[rune]int, lnRunes - i)
		}

		memo[runes[i]] = i
	}

	// Check the last run
	if i-start > longestEnd-longestStart {
		longestStart = start
		longestEnd = i - 1
	}

	if longestStart > lnRunes-1 {
		longestStart = lnRunes - 1
	}

	if lnRunes > 0 {
		out = string(runes[longestStart : longestEnd+1])
	}
	return
}
