package main

import "fmt"

func main() {
	fmt.Println("Welcome to the methods in Golang")

	shaad := User{"Shaad", "callshaad@gmail.com", true, 19}

	fmt.Println(shaad)

	fmt.Printf("Shaad details are %+v\n", shaad)
	fmt.Printf("Shaad details are %#v\n", shaad)

	shaad.GetStatus()

	shaad.NewMail()
	fmt.Println("The original Email :", shaad.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user active? : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "textshaad@gmail.com"
	fmt.Println("\nEmail of this user is : ", u.Email)
}

