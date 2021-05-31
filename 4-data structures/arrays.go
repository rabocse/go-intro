package main

import "fmt"

func main() {

	var myFirstArray [3]string

	myFirstArray[0] = "The"
	myFirstArray[1] = "Coding"
	myFirstArray[2] = "Hour"

	fmt.Println(" ")
	fmt.Println("==========================================================")

	fmt.Println(myFirstArray)

	fmt.Println(" ")
	fmt.Println("==========================================================")

	fmt.Println(myFirstArray[0])
	fmt.Println(myFirstArray[1])
	fmt.Println(myFirstArray[2])

	fmt.Println(" ")
	fmt.Println("================= Array Literal =========================================")

	mySecondArray := [3]string{"Cisco ", "Krakow ", "Rocks "}

	fmt.Println(mySecondArray)
	fmt.Print(mySecondArray[0])
	fmt.Print(mySecondArray[1])
	fmt.Print(mySecondArray[2])

}

// An array's length is part of its type, so arrays cannot be resized !!!!
