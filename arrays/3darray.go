package main

import "fmt"

func main() {
	sample := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	fmt.Printf("Length of first dimension: %d\n", len(sample))
	fmt.Printf("Length of second dimension: %d\n", len(sample[0]))
	fmt.Printf("Length of third dimension: %d\n", len(sample[0][0]))

	fmt.Printf("Overall Dimension of the array: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))
	fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0])*len(sample[0][0]))

	// The range keyword is used in for loop to iterate over items of an array, slice, channel or map. With array and slices,
	// it returns the index of the item as integer. With maps, it returns the key of the next key-value pair.

	for _, first := range sample {
		for _, second := range first {
			for _, value := range second {
				fmt.Println(value)
			}
		}
	}
}
