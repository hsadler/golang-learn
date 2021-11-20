package main

import (
	"fmt"
	"time"
)

func hello1(done chan bool) {
	fmt.Println("hello1 goroutine")
	done <- true
}

func hello2(done chan bool) {
	fmt.Println("hello2 go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello2 go routine awake and going to write to done")
	done <- true
}

func main() {

	// example of receiving data from a goroutine via a channel
	done1 := make(chan bool)
	go hello1(done1)
	result := <-done1
	fmt.Printf("hello1 returned: %v\n", result)

	// example of channel reads blocking execution until the channel write occurs
	done2 := make(chan bool)
	fmt.Println("Main going to call hello2 go goroutine")
	go hello2(done2)
	<-done2
	fmt.Println("Main received data from hello2")

}
