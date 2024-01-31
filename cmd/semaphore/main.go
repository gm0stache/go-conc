package main

import (
	"fmt"

	"github.com/gm0stache/go-conc/internal/sync"
)

func main() {
	s := sync.NewSemaphore(10)
	for i := 0; i < 5000; i++ {
		go doWork(i, s)
		fmt.Println("waiting for child goroutine")
		s.Aquire()
		fmt.Println("child goroutine finished")
	}
}

func doWork(idx int, s *sync.Semaphore) {
	fmt.Printf("work started... (goroutine %d)", idx)
	fmt.Printf("work finished... (goroutine %d)", idx)
	s.Release()
}
