package ratelimit

import (
	"math"
	"time"
)

type FakeBucket struct {}

// NewBucket returns a new fake token bucket.
func NewFakeBucket() *FakeBucket {
	return &FakeBucket{}
}

// Wait takes count tokens from the bucket, waiting until they are available.
func (tb *FakeBucket) Wait(count int64) {}

// WaitMaxDuration is like Wait.
func (tb *FakeBucket) WaitMaxDuration(count int64, maxWait time.Duration) bool {
	return true
}

// Take takes count tokens from the bucket without blocking. It returns
// the time that the caller should wait until the tokens are actually
// available.
//
// Note that if the request is irrevocable - there is no way to return
// tokens to the bucket once this method commits us to taking them.
func (tb *FakeBucket) Take(count int64) time.Duration {
	return 0
}

// TakeMaxDuration is like Take, except that
// it will only take tokens from the bucket if the wait
// time for the tokens is no greater than maxWait.
//
// If it would take longer than maxWait for the tokens
// to become available, it does nothing and reports false,
// otherwise it returns the time that the caller should
// wait until the tokens are actually available, and reports
// true.
func (tb *FakeBucket) TakeMaxDuration(count int64, maxWait time.Duration) (time.Duration, bool) {
	return 0, true
}

// TakeAvailable takes up to count immediately available tokens from the
// bucket. It returns the number of tokens removed, or zero if there are
// no available tokens. It does not block.
func (tb *FakeBucket) TakeAvailable(count int64) int64 {
	return math.MaxInt64
}

// Available returns the number of available tokens.
func (tb *FakeBucket) Available() int64 {
	return math.MaxInt64
}

// Capacity returns the capacity that the bucket was created with.
func (tb *FakeBucket) Capacity() int64 {
	return math.MaxInt64
}

// Rate returns the fill rate of the bucket, in tokens per second.
func (tb *FakeBucket) Rate() float64 {
	return math.MaxFloat64
}