package main

import "fmt"

func defer1() {

	// does not run until after outer function returns
	defer fmt.Println("world")

	fmt.Println("hello")

}

func defer2() {

	fmt.Println("counting")

	// multiple defers are pushed onto a stack resulting in last-in-first-out
	// execution
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

func main() {
	defer1()
	defer2()
}
