package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d <len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(i)
	// 	fmt.Println(days[i])
	// } 

	// range - loops over the array or slice

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue <= 10 {

		if rougueValue == 5 {
			rougueValue++
			continue
		} else if rougueValue == 8 {
			break
		}
		
		fmt.Println(rougueValue)
		rougueValue++
	}
}
