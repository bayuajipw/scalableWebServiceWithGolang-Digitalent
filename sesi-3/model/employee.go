package model

import "fmt"

type Employee struct {
	Person
	Division string
}

func (e *Employee) Greet() {
	fmt.Println("Employee name: ", e.Person.Name)
	fmt.Println("Employee age: ", e.Person.Age)
	e.updateName("Jason")
	fmt.Println("Employee name updated: ", e.Person.Name)
}

func (e *Employee) updateName(name string) {
	e.Person.Name = name
}
