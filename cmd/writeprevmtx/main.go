package main

import (
	"fmt"
	"time"

	mySync "github.com/gm0stache/go-conc/internal/sync"
)

// This program demonstrates how a 'write-starvation' can be prevented using a write favoring mutex.
func main() {
	rwMtx := mySync.NewWriterReaderMutex()
	for i := 0; i < 2; i++ {
		go func() {
			for {
				rwMtx.ReadLock()
				time.Sleep(1 * time.Second)
				fmt.Println("read done")
				rwMtx.ReadUnlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	rwMtx.WriteLock()
	fmt.Println("write finished")
}
