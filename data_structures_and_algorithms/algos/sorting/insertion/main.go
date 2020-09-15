package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	size := 20
	array := rand.Perm(size)
	fmt.Print(array)
	insertionSort(array, size)
	fmt.Print(array)
}

func insertionSort(arr []int, size int) {
	for i := 1; i < size; i++ {
		value := arr[i]
		before := i - 1

		for {
			if before < 0 || value > arr[before] {
				break
			}
			arr[before+1] = arr[before]
			before--
		}
		arr[before+1] = value
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
