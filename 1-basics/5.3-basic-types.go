package main

import (
	"fmt"
	"math/cmplx"
)

var ( //  variable declaration "factored" into blocks
	CiscoRocks bool       = true
	tempKrakow float32    = 32.5
	MaxInt     uint64     = 1<<64 - 1
	z          complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("TYPE: %T 			VALUE: %v\n", CiscoRocks, CiscoRocks)
	fmt.Printf("TYPE: %T 			VALUE: %v\n", tempKrakow, tempKrakow)
	fmt.Printf("TYPE: %T			VALUE: %v\n", z, z)

}
