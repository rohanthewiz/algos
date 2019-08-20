package array_ints

import "fmt"

// Task: find indexes of item pairs across two arrays adding up to a given total
func UniqArrayInts(arr []int) []int {
	umap := make(map[int]bool) // unique map

	// For first element
	lnArr := len(arr)

	for i := 0; i < lnArr; i++ {
		if _, ok := umap[arr[i]]; ok {
			// remove the element from the array
			arr[i] = arr[lnArr-1] // copy from the end
			lnArr -= 1            // upper limit is one less
			i -= 1                // keep the same index on the next iteration
			continue
		}
		umap[arr[i]] = true // store the number in the map
	}

	arr = arr[:lnArr]
	fmt.Println(arr)
	fmt.Println("New length is", len(arr))
	return arr
}
