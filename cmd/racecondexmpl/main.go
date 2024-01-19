package main

import (
	"fmt"
	"time"
)

func main() {
	// want to see that this code breaks due to race conditions introduced by parallelizing work?
	// uncommend the code below to run the whole prog on a single CPU core.
	// runtime.GOMAXPROCS(1)

	money := 0

	go stingy(&money)
	go greedy(&money)

	time.Sleep(2 * time.Second)
	fmt.Println("money: ", money)
}

func stingy(money *int) {
	for i := 0; i < 100000; i++ {
		*money += 10
	}
	fmt.Println("stingy done")
}

func greedy(money *int) {
	for i := 0; i < 100000; i++ {
		*money -= 10
	}
	fmt.Println("greedy done")
}
