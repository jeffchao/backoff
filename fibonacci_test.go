package backoff

import (
	"errors"
	"testing"
	"time"
)

func TestNextFibonacciBackoff(t *testing.T) {
	f := Fibonacci()
	f.Interval = 1 * time.Second
	f.MaxRetries = 7

	expectedRetries := []int{1, 2, 3, 4, 5, 6, 7, 7, 7}
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

func TestRetryFibonacci(t *testing.T) {
	f := Fibonacci()
	f.Interval = 1 * time.Millisecond
	f.MaxRetries = 5

	test := func() error {
		return errors.New("an error occurred")
	}
	f.Retry(test)

	if f.Retries != f.MaxRetries {
		t.Errorf("f.Retries does not match f.MaxRetries: got %d, expected %d", f.Retries, f.MaxRetries)
	}

	if f.Retries > f.MaxRetries {
		t.Errorf("overflow: retries %d greater than maximum retries %d", f.Retries, f.MaxRetries)
	}

	test = func() error {
		return nil
	}
	f.Reset()
	err := f.Retry(test)

	if err != nil {
		t.Errorf("failure in retry logic. expected success but got a failure: %+v", err)
	}

	retries := 0
	f.Reset()
	test = func() error {
		if retries == 0 {
			retries++
			return errors.New("an error occurred")
		}
		return nil
	}

	f.Retry(test)

	if err != nil {
		t.Errorf("failure in retry logic. expected success but got a failure: %+v", err)
	}
}

func TestResetFibonacci(t *testing.T) {
	f := Fibonacci()
	f.Interval = 1 * time.Second
	f.MaxRetries = 5

	for i := 0; i < 4; i++ {
		f.Next()
	}

	assertEquals(t, f.Retries, 4)
	assertEquals(t, f.Delay, time.Duration(2*time.Second))
	f.Reset()
	assertEquals(t, f.Retries, 0)
	assertEquals(t, f.Delay, time.Duration(0*time.Second))
}
