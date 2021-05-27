package main

import (
	"fmt"
)

func split(sum int) (x, y int) { // x and y are the name return values

	x = sum * 4 / 9
	y = sum - x

	return // This returns the named return values
}

func main() {

	fmt.Println(split(200))
}
