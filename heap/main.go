package main

import "fmt"

type MaxHeap struct {
	array []int
}

// Add a value to the heap
func (h *MaxHeap) Push(val int) {
	h.array = append(h.array, val)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h MaxHeap) ParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h MaxHeap) LeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h MaxHeap) RightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MaxHeap) Swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

// Correct the heap upwards after value change at index
func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.array[h.ParentIndex(idx)] < h.array[idx] { // as long as the parent val < current val
		h.Swap(idx, h.ParentIndex(idx))
		idx = h.ParentIndex(idx)
	}
}

// Get the largest val and remove from the heap
// Return -1 if the heap is empty
func (h *MaxHeap) Pop() (largest int) {
	if len(h.array) > 0 {
		largest = h.array[0]

		lastIdx := len(h.array) - 1
		h.array[0] = h.array[lastIdx] // move the last to the first
		h.array = h.array[:lastIdx]   // chop the last
		h.MaxHeapifyDown(0)

		return largest
	}
	fmt.Println("cannot extract because array length is 0")
	return -1
}

// Correct the Heap working down the tree
// This usually happens after a Pop
func (h *MaxHeap) MaxHeapifyDown(idx int) {
	lastIdx := len(h.array) - 1
	l, r := h.LeftChildIndex(idx), h.RightChildIndex(idx)
	var childToCompare int

	// loop while node has at least one child
	for l <= lastIdx {
		// Figure which child is the larger (The tree fills left to right)
		if l == lastIdx || h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[idx] < h.array[childToCompare] { // child is larger
			h.Swap(idx, childToCompare)
			idx = childToCompare
			l, r = h.LeftChildIndex(idx), h.RightChildIndex(idx)
		} else { // no child is greater than the parent
			break
		}
	}
}

func main() {
	// Todo add via array
	heap := MaxHeap{}
	heap.Pop() // this should not panic
	heap.Push(44)
	heap.Push(50)
	heap.Push(24)
	fmt.Println(heap)
	heap.Push(70)
	fmt.Println(heap)
	heap.Push(23)
	fmt.Println(heap)
	heap.Push(24)
	fmt.Println(heap)
	heap.Pop()
	fmt.Println(heap)
	heap.Push(30)
	fmt.Println(heap)
	heap.Push(35)
	fmt.Println(heap)
	heap.Push(38)
	fmt.Println(heap)
	heap.Pop()
	fmt.Println(heap)
	heap.Pop()
	fmt.Println(heap)
	heap.Pop()
	fmt.Println(heap)
}
