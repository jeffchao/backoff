package backoff

import (
	"math"
	"time"
)

type ExponentialBackoff struct {
	Retries             int
	MaxRetries          int
	Delay               time.Duration
	Interval            time.Duration
}

func Exponential() *ExponentialBackoff {
	return &ExponentialBackoff{
		Retries:             0,
		MaxRetries:          5,
		Delay:               time.Duration(0),
		Interval:            time.Duration(1 * time.Second),
	}
}

func (self *ExponentialBackoff) Next() bool {
	self.Retries++

	if self.Retries >= self.MaxRetries {
		return false
	}

	self.Delay = time.Duration(math.Pow(2, float64(self.Retries))-1) * self.Interval

	return true
}

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

func (self *ExponentialBackoff) Reset() {
	self.Retries = 0
	self.Delay = time.Duration(0 * time.Second)
}
