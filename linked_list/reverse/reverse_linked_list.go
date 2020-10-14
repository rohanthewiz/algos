package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

// Remove the node after n, returning the value of the removed node
// Return -1 in case of error
func (l *LinkedList) RemoveAfter(n *node) (val int) {
	if n == nil || n.next == nil {
		return -1
	}
	val = n.next.data // sav

	n.next = n.next.next // remove
	return
}

// Insert a value at the top of the linked list
// -1 is an error indicator
func (l *LinkedList) InsertAtHead(val int) {
	l.head = &node{next: l.head, data: val}
}

func (l *LinkedList) InsertValsAtHead(vals []int) {
	for _, val := range vals {
		l.InsertAtHead(val)
	}
}

func (l LinkedList) Print() {
	curr := l.head
	fmt.Print("H")
	for curr != nil {
		fmt.Print("->", curr.data)
		curr = curr.next
	}
	fmt.Println("")
}

// Algo: move successive values into head starting at head.next
func (l *LinkedList) Reverse() {
	if l.head == nil {
		return
	}

	curr := l.head
	for curr.next != nil {
		val := l.RemoveAfter(curr)
		if val == -1 {
			return
		}
		l.InsertAtHead(val)
	}
}

func main() {
	ll := LinkedList{}
	ll.Print()
	ll.InsertAtHead(2)
	ll.InsertValsAtHead([]int{4, 13, 11, 5, 8})
	ll.Print()
	ll.Reverse()
	ll.Print()
}
