package main

import "fmt"

func main() {
	// define a channel variable, then give it a channel value
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	// or, the shorthand version just like maps and slices
	b := make(chan int)
	fmt.Printf("Type of b is %T\n", b)
}
