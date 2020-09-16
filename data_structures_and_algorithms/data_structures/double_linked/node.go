package main

type node struct {
	value element
	next  *node
	prev  *node
}

func newNode(value element) *node {
	n := new(node)
	n.value = value
	n.next = nil
	n.prev = nil
	return n
}

func (n *node) getValue() element {
	return n.value
}

func (n *node) getNextNode() *node {
	return n.next
}

func (n *node) getPrevNode() *node {
	return n.prev
}
func (n *node) setValue(e element) {
	n.value = e
}
func (n *node) setNextNode(next *node) {
	n.next = next
}

func (n *node) setPrevNode(prev *node) {
	n.prev = prev
}
