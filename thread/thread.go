package thread

import (
	"sync/atomic"
)

type Pool struct {
	maxSize  int64
	currSize atomic.Int64
	signal   chan struct{}
}

type threadFunc func(...interface{})

func NewPool(max int64) *Pool {
	t := &Pool{
		maxSize:  max,
		currSize: atomic.Int64{},
		signal:   make(chan struct{}),
	}
	t.currSize.Store(0)
	return t
}

func (t *Pool) Start(exec threadFunc, params ...interface{}) {
	t.lock()
	go func(p ...interface{}) {
		exec(params...)
		t.release()
	}(params)
}

func (t *Pool) release() {
	if t.currSize.Add(-1) == t.maxSize-1 {
		go func() {
			t.signal <- struct{}{}
		}()
	}
}

func (t *Pool) lock() {
	if t.currSize.Load() >= t.maxSize {
		<-t.signal
	}
	t.currSize.Add(1)
}
