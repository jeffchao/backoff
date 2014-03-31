package backoff

import (
	"time"
)

type MILDBackoff struct {
	Retries    int
	MaxRetries int
	Delay      time.Duration
	Interval   time.Duration // time.Second, time.Millisecond, etc.
	Slots      []time.Duration
}

// MILD creates a new instance of MILDBackoff.
func MILD() *MILDBackoff {
	return &MILDBackoff{
		Retries:    0,
		MaxRetries: 5,
		Delay:      time.Duration(0),
	}
}

/*
Next gets the next backoff delay. This method will increment the retries and check
if the maximum number of retries has been met. If this condition is satisfied, then
the function will return. Otherwise, the next backoff delay will be computed.

The MILD backoff delay is computed as follows:
`n = min(1.5 * n, len(slots)) upon failure; n = max(slots(c) - 1, 0) upon success;
n(0) = 0, n(1) = 1`
where `n` is the backoff delay, `c` is the retry slot, and `slots` is an array of retry delays.

This means a method must repeatedly succeed until `slots` is empty for the overall
backoff mechanism to terminate. Conversely, a repeated number of failures until the
maximum number of retries will result in a failure.

Example, given a 1 second interval, with max retries of 5:

  Retry #        Backoff delay (in seconds)       success/fail
    1                   1                             fail
    2                   1.5                           fail
    3                   1                             success
    4                   1.5                           fail
    5                   2.25                          fail

  Retry #        Backoff delay (in seconds)       success/fail
    1                   1                             fail
    2                   1.5                           fail
    3                   1                             success
    4                   0                             success
    5                   -                             success
*/
func (self *MILDBackoff) Next() bool {
	if self.Retries >= self.MaxRetries {
		return false
	}

	self.increment()

	return true
}

/*
Retry will retry a function until the maximum number of retries is met. This method expects
the function `f` to return an error. If the failure condition is met, this method
will surface the error outputted from `f`, otherwise nil will be returned as normal.
*/
func (self *MILDBackoff) Retry(f func() error) error {
	err := f()

	if err == nil {
		return nil
	}

	for self.Next() {
		if err := f(); err == nil {
			if len(self.Slots) == 0 {
				return nil
			}
			self.decrement()
		}

		time.Sleep(self.Delay)
	}

	return err
}

func (self *MILDBackoff) increment() {
	self.Retries++

	if self.Delay == 0 {
		self.Delay = time.Duration(1 * self.Interval)
	} else {
		self.Delay = self.Delay + (self.Delay / 2)
	}

	self.Slots = append(self.Slots, self.Delay)
}

func (self *MILDBackoff) decrement() {
	copy(self.Slots[len(self.Slots)-1:], self.Slots[len(self.Slots):])
	self.Slots[len(self.Slots)-1] = time.Duration(0 * self.Interval)
	self.Slots = self.Slots[:len(self.Slots)-1]
	self.Retries--
	self.Delay = self.Slots[len(self.Slots)-1]
}

// Reset will reset the retry count, the backoff delay, and backoff slots back to its initial state.
func (self *MILDBackoff) Reset() {
	self.Retries = 0
	self.Delay = time.Duration(0 * time.Second)
	self.Slots = nil
	self.Slots = make([]time.Duration, 0, self.MaxRetries)
}
