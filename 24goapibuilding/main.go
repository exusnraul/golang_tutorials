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

// model for course
type Course struct {
	CourseId    string  `json:"curseid"`
	Coursename  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var courses []Course

//middlewares, helpers

func IsEmpty(c *Course) bool {
	// return c.CourseId == "" && c.Coursename == ""
	return c.Coursename == ""
}

func main() {
	fmt.Println("Api works")
	r := mux.NewRouter()
	//seding of data
	courses = append(courses, Course{CourseId: "2", Coursename: "Python", CoursePrice: 299, Author: &Author{Fullname: "Rahul Saha", Website: "google"}})

	courses = append(courses, Course{CourseId: "4", Coursename: "GoLang", CoursePrice: 299, Author: &Author{Fullname: "Rahul Saha", Website: "google"}})

	//routing
	r.HandleFunc("/", serveHome).Methods{"GET"}
	r.HandleFunc("/courses", getAllcourses).Methods{"GET"}
	r.HandleFunc("/course/{id}", getOnecourse).Methods{"GET"}
	r.HandleFunc("/course", createonecourse).Methods{"POST"}
	r.HandleFunc("/course/{id}", updateonecourse).Methods{"PUT"}
	r.HandleFunc("/course/{id}", deleteonecourse).Methods{"DELETE"}
	//lsten to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to GoLang</h1>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loopthrough courses, find matching id and return resp
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given Id")
	return
}

func createonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some Data")

	}

	//what abut {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside Json")
		return
	}

	//generate uniqueid and string
	//append courses into courses

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	// loop, id ., remove , add with my ID
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	//first - grab id from request
	params := mux.Vars(r)

	// loop , id ,remove (index,index)
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			break
		}

	}
}
