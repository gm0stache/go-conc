// Package sync re-implements some synchronisation mechanisms
// for educational purposes.
package sync

import "sync"

// ReadWriteMutex enables restricting read/write access to a resource individually.
type ReadWriteMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock()
	defer rw.readersLock.Unlock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
}

func (rw *ReadWriteMutex) TryReadLock() bool {
	gotLock := rw.readersLock.TryLock()
	if gotLock {
		rw.readersCounter++
		if rw.readersCounter == 1 {
			rw.globalLock.Lock()
		}
	}
	return gotLock
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	defer rw.readersLock.Unlock()
	if rw.readersCounter == 1 {
		rw.globalLock.Unlock()
	}
	rw.readersCounter--
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadWriteMutex) TryLock() bool {
	return rw.globalLock.TryLock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}
