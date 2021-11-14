package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // channel must be closed or panic!
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	ch1 := make(chan int)
	go sum(s[:len(s)/2], ch1)
	go sum(s[len(s)/2:], ch1)
	x, y := <-ch1, <-ch1 // receive from c
	fmt.Println(x, y, x+y)

	// buffered channels
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	// ch <- 3 // panic!
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// receiving from a channed until closed
	ch3 := make(chan int, 10)
	go fibonacci(cap(ch3), ch3)
	for i := range ch3 {
		fmt.Println(i)
	}
}
