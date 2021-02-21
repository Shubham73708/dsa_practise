package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := generateSlice(20)
	fmt.Println(" Unsorted \n", slice)
	selectionsort(slice)
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

func selectionsort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

// 990500JB00000025
