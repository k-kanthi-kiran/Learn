package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

//Fake DB
var courses []Course

//helper/middleware methods -file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

//controllers - Eventually have to go to another file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by kanthi </h1)"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	//fetching the id from the request
	params := mux.Vars(r)

	//looping through the courses which are slices.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No Course found  with the given Id")
		return
	}

}

//Create of resource
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//Scenario when no body is sent in the request
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//Scenario when only {} as body is sent in the request
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate Unique Id
	//Convert the Id to String
	//Append the course to the courses.
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

//Updation
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//Scenario when no body is sent in the request

	//fetching the information from request
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["Id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["Id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

//Delete one Course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
