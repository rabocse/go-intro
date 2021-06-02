# __Intro to Go__
---

![Intro to Go](https://www.hardwinsoftware.com/blog/wp-content/uploads/2018/02/golang-gopher.png)   

// Picture needs to be replaced !!!! 


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

```

__type-conversions.go__


__constants.go__    

__numeric-constants.go__


 ---
### __2-Flowcontrol__

1. for.go
    1. for-continued.go
    2. while.go
    3. forever.go

2. if.go
    1. if-with-a-short-declaration.go

3. switch.go

4. functions.go

5. Defer (pending)
    1. Stacking defers








