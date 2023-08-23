package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model

type Course struct {
	ID          string  `json:"id"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Website  string `json:"website"`
}

// Fake DB

var courses []Course

// Middleware | Helper

func (c *Course) IsEmpty() bool {
	// return c.ID == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hello to APIs!")

	r := mux.NewRouter()

	// Fake Data

	courses = append(courses, Course{
		ID:          "1",
		CourseName:  "Golang",
		CoursePrice: 100,
		Author: &Author{
			Fullname: "John Doe",
			Email:    "johndoe@gmail.com",
			Website:  "johndoe.com",
		},
	})

	courses = append(courses, Course{
		ID:          "2",
		CourseName:  "Python",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "Jane Doe",
			Email:    "janedoe@gmail.com",
			Website:  "janedoe.com",
		},
	})

	// Routes

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", r))

}

// Controller

// Serve Home Route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

// Serve Courses Route
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a single course")

	w.Header().Set("Content-Type", "application/json")

	// Grab ID from request

	params := mux.Vars(r)

	// Loop through courses and find with ID

	for _, course := range courses {
		if course.ID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with that ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a single course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if : {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside body")
		return
	}

	// Generate unique ID, string
	// Append course into courses

	rand.Seed(time.Now().UnixNano())
	course.ID = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a single course")
	w.Header().Set("Content-Type", "application/json")


	// First - Grab ID from request

	params := mux.Vars(r)

	// Second - Loop through courses and find with ID

	for index, course := range courses {
		if course.ID == params["id"] {
			// Third - Update course

			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with that ID")
}

func deleteOneCourse (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a single course")
	w.Header().Set("Content-Type", "application/json")

	// First - Grab ID from request
	params := mux.Vars(r)

	// Second - Loop through courses and find with ID
	for index, course := range courses {
		if course.ID == params["id"] {
			// Third - Delete course
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode("Course deleted")
}