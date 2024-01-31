package sync

import "sync"

type WeightedSemaphore struct {
	cond           *sync.Cond
	permits        int
	grantedPermits int
}

func NewWeightedSemaphore(permits int) *WeightedSemaphore {
	return &WeightedSemaphore{
		cond:           sync.NewCond(&sync.Mutex{}),
		permits:        permits,
		grantedPermits: 0,
	}
}

func (s *WeightedSemaphore) Aquire(permits int) {
	s.cond.L.Lock()
	for (s.grantedPermits + permits) >= s.permits {
		s.cond.Wait()
	}
	s.grantedPermits += permits
	s.cond.L.Unlock()
}

func (s *WeightedSemaphore) Release(permits int) {
	s.cond.L.Lock()
	s.grantedPermits -= permits
	s.cond.Signal()
	s.cond.L.Unlock()
}
