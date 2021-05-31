package main

import (
	"fmt"
)

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

func main() {

	var i, j int = 1, 2

	k := 3 // Declaring and assining with "short assigment statement". It can be used ONLY inside a function !

	c, python, java := true, false, "no!"

	fmt.Println(i, j, c, k, python, java)

	fmt.Printf("The type for k is %T", k)
}
