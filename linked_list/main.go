package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

// Make a new node the head of the list
func (l *LinkedList) Prepend(n *node) {
	n.next = l.head // the previous head becomes the next of the new node
	l.head = n      // make the new node the head
}

func (l LinkedList) Print() {
	currNode := l.head

	fmt.Print("H")
	for currNode != nil {
		fmt.Printf("-> %d", currNode.data)
		currNode = currNode.next
	}
	fmt.Println("")
}

// Search for node by value and delete
// Return error if not found
func (l *LinkedList) DeleteByValue(val int) (found bool) {
	// We don't know what's ahead so start both pointers at the same point
	currNode := l.head
	prevNode := l.head // we need to track this since the current node has no knowledge of the prev node, only the next

	for currNode != nil {
		if currNode.data == val {
			found = true
			if currNode == l.head { // special case of first node since we start off with a tie
				l.head = currNode.next // just make the next node the head
			} else {
				prevNode.next = currNode.next // join the prev with the next - current will be garbage collected
			}
			break
		}
		prevNode = currNode
		currNode = currNode.next
	}

	return
}

// Search for node by value and delete
// Return error if not found
func (l *LinkedList) DeleteByValue2(val int) (found bool) {
	if l.head == nil { // empty linked list
		return
	}
	if l.head.data == val { // val we are deleting is the head
		l.head = l.head.next
		return true
	}

	for currNode := l.head; currNode.next != nil; currNode = currNode.next {
		if currNode.next.data == val {
			found = true
			currNode.next = currNode.next.next // remove the next
			return
		}
	}
	return
}

// The address of the thing holding the ptr to the next thing is the previous thing
// In the case of head it is the address of the linked list attribute head
// In the case of any other node, it is the address of the "next" attribute holding the ptr to the next node, i.e. the [prev] node
func (l *LinkedList) DeleteByValue3(val int) (found bool) {
	for holderPtr := &l.head; *holderPtr != nil; holderPtr = &(*holderPtr).next {
		if (*holderPtr).data == val {
			*holderPtr = (*holderPtr).next // delete
			found = true
			return
		}
	}
	return
}

func main() {
	ll := LinkedList{}
	ll.Prepend(&node{data: 29})
	ll.Prepend(&node{data: 17})
	ll.Prepend(&node{data: 3})
	ll.Print()
	_ = ll.DeleteByValue3(17)
	_ = ll.DeleteByValue3(3)
	ll.Print()
	ll.Prepend(&node{data: 13})
	ll.Print()
	_ = ll.DeleteByValue3(29)
	ll.Print()
	_ = ll.DeleteByValue3(13) // deleting the only node should not panic
	ll.Print()
	found := ll.DeleteByValue(10) // deleting a non existent value should not panic
	if !found {
		fmt.Println("10 was not found")
	}
}
