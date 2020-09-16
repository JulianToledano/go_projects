package main

import (
	"fmt"
)

func main() {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)

	ll := newLinkedList()
	// for i := 0; i < 100; i++ {
	// 	ll.insertNodo(*newElement(r1.Intn(100)))
	// }
	// ll.printList()
	// a := ll.findNode(newElement(r1.Intn(100)))
	// fmt.Println(a)

	z := []int{10, 20, 30, 40, 50}
	for _, v := range z {
		e := newElement(v)
		ll.insertNodo(*e)
	}
	fmt.Println("-----------------------------")
	ll.printList()
	fmt.Println("-----------------------------")
	ll.delete(newElement(50))
	ll.printList()

}
