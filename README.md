backoff
=======

[![Build Status](https://travis-ci.org/jeffchao/backoff.svg?branch=master)](https://travis-ci.org/jeffchao/backoff)

Backoff algorithms are used to space out repeated retries of the same block of code. By gradually decreasing the retry rate, backoff algorithms aim to avoid congestion. This Go backoff library provides a set of backoff algorithms well-known in the computer networks research community. These algorithms are acceptable for use in generalized computation. All algorithms use a uniform distribution of backoff times.

This library provides the following backoff algorithms:

1. `Exponential`
2. `Fibonacci`
3. `MILD`
4. `PLEB`
5. `PFB`

each with properties:

```go
type <all backoff algorithms> struct {
        Retries    int
        MaxRetries int
        Delay      time.Duration
        Interval   time.Duration    // time.Second, time.Millisecond, etc.
}
```



all of which are based off a `Backoff` interface defined as:

```go
type Backoff interface {
        Next() bool
        Retry(func() error) error  // Surface the resulting error to the user.
        Reset()
}
```

so you may define additional backoff algorithms of your choice.

### Exponential Backoff

Gradually decrease the retry rate in an exponential manner with base 2. The algorithm is defined as `n = 2^c - 1` where `n` is the backoff delay and `c` is the number of retries.

Example usage:

```go
e := backoff.Exponential()
e.Interval = 1 * time.Millisecond
e.MaxRetries = 5

fooFunc := func() error {
        // Do some work here
}

err := e.Retry(fooFunc)
e.Reset()
```

### Fibonacci Backoff

Gradually decrease the retry rate using a fibonacci sequence. The algorithm is defined as `n = fib(c - 1) + fib(c - 2); f(0) = 0, f(1) = 1; n >= 0` where `n` is the backoff delay and `c` is the retry slot.

```go
f := backoff.Fibonacci()
f.Interval = 1 * time.Millisecond
f.MaxRetries = 5

fooFunc := func() error {
        // Do some work here
}

err := f.Retry(fooFunc)
f.Reset()
```

Additionally, the `FibonacciBackoff` struct exposes its delay slots in the form of a slice of `time.Duration`.

```go
log.Printf("%+v", f.Slots)
```

### MILD - Multiplicative Increase and Linear Decrease

WIP

```go
WIP
```

### PLEB - Pessimistic Linear-Exponential Backoff

WIP

```go
WIP
```

### PFB - Pessimistic Fibonacci Backoff

WIP

```go
WIP
```

## Author

Jeff Chao, @thejeffchao, http://thejeffchao.com
