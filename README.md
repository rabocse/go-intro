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

__variable-with-initializers.go__ 

__short-variable-declaration.go__

__basic-types.go__

__zero-values.go__ 

__type-conversions.go__

__type-inference.go__ 

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








