//Bubble Sort is the simplest sorting algorithm that works by
//repeatedly swapping the adjacent elements if they are in wrong order.

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n Unsorted \n", slice)
	bubblesort(slice)
	fmt.Println("\n Sorted \n", slice, "\n")
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
