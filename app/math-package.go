
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	// using a sub-package
	fmt.Println("My favorite number is", rand.Intn(10))

	// interpolated string print
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// accessing a package exported name (must be capitalized)
	fmt.Println(math.Pi)

}
