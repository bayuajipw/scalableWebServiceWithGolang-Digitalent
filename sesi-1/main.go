package main

import "fmt"

func main() {
	var animeName, charName string = "MASHLE", "Lance"
	var charAge int = 22

	charDorm := "Lion"

	const pi = 1.27

	fmt.Println("Hello, World!")
	println(pi)
	fmt.Println(animeName)
	fmt.Printf("His name is %s and he is %v. He got a dorm named %s", charName, charAge, charDorm)
}