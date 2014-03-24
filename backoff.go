// Utility to retry functions with support from a variety of backoff algorithms.
package backoff

type Backoff interface {
	// Compute and return the next backoff delay.
	Next() bool
	// Retry a function until an error or backoff delay condition is met.
	Retry(func() error) error
	// Reset the backoff delay to its initial value.
	Reset()
}
