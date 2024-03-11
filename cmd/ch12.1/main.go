// main demonstrates how to use an atomic variable.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func stingy(money *int32) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(money, 10)
	}
	fmt.Println("stingy done.")
}

func spendy(money *int32) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt32(money, -10)
	}
	fmt.Println("spendy done.")
}

func main() {
	money := int32(100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		stingy(&money)
	}()
	go func() {
		defer wg.Done()
		spendy(&money)
	}()
	wg.Wait()
	fmt.Println("money:", atomic.LoadInt32(&money))
}
