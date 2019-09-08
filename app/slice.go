
package main

import "fmt"


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {


	// create array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// create slice as a subset of array
	var primes_slice []int = primes[1:4]
	fmt.Println(primes_slice)


	// DEMONSTRATION: Slices don't copy array values. They are references.

	// create string array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// create two slices referencing array values
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// reassign reference to array via slice index
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)


	// DEMONSTRATION: Slice literals

	// create slice using literal (slice pointing to an array underneath)
	int_slice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(int_slice)

	// bool slice literal
	bool_slice := []bool{true, false, true, true, false, true}
	fmt.Println(bool_slice)

	// struct slice literal
	struct_slice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(struct_slice)


	// DEMONSTRATION: Slice default indexes

	// create slice
	s := []int{2, 3, 5, 7, 11, 13}

	// without default indexes
	s = s[1:4]
	fmt.Println(s)

	// defaut start index (defaults to 0)
	s = s[:2]
	fmt.Println(s)

	// default end index (defaults to length of what's being sliced)
	s = s[1:]
	fmt.Println(s)


	// DEMONSTRATION: Slice length and capacity

	sl := []int{2, 3, 5, 7, 11, 13}
	printSlice(sl)

	// Slice the slice to give it zero length.
	sl = sl[:0]
	printSlice(sl)

	// Extend its length.
	sl = sl[:4]
	printSlice(sl)

	// Drop its first two values.
	sl = sl[2:]
	printSlice(sl)

	// CANNOT DO THIS: Extend it's length beyond capacity
	// sl = sl[:5]
	// printSlice(sl)


	// DEMONSTRATION: nil slice

	var zero_value_slice []int
	fmt.Println(
		zero_value_slice,
		len(zero_value_slice),
		cap(zero_value_slice),
	)
	if zero_value_slice == nil {
		fmt.Println("nil!")
	}


}
