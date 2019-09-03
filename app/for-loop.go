
package main

import "fmt"

func main() {


	// The basic for loop has three components separated by semicolons:

		// the init statement: executed before the first iteration
		// the condition expression: evaluated before every iteration
		// the post statement: executed at the end of every iteration

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	// init and post statements are optional
	sum2 := 1
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)


	// dropped semicolons
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println(sum3)


	// infinite loop
	for {
		fmt.Println("im in an infinite loop...")
	}


}
