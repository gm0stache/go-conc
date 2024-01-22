package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	money := 0

	go stingy(&money)
	go greedy(&money)

	time.Sleep(2 * time.Second)
	fmt.Println("money: ", money)
}

func stingy(money *int) {
	for i := 0; i < 100000; i++ {
		saveMoneyUpdate(money, 10)
	}
	fmt.Println("stingy done")
}

func greedy(money *int) {
	for i := 0; i < 100000; i++ {
		saveMoneyUpdate(money, -10)
	}
	fmt.Println("greedy done")
}

var mtx sync.Mutex

func saveMoneyUpdate(ptr *int, newVal int) {
	mtx.Lock()
	defer mtx.Unlock()

	*ptr += newVal
}
