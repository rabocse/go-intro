package main

import (
	"fmt"
)

func main() {

	fmt.Println(math.pi)
}

// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

// Unexported ? Exported ? What ? Same than private/public.

// Example.....
