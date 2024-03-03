package main

import (
	"fmt"

	"github.com/bayuajipw/sesi-3/model"
)

// // ------- Function ---------
// func main() {
// 	name := "Kurai"
// 	age := 23

// 	greet(name, age)
// }

// func greet(name string, age int) {
// 	fmt.Printf("Hi, I'm %s and I'm %v years old", name, age)
// }

// // ------- Function (Return) ---------
// func main() {
// 	fmt.Println(greet("Group 1 members:", []string{"Biawak, Tzy, Paradise"}))
// }

// func greet() {
// 	var joinString = strings.Join()
// }

// // ------- Function (Returning multiple values) -------
// func main() {
// 	diameter := float64(14)
// 	area, circumference := calculateCircle(diameter)

// 	fmt.Printf("Circle width: %v\n", diameter)
// 	fmt.Printf("Circle area: %v\n", area)
// 	fmt.Printf("Circle circumference: %v\n", circumference)
// }

// func calculateCircle(diameter float64) (area float64, circumference float64) {
// 	const pi = 3.14

// 	area = (pi * math.Pow(diameter, 2)) / 4
// 	circumference = pi * diameter

// 	return
// }

// // ------- Function (Variadic function 1) -------
// func main() {
// 	studentList := print("Kurai", "Biawak", "Crusher")

// 	fmt.Printf("%s", studentList)
// }

// func print(names ...string) []map[string]string {
// 	result := []map[string]string{}

// 	for i, name := range names {
// 		key := fmt.Sprintf("student-%v", i+1)
// 		result = append(result, map[string]string{
// 			key: name,
// 		})
// 	}

// 	return result
// }

// func main() {
// 	total := sum(5, 5, 2, 3, 5, 1, 2, 2, 10, 5)

// 	fmt.Println("Sum: ", total)
// }

// func sum(numbers ...int) int {
// 	result := 0

// 	for _, num := range numbers {
// 		result += num
// 	}

// 	return result
// }

// // ------- Function (Variadic function 3) -------
// func main() {
// 	profile("Kurai", "Mashle", "OPM", "AOT")
// }

// func profile(name string, favAnime ...string) {
// 	mergefavAnime := strings.Join(favAnime, ", ")

// 	fmt.Printf("Hi, my name is %s\n", name)
// 	fmt.Printf("and my fav Animes are: %s", mergefavAnime)
// }

// // ------- Closure(Declare closure in variable) -------
// func main() {
// 	evenNumbers := func(numbers ...int) []int {
// 		res := []int{}

// 		for _, num := range numbers {
// 			if num%2 == 0 {
// 				res = append(res, num)
// 			}
// 		}

// 		return res
// 	}

// 	fmt.Println(evenNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9))
// }

// // ------- Closure(IIFE) -------
// func main() {
// 	func(str string) {
// 		temp := ""

// 		for i := len(str) - 1; i >= 0; i-- {
// 			temp += string(str[i])
// 		}

// 		fmt.Println("Is palindrome?", str == temp)
// 	}("rusak kasur")
// }

// // ------- Closure (Closure as a return value) -------
// func main() {
// 	studentLists := []string{"Ogi", "Marsha", "Julian"}

// 	find := findStudent(studentLists)

// 	fmt.Println(find("Marsha"))
// }

// func findStudent(students []string) func(string) string {
// 	return func(s string) string {
// 		student := ""
// 		position := -1

// 		for i, stud := range students {
// 			if strings.ToLower(s) == strings.ToLower(stud) {
// 				student = stud
// 				position = i
// 			}
// 		}

// 		if position < 0 {
// 			return fmt.Sprintln("Student not found")
// 		}

// 		return fmt.Sprintf("Found student with name %s in position %v", student, position)
// 	}

// }

// // ------- Closure (Callback) -------
// func main() {
// 	numbers := []int{1, 4, 5, 5, 2}

// 	find = findOddNumbers(numbers, func(numb int) bool {
// 		return numb%2 != 0
// 	})

// 	fmt.Println("Total odd number: ", find)
// }

// func findOddNumbers(number []int, callback isOddNumber) int {
// 	totalOddNumber := 0

// 	for _, num := range numbers {
// 		if callback(num) {
// 			totalOddNumber++
// 		}
// 	}

// 	return totalOddNumber
// }

// // ------- Pointer & Pointer (Memory Address) & Pointer (Changing value through pointer) -------
// func main() {
// 	var firstNumber int = 4
// 	var secondNumber *int = &firstNumber

