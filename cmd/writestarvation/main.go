package main

import (
	"fmt"
	"time"

	"github.com/gm0stache/go-conc/sync"
)

// main demonstrates 'write-starvation' scenario,
// where a 'write' lock can never be obtained because the 'read' lock is never freed.
func main() {
	rwMtx := sync.ReadWriteMutex{}
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
