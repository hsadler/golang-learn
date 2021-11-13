package main

import "fmt"

// variable declaration at package level
var x, y bool

// short assignment not avail at package level
// k := 3 // throws error

func main() {

	// assigned default values
	var i int
	var f float32
	var s string

	fmt.Printf(
		"%v %v %v %v %q\n",
		x, y, i, f, s,
	)

	// initializers
	var p, j int = 1, 2

	// type assignment not needed
	var c, python, java = true, false, "no!"

	fmt.Println(p, j, c, python, java)

	// short assignment statement
	k := 3

	fmt.Println(k)

}
