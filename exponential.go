package backoff

import (
	"math"
	"time"
)

type ExponentialBackoff struct {
	Retries    int
	MaxRetries int
	Delay      time.Duration
	Interval   time.Duration // time.Second, time.Millisecond, etc.
}

// Creates a new instance of ExponentialBackoff.
func Exponential() *ExponentialBackoff {
	return &ExponentialBackoff{
		Retries:    0,
		MaxRetries: 5,
		Delay:      time.Duration(0),
		Interval:   time.Duration(1 * time.Second),
	}
}

/*
Gets the next backoff delay. This method will increment the retries and check
if the maximum number of retries has been met. If this condition is satisfied, then
the function will return. Otherwise, the next backoff delay will be computed.

The exponential backoff delay is computed as follows:
`n = 2^c - 1` where `n` is the backoff delay and `c` is the number of retries.

Example, given a 1 second interval:

  Retry #        Backoff delay (in seconds)
    0                   0
    1                   1
    2                   3
    3                   7
    4                   15
*/
func (self *ExponentialBackoff) Next() bool {
	self.Retries++

	if self.Retries >= self.MaxRetries {
		return false
	}

	self.Delay = time.Duration(math.Pow(2, float64(self.Retries))-1) * self.Interval

	return true
}

/*
Retries a function until the maximum number of retries is met. This method expects
the function `f` to return an error. If the failure condition is met, this method
will surface the error outputted from `f`, otherwise nil will be returned as normal.
*/
func (self *ExponentialBackoff) Retry(f func() error) error {
	err := f()

	if err == nil {
		return nil
	}

	for self.Next() {
		if err = f(); err == nil {
			return nil
		}

		time.Sleep(self.Delay)
	}

	return err
}

// Resets the retry count and the backoff delay back to its initial state.
func (self *ExponentialBackoff) Reset() {
	self.Retries = 0
	self.Delay = time.Duration(0 * time.Second)
}
