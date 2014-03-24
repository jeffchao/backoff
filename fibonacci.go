package backoff

import (
	"time"
)

type FibonacciBackoff struct {
	Retries    int
	MaxRetries int
	Delay      time.Duration
	Interval   time.Duration
}

func Fibonacci() *FibonacciBackoff {
	return &FibonacciBackoff{
		Retries:    0,
		MaxRetries: 5,
		Delay:      time.Duration(0),
  }
}
