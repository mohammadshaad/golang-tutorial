package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)
	fmt.Println(languages)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages : ", languages)

	fmt.Println("JS : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages : ", languages)

	// Iterating over maps

	for _, value := range languages {
		fmt.Printf("For key, value is %v\n", value)
	}
}