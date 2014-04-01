package backoff

import (
	"testing"
	"time"
)

func TestNextFibonacciBackoff(t *testing.T) {
	f := Fibonacci()
	f.Interval = 1 * time.Second
	f.MaxRetries = 7

	expectedRetries := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedDelays := []time.Duration{0, 1, 1, 2, 3, 5, 8, 8, 8}
	for i, v := range expectedDelays {
		expectedDelays[i] = v * time.Second
	}

	for i, expected := range expectedRetries {
		f.Next()
		assertEquals(t, expected, f.Retries)
		assertEquals(t, expectedDelays[i], f.Delay)
	}
}
