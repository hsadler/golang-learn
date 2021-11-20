package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {

	// exampe of buffered channel
	ch1 := make(chan string, 2)
	ch1 <- "naveen"
	ch1 <- "paul"
	// causes panic: cannot fill a channel past its buffer capacity
	// ch1 <- "harry"
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	// causes panic: cannot request more from a channel when buffer is empty
	// fmt.Println(<-ch1)

	// example of blocking channel send due to filled buffer
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(time.Second)
	for v := range ch2 {
		fmt.Println("read value", v, "from ch")
		time.Sleep(time.Second)
	}

	// example of reading from a closed channel until empty, then the channel
	// reports itself as closed
	ch3 := make(chan int, 5)
	ch3 <- 5
	ch3 <- 6
	close(ch3)
	n, open := <-ch3
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch3
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch3
	fmt.Printf("Received: %d, open: %t\n", n, open)

	// range loop version of the previous example
	ch4 := make(chan int, 5)
	ch4 <- 5
	ch4 <- 6
	close(ch4)
	for n := range ch4 {
		fmt.Println("Received:", n)
	}

	// example of length and capacity api for channels
	ch5 := make(chan string, 3)
	ch5 <- "naveen"
	ch5 <- "paul"
	fmt.Println("capacity is", cap(ch5))
	fmt.Println("length is", len(ch5))
	fmt.Println("read value", <-ch5)
	fmt.Println("new length is", len(ch5))

}
