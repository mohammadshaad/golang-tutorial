package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	result := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your full name : ")

	// ReadString will block until the delimiter is entered
	// Delimiter is defined by the user
	// comma ok || err ok

	input, _ := result.ReadString('\n')

	fmt.Println("Hello, Mr. " + input + "Welcome to the world of GoLang")
	// fmt.Printf("Hello, Mr. %sWelcome to the world of GoLang", input)
	fmt.Printf("\nThe type of the variable is %T", input)

} 