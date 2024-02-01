package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const parallelExecs = 4

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(parallelExecs)
	for i := 0; i < parallelExecs; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
	fmt.Println("all done.")
}

func doWork(idx int, wg *sync.WaitGroup) {
	rnd := rand.Intn(5)
	time.Sleep(time.Duration(rnd) * time.Second)
	fmt.Printf("%d has finished.\n", idx)
	wg.Done()
}
