package backoff

import "testing"

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("error. got %d, expected: %d", actual, expected)
	}
}
