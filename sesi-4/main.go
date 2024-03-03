package main

import (
	"fmt"
	"sync"
)

// // ============== Interface ==============
// // Contoh penggunaan Interface di Real Project (Project dari Coach HackTiv8): github.com/bxcodec/go-clean-arch/blob/master/internal/rest/article.go
// type shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type circle struct {
// 	radius float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * math.Pow(c.radius, 2) //pangkat 2
// }

// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func (c circle) volume() float64 {
// 	return math.Pi * math.Pow(c.radius, 3)
// }

// type rectangle struct {
// 	width, length float64
// }

// func (r rectangle) area() float64 {
// 	return r.width * r.length
// }

// func (r rectangle) perimeter() float64 {
// 	return 2 * (r.length * r.width)
// }

// func main() {
// 	var lingkaran shape = circle{radius: 7}
// 	var persegiPanjang shape = rectangle{width: 4, length: 5}

// 	fmt.Printf("Type of lingkaran is: %T\n", lingkaran)
// 	fmt.Println("Circle area is", lingkaran.area())
// 	fmt.Println("Circle perimeter is", lingkaran.perimeter())

// 	fmt.Printf("Type of rectangle is: %T\n", persegiPanjang)
// 	fmt.Println("Rectangle area is", persegiPanjang.area())
// 	fmt.Println("Rectangle perimeter is", persegiPanjang.perimeter())

// 	fmt.Println("Volume of lingkaran is", lingkaran.(circle).volume())

// 	var phoneNumber interface{}

// 	phoneNumber = "6281245678903"
// 	fmt.Println("Phone number data type", reflect.TypeOf(phoneNumber))
// 	fmt.Println("Phone number", phoneNumber.(string))

// 	phoneNumber = 6281245678903
// 	fmt.Println("Phone number data type", reflect.TypeOf(phoneNumber))
// 	fmt.Println("Phone number", phoneNumber.(int))
// }

// // ============== Empty Interface & Empty Interface (Type Assertion) ==============
// func main() {
// 	var phoneNumber interface{}

// 	phoneNumber = []string{"6281245678903", "6281245678903"}

// 	val, ok := phoneNumber.(string)
// 	if !ok {
// 		fmt.Println("phoneNumber is not a string")
// 		return
// 	}

// 	val += " success assert"
// 	fmt.Println("Value of val", val)
// }

// // ============== Empty Interface (Map & Slice with Empty Interface) ==============
// func main() {
// 	test := map[string]interface{}{
// 		"username":   "mantap123",
// 		"age":        23,
// 		"fullName":   "Kurai",
// 		"postalCode": 111111,
// 		"isVerified": false,
// 	}

// 	fmt.Printf("%+v", test)
// }

// // ============== Reflect ==============
// func main() {
// 	age := 24

// 	reflectValue := reflect.ValueOf(age)
// 	reflectType := reflect.TypeOf(age)

// 	fmt.Printf("reflectValue: %v\n", reflectValue)
// 	fmt.Printf("reflectType: %v\n", reflectType)

// 	if reflectValue.Kind() == reflect.Int {
// 		fmt.Println("Age is an int")
// 	}

// 	fmt.Println(reflectValue.Interface())
// 	fmt.Println(reflectValue.Interface().(int))
// }

// // ============== Reflect - Identifying Method Information ==============
// type student struct {
// 	Name string
// 	ID   uint
// }

// func (s *student) SetName(name string) {
// 	s.Name = name
// }

// func main() {
// 	s1 := &student{Name: "Jokowi", ID: 1}
// 	fmt.Println("Student name:", s1.Name)

// 	reflectValue := reflect.ValueOf(s1)
// 	method := reflectValue.MethodByName("SetName")
// 	method.Call([]reflect.Value{
// 		reflect.ValueOf("Bumi"),
// 	})

// 	fmt.Println("Student name:", s1.Name)
// }

// // ============== Concurrency - Goroutines ==============
// func main() {
// 	fmt.Println("Main function started")

// 	fmt.Println("Before goroutine")
// 	go func() {
// 		fmt.Println("Hi from goroutine")
// 	}()

// 	fmt.Println("Main function finished")
// 	time.Sleep(300 * time.Millisecond)
// }

// // ============== Goroutines (Asynchronous process #4) ==============
// func main() {
// 	fmt.Println("Main function started")

// 	go firstProcess(10)
// 	secondProcess(10)

// 	fmt.Println("Active goroutine:", runtime.NumGoroutine())

// 	fmt.Println("Main function finished")
// }

// func firstProcess(index int) {
// 	fmt.Println("First process started")
// 	for i := 1; i <= index; i++ {
// 		fmt.Println("i=", i)
// 	}
// 	fmt.Println("First process finished")
// }

// func secondProcess(index int) {
// 	fmt.Println("Second process started")
// 	for i := 1; i <= index; i++ {
// 		fmt.Println("j=", i)
// 		time.Sleep(2 * time.Millisecond)
// 	}
// 	fmt.Println("Second process finished")
// }

// // ============== Waitgroup ==============
func main() {
	fruits := []string{"Mangga", "Durian", "Apel", "Melon"}

	var wg sync.WaitGroup

	fmt.Println("Before goroutine")

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}

	fmt.Println("Before wait")
	wg.Wait()
	fmt.Println("After goroutine")
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index = %v, fruit = %v\n", index, fruit)
	wg.Done()
}
