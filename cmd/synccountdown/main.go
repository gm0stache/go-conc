package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mtx *sync.Mutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		mtx.Lock()
		*seconds -= 1
		mtx.Unlock()
	}
}

func main() {
	var mtx sync.Mutex
	count := 5
	go countdown(&count, &mtx)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}
