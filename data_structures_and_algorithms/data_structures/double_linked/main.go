package main

import (
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	ll := newLinkedList()
	for i := 0; i < 100; i++ {
		e := newElement(r1.Intn(100))
		ll.insertNodo(*e)
	}
	ll.printList()
}
