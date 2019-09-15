
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// attach a method to a type with the receiver arg, in this case '(v Vertex)'
func (v Vertex) AbsMethod() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs() defined as a normal function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// attach method with receiver as pointer to vertex struct
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	// NOTES ABOUT METHODS:
		// - can only be declared in same package as the defined type
		// - can be attached to any user-defined type, but no built-in types
		// - can be attached to struct via pointer (more common way since
		//		allows mutations to the struct)
		// - calls to v.Scale() are auto-converted to (&v).Scale()
		// - advice: don't mix value and pointer receivers for each struct type

	v := Vertex{3, 4}

	// method call
	fmt.Println(v.AbsMethod())

	// regular function call
	fmt.Println(AbsFunc(v))

	// call to pointer method
	v.Scale(10)
	fmt.Println(v.AbsMethod())

}
