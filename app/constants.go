
package main

import "fmt"

const Pi = 3.14

func main() {

	// constants can be string, bool, or numeric
	// cannot use the ':=' syntax

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}
