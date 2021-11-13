package main

import (
	"fmt"
	"math"
)

func main() {

	// DEMONSTRATION: function callbacks

	// define arg as callback input
	compute := func(
		fn func(float64, float64) float64,
	) float64 {
		return fn(3, 4)
	}

	// set var to function definition
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// use bare
	fmt.Println(hypot(5, 12))

	// use as callback to another function
	fmt.Println(compute(hypot))

	// use standard lib function as callback to another function
	fmt.Println(compute(math.Pow))

	// DEMONSTRATION: function closures

	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}
	// each closure function persists the 'sum' var state
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	// fibonacci is a function that returns a function that returns an int.
	fibonacci := func() func() int {
		first_time := true
		curr := 0
		prev := 0
		return func() int {
			if first_time {
				first_time = false
				return curr
			} else if curr == 0 {
				curr = 1
			} else {
				next := prev + curr
				prev = curr
				curr = next
			}
			return curr
		}
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
