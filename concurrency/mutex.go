package main

import (
	"fmt"
	"sync"
)

var x = 0
var y = 0
var z = 0

func incrementX(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func incrementY(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	m.Unlock()
	wg.Done()
}

func incrementZ(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	<-ch
	wg.Done()
}

func main() {

	// example of race condition due to concurrency of goroutines
	var w1 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w1.Add(1)
		go incrementX(&w1)
	}
	w1.Wait()
	fmt.Println("final value of x", x)

	// prevent race condition by adding mutex to shared resource mutation
	var w2 sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w2.Add(1)
		go incrementY(&w2, &m)
	}
	w2.Wait()
	fmt.Println("final value of y", y)

	// example using channels (though this is confusing compared to mutex)
	var w3 sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w3.Add(1)
		go incrementZ(&w3, ch)
	}
	w3.Wait()
	fmt.Println("final value of z", z)

}
