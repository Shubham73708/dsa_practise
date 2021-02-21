// Insertion sort is a simple sorting algorithm that works similar to the way you sort playing cards in your hands.
// The array is virtually split into a sorted and an unsorted part.
// Values from the unsorted part are picked and placed at the correct position in the sorted part.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := generateSlice(20)
	fmt.Println(" Unsorted \n", slice)
	insertionsort(slice)
	fmt.Println(" Sorted \n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		// Pick ith element and keep swapping with i-1th element if nums[i] < nums[i-1]
		j := i
		for j > 0 {
			// If value at index j is smaller than the one at j-1, swap them
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
