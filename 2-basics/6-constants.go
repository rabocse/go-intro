package main

import "fmt"

func main() {

	const myAge = 31
	const Pi = 3.14

	fmt.Printf("I am %v years old\n", myAge)
	fmt.Printf("This is the value of Pi: %v\n", Pi)

	const Truth = true
	fmt.Println("Cisco rocks ?", Truth)
}

// := cannot be used when declaring a constant.
