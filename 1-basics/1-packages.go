package main // What is a package ? What is "main" ? Is it the equivalent of a module ?... I think so...

import ( // Parethesis or individual imports.
	"fmt"       // What are these packages ?
	"math/rand" // Why / in the name of the package ?
)

func main() {

	fmt.Println("My number favorite number is", rand.Intn(10)) // Does the "P" in Println must be capital ?
}

// No need for ; , as a matter of fact the compiler adds them.
