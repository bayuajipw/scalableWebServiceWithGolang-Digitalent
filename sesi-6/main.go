package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PORT = ":8080"

// // ============== WEB SERVER ==============
// func main() {
// 	http.HandleFunc("/", greet)

// 	http.ListenAndServe(PORT, nil)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hi, from endpoint!"
// 	fmt.Fprint(w, msg)
// }

// // ============== API ==============

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees []Employee

func init() {
	employees = []Employee{
		{
			ID:       1,
			Name:     "Windah",
			Age:      23,
			Division: "Youtuber",
		},
		{
			ID:       2,
			Name:     "Basudara",
			Age:      28,
			Division: "Creative",
		},
	}
}

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	err := http.ListenAndServe(PORT, nil)
	if err == nil {
		fmt.Println("Service is running on port:", PORT)
	} else {
		fmt.Println("Failed to run service, err::", err)
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templ, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, fmt.Sprint("failed to parse template html, err: ", err), http.StatusInternalServerError)
			return
		}

		templ.Execute(w, employees)

		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
