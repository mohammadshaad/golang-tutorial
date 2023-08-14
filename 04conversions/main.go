package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main () {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// comma ok || err ok
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ",input)

	numRating, err  := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating : ", numRating + 1)
	}

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	// var rating int
	// fmt.Scan(&rating)

	// fmt.Println("Thanks for rating, ",rating)

	// switch rating {
	// case 5:
	// 	fmt.Println("Thanks for rating 5 stars")
	// case 4:
	// 	fmt.Println("Thanks for rating 4 stars")
	// case 3:
	// 	fmt.Println("Thanks for rating 3 stars")
	// case 2:
	// 	fmt.Println("Thanks for rating 2 stars")
	// case 1:
	// 	fmt.Println("Thanks for rating 1 stars")
	// default:
	// 	fmt.Println("Invalid rating")
	// }
}