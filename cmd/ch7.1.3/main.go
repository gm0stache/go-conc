package main

import "time"

// demonstrate what happens if a channel is read from, but never written to...
func main() {
	ch := make(chan string)
	go writeMsg(ch)
	println("received:", <-ch)
}

func writeMsg(ch chan string) {
	time.Sleep(time.Second * 7)
}
