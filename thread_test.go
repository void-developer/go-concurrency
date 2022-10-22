package go_concurrency

import (
	"github.com/void-developer/go-concurrency/task"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestWithinPoolLimits(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(10000)
	res := atomic.Int64{}
	for i := int64(0); i < 10000; i++ {
		pool.Start(func(p ...interface{}) {
			res.Add(p[0].(int64))
			wg.Done()
		}, i)
	}
	wg.Wait()
	if res.Load() != 49995000 {
		t.Error("Sum should be 49995000, but got ", res.Load())
	}
}

func TestExceedPoolLimits(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(10000)
	res := atomic.Int64{}
	for i := int64(0); i < 10000; i++ {
		pool.Start(func(p ...interface{}) {
			res.Add(p[0].(int64))
			wg.Done()
		}, i)
	}
	wg.Wait()
	if res.Load() != 49995000 {
		t.Error("Sum should be 49995000, but got ", res.Load())
	}
}

func TestBelowPoolLimits(t *testing.T) {
	pool := NewPool(10000)
	wg := sync.WaitGroup{}
	wg.Add(9000)
	res := atomic.Int64{}
	for i := int64(0); i < 9000; i++ {
		pool.Start(func(p ...interface{}) {
			res.Add(p[0].(int64))
			wg.Done()
		}, i)
	}
	wg.Wait()
	if res.Load() != 40495500 {
		t.Error("Sum should be 40495500, but got ", res.Load())
	}
}

type CustomObject struct {
	Name string
	Val  int
}

func TestWithCustomObject(t *testing.T) {
	pool := NewPool(10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	res := 0
	pool.Start(func(p ...interface{}) {
		res += p[0].(CustomObject).Val
		wg.Done()
	}, CustomObject{Name: "CustomObject", Val: 10}, 0)
	wg.Wait()
	if res != 10 {
		t.Error("Sum should be 10, but got ", res)
	}
}

func TestExceedNonBlockingFunctions(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(9)
	sum := 0
	for i := 0; i < 9; i++ {
		if !pool.StartWithOptions(task.Options{Blocking: false}, func(p ...interface{}) {
			time.Sleep(1 * time.Second)
			sum += p[0].(int)
			wg.Done()
		}, i) {
			sum += i
			wg.Done()
		}
	}
	wg.Wait()
	if sum != 36 {
		t.Error("Sum should be 36, but got ", sum)
	}
}

func TestTimedOutSchedule(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(6)
	res := atomic.Int64{}
	for i := int64(0); i < 6; i++ {
		if !pool.StartWithOptions(task.Options{Blocking: false, ScheduleTimeout: 1 * time.Second}, func(p ...interface{}) {
			time.Sleep(5 * time.Second)
			res.Add(p[0].(int64))
			wg.Done()
		}, i) {
			wg.Done()
		}
	}
	wg.Wait()
	if res.Load() != 3 {
		t.Error("Sum should be 3, but got ", res.Load())
	}
}
