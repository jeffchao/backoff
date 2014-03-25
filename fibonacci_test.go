package backoff

import (
	"testing"
	"time"
)

func TestNextFibonacciBackoff(t *testing.T) {
	f := Fibonacci()
	f.Interval = 1 * time.Second
	f.MaxRetries = 9

	expectedRetries := []int{1, 2, 3, 4, 5, 6, 7}
	expectedDelays := []time.Duration{1, 1, 2, 3, 5, 8, 13, 13, 13}
	for i, v := range expectedDelays {
		expectedDelays[i] = v * time.Second
	}

	for i, expected := range expectedRetries {
		f.Next()
		assertEquals(t, expected, f.Retries)
		assertEquals(t, expectedDelays[i], f.Delay)
	}
}
