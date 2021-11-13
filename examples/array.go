package main

import "fmt"

func main() {

	// create using type
	var arr1 [3]int
	fmt.Println(arr1)

	// create with literal
	arr2 := [3]int{0, 0, 0}
	fmt.Println(arr2)

	// create with type and assign indexes
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// create with literal and indexes assigned
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

}
