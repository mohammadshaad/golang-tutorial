package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string   `json:"coursename"`
	Price     float64  `json:"price"`
	IsFree    bool     `json:"is_free"`
	Plateform string   `json:"platform"`
	Password  string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to 22bitmorejson")

	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	myCourse := []course{
		{"ReactJS Bootcamp", 299, false, "Udemy", "123456", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, false, "Udemy", "123456", []string{"web-dev", "js"}},
		{"Go Bootcamp", 199, false, "Udemy", "123456", nil},
	}

	// package this data as JSON data

	finalJSON, err := json.MarshalIndent(myCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", finalJSON)
	fmt.Println(string(finalJSON))

}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"price": 299,
		"is_free": false,
		"platform": "Udemy",
		"tags": ["web-dev","js"]
	}
	
	`)

	var course course // empty struct

	checkvalid := json.Valid(jsonDataFromWeb)

	if checkvalid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &course)
		fmt.Printf("%#v\n", course)

	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}
}
