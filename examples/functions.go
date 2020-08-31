
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

// multi-result function
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


func main() {

	fmt.Println(add(42, 13))

	fmt.Println(add_again(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	x, y := split(17)
	fmt.Println(x, y)

}
