package main

import (
	"fmt"
	"math/rand"
	"time"
)

type arr struct {
	size     int
	contents []int
}

func main() {
	a := newArr(50)
	fmt.Printf("%T\n", &a.contents[1])
	fmt.Println(a)
	a.bubbleSort()
	fmt.Println(a)
}

func newArr(size int) *arr {
	rand.Seed(time.Now().Unix())
	a := new(arr)
	a.size = size
	a.contents = rand.Perm(size)
	return a
}

func (a *arr) bubbleSort() {
	for i := 0; i < a.size; i++ {
		for j := 0; j < (a.size - i - 1); j++ {
			if a.contents[j] > a.contents[j+1] {
				swap(&a.contents[j], &a.contents[j+1])
			}
		}
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
