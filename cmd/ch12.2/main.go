package main

import (
	"fmt"
	"sync"
)

type flight struct {
	origin, dest string
	seatsLeft    int
	locker       sync.Locker
}

func book(flights []*flight, seatsToBook int) bool {
	for _, f := range flights {
		f.locker.Lock()
	}
	for _, f := range flights {
		if f.seatsLeft < seatsToBook {
			fmt.Println("not enough seats available")
			return false
		}
		f.seatsLeft -= seatsToBook
	}
	for _, f := range flights {
		f.locker.Unlock()
	}
	return true
}

func newFlight(origin, dest string) *flight {
	return &flight{
		origin:    origin,
		dest:      dest,
		seatsLeft: 200,
		locker:    newSpinLock(),
	}
}
