package recursive_search

import (
	"fmt"
)

//func main() {
//	nums := []int{-1, 0, 3, 5}
//	target := 0
//	fmt.Println(rSearch(nums, 0, len(nums), target))
//}

// Recursive concurrent search
// nums - array of ints to search, lIdx - lower index, uIdx - upper index,
// result - index of found target else -1
func rSearch(nums []int, lIdx, uIdx, target int) (result int) {
	fmt.Println(lIdx, ",", uIdx)

	// Safeties
	if lIdx < 0 || lIdx >= uIdx {
		fmt.Println("Invalid indices", lIdx, uIdx)
		return -1
	}
	// One item left - is it the target?
	if uIdx - lIdx == 1 {
		if nums[lIdx] == target {
			return lIdx
		}
		return -1
	}

	resultChan := make(chan int, 2)
	mid := lIdx + (uIdx - lIdx)/2

	go func(l, u int) {
		resultChan <- rSearch(nums, l, u, target)
	}(lIdx, mid)

	go func(l, u int) {
		resultChan <- rSearch(nums, l, u, target)
	}(mid, uIdx)

	for i := 0; i < 2; i++ {
		select {
		case r := <-resultChan :
			if r != -1 {
				return r
			}
		}
	}
	return -1
}
