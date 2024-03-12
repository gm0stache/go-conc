package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock int32

func newSpinLock() sync.Locker {
	return new(spinLock)
}

func (sl *spinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
		// explicitly yield execution, enables other goroutines to possibly unlock the spinlock
		runtime.Gosched()
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreInt32((*int32)(sl), 0)
}
