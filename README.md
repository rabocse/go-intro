# __Intro to Go__
---

![Intro to Go](https://www.hardwinsoftware.com/blog/wp-content/uploads/2018/02/golang-gopher.png)   

PENDING LIST:

// Picture needs to be replaced !!!! 
// Use a more adecuate Rune example instead of 'K'. Also explain why '' instead of "".
---
## Index ???? 
---



---
### __1-Basics__
---

__hello.go__

```go

package main

import "fmt"

func main() {

	fmt.Println("Hello, Hola")
}
```


__packages.go__ 
   

```go
package main 

import ( 
	"fmt"      
	"math/rand" 
)

func main() {

	fmt.Println("My favorite number is", rand.Intn(10)) 
}

```


__imports.go__

```go

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("Now you have %g problems. \n ", math.Sqrt(16)) 
}
```

   
__exported-name.go__

```go

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Pi)
}

```


    
__functions.go__ 


```go

package main

import (
	"fmt"
)

func add(x int, y int) int { 

	return x + y

}

func main() {

	fmt.Println(add(3, 15))
}


```

__multiple-results.go__


```go 

package main

import "fmt"

func calc(x int, y int) (a int, b string) {

	a = x + y

	if a > 0 {

		b = "The result is a positive number"
		return a, b
	} else if a < 0 {
		b = "The result is a negative number"
		return a, b
	} else {

		b = "The result is zero"
		return a, b
	}
}

func main() {

	total, info := calc(-5, 5)

	fmt.Println("========================================")
	fmt.Printf("The result was: %v \n", total)
	fmt.Println("========================================")
	fmt.Println(info)
}

```


         
__variables.go__

```go
package main

import "fmt"

// Variables are declared but not manually initialized (nothing was assigned)

var NumberOfApples int

var NumberOfOranges int

var tempKrakow float32

var tempLisbon float32

var carColor string

var wallColor string

var CiscoRocks bool

var KrakowRocks bool

func main() {

	var NumberOfPears int
	var tempBrussels float32
	var bloodColor string
	var LisbonRocks bool

	fmt.Println("============= Apples and Oranges =============================")
	fmt.Println(NumberOfApples)
	fmt.Println(NumberOfOranges)
	fmt.Println("============= Temp. Krakow and Lisbon =========================")
	fmt.Println(tempKrakow)
	fmt.Println(tempLisbon)
	fmt.Println("============= Car and Wall Color =============================")
	fmt.Println(carColor)
	fmt.Println(wallColor)
	fmt.Println("============= Do they rock ? ===================================")
	fmt.Println(CiscoRocks)
	fmt.Println(KrakowRocks)
	fmt.Println("============= Fruit ===========================================")
	fmt.Println(NumberOfPears)
	fmt.Println("============= Temp ============================================")
	fmt.Println(tempBrussels)
	fmt.Println("============= Color ===========================================")
	fmt.Println(bloodColor)
	fmt.Println("============= Do they rock ? ===================================")
	fmt.Print(LisbonRocks)

}


```


__variable-with-initializers.go__ 


```go
package main

import "fmt"

// Variables are declared and initialized.

var NumberOfApples int = 5

var NumberOfOranges int = 10

var tempKrakow float32 = 32.5

var tempLisbon float32 = 25.7

var carColor string = "White"

var wallColor string = "Green"

var CiscoRocks bool = true

var KrakowRocks bool = false

func main() {

	var NumberOfPears int = 7
	var tempBrussels float32 = 21.3
	var bloodColor string = "Red"
	var LisbonRocks bool = false

	fmt.Println("============= Apples and Oranges =============================")
	fmt.Println(NumberOfApples)
	fmt.Println(NumberOfOranges)
	fmt.Println("============= Temp. Krakow and Lisbon ========================")
	fmt.Println(tempKrakow)
	fmt.Println(tempLisbon)
	fmt.Println("============= Car and Wall Color =============================")
	fmt.Println(carColor)
	fmt.Println(wallColor)
	fmt.Println("============= Cisco & Krakow. Do they rock ? =================")
	fmt.Println(CiscoRocks)
	fmt.Println(KrakowRocks)
	fmt.Println("============= Pears ==========================================")
	fmt.Println(NumberOfPears)
	fmt.Println("============= Temp. Brussels ==================================")
	fmt.Println(tempBrussels)
	fmt.Println("============= Blood Color =====================================")
	fmt.Println(bloodColor)
	fmt.Println("============= Lisbon rocks ? ==================================")
	fmt.Print(LisbonRocks)

}

```

__zero-values.go__

```go

package main

import (
	"fmt"
)

func main() {

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

```


__short-variable-declaration.go__


```go

package main

import "fmt"

// NumberOfCars := 2

func main() {

	// short-declaration notation ( := ) is used instead of "var" keyword.

	NumberOfApples := 5
	NumberOfOranges := 10
	tempKrakow := 32.5
	tempLisbon := 25.7
	carColor := "White"
	wallColor := "Green"
	CiscoRocks := true
	KrakowRocks := false
	NumberOfPears := 7
	tempBrussels := 21.3
	bloodColor := "Red"
	LisbonRocks := false

	fmt.Println("============= Apples and Oranges =============================")
	fmt.Println(NumberOfApples)
	fmt.Println(NumberOfOranges)
	fmt.Println("============= Temp. Krakow and Lisbon ========================")
	fmt.Println(tempKrakow)
	fmt.Println(tempLisbon)
	fmt.Println("============= Car and Wall Color =============================")
	fmt.Println(carColor)
	fmt.Println(wallColor)
	fmt.Println("============= Cisco & Krakow. Do they rock ? =================")
	fmt.Println(CiscoRocks)
	fmt.Println(KrakowRocks)
	fmt.Println("============= Pears ==========================================")
	fmt.Println(NumberOfPears)
	fmt.Println("============= Temp. Brussels ==================================")
	fmt.Println(tempBrussels)
	fmt.Println("============= Blood Color =====================================")
	fmt.Println(bloodColor)
	fmt.Println("============= Lisbon rocks ? ===================================")
	fmt.Print(LisbonRocks)

}


```

__type-inference.go__ 

```go

package main

import (
	"fmt"
)

func main() {

	value := 29 // change me !

	fmt.Printf("The variable \"value\" is of type: %T \n", value)
}


```


__basic-types.go__


```go

package main

import (
	"fmt"
)

// Boolean

var summerRocks bool = true
var winterRocks bool = false

var (

	// Integers (Signed Integers)
	value1 int   = 1990
	value2 int8  = 127
	value3 int16 = 1990
	value4 int32 = 1990
	value5 int64 = 1990

	// Integers (Unsigned Integers)
	value6  uint    = 255
	value7  uint8   = 255
	value8  uint16  = 255
	value9  uint32  = 255
	value10 uint64  = 1<<64 - 1
	value11 uintptr = 255

	// Floats
	value12 float32 = 37.5
	value13 float64 = 37.5

	// Byte and Rune are aliases for uint8 and int32 respectively.
	value14 byte = 255
	value15 rune = 'k'

	// Complex numbers
	value16 complex64  = 2 + 0.5i
	value17 complex128 = 2 + 0.7i
)

func main() {

	fmt.Println(" ")
	fmt.Println("=============== Booleans =========================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", summerRocks, summerRocks)
	fmt.Printf("VALUE: %v				TYPE: %T \n", winterRocks, winterRocks)
	fmt.Println(" ")
	fmt.Println("=============== Integers (Signed Integers) =======================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value1, value1)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value2, value2)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value3, value3)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value4, value4)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value5, value5)
	fmt.Println(" ")
	fmt.Println("============== Integers (Unsigned Integers ) ======================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value6, value6)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value7, value7)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value8, value8)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value9, value9)
	fmt.Printf("VALUE: %v		TYPE: %T \n", value10, value10)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value11, value11)
	fmt.Println(" ")
	fmt.Println("============== Floats ==============================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value12, value12)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value13, value13)
	fmt.Println(" ")
	fmt.Println("============== Byte and Rune =======================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value14, value14)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value15, value15)
	fmt.Println(" ")
	fmt.Println("============== Complex Numbers =====================================")
	fmt.Printf("VALUE: %v				TYPE: %T \n", value16, value16)
	fmt.Printf("VALUE: %v				TYPE: %T \n", value17, value17)

}



```

__type-conversions.go__

```go

package main

import (
	"fmt"
)

func main() {

	var value1 int = 150

	var value2 int = 30

	var value3 float32 = 30.5

	result1 := value1 + value2

	result2 := value1 + int(value3) // Remove the conversion for value3 and the program will not compile.

	fmt.Println(result1)
	fmt.Println(result2)

}


```


__constants.go__    


```go

package main

import "fmt"

func main() {

	const myAge = 31
	const Pi = 3.14
	const ciscoRocks = true

	fmt.Printf("I am %v years old\n", myAge)
	fmt.Printf("This is the value of Pi: %v\n", Pi)
	fmt.Println("Cisco rocks ?", ciscoRocks)

}


```

 ---
### __2-Flowcontrol__


__for.go__

```go



```



__for-continued.go__

```go



```


__while.go__

```go



```

__forever.go__

```go



```

__if.go__

```go



```

__if-with-a-short-declaration.go__

```go



```


__switch.go__

```go



```

__functions.go__

```go



```

__Defer (pending)__

```go



```

__Stacking defers__

```go



```


