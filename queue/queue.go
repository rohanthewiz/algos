package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type Queue struct {
	head *node
	tail *node
}

// Insert the node at the top of the Linked List
func (q *Queue) Enqueue(newNode *node) {
	prevHead := q.head

	// connect up new node
	newNode.next = prevHead // right side
	q.head = newNode        // left side
	// Let's leave newNode.prev to nil when at the top of the linked list
	// - it's one way to check we are at the top

	// connect up previous head node
	if prevHead != nil {
		prevHead.prev = newNode
	}

	// connect up tail if queue was empty
	if prevHead == nil {
		q.tail = newNode
	}
}

func (q *Queue) Print() {
	currNode := q.head
	fmt.Print("H")

	for currNode != nil {
		if currNode.prev == nil {
			fmt.Printf("->%d", currNode.data)
		} else {
			fmt.Printf("<->%d", currNode.data)
		}
		currNode = currNode.next
	}
	fmt.Print("<-T")
	fmt.Println("")
}

func (q *Queue) Dequeue() (lastNode *node) {
	if q.tail == nil {
		return
	}

	lastNode = q.tail

	// Rewire tail and penultimate node
	q.tail = lastNode.prev // even if it is nil as in the case of one item in queue
	if q.tail != nil {
		q.tail.next = nil // make sure to terminate
	} else {
		// we are empty -- make sure the head knows it
		q.head = nil
	}

	return
}

func main() {
	que := Queue{}
	que.Print()
	que.Enqueue(&node{data: 17})
	que.Print()
	que.Enqueue(&node{data: 10})
	que.Enqueue(&node{data: 12})
	que.Enqueue(&node{data: 5})
	que.Print()

	for i := 0; i < 5; i++ {
		n := que.Dequeue()
		if n != nil {
			fmt.Println("Dequeued:", n.data)
		}
		que.Print()
	}
}
