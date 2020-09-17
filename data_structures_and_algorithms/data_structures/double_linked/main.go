package main

import (
	"fmt"
)

func main() {

	ll := newLinkedList()

	z := []int{10, 20, 30, 40, 50, 60, 70, 80}
	for _, v := range z {
		e := newElement(v)
		ll.insertNodo(*e)
	}
	fmt.Println("-----------------------------")
	ll.printList()
	fmt.Println("-----------------------------")
	ll.delete(newElement(80))
	ll.printList()
	fmt.Println("-----------------------------")
	ll.delete(newElement(10))
	ll.printList()

}
