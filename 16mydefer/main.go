package main

import "fmt"

func main() {

	// Last in first out for defer
	// defer will execute after the function is finished
	// defer will execute in reverse order
	// defer is used to close the file, close the database connection, unlock the resource

	defer fmt.Println("This is the first line")
	defer fmt.Println("This is the second line")
	defer fmt.Println("This is the third line")
	fmt.Println("Welcome to the defer in Golang")

	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
