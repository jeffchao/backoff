package backoff

import (
	"testing"
	"time"
)

func TestNextMILDBackoff(t *testing.T) {
	f := MILD()
	f.Interval = 1 * time.Second
	f.MaxRetries = 5

	expectedRetries := []int{1, 2, 3, 4, 5}
	expectedDelays := []time.Duration{}

	expectedDelays = append(expectedDelays, time.Duration(1*f.Interval))
	expectedDelays = append(expectedDelays, time.Duration(expectedDelays[len(expectedDelays)-1]+(expectedDelays[len(expectedDelays)-1]/2)))
	expectedDelays = append(expectedDelays, time.Duration(expectedDelays[len(expectedDelays)-1]+(expectedDelays[len(expectedDelays)-1]/2)))
	expectedDelays = append(expectedDelays, time.Duration(expectedDelays[len(expectedDelays)-1]+(expectedDelays[len(expectedDelays)-1]/2)))
	expectedDelays = append(expectedDelays, time.Duration(expectedDelays[len(expectedDelays)-1]+(expectedDelays[len(expectedDelays)-1]/2)))
	expectedDelays = append(expectedDelays, expectedDelays[len(expectedDelays)-1])
	expectedDelays = append(expectedDelays, expectedDelays[len(expectedDelays)-1])

	for i, expected := range expectedRetries {
		f.Next()
		assertEquals(t, expected, f.Retries)
		assertEquals(t, expectedDelays[i], f.Delay)
	}
}

func TestResetMILD(t *testing.T) {
	m := MILD()
	m.Interval = 1 * time.Second
	m.MaxRetries = 5

	for i := 0; i < 4; i++ {
		m.Next()
	}

	assertEquals(t, m.Retries, 4)
	m.Reset()
	assertEquals(t, m.Retries, 0)
	assertEquals(t, m.Delay, time.Duration(0*time.Second))
}

func TestDecrementMILD(t *testing.T) {
	m := MILD()
	m.Interval = 1 * time.Second
	m.MaxRetries = 5

	m.Next()
	assertEquals(t, m.Retries, 1)
	assertEquals(t, len(m.Slots), 1)
	m.Next()
	assertEquals(t, m.Retries, 2)
	assertEquals(t, len(m.Slots), 2)
	m.decrement()
	assertEquals(t, m.Retries, 1)
	assertEquals(t, len(m.Slots), 1)
	m.decrement()
	assertEquals(t, m.Retries, 0)
	assertEquals(t, len(m.Slots), 0)

}
