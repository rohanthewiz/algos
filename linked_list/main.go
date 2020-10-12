package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

// Make a new node the head of the list
func (l *LinkedList) Prepend(n *node) {
	n.next = l.head // the previous head becomes the next of the new node
	l.head = n      // make the new node the head
	l.length++
}

func (l LinkedList) Print() {
	currNode := l.head

	fmt.Print("H")
	for i := 0; i < l.length; i++ {
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

	for i := 0; i < l.length; i++ {
		if currNode.data == val {
			found = true
			if currNode == l.head { // special case of first node since we start off with a tie
				l.head = currNode.next // just make the next node the head
			} else {
				prevNode.next = currNode.next // join the prev with the next - current will be garbage collected
			}
			l.length--
			break
		}
		prevNode = currNode
		currNode = currNode.next
	}

	return
}

func main() {
	ll := LinkedList{}
	ll.Prepend(&node{data: 29})
	ll.Prepend(&node{data: 17})
	ll.Prepend(&node{data: 3})
	fmt.Printf("ll->%#v\n", ll)
	_ = ll.DeleteByValue(17)
	ll.Print()
	_ = ll.DeleteByValue(3)
	ll.Print()
	ll.Prepend(&node{data: 13})
	ll.Print()
	_ = ll.DeleteByValue(29)
	ll.Print()
	found := ll.DeleteByValue(23)
	if !found {
		fmt.Println("23 was not found")
	}
	ll.Print()

	found = ll.DeleteByValue(13)
	if !found {
		fmt.Println("13 was not found")
	}
	ll.Print()
}
