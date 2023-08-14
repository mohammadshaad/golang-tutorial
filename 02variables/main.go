package main

import "fmt"

// var username string = "shaad" // This will work
// jwtToken := "asdasdasd" // This will not work
// var jwtToken = "asdasdasd" // This will work

const LoginToken string = "asdasdasd" // Public variable because it starts with capital letter

func main () {
	var username string = "shaad"

	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n\n", username)

	var isLoggedIn bool = false

	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n\n", isLoggedIn)

	var smallVal uint16 = 256

	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n\n", smallVal)

	var smallFloat float64 = 256.9230123

	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n\n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n\n", anotherVariable)

	// Implicit type

	var website = "shaad.pro"
	fmt.Println(website)

	// Shorthand
	// Only works inside a function
	// Outside a function, we use var
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n\n", LoginToken)
}