package main

import "fmt"

func sendData1(sendch chan<- int) {
	sendch <- 10
}

func sendData2(sendch chan<- int) {
	sendch <- 20
}

func main() {

	// example of unidirectional channel
	sendch := make(chan<- int)
	go sendData1(sendch)
	// not possible since sendch cannot be used for reception
	// received := <-sendch

	// example of channel type conversion
	chnl := make(chan int)
	go sendData2(chnl)
	fmt.Println(<-chnl)
}
