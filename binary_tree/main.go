package main

import "fmt"

type node struct {
	key         int
	left, right *node
}

// We will insert values greater than the node's key to the right
// Less than or equal to will go left
func (n *node) Insert(key int) {
	if key > n.key { // greater goes right
		if n.right == nil {
			n.right = &node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else { // less than or equal to goes left
		if n.left == nil {
			n.left = &node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

func (n node) Print() {
	l, r := -1, -1
	if n.left != nil {
		l = n.left.key
	}
	if n.right != nil {
		r = n.right.key
	}
	fmt.Printf("%d->(%d, %d) ", n.key, l, r)
	if n.left != nil {
		n.left.Print()
	}
	if n.right != nil {
		n.right.Print()
	}
}

func (n node) Search(key int) (foundNode *node) {
	if n.key == key {
		return &n
	}
	if key > n.key { // greater will go right
		if n.right != nil { // try the right node
			if ret := n.right.Search(key); ret != nil {
				return ret
			}
		}
	} else { // go left
		if n.left != nil { // try the left node
			if ret := n.left.Search(key); ret != nil {
				return ret
			}
		}
	}
	return
}

func main() {
	tree := node{key: 100}
	tree.Insert(102)
	tree.Insert(98)
	tree.Insert(68)
	tree.Insert(75)
	tree.Print()
	println()

	ret := tree.Search(68)
	if ret != nil {
		ret.Print()
		println()
	}
}
