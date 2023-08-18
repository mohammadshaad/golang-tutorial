package main

import "fmt"

func main() {
	fmt.Println("Welcome to the functions in Golang")
	
	subMain()

	result := adder(3, 5)

	fmt.Println("Result is :", result)

	proTextResult, proTotalResult := proAdder(3, 5, 6, 7, 8, 9, 10)

	fmt.Println(proTextResult, proTotalResult)
}


// variadic functions
func proAdder (values ...int) (string, int) {
	total := 0
	for _, value := range values {
		total += value
	}

	return "The calculated sum is ", total
}

// function signtarures - func name (params) return type
func adder (a int, b int) int {
	return a + b
}

func subMain () {
	fmt.Println("Another Salam from Golang")
}
