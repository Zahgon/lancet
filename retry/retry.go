package retry

import (
	"context"
	"time"
)

const (
	DefaultRetryTimes = 5

	DefaultRetryLinearInterval = time.Second * 3
)

type RetryConfig struct {
	context         context.Context
	retryTimes      uint
	backoffStrategy BackoffStrategy
}

type RetryFunc func() error

type Option func(*RetryConfig)

func RetryTimes(n uint) Option { _ = "STUB: not implemented"; return *new(Option) }

func RetryWithCustomBackoff(backoffStrategy BackoffStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func RetryWithLinearBackoff(interval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func RetryWithExponentialWithJitterBackoff(interval time.Duration, base uint64, maxJitter time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Context(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

func Retry(retryFunc RetryFunc, opts ...Option) error { _ = "STUB: not implemented"; return nil }

type BackoffStrategy interface {
	CalculateInterval() time.Duration
}

type linear struct {
	interval time.Duration
}

func (l *linear) CalculateInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type exponentialWithJitter struct {
	base      time.Duration
	interval  time.Duration
	maxJitter time.Duration
}

func (e *exponentialWithJitter) CalculateInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type shiftExponentialWithJitter struct {
	interval  time.Duration
	maxJitter time.Duration
	shifter   uint64
}

func (e *shiftExponentialWithJitter) CalculateInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func jitter(maxJitter time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
