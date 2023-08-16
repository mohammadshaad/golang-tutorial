package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance in golang
	// no super or parent keyword in golang
	// no this or self keyword in golang
	// no constructor in golang
	// no super() or parent() in golang
	// no method overloading in golang
	// no method overriding in golang
	// no virtual keyword in golang
	// no implements keyword in golang
	// no extends keyword in golang
	// no throws keyword in golang
	// no try catch finally in golang

	shaad := User{"Shaad", "callshaad@gmail.com", true, 19}

	fmt.Println(shaad)

	fmt.Printf("Shaad details are : %+v\n", shaad)

	fmt.Printf("Name is %v and Email is %v\n", shaad.Name, shaad.Email)
}

// struct is a user defined type
// we can create a struct by using struct keyword
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// first letter capital - public
