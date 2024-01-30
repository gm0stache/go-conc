package sync

import (
	"sync"
)

// WriteReadMutex prevers 'write' access over 'read' access.
type WriteReadMutex struct {
	readersCounter int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

func NewWriterReaderMutex() *WriteReadMutex {
	return &WriteReadMutex{
		readersCounter: 0,
		writersWaiting: 0,
		writerActive:   false,
		cond:           sync.NewCond(&sync.Mutex{}),
	}
}

func (wrm *WriteReadMutex) ReadLock() {
	wrm.cond.L.Lock()
	for wrm.writersWaiting > 0 || wrm.writerActive {
		wrm.cond.Wait()
	}
	wrm.readersCounter++
	wrm.cond.L.Unlock()
}

func (wrm *WriteReadMutex) ReadUnlock() {
	wrm.cond.L.Lock()
	wrm.readersCounter--
	if wrm.readersCounter == 0 {
		wrm.cond.Broadcast()
	}
	wrm.cond.L.Unlock()
}

func (wrm *WriteReadMutex) WriteLock() {
	wrm.cond.L.Lock()
	wrm.writersWaiting++
	for wrm.readersCounter > 0 || wrm.writerActive {
		wrm.cond.Wait()
	}
	wrm.writersWaiting--
	wrm.writerActive = true
	wrm.cond.L.Unlock()
}

func (wrm *WriteReadMutex) WriteUnlock() {
	wrm.cond.L.Lock()
	wrm.writerActive = false
	wrm.cond.Broadcast()
	wrm.cond.L.Unlock()
}
