package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"} // we don't need to specify the size of the slice
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // this will remove the first element from the slice
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // the edges are non-inclusive
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // this will give an error as the size of the slice is 4
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 777) // resizing the slice
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // sorting the slice
	fmt.Println(highScores)

	// slicing a slice

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // ... is used to append the slice
	fmt.Println(courses)

}

