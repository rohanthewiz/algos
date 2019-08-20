package array_ints

// Task: find indexes of item pairs across two arrays adding up to a given total

// For first element
// - scan the array looking for its compliment or
// 	storing the numbers and their indexes

// For the second onwards
// Simply check if the compliment of the current number is known to our map

// Remove from the input array writing over the last item in the array
//		- Shorten the array (len) by 1
