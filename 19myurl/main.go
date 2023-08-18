package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://mohammadshaad.github.io/about?name=shaad&age=21#education"

func main() {
	fmt.Println("Welcome to the URLs in Golang")

	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	checkNilError(err)

	fmt.Printf("Result is of type: %T\n", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Fragment)
	fmt.Println(result.User)

	queryparams := result.Query()
	fmt.Printf("Query params are of type: %T\n", queryparams)

	fmt.Println(queryparams["name"])

	for _, value := range queryparams {
		fmt.Println(value)
	}

	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "mohammadshaad.github.io",
		Path: "/about",
		RawQuery: "name=shaad",
		Fragment: "education",
	}

	fmt.Println(partsOfUrl)
}


func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}