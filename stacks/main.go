package main

import "fmt"

type Stack struct {
	items []int
}

// Instantiate the array with a cap of 8
func (s *Stack) Init() {
	s.items = make([]int, 0, 8)
}

// Push
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s Stack) Print() {
	for i, item := range s.items {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Print(item)
	}
	fmt.Println("")
}

// Pop
func (s *Stack) Pop() (item int) {
	lastIdx := len(s.items) - 1

	item = s.items[lastIdx] // we'll return the last item

	s.items = s.items[:lastIdx] // chop the last element of the slice
	return
}

func main() {
	stk := Stack{}
	stk.Init()

	stk.Push(3)
	stk.Push(5)
	stk.Push(7)
	stk.Push(1)
	stk.Print()
	fmt.Println(stk.Pop())
	stk.Print()
	fmt.Println(stk.Pop())
	stk.Print()
}
