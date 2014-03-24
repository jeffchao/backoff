package backoff

import (
	"errors"
	"testing"
	"time"
)

func TestNextExponentialBackoff(t *testing.T) {
	e := Exponential()
	e.Interval = 1 * time.Second
	e.MaxRetries = 5

	expectedRetries := []int{1, 2, 3, 4, 5, 6, 7}
	expectedDelays := []time.Duration{1, 3, 7, 15, 15, 15, 15}
	for i, v := range expectedDelays {
		expectedDelays[i] = v * time.Second
	}

	for i, expected := range expectedRetries {
		e.Next()
		assertEquals(t, expected, e.Retries)
		assertEquals(t, expectedDelays[i], e.Delay)
	}
}

func TestRetryExponential(t *testing.T) {
	e := Exponential()
	e.Interval = 1 * time.Millisecond
	e.MaxRetries = 5

	retries := 0

	test := func() error {
		retries++
		return errors.New("an error occurred")
	}
	e.Retry(test)

	if retries != e.Retries {
		t.Errorf("retries count does not match e.Retries: got %d, expected %d", retries, e.Retries)
	}

	if e.Retries > e.MaxRetries {
		t.Errorf("overflow: retries %d greater than maximum retries %d", e.Retries, e.MaxRetries)
	}

	e.Reset()
	retries = 0

	test = func() error {
		retries++
		return nil
	}

	err := e.Retry(test)

	if e.Retries > 0 && err != nil {
		t.Errorf("failure in retry logic. expected success but got a failure: %+v", err)
	}
}

func TestResetExponential(t *testing.T) {
	e := Exponential()
	e.Interval = 1 * time.Second
	e.MaxRetries = 5

	e.Next()
	assertEquals(t, e.Retries, 1)
	assertEquals(t, e.Delay, time.Duration(1*time.Second))
	e.Reset()
	assertEquals(t, e.Retries, 0)
	assertEquals(t, e.Delay, time.Duration(0*time.Second))
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("error. got %d, expected: %d", actual, expected)
	}
}
