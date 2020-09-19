package main

type node struct {
	e    element
	next *node
}

func (n *node) getNext() *node {
	return n.next
}

func (n *node) setNext(next *node) {
	n.next = next
}

func (n *node) getElement() element {
	return n.e
}
func (n *node) setElement(e element) {
	n.e = e
}
