package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (ih intHeap) Len() int {
	return len(ih)
}

func (ih intHeap) Less(i, j int) bool {
	return ih[i] < ih[j]
}

func (ih intHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *intHeap) Push(heapintf interface{}) {
	*ih = append(*ih, heapintf.(int))
}

func (ih *intHeap) Pop() interface{} {
	prev := *ih // save

	lastIdx := len(prev) - 1
	lastVal := prev[lastIdx]

	*ih = prev[0:lastIdx] // truncate

	return lastVal
}

func main() {
	iHeap := &intHeap{1, 4, 5}

	heap.Init(iHeap)
	heap.Push(iHeap, 2)
	fmt.Printf("The heap is: %#v\n", iHeap)

	for iHeap.Len() > 0 {
		heap.Pop(iHeap)
		fmt.Printf("The heap is: %#v\n", iHeap)
	}
}
