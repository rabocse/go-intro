package main

import (
	"fmt"
)

func main() {

	value1 := 20.5

	value2 := 20.555555555555555555555555555555555555555555555555

	var value3 float32 = 20.

	fmt.Printf("The type for \"value1\" is: %T\n", value1)
	fmt.Printf("The type for \"value2\" is: %T\n", value2)
	fmt.Printf("The type for \"value3\" is: %T\n", value3)
}
