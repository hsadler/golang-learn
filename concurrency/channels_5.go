package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {

	// example of channel reception until channel close
	ch1 := make(chan int)
	go producer(ch1)
	for {
		v, ok := <-ch1
		fmt.Println("Received ", v, ok)
		if ok == false {
			break
		}
	}

	// example of range loop with channel (automatically exits loop when channel
	// is closed)
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received ", v)
	}

}
