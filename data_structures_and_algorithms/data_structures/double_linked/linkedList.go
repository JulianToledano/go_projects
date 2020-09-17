package main

import "fmt"

type linkedList struct {
	root *node
	size int
}

func newLinkedList() *linkedList {
	return &linkedList{
		root: nil,
		size: 0,
	}
}

func (ll *linkedList) getTail() *node {
	return ll.root.prev
}

func (ll *linkedList) insertNodo(e element) {
	n := newNode(e)
	if ll.root == nil {
		ll.root = n
		ll.root.prev = n
	} else {
		tail := ll.getTail()
		tail.next = n
		n.prev = tail
		ll.root.prev = n
	}
	ll.size++
}

func (ll *linkedList) findNode(e *element) *node {
	return ll.findNodeRecursive(ll.root, e)
}

func (ll *linkedList) findNodeRecursive(n *node, e *element) *node {
	if e.value == n.value.value {
		return n
	}
	if n.next == nil {
		return nil
	}
	return ll.findNodeRecursive(n.next, e)
}

func (ll *linkedList) delete(e *element) {
	n := ll.findNode(e)
	if n != nil {
		if n == ll.root {
			n.next.prev = ll.getTail()
			ll.root = n.next
		} else {
			n.prev.next = n.next
			if n.next == nil {
				ll.root.prev = n.prev
			} else {
				n.next.prev = n.prev
			}
		}
	}
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
