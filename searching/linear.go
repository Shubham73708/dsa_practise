// linear search or sequential search is a method for finding an element within a list.
// It sequentially checks each element of the list until a match is found or the whole list has been searched.

package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	//	The range keyword is used in for loop to iterate over items of an array, slice, channel or map. With array and slices,
	//	it returns the index of the item as integer. With maps, it returns the key of the next key-value pair.
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println(linearsearch(items, 58))
}
