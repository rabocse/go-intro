package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 2 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {

	fmt.Printf("%T , %v \n", Big, Big)

	fmt.Printf("%T, %v \n", Small, Small)

}

// This will overflow the int.
