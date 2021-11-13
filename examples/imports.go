package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome!")
	fmt.Println("The time is", time.Now())

	// using a sub-package
	// initialize the default source for rand: https://pkg.go.dev/math/rand
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))

	// interpolated string print
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// accessing a package exported name (must be capitalized)
	fmt.Println(math.Pi)

}
