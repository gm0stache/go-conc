package sync

import "sync"

type Channel[T any] struct {
	bufferSize int
	buffer     []T
	sendCond   sync.Cond
	recvCond   sync.Cond
}

func NewChannel[T any](buffSize int) *Channel[T] {
	ch := &Channel[T]{
		bufferSize: buffSize,
		buffer:     make([]T, buffSize),
		sendCond:   *sync.NewCond(&sync.Mutex{}),
		recvCond:   *sync.NewCond(&sync.Mutex{}),
	}
	ch.recvCond.L.Lock()
	return ch
}

func (ch *Channel[T]) Send(t T) {
	ch.sendCond.L.Lock()
	defer ch.sendCond.L.Unlock()
	if len(ch.buffer) == ch.bufferSize {
		ch.recvCond.Wait()
	}
	ch.buffer = append(ch.buffer, t)
	ch.sendCond.Signal()
}

func (ch *Channel[T]) Receive() T {
	ch.recvCond.L.Lock()
	defer ch.recvCond.L.Unlock()
	if len(ch.buffer) == 0 {
		ch.sendCond.Wait()
	}
	t := ch.buffer[0]
	ch.buffer = ch.buffer[1:]
	ch.recvCond.Signal()
	return t
}
