package main

import (
	"fmt"
	"time"
)

func decrementCntr(ctr *int) {
	for *ctr > 0 {
		time.Sleep(1 * time.Second)
		*ctr -= 1
	}
}

func main() {
	counter := 5
	go decrementCntr(&counter)
	for counter > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("coutner: ", counter)
	}
}
