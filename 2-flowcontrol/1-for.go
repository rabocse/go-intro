package main

import "fmt"

func main() {

	fmt.Println("=============== COUNT ===========================")

	sum := 0

	for i := 0; i < 10; i++ {

		sum += i // sum = sum + i

		fmt.Println(i)
	}

	fmt.Println("=============== SUM ===========================")
	fmt.Println(sum)

}

/* for INIT ; CONDITION; POST

INIT: Executed before the first iteration
CONDITION: Evaluated before every iteration
POST: Executed at the end of every iteration



Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

*/
