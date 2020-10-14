package main

import "fmt"

// Given a linked list. Group even indexed items at the start of the linked linked list followed by odd indices

// Solution algorithm: (we will move even indexed values from a lower pointer to an upper destination)
// Upper pointer points to where I want to land nodes. We will advance this by one after the copy.
// Lower pointer points to the node whose next node we want to move.
// After the move we only have to advance one, because the next node was removed

// 43->15->13->20->6->17->18->22->33->11->Null
// u->43, l->15 -- l.next is 13, so move it to after upr (43), then advance upr and lwr
// 43->13->15->20->6->17->18->22->33->11->Null
// u ->13, l->20 -- l.next is 6, move it to after upr (13), then advance upr and lwr
// 43->13->6->15->20->17->18->22->33->11->Null
// u -> 6, l->17 -- rinse and rpt
// 43->13->6->18->15->20->17->22->33->11->Null
// 43->13->6->18->33->15->20->17->22->11->Null

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

// Move items with even indices to the start of the Linked List and those with odd indices to the end  maintaining order
func (l *LinkedList) GroupEvenOdd() {
	if l.head == nil || l.head.next == nil {
		return
	}

	upr := l.head      // dest
	lwr := l.head.next // src

	for lwr.next != nil {
		// remove and insert  after upr
		v := l.RemoveNext(lwr)
		l.InsertAfter(upr, v)

		upr = upr.next
		lwr = lwr.next
	}
}

// Add at head
func (l *LinkedList) InsertAtHead(vals []int) {
	for _, val := range vals {
		l.head = &node{next: l.head, data: val}
	}
}

func (l LinkedList) InsertAfter(curr *node, val int) {
	newNode := &node{next: curr.next, data: val}
	curr.next = newNode
}

// Return -1 if not found
func (l *LinkedList) RemoveNext(curr *node) (val int) {
	if curr.next == nil {
		return -1
	}
	val = curr.next.data

	curr.next = curr.next.next // skip the node after current
	return
}

func (l LinkedList) Print() {
	currNode := l.head
	fmt.Print("H")
	for currNode != nil {
		fmt.Print("->", currNode.data)
		currNode = currNode.next
	}
	fmt.Println("")
}

func main() {
	ll := LinkedList{}
	ll.Print() // should not panic
	ll.InsertAtHead([]int{3, 5, 8, 9, 7, 12, 22, 19, 39})
	ll.Print() // should not panic
	val := ll.RemoveNext(ll.head)
	if val != -1 {
		fmt.Println("Removed", val)
	}
	ll.Print() // should not panic
	ll.GroupEvenOdd()
	ll.Print()
}
