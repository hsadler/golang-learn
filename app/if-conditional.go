
package main

import (
	"fmt"
)

func yes_or_no(yes bool) string {
	// example of 'if' conditional use
	if yes {
		return "yes"
	}
	return "no"
}

func main() {
	fmt.Println(yes_or_no(true), yes_or_no(false))
}

