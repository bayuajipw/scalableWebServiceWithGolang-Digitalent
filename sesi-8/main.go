package main

import (
	"fmt"
	"net/url"
)

// //========================= DECODING & PARSING JSON =========================

type Employee struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// func main() {
// 	// jsonString := `
// 	// 	{
// 	// 		"full_name": "Windah Basudara",
// 	// 		"email": "windah@basudara.gg",
// 	// 		"age": 30
// 	// 	}
// 	// `

// 	jsonString := `
// 		[
// 			{
// 				"full_name": "Windah Basudara",
// 				"email": "windah@basudara.gg",
// 				"age": 30
// 			},
// 			{
// 				"full_name": "Dean KT",
// 				"email": "Dean@jkt48.gg",
// 				"age": 50
// 			}
// 		]
// 	`

// 	// var employee Employee
// 	// var employee map[string]interface{}
// 	// var employee interface{}
// 	var employees []Employee

// 	// err := json.Unmarshal([]byte(jsonString), &employee)
// 	// if err != nil {
// 	// 	fmt.Println("Failed to unmarshal json data, err:", err)
// 	// 	return
// 	// }

// 	err := json.Unmarshal([]byte(jsonString), &employees)
// 	if err != nil {
// 		fmt.Println("Failed to unmarshal json data, err:", err)
// 		return
// 	}

// 	// fmt.Println("full_name:", employee.Fullname)
// 	// fmt.Println("email:", employee.Email)
// 	// fmt.Println("age:", employee.Age)

// 	// fmt.Println("full_name:", employee["full_name"])
// 	// fmt.Println("email:", employee["email"])
// 	// fmt.Println("age:", employee["age"])

// 	// result := employee.(map[string]interface{})

// 	// fmt.Println("full_name:", result["full_name"])
// 	// fmt.Println("email:", result["email"])
// 	// fmt.Println("age:", result["age"])

// 	for i, employee := range employees {
// 		fmt.Printf("Employee-%d: %+v\n", i+1, employee)
// 	}
// }

// func main() {
// 	// employee := Employee{
// 	// 	Fullname: "Windah Basudara",
// 	// 	Email:    "windah@basudara.gg",
// 	// 	Age:      30,
// 	// }

// 	// employee := map[string]interface{}{
// 	// 	"full_name": "Marsha",
// 	// 	"email":     "marsha@lenathea.com",
// 	// 	"age":       18,
// 	// }

// 	employee := []Employee{
// 		{
// 			Fullname: "Monkey Luffy",
// 			Email:    "monkey@luffy.gg",
// 			Age:      17,
// 		},
// 		{
// 			Fullname: "Avatar Aang",
// 			Email:    "avatar@aang.gg",
// 			Age:      12,
// 		},
// 	}

// 	jsonString, err := json.Marshal(&employee)
// 	if err != nil {
// 		fmt.Println("Failed to marshal, err:", err)
// 	}

// 	fmt.Println("Result from json marshal:", string(jsonString))
// }

// // ========================= URL PARSING =========================

func main() {
	urlString := "https://gg-geming.com:8080/user/1?lang=id&curr=USD"

	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Failed to parse url string, err:", err)
		return
	}

	fmt.Println("Full URL:", urlString)

	fmt.Println("URL Protocol:", url.Scheme)
	fmt.Println("URL host:", url.Host)
	fmt.Println("URL path/endpoint:", url.Path)

	fmt.Println("URL query params:", url.Query().Encode())

	fmt.Println("Currency query params:", url.Query().Get("curr"))
	fmt.Println("Language query params:", url.Query().Get("lang"))
}
