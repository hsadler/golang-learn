
package main

import "fmt"

// basic example
func add(x int, y int) int {
	return x + y
}

// omitted type function
func add_again(x, y int) int {
	return x + y
}

// multi-result funtion
func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	fmt.Println(add(42, 13))

	fmt.Println(add_again(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
