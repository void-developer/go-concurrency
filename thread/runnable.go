package thread

import "github.com/void-developer/go-heaps/src/heaps/types"

type Runnable interface {
	types.Comparable
}

type Task struct {
	Id           int
	Function     func(...interface{})
	BasePriority int
}

func (t Task) Compare(other types.Comparable) int {
	return t.BasePriority - other.(Task).BasePriority
}

func (t Task) IsNull() bool {
	return t.Id == -1
}

func (t Task) Equals(other types.Comparable) bool {
	return t.Id == other.(Task).Id
}
