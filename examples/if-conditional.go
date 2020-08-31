
package main

import (
	"fmt"
	"math"
)


// examples of 'if' conditional use
func yes_or_no(yes bool) string {
	if yes {
		return "yes"
	}
	return "no"
}

func pow(x, n, lim float64) float64 {
	// with pre-conditional statement
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func hot_or_cold(temp float64) string {
	// with 'else if' and 'else'
	if temp > 90 {
		return "hot"
	} else if temp < 70 {
		return "cold"
	} else {
		return "just right"
	}
}


func main() {
	fmt.Println(yes_or_no(true), yes_or_no(false))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	fmt.Println(hot_or_cold(99), hot_or_cold(60), hot_or_cold(72))
}

