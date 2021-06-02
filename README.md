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

0. __hello.go__

```go

package main

import "fmt"

func main() {

	fmt.Println("Hello, Hola")
}
```


1. __packages.go__ 
   

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


2. __imports.go__

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

   
3. __exported-name.go__

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


    
4. __functions.go__ 


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

__4.1_multiple-results.go__


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



         
1. variables.go 
    1. variable-with-initializers.go 
    2. short-variable-declaration.go 
    3. basic-types.go
    4. zero-values.go 
    5. type-conversions.go
    6. type-inference.go 
    7. constants.go    
    8. numeric-constants.go

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








