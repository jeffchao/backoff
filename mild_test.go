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
