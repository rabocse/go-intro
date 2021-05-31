package main

import "fmt"

var c bool = true

var python, java bool = false, true // "var" keyword to declare at package level

func main() {

	var ruby string = "Ruby"

	var i int = 10 // "var" to decleare at function level

	fmt.Println(i, c, python, java, ruby)
}

// More important, "var" creates the variable with a "zero value"
