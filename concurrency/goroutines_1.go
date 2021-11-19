package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	// uncomment the line below to to allow time for the hello() print to make
	// it to the console
	// time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
