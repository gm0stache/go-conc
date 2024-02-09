package main

import (
	"sync"
	"time"
)

func sender(msgs chan<- int) {
	for i := 0; i < 7; i++ {
		currentBuffrSize := len(msgs)
		println(time.Now().Format("15:00:00"), "sending:", i, "buffer size:", currentBuffrSize)
		msgs <- i
	}
	close(msgs)
}

func receiver(msgs <-chan int, wg *sync.WaitGroup) {
	for msg := range msgs {
		time.Sleep(time.Second)
		println("received:", msg)
	}
	wg.Done()
}

// main demonstrates simple msg buffering in channels
func main() {
	bffrSize := 3
	msgs := make(chan int, bffrSize)
	go sender(msgs)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go receiver(msgs, &wg)
	wg.Wait()
}
