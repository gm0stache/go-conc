package main

import (
	"math/rand"
	"time"

	"github.com/gm0stache/go-conc/internal/sync"
)

func main() {
	b := sync.NewBarrier(3)

	go someWork("first", rand.Intn(3), b)
	go someWork("second", rand.Intn(3), b)

	b.Wait()
	println("the end.")
}

func someWork(name string, duration int, b *sync.Barrier) {
	println(name, "started at:", time.Now().String())
	time.Sleep(time.Duration(duration) * time.Second)
	println(name, "ended at:", time.Now().String())
	b.Wait()
}
