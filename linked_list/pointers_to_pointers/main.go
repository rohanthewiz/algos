package main

import (
	"fmt"
)

// Exercising Pointers in Linked Lists
func main() {
	// BRIEF INTRO TO POINTERS
	a := 2
	p := &a // p holds the address of the variable -- it points to the container of the value
	fmt.Println(*p)

	// struct declaration and instantiation
	o := struct {
		attr1 int
		attr2 string
	}{attr1: 2, attr2: "hello"}
	fmt.Println(o)

	// get a pointer to the attr2
	op := &o.attr2
	fmt.Println(*op)

	// POINTERS TO POINTERS IN LINKED LISTS
	// A Linked list is basically one thing pointing to the next

	type node struct {
		data int
		next *node
	}

	// create a linked list via literal struct
	ll := struct{ head *node }{head: &node{data: 1}}

	// add another node -- just for fun
	ll.head.next = &node{data: 2}

	// Print out the list
	// For loop dissection:
	// - init lp with a ptr to the head attribute of the linked list -- this is the starting container
	// - while the container has a pointer which doesn't point to nil
	// - at the end of each iteration get the address of the next container (next container is always a node in this example)
	// This works because lp is tracking the address of the container (head or node)
	// Remember that in Go, first level pointers don't have to be deferenced
	// - If lp points to head, `(*lp)` will return the value contained by head which is a pointer to the head node
	// We can however still get to this node's data without another deference `(*lp).data` => `pointer.data` -- thanks Go!
	// - If lp points to a node, `(*lp)` will deference to the node itself: `(*lp).data` => `node.data`.
	// This is also the case of `(*lp).next`. For head it is `pointer.next`, or a node it is `node.next`
	//
	// Why would we do this? If we needed to mutate the linked lists, like remove an element from any position,
	// or whatever else you need to do, you have pointer access to everything!
	// So this is a good way in Go, to iterate a linked list for mutation.

	fmt.Print("H")
	i := 0
	for lp := &ll.head; *lp != nil; lp = &((*lp).next) {
		fmt.Print("->", (*lp).data)
		// Proving the power of pointers - change out the list of nodes in the middle of the loop
		if i == 0 {
			ll.head = &node{data: 7, next: &node{data: 8, next: &node{data: 9}}}
		}
		i++
	}
	fmt.Println()

	// Okay now just print it the simple way
	fmt.Print("H")
	for nde := ll.head; nde != nil; nde = nde.next {
		fmt.Print("->", nde.data)
	}
	fmt.Println()
}
