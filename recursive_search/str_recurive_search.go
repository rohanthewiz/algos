package recursive_search

// Returns index of found string else -1
func StartStrRecursiveSearch(arr []string, target string) (result int) {
	return strRecurSearch(arr, target, 0, len(arr))
}

// Recursive concurrent search
// arr - array of string to search, lIdx - lower index, uIdx - upper index,
// result - index of found target else -1
func strRecurSearch(arr []string, target string, lIdx, uIdx int) (result int) {
	const batchSize = 512

	//log.Println(lIdx, ",", uIdx)

	// Safeties
	if lIdx < 0 || lIdx >= uIdx {
		//log.Println("Invalid indices", lIdx, uIdx)
		return -1
	}

	// If at or less than the batchSize then let the machine scan
	if uIdx-lIdx <= batchSize {
		for i := lIdx; i < uIdx; i++ {
			if arr[i] == target {
				//log.Println("Found in this batch (index):", lIdx, ",", uIdx)
				return i
			}
		}
		return -1
	}

	// TODO - Alleviate perf hit from memory alloc here - use some kind of channel pool
	resultChan := make(chan int, 2)
	mid := lIdx + (uIdx-lIdx)/2

	go func(l, u int) {
		resultChan <- strRecurSearch(arr, target, l, u)
	}(lIdx, mid)

	go func(l, u int) {
		resultChan <- strRecurSearch(arr, target, l, u)
	}(mid, uIdx)

	for i := 0; i < 2; i++ {
		select {
		case r := <-resultChan:
			if r != -1 { // found!
				return r
			}
		}
	}
	return -1
}
