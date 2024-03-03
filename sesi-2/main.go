package main

import "fmt"

func main() {
	// isAuthorized := false

	// if isAuthorized {
	// 	fmt.Println("Forbidden")
	// } else {
	// 	fmt.Println("You are authorized!")
	// }

	// thisYear := 2023

	// if thisYear == 2024 {
	// 	fmt.Println("This is dragon year!")
	// } else {
	// 	fmt.Println("No, this is not a dragon year!")
	// }

	// thisYear := 2024

	// if age := thisYear - 2000; age < 17 {
	// 	fmt.Println("Tidak Bisa buat KTP")
	// } else {
	// 	fmt.Println("Bisa buat KTP")
	// }

	// score := 10
	// // ============= Switch Case =============
	// switch {
	// case (score >= 8) && (score <= 10):
	// 	fmt.Println("Perfect")
	// case score >= 5:
	// 	fmt.Println("Normal")
	// case score < 5:
	// 	fmt.Println("Not Bad")
	// }

	// // ============= If Conditions =============
	// if score == 10 {
	// 	fmt.Println("Perfect")
	// } else if (score >= 8) && (score <= 9) {
	// 	fmt.Println("Good")
	// } else if score >= 5 {
	// 	fmt.Println("Not Bad")
	// } else {
	// 	fmt.Println("Bad")
	// }

	// // ============= Nested Conditions =============
	// // If conditions dan Switch case bisa digabung
	// if score >= 8 {
	// 	switch score {
	// 	case 10:
	// 		fmt.Println("Perfect")
	// 	case 9:
	// 		fmt.Println("Almost Perfect")
	// 	default:
	// 		fmt.Println("Good Job")
	// 	}
	// } else if score >= 5 {
	// 	fmt.Println("Poor")
	// } else {
	// 	fmt.Println("Not Good at All")
	// }

	// // ============= For Loop =============
	// for i := 0; i < 3; i++ {
	// 	fmt.Printf("Looping ke-%v\n", i+1)
	// }

	// // ============= Loopings (break and continue keywords) =============
	// i := 0

	// for {
	// 	if i < 3 {
	// 		fmt.Printf("Looping ke-%v\n", i+1)
	// 		i++
	// 		// break
	// 		continue
	// 	}
	// }

	// // ============= Nested Loop =============
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print("*", " ")
	// 	}

	// 	fmt.Println()
	// }

	// for i := 0; i < 5

	// // ============= Looping (Labels) =============
	// outerloop:
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("Perulangan ke-", i)
	// 		for j := 0; j < 5; j++ {
	// 			if i == 3 {
	// 				break outerloop
	// 			}

	// 			fmt.Print(j, " ")
	// 		}

	// 		fmt.Println()
	// 	}

	// // ============= Array =============
	// Bedanya dengan Slice adalah Array datanya sudah fix
	// Jika ditanya lebih sering menggunakan yang mana, jawabannya adalah Slice
	// favouriteFruits := [3]string{"Apple", "Melon"}
	// favouriteFruits[0] = "Grape"
	// favouriteFruits[2] = "Mango"

	// // fmt.Println(favouriteFruits[1])

	// for _, favfavouriteFruit := range favouriteFruits {
	// 	fmt.Println("My fav fruit is", favfavouriteFruit)
	// }

	// // ============= Array(Multidimensional Array) =============
	// sudoku := [3][3]int{{1, 3, 3}, {2, 4, 7}, {5, 9, 5}}

	// for i := 0; i < len(sudoku); i++ {
	// 	for j := 0; j < len(sudoku[i]); j++ {
	// 		fmt.Print(sudoku[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	// // ============= Slice =============
	// favouriteFruits := []string{"Apple", "Mango", "Grape"}

	// fmt.Printf("%#v", favouriteFruits)
	// favouriteFruits := make([]string, 2)
	// favouriteFruits = append(favouriteFruits, "Banana", "Grape", "Mango")

	// fmt.Printf("%#v", favouriteFruits)

	// // ============= Slice (append function with ellipsis) =============
	// favouriteFruits := []string{"Grape", "Apple"}
	// anotherFavouriteFruits := []string{"Berry", "Mango"}

	// favouriteFruits = append(favouriteFruits, anotherFavouriteFruits...)
	// fmt.Printf("%#v", favouriteFruits)

	// // ============= Slice (copy function) =============
	// favouriteFruits := []string{"Grape", "Apple"}
	// anotherFavouriteFruits := []string{"Berry", "Mango"}

	// copyCount := copy(favouriteFruits, anotherFavouriteFruits)
	// fmt.Printf("favouriteFruits ==> %#v\n", favouriteFruits)
	// fmt.Printf("anotherFavouriteFruits ==> %#v\n", anotherFavouriteFruits)
	// fmt.Printf("copyCount ==> %#v", copyCount)

	// // ============= Slice (slicing) =============
	// favouriteFruits := []string{"Grape", "Apple"}
	// anotherFavouriteFruits := []string{"Berry", "Mango"}

	// favouriteFruits = append(favouriteFruits, anotherFavouriteFruits...)

	// fmt.Printf("Hasil slicing %#v", favouriteFruits[0:3])

	// // ============= Slice (Combining slicing and append) =============
	// favouriteFruits := []string{"Grape", "Apple"}
	// anotherFavouriteFruits := []string{"Berry", "Mango"}

	// favouriteFruits = append(favouriteFruits[:1], anotherFavouriteFruits...)

	// fmt.Printf("Hasil slicing %#v", favouriteFruits)

	// // ============= Slice (Backing Array) =============
	// favouriteFruits := []string{"Grape", "Apple"}
	// anotherFavouriteFruits := []string{"Berry", "Mango"}

	// favouriteFruits = append(favouriteFruits[:1], anotherFavouriteFruits...)

	// emptySlice := []string{"test"}

	// fmt.Printf("Hasil slicing %#v\n", favouriteFruits)
	// fmt.Printf("Hasil length %#v\n", len(favouriteFruits))
	// fmt.Printf("Hasil capacity %#v", cap(favouriteFruits))
	// fmt.Printf("Hasil capacity %#v", cap(emptySlice))

	// // ============= Map =============
	// Sama kayak Array & Slice, bedanya bisa nyimpan lebih dari 1 data dan ada keyvalue pairs
	// userInfo := map[string]string{
	// 	"name":    "Kurai",
	// 	"address": "Jakarta",
	// 	"age":     "23",
	// }

	// userInfo["name"] = "Marsha"
	// userInfo["phoneNumber"] = "08111111111"

	// fmt.Printf("%#v\n", userInfo)
	// fmt.Printf("My name is: %v\n", userInfo["name"])

	// // ============= Map (Looping with map) =============
	// userInfo := map[string]string{
	// 	"name":    "Kurai",
	// 	"address": "Jakarta",
	// 	"age":     "23",
	// }

	// userInfo["name"] = "Marsha"
	// userInfo["phoneNumber"] = "08111111111"

	// fmt.Printf("%#v\n", userInfo)

	// // delete(userInfo, "name")

	// for key, value := range userInfo {
	// 	fmt.Printf("%v: %v\n", key, value)
	// }

	// val, exist := userInfo["name"]
	// if !exist {
	// 	fmt.Println("key 'name' not found")
	// } else {
	// 	fmt.Println(val)
	// }

	// // ============= Map (Combining slice with map) =============
	// person := []map[string]string{
	// 	{"name": "Kurai", "address": "Jakarta"},
	// 	{"name": "Bima", "address": "Jakarta"},
	// 	{"name": "Ilham", "address": "Jakarta"},
	// }

	// fmt.Print(person)

	// // ============= Aliase =============
	// var angka uint8 = 1
	// var angkaByte byte

	// angkaByte = angka
	// _ = angkaByte

	// type status string

	// var created status = "created"

	// fmt.Println(created)

	// // ============= Strings In Depth - Looping Over String (byte-by-byte) =============
	city := "Jakarta"

	// for _, c := range city {
	// 	fmt.Println(c)
	// 	fmt.Println(string(c))
	// }

	// // ============= Strings In Depth - Looping Over String (rune-by-rune) =============
	fmt.Println(len(city))
}
