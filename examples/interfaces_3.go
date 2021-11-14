package main

import "fmt"

func main() {

	var i interface{}
	describe(i)

	// all types implement the empty interface
	i = 42
	describe(i)
	i = "hello"
	describe(i)
	i = true
	describe(i)
	i = []string{"hi"}
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
