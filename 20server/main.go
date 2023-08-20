package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the server side")

	// PerformGetRequest("http://localhost:8080/get")

	// PerformPostJsonRequest("http://localhost:8080/post")
	PerformPostFormRequest("http://localhost:8080/postform")

}

func PerformGetRequest(myurl string) {

	response, err := http.Get(myurl)

	checkNilError(err)

	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	fmt.Println("Response content length:", response.ContentLength)

	var responseString strings.Builder

	result, err := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(result)

	checkNilError(err)

	fmt.Println("ByeCount is :", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println("Response body:", string(result))
	// fmt.Printf("Response body type: %T\n", result)
}

func PerformPostJsonRequest(myurl string) {

	// fake json payload

	requestBody := strings.NewReader(`
	{
		"username": "admin",
		"password": "admin",
		"email": "callshaad@gmail.com",
		"phone": "1234567890",
		"address": "123, abc street, xyz city, 12345"
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	checkNilError(err)

	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(result))
}

func PerformPostFormRequest(myurl string) {

	data := url.Values{}

	data.Add("username", "admin")
	data.Add("password", "admin")
	data.Add("email", "callshaad@gmail.com")
	data.Add("phone", "1234567890")
	data.Add("address", "123, abc street, xyz city, 12345")

	response, err := http.PostForm(myurl, data)

	checkNilError(err)

	defer response.Body.Close()

	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
