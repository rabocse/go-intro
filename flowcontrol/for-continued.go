package main

import "fmt"

/*

"for" loop without "init" and "post"

*/

func main() {

	fmt.Println("======================= LOOPING... ===========================")

	sum := 1

	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println("======================= TOTAL =================================")
	fmt.Println(sum)
}
