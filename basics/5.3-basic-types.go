package main

import (
	"fmt"
	"math/cmplx"
)

var ( //  variable declaration "factored" into blocks
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("TYPE: %T 			VALUE: %v\n", ToBe, ToBe)
	fmt.Printf("TYPE: %T 			VALUE: %v\n", MaxInt, MaxInt)
	fmt.Printf("TYPE: %T		VALUE: %v\n", z, z)

}
