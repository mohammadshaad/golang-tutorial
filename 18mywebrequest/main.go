package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://mohammadshaad.github.io"

func main() {
	fmt.Println("Welcome to the Web Requests in Golang")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Close the response body after the function ends

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
