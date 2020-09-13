package binary_search

// Synopsis
// Given: a sorted array
// Iteratively search for a target value
// - take the lower half of the array
// - if the mid ele is = target, we found the target element
// - else if the mid ele is > target, next look in the lower half of the array
// - else (the mid ele is < target) next look in the upper half of the array
// Returns index of found element or -1 if not found
func IterativeBinarySearchInt(arrInt []int, target int) (index int) {
	lnArr := len(arrInt)
	lwr := 0
	upr := lnArr - 1

	for lwr <= upr { // = here enables mid to coincide with the index of upr
		mid := (upr + lwr) / 2 // integer division, so mid will have an affinity for the lower index

		if arrInt[mid] == target { // yay! we found it
			return mid
		}

		if arrInt[mid] > target { // too high - go low
			upr = mid - 1
		} else { // too low - go high
			lwr = mid + 1
		}
	}

	return -1
}

// Synopsis
// Given: a sorted array
// Iteratively search for a target value
// - take the lower half of the array
// - if the mid ele is = target, we found the target element
// - else if the mid ele is > target, next look in the lower half of the array
// - else (the mid ele is < target) next look in the upper half of the array
// Returns index of found element or -1 if not found
func IterativeBinarySearchString(arrInt []string, target string) (index int) {
	lnArr := len(arrInt)
	lwr := 0
	upr := lnArr - 1

	for lwr <= upr { // = here enables mid to coincide with the index of upr
		mid := (upr + lwr) / 2 // integer division, so mid will have an affinity for the lower index

		if arrInt[mid] == target { // yay! we found it
			return mid
		}

		if arrInt[mid] > target { // too high - go low
			upr = mid - 1
		} else { // too low - go high
			lwr = mid + 1
		}
	}

	return -1
}
