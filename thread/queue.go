package thread

import heaps "github.com/void-developer/go-heaps"

type RunQueue struct {
	heaps.Heap[Task]
}

func InitRunQueue() RunQueue {
	return RunQueue{heaps.Heap[Task]{Tree: make([]Task, 1000), Type: heaps.MaxQ}}
}

func (rq *RunQueue) Add(task Task) {
	task.Id = rq.Heap.Size + 1
	rq.Heap.Push(task)
}

func (rq *RunQueue) Remove() Task {
	return rq.Heap.Pop()
}
