package main

import (
	"fmt"
)

func main() {

	x := 0 // Play with "x"

	if x < 0 {

		fmt.Println("It is a negative number")
	} else if x > 0 {
		fmt.Println("it is a positive number")
	} else {
		fmt.Println("It is zero !")
	}

}

// Similar than "for" keyword
// After "if" keyword, there is no need of parenthesis for the expression.
// Braces are required to define scope.
