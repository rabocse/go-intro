package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4
	var z float64 = math.Sqrt(float64(x*x + y*y)) // If not converted, the program will not compile.
	var f uint = uint(f)

	fmt.Println(x, y, z)

}
