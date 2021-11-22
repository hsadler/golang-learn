package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(time.Second)
	ch <- "from server2"

}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {

	// example of waiting for a single channel to return
	// the channel that returns first gets selected for print
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// example of default select case
	// inclusion of default case ensures that the select statement doesn't block
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
