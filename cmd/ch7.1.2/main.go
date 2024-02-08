package main

import "time"

// demonstrate what happens if a channel is written to, but never read from...
func main() {
	ch := make(chan string)
	go readMsg(ch)
	ch <- "hello?"
}

func readMsg(ch chan string) {
	time.Sleep(time.Second * 7)
}
