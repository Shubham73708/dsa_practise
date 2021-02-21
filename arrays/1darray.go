// an array is a numbered sequence of elements of a specific length.

package main

import "fmt"

func main() {
	var myarr [3]string

	myarr[0] = "shubh"
	myarr[1] = "shahi"
	myarr[2] = "shubham"

	fmt.Println("Elements of array:")
	fmt.Println("Elements 1", myarr[0])
	fmt.Println("Elements 2", myarr[1])
	fmt.Println("Elements 3", myarr[2])
}
