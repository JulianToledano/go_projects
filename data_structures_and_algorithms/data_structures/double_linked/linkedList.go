package main

import "fmt"

type linkedList struct {
	root *node
}

func newLinkedList() *linkedList {
	return &linkedList{
		root: nil,
	}
}

func (ll *linkedList) getTail() *node {
	return ll.root.prev
}

func (ll *linkedList) insertNodo(e element) {
	n := newNode(e)
	if ll.root == nil {
		fmt.Println("Inserting root")
		ll.root = n
		ll.root.prev = n
	} else {
		tail := ll.getTail()
		tail.next = n
		n.prev = tail
		ll.root.prev = n
	}
}

func (ll *linkedList) findRecursive(v int, n *node) *node {
	if n.value.value == v {
		return n
	}
	if n.next == nil {
		return nil
	}
	return ll.findRecursive(v, n.next)
}

func (ll *linkedList) find(v int) element {
	n := ll.findRecursive(v, ll.root)
	return n.value
}

func (ll *linkedList) _printList(n *node) {
	for {
		fmt.Println(n)
		if n.next == nil {
			break
		}
		n = n.next
	}
}

func (ll *linkedList) printList() {
	ll._printList(ll.root)
}
