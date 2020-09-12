package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range numbers {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i], "  is an even number")
		} else {
			fmt.Println(numbers[i], " is not an even number")
		}
	}
}
