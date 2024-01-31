package sync

import "sync"

type Semaphore struct {
	cond           *sync.Cond
	permits        int
	grantedPermits int
}

func NewSemaphore(permits int) *Semaphore {
	return &Semaphore{
		cond:           sync.NewCond(&sync.Mutex{}),
		permits:        permits,
		grantedPermits: 0,
	}
}

func (s *Semaphore) Aquire() {
	s.cond.L.Lock()
	for s.grantedPermits == s.permits {
		s.cond.Wait()
	}
	s.grantedPermits++
	s.cond.L.Unlock()
}

func (s *Semaphore) Release() {
	s.cond.L.Lock()
	s.grantedPermits--
	s.cond.Broadcast()
	s.cond.L.Unlock()
}
