/*
MATEJ OSTADAL
KMI UPOL
2022
*/

package binomialheaps

import (
	"fmt"
)

// node structure
type BinomialNode struct {
	Key     int
	child   *BinomialNode // pointer to the most left child
	sibling *BinomialNode // right sibling, in case of root node, pointer to the next root
	degree  int
	parent  *BinomialNode // pointer to the parent node, nil for root nodes
}

// heap structure (heap is a linked list of roots, this is just a pointer to the first one)
type BinomialHeap struct {
	head *BinomialNode // pointer to the most left tree
}

/*
EXPLANATION OF PRINTING

Heap tree roots are left without spacing from the left.
Each tab (spacing from the left) means we went to a greater depth (to nodes children).
Siblings are always on the same level of spacing.

JUST FOR DEMONSTRATION:

ROOT R1
	R1.CHILD1
		R1.CHILD1.CHILD
	R1.CHILD2
		R1.CHILD2.CHILD
ROOT R2
	R2.CHILD1
		R2.CHILD1.CHILD
	R2.CHILD2

	...
*/

// inserts new node with a given key
func (h *BinomialHeap) Insert(value int) {

	x := BinomialNode{Key: value} // new node with the given value
	h.InsertNode(&x)
}

// inserts node to the heap
func (h *BinomialHeap) InsertNode(x *BinomialNode) {

	h1 := MakeBinoHeap(x)
	h2 := h

	h.head = union(h1, h2) // the heap is set to its union with h1
}

// changes the priority of the node bubbles to follow the min-heap rule
func (h *BinomialHeap) DecreaseKey(x *BinomialNode, k int) {

	if k < x.Key && k > 0 {
		x.Key = k
		y := x
		z := y.parent

		for z != nil && y.Key < z.Key {

			y.Key, z.Key = z.Key, y.Key

			y = z
			z = y.parent
		}
	}
}

// removes and returns a node with a minimal key
func (h *BinomialHeap) ExtractMin() *BinomialNode {

	if h.head == nil {
		return nil
	}

	x := h.removeMinList()

	children_heap := reverseChildren(x.child)

	h.head = union(h, children_heap)

	return x
}

// removes the node from the list from the heap
func (h *BinomialHeap) removeMinList() *BinomialNode {

	x, x_left := h.MinimumNode()

	if x == h.head {
		h.head = x.sibling
	} else {
		x_left.sibling = x.sibling
	}

	return x
}

// returns the node with the minimal key and the node pointing to it
func (h *BinomialHeap) MinimumNode() (*BinomialNode, *BinomialNode) {

	if h.head == nil {
		return nil, nil
	}

	min := h.head
	x := min.sibling // node we iterate with
	x_left := min
	var min_left *BinomialNode // node having min as the sibling

	for x != nil {
		if x.Key < min.Key {
			min = x
			min_left = x_left // since we found new min node, we must set node pointing at min correctly
		}

		x_left = x
		x = x.sibling
	}

	return min, min_left
}

// merges two heaps into one based on their degree
func merge(h1, h2 *BinomialHeap) *BinomialNode {

	h := BinomialHeap{}

	if h1.head == nil {
		return h2.head
	}

	if h2.head == nil {
		return h1.head
	}

	h1_next := h1.head
	h2_next := h2.head

	// first, we need to set the head of the new heap correctly
	if h1_next.degree <= h2_next.degree {

		h.head = h1_next
		h1_next = h1_next.sibling

	} else {

		h.head = h2_next
		h2_next = h2_next.sibling

	}

	// last is the last conneted node in the heap
	last := h.head

	// repeat until one of the heaps has no root nodes left
	for h1_next != nil && h2_next != nil {

		if h1_next.degree <= h2_next.degree {

			last.sibling = h1_next
			h1_next = h1_next.sibling

		} else {

			last.sibling = h2_next
			h2_next = h2_next.sibling

		}

		last = last.sibling
	}

	// connect the rest to the heap
	if h1_next != nil {
		last.sibling = h1_next
	} else {
		last.sibling = h2_next
	}

	return h.head
}

// makes a heap out of two heaps correctly
func union(h1, h2 *BinomialHeap) *BinomialNode {

	h := merge(h1, h2)

	if h == nil {
		return nil // trivial
	}

	var prev_x *BinomialNode
	x := h

	next_x := x.sibling

	for next_x != nil {
		if (x.degree != next_x.degree) || (next_x.sibling != nil && next_x.sibling.degree == x.degree) {

			prev_x = x // if ok, simply move
			x = next_x

		} else { // else => correction
			if x.Key <= next_x.Key { // links them together and omits the second one in the list of roots
				x.sibling = next_x.sibling
				x.link(next_x)
			} else {
				if prev_x == nil { // checks if it is not the edge case
					h = next_x
				} else {
					prev_x.sibling = next_x
				}
				next_x.link(x) // links them correctly
				x = next_x
			}
		}
		next_x = x.sibling
	}

	return h
}

// reverses the list starting with a certain node
func reverseChildren(orig_root *BinomialNode) *BinomialHeap {

	x := orig_root
	var last_node *BinomialNode

	for x != nil { // reversing the linked list of children
		x_right := x.sibling
		x.sibling = last_node

		x.parent = nil // we wanna create an independent heap

		last_node = x
		x = x_right
	}

	return MakeBinoHeap(last_node)
}

// connects node y as a child of z
func (z *BinomialNode) link(y *BinomialNode) {

	y.parent = z
	y.sibling = z.child
	z.child = y
	z.degree += 1
}

// printing functions (depth)
func printTree(t *BinomialNode, space string) {

	if t != nil {
		printNode(t, space)
	}

	for child := t.child; child != nil; child = child.sibling {
		printTree(child, space+"\t")
	}
}

func (h *BinomialHeap) PrintHeap() {

	for x := h.head; x != nil; x = x.sibling {
		printTree(x, "")
	}

}

func printNode(x *BinomialNode, space string) {

	fmt.Printf(space+"NODE: key: %d, degree: %d, child: %p, parent: %p \n", x.Key, x.degree, x.child, x.parent)

}

func MakeBinoHeap(head *BinomialNode) *BinomialHeap {
	h := &BinomialHeap{head: head}
	return h
}

func MakeBinoNode(value int) *BinomialNode {
	n := &BinomialNode{Key: value}
	return n
}
