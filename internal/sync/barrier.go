package sync

import "sync"

type Barrier struct {
	cond    sync.Cond
	waiting int
	size    int // The total number of expected 'waiters'.
}

func NewBarrier(size int) *Barrier {
	return &Barrier{
		cond:    *sync.NewCond(&sync.Mutex{}),
		waiting: 0,
		size:    size,
	}
}

func (b *Barrier) Wait() {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.waiting++
	if b.waiting < b.size {
		b.cond.Wait()
	} else {
		b.waiting = 0
		b.cond.Broadcast()
	}
}
