package main

import (
	"fmt"

	"github.com/gm0stache/go-conc/internal/sync"
)

func main() {
	s := sync.NewWeightedSemaphore(100)
	for i := 0; i < 5000; i++ {
		go doWork(i, s)
		fmt.Println("waiting for child goroutine")
		s.Aquire(5)
		fmt.Println("child goroutine finished")
	}
}

func doWork(idx int, s *sync.WeightedSemaphore) {
	fmt.Printf("work started... (goroutine %d)\n", idx)
	fmt.Printf("work finished... (goroutine %d)\n", idx)
	s.Release(5)
}
