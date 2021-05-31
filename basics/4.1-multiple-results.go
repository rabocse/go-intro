package main

import (
	"fmt"
)

func swap(x, y string) (string, string) { // Notice how the return was defined as two strings (string, string) without a names

	return y, x // A function can return any number of results.

}

func main() {

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
