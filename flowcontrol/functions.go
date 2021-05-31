package main

import (
	"fmt"
)

// 1- Function with return.
func double(x float32) float32 {

	return x * 2
}

// 2- Function without return.
func doublePrint(x float32) {

	fmt.Println(x * 2)
}

// 3- Func with two arguments and one return
func sum(a float32, b float32) float32 {

	return a + b
}

// 4- func with two arguments and no return
func sub(a float32, b float32) {

	fmt.Println(a - b)
}

// 5-  Function with multiple returns
func display(a float32, b float32) (c, d float32) {

	item1 := a
	item2 := b

	return item1, item2
}

func main() {

	fmt.Println("============================ 1- double ============================")

	y := double(3.5)
	fmt.Println(y)

	fmt.Println("============================ 2- doublePrint ======================")

	doublePrint(10)

	fmt.Println("============================ 3- sum ==============================")

	totalSum := sum(35, 67)
	fmt.Println(totalSum)

	fmt.Println("============================ 4- sub ==============================")

	sub(100, 10)

	fmt.Println("============================ 5- display ===========================")

	result1, result2 := display(29, 31)

	fmt.Println(result1)
	fmt.Println(result2)

}
