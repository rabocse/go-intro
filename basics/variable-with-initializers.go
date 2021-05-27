package main

import (
	"fmt"
)

var i, j int = 1, 2 // One initializer per variable.

var k = "guess what?"

func main() {

	var c, python, java = true, false, "no!"

	fmt.Println(i, j, c, python, java)

	fmt.Printf("The type for k is %T \n ", k) // Printing the type

}
