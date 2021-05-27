package main

import "fmt"

var c, python, java bool // "var" keyword to declare at package level

func main() {
	var i int // "var" to decleare at function level

	fmt.Println(i, c, python, java)
}

// More important, "var" creates the variable with a "zero value"
