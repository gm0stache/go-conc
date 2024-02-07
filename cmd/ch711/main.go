package main

func main() {
	ch := make(chan string)
	go readMsg(ch)
	ch <- "hello"
	ch <- "there"
	ch <- "stop"
}

func readMsg(ch chan string) {
	msg := ""
	for msg != "stop" {
		msg = <-ch
		println("received:", msg)
	}
}
