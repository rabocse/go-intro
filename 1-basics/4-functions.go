package main

import (
	"fmt"
)

func add(x int, y int) int { // function with two arguments and one return // If both parameters share a type, you only need to specify the last one.

	return x + y

}

func main() {

	fmt.Println(add(3, 15))
}

// Read article about Go's declartion syntax by Rob Pike. // Comparison with C syntax.
