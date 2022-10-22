package task

import "time"

type Options struct {
	Blocking        bool
	ScheduleTimeout time.Duration
}
