# Go Concurrency

This go library implements a 'thread pool' pattern for concurrent execution of go routines.
The pool has a max parallel parameter that allows the user to specify the maximum of concurrent executions permitted.
When reaching the threshold the pool will stop accepting new thread schedules by guarding the start of the go routine until notified that a new free spot is available.

As of right now the 'waiting' part is left to the caller. Therefor if the pool is fully used, it's the caller that'll be put on wait until it's possible to schedule the requested go-routine

In code:
```go
func (t *Pool) Start(exec threadFunc, params ...interface{}) {
    t.lock() //if the pool is maxed out the executor will not go pass this point until further notice
    //...
}
```

So that in the caller function

```go
fmt.Println("Started")
pool.Start(func(p ...interface{}) {
	//...
}
// if the pool is maxed out the caller will be put on wait here
fmt.Println("Ended")
```

## Installation

```bash
go get github.com/void-developer/go-concurrency
```

## Usage

To create a pool with a maximum number of parallel executions

```go
package x
import "github.com/void-developer/go-concurrency"

pool := thread.NewPool(3)
```

Scheduling a task with the pool is done by calling the `Start` method
```go
pool.Start(func(p ...interface{}) {
    //...
    }, i,x,y,z)
```

the function that has to be scheduled must have an array of interfaces as input, implementing the type:

```go
type threadFunc func(params ...interface{})
```

In fact, the start function has as input parameters:
* the function to schedule
* an array of parameters to pass to the scheduled function

To access the parameters within the function, it's possible to use the specific known index position of the desired parameters
This is an example from the package tests

```go
for i := 0; i < 3; i++ {
		pool.Start(func(p ...interface{}) {
			start := time.Now()
			fmt.Println(fmt.Sprintf("Executing %d at %d", p[0], start.Nanosecond()))
			time.Sleep(3 * time.Second)
			wg.Done()
		}, i)
	}
```