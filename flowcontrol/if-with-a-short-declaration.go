package main

import "fmt"

// printB4 returns a "delimiter string" ( =================================== )
func printB4() string {

	y := "=========================="

	return y
}

func main() {

	var x int = 50 // Change the value to play with the condition x > 100

	if z := printB4(); x > 100 {

		fmt.Println(z) // Only reason why we can print "z" is because is executed before the condition  (x > 100). In other words, it does not depend on it.
		fmt.Println("Condition was \"TRUE\"")

	} else {
		fmt.Println(z) // Same here than previous "z" printed
		fmt.Println("Condition was \"FALSE\"")
	}

	// fmt.Println(z)   Variables declared by the statement are only in scope until the end of the if.

}
