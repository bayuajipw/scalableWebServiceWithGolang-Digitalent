package main

import "fmt"

// // =========== Variables ============
// func main() {
// 	var animeName, charName string = "MASHLE", "Lance"
// 	var charAge int = 22

// 	charDorm := "Lion"

// 	const pi = 1.27

// 	// _ = animeName

// 	fmt.Println("Hello, World!")
// 	println(pi)
// 	fmt.Println(animeName)
// 	fmt.Printf("His name is %s and he is %v. He got a dorm named %s", charName, charAge, charDorm)
// }

// // =========== Data Types ============
// Tipe data pada bahasa Go terbagi menjadi 4 kategori dengan detail seperti berikut ini:
// 1.Basic Type: number, string, boolean.
// 2.Aggregate Type: array dan struct.
// 3.Reference Type: slice, pointer, map, function, channel
// 4.Interface Type: interface
// func main() {
// var tahunNaga int = 2024
// var hargaBeras float32 = 1990.5
// var balance int64 = -1000
// var point uint = 10
// var pi float32 = 3.144444444
// isAdmin := true
// isAdmin = false

// var name string
// var isAuthorized bool
// fmt.Println(name, nil, isAuthorized)

// fmt.Println(tahunNaga, hargaBeras, balance, point, pi)
// fmt.Printf("Harga beras terupdate: %.3f\n", hargaBeras)
// fmt.Printf("Nilai pi: %.4f\n", pi)
// fmt.Printf("Is current user an admin? %v", isAdmin)
// fmt.Printf("Is current user an admin? %v", isAdmin)
// }

// // =========== Constants & Operators ============
func main() {
	// const pi float32 = 3.34 // tak bisa diubah
	// // pi = 1.2

	// fmt.Println(pi)

	// firstExam := 50
	// secondExam := 100
	// finalScore := (firstExam + secondExam) / 2

	// fmt.Printf("Nilai anda: %v\n", finalScore)
	// fmt.Printf("Apakah nilai ujian pertama lebih besar? %v\n", firstExam > finalScore)
	// fmt.Printf("Apakah nilai akhir sama dengan nilai ujian pertama? %v", finalScore == firstExam)

	// firstName := "Marsha"
	// lastName := "Lenathea"
	// fmt.Printf("\n%v", firstName != lastName)

	isAuthorized := true
	isAdmin := false
	canAccess := isAuthorized != isAdmin

	fmt.Printf("Can access? %v\n", canAccess)
}
