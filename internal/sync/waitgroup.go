package sync

import "sync"

type WaitGroup struct {
	cond      sync.Cond
	groupSize int
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		cond:      *sync.NewCond(&sync.Mutex{}),
		groupSize: 0,
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.groupSize += delta
}

func (wg *WaitGroup) Wait() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
}

func (wg *WaitGroup) Done() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
}
