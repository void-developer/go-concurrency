package main

import (
	"sync/atomic"
)

type ThreadPool struct {
	maxSize  int64
	currSize atomic.Int64
	signal   chan struct{}
}

type threadFunc func(...interface{})

func NewPool(max int64) ThreadPool {
	t := ThreadPool{
		maxSize:  max,
		currSize: atomic.Int64{},
		signal:   make(chan struct{}),
	}
	t.currSize.Store(0)
	return t
}

func (t *ThreadPool) Start(exec threadFunc, params ...interface{}) {
	t.lock()
	go func(p ...interface{}) {
		exec(params)
		t.release()
	}(params)
}

func (t *ThreadPool) release() {
	if t.currSize.Add(-1) == t.maxSize-1 {
		go func() {
			t.signal <- struct{}{}
		}()
	}
}

func (t *ThreadPool) lock() {
	if t.currSize.Load() >= t.maxSize {
		<-t.signal
	}
	t.currSize.Add(1)
}
