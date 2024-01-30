package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mtx := sync.Mutex{}
	cond := sync.NewCond(&mtx)

	msg := ""

	go writeMsg(&msg, cond)
	go readMsg(&msg, cond)

	time.Sleep(10 * time.Second)
	fmt.Println("Done.")
}

func writeMsg(msg *string, cond *sync.Cond) {
	cond.L.Lock()
	// for i := 0; i < 100; i++ {
	*msg = fmt.Sprintf("Msg %d", 1)
	fmt.Printf("Send message: %q\n", *msg)
	cond.Signal()
	// }
	cond.L.Unlock()
}

func readMsg(msg *string, cond *sync.Cond) {
	cond.L.Lock()
	lastMsg := ""
	for i := 0; i < 10000; i++ {
		for *msg != lastMsg {
			cond.Wait()
		}
		fmt.Printf("Received message: %q\n", *msg)
		lastMsg = *msg
	}
	cond.L.Unlock()
}