// 	fmt.Println("firstnumber: ", firstNumber)
// 	fmt.Println("firstnumber memory address: ", &firstNumber)
// 	fmt.Println("secondNumber: ", *secondNumber)
// 	// firstNumber = 1
// 	*secondNumber = 1

// 	fmt.Println("firstnumber after update: ", firstNumber)
// 	fmt.Println("secondNumber memory address: ", secondNumber)
// }

// // ------- Pointer (Pointer as a parameter) -------
// func main() {
// 	number := 10
// 	fmt.Println("Number init: ", number)

// 	change(&number)
// 	fmt.Println("Number after change: ", number)

// }

// func change(num *int) {
// 	*num = 1
// }

// // ------- Struct -------
// type Person struct {
// 	name string
// 	age  int
// }

// type Employee struct {
// 	person   Person
// 	division string
// }

// func main() {
// 	employee := Employee{
// 		name:     "John",
// 		age:      35,
// 		division: "Technology",
// 	}
// 	employee.age = 32
// 	employee.division = "Psychology"

// 	fmt.Println("Employee name: ", employee.name)
// 	fmt.Println("Employee age: ", employee.age)
// 	fmt.Println("Employee division: ", employee.division)
// }

// // ------- Struct (Pointer to a struct) -------
// func main() {
// 	firstEmployee := Employee{name: "Fajar", age: 24, division: "IT"}
// 	secondEmployee := &firstEmployee

// 	fmt.Println("firstEmployee name: ", firstEmployee.name)
// 	fmt.Println("secondEmployee name: ", secondEmployee.name)

// 	secondEmployee.name = "Faqih"

// 	fmt.Println("firstEmployee name: ", firstEmployee.name)
// 	fmt.Println("secondEmployee name: ", secondEmployee.name)
// }

// // ------- Struct (Embedded struct) -------
// func main() {
// 	firstEmployee := Employee{
// 		person: Person{
// 			name: "Katara",
// 			age:  16,
// 		},
// 		division: "Water Bender",
// 	}

// 	// firstEmployee.person.name = "Aang"
// 	// firstEmployee.person.age = 100
// 	// firstEmployee.division = "Avatar"

// 	fmt.Printf("Employee is %+v", firstEmployee)
// }

// // ------- Struct (Anonymous struct) -------
// func main() {
// 	employee := struct {
// 		person   Person
// 		division string
// 	}{
// 		person: Person{
// 			name: "Bumi",
// 			age:  60,
// 		},
// 		division: "Earth Bender",
// 	}

// 	fmt.Printf("Employee is %+v", employee)
// }

// // ------- Struct (Slice of struct) -------
// func main() {
// 	employees := []Employee{
// 		{
// 			person: Person{
// 				name: "Marsha",
// 				age:  17,
// 			},
// 			division: "Center",
// 		},
// 		{
// 			person: Person{
// 				name: "Shani",
// 				age:  22,
// 			},
// 			division: "Captain",
// 		},
// 	}

// 	employees = append(employees, Employee{})

// 	employees[2].person.name = "Bima"
// 	employees[2].person.age = 27
// 	employees[2].division = "Hero"

// 	for _, employee := range employees {
// 		fmt.Printf("Employee data: %v\n", employee)
// 	}
// }

// // ------- Struct (Slice of anonymous struct) -------
// func main() {
// 	employees := []struct {
// 		person   Person
// 		division string
// 	}{
// 		{
// 			person: Person{
// 				name: "Marsha",
// 				age:  17,
// 			},
// 			division: "Center",
// 		},
// 		{
// 			person: Person{
// 				name: "Shani",
// 				age:  22,
// 			},
// 			division: "Captain",
// 		},
// 	}

// 	employees = append(employees, Employee{})

// 	employees[2].person.name = "Bima"
// 	employees[2].person.age = 27
// 	employees[2].division = "Hero"

// 	for _, employee := range employees {
// 		fmt.Printf("Employee data: %v\n", employee)
// 	}
// }

// // ------- Struct (Pointer struct) -> jawaban dari ada yang nanya -------
// type Person struct {
// 	name string
// 	age  int
// }

// type employeeStatus string

// const (
// 	FulltimeEmployee employeeStatus = "fulltime"
// 	ContractEmployee employeeStatus = "contract"
// )

// func (e *Employee) GetPerson() Person {
// 	return e.person
// }

// func (e *Employee) UpdateToFulltime() {
// 	e.status = FulltimeEmployee
// }

// // ------- Exported & Unexported -------
func init() {
	fmt.Println("Hi, from init func!")
}

func main() {
	employee := model.Employee{
		Person: model.Person{
			Name: "Gibran",
			Age:  90,
		},
		Division: "Kang Ngibul",
	}

	employee.Greet()

}
