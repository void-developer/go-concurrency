package thread

import "testing"

func TestSimpleTaskQueue(t *testing.T) {
	rq := InitRunQueue()
	rq.Add(Task{BasePriority: 1})
	rq.Add(Task{BasePriority: 2})
	rq.Add(Task{BasePriority: 3})
	rq.Add(Task{BasePriority: 4})
	rq.Add(Task{BasePriority: 5})

	for i := 5; i > 0; i-- {
		task := rq.Remove()
		if task.BasePriority != i || task.Id != i {
			t.Errorf("Expected %d, got %d (id=%d)", i, task.BasePriority, task.Id)
		}
	}
}

func TestUnsortedTaskQueue(t *testing.T) {
	rq := InitRunQueue()
	rq.Add(Task{BasePriority: 5})
	rq.Add(Task{BasePriority: 3})
	rq.Add(Task{BasePriority: 1})
	rq.Add(Task{BasePriority: 4})
	rq.Add(Task{BasePriority: 2})

	for i := 5; i > 0; i-- {
		task := rq.Remove()
		if task.BasePriority != i {
			t.Errorf("Expected %d, got %d (id=%d)", i, task.BasePriority, task.Id)
		}
	}
}
