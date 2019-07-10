package ratelimit

import "time"

type Bucket interface {
	Wait(count int64)
	WaitMaxDuration(count int64, maxWait time.Duration) bool
	Take(count int64) time.Duration
	TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool)
	TakeAvailable(count int64) int64
	Available() int64
	Capacity()
	Rate()
}
