package main

import (
	"fmt"
)

func main() {

	fmt.Println(" ")
	fmt.Println("=============================== Switch ===========================================")

	// Switch case

	i := 1

	fmt.Print("Write ", i, " as ")

	switch i {

	case 1:
		fmt.Println("\"one\"")
	case 2:
		fmt.Println("\"two\"")
	case 3:
		fmt.Println("three")
	}

	fmt.Println(" ")
	fmt.Println("================= Switch with no condition ==========================================")

	// Switch with no condition = switch true

	number := 100

	switch true {

	case number < 100:
		fmt.Println("The number is less than 100")

	case number > 100:
		fmt.Println("The number is higher than 100")
	default:
		fmt.Println("The number is 100")
	}

}
