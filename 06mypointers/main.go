package main

import "fmt"

func main () {
	fmt.Println("Welcome to the playground!")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber // & - Reference operator

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr) // * - Dereference operator

	*ptr = *ptr * 2
	fmt.Println("New Value is ", myNumber)
}