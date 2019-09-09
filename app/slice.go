
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


	// DEMONSTRATION: using 'make' to create dynamically sized arrays

	// make() arg2 specifies length
	a1 := make([]int, 5)
	printSlice(a1)

	// make() arg3 specifies capacity
	b1 := make([]int, 0, 5)
	printSlice(b1)

	c1 := b1[:2]
	printSlice(c1)

	d1 := c1[2:5]
	printSlice(d1)


	// DEMONSTRATION: appending to a slice

	var sli []int
	printSlice(sli)

	// append works on nil slices
	sli = append(sli, 0)
	printSlice(sli)

	// slice grows as needed
	sli = append(sli, 1)
	printSlice(sli)

	// add more than one element at a time
	sli = append(sli, 2, 3, 4)
	printSlice(sli)


	// DEMONSTRATION: using 'range' for for-loop iteration

	// for index and value in a slice
	var my_slice = []string{"a", "b", "c"}
	for i, v := range my_slice {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	// use a slice simply for the range iteration
	for index, _ := range make([]int, 10) {
		fmt.Printf("index: %v\n", index)
	}

	// or if you only want just the values
	for _, value := range []string{"one", "two", "three"} {
		fmt.Printf("value: %v\n", value)
	}


}
