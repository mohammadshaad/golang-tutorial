package main

import "fmt"

func main () {
	fmt.Println("Welcome to the ifelse in golang")

	loginCount := 10
	var result string 

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Special User"
	} else {
		result = "Guest User"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

	// if err != nil {
	// 	fmt.Println("Error Occured")
	// }

}