package thread

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWithinPoolLimits(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		pool.Start(func(p ...interface{}) {
			start := time.Now()
			fmt.Println(fmt.Sprintf("Executing %d at %d", p[0], start.Nanosecond()))
			time.Sleep(3 * time.Second)
			wg.Done()
		}, i)
	}
	wg.Wait()
}

func TestExceedPoolLimits(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(9)
	for i := 0; i < 9; i++ {
		pool.Start(func(p ...interface{}) {
			start := time.Now()
			wait := time.Duration(rand.Intn(10000)) * time.Millisecond
			fmt.Println(fmt.Sprintf("Executing %d at %d, waiting %d", p[0], start.Nanosecond(), wait.Milliseconds()))
			time.Sleep(wait)
			fmt.Println(fmt.Sprintf("Finished %d at %d", p[0], time.Now().Nanosecond()))
			wg.Done()
		}, i)
	}
	wg.Wait()
}

func TestBelowPoolLimits(t *testing.T) {
	pool := NewPool(3)
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i++ {
		pool.Start(func(p ...interface{}) {
			start := time.Now()
			fmt.Println(fmt.Sprintf("Executing %d at %d", p[0], start.Nanosecond()))
			time.Sleep(3 * time.Second)
			wg.Done()
		}, i)
	}
	wg.Wait()
}
