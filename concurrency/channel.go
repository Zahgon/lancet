package concurrency

import (
	"context"
)

type Channel[T any] struct {
}

func NewChannel[T any]() *Channel[T] { _ = "STUB: not implemented"; return nil }

func (c *Channel[T]) Generate(ctx context.Context, values ...T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) Repeat(ctx context.Context, values ...T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) RepeatFn(ctx context.Context, fn func() T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) Take(ctx context.Context, valueStream <-chan T, number int) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) FanIn(ctx context.Context, channels ...<-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) Tee(ctx context.Context, in <-chan T) (<-chan T, <-chan T) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Channel[T]) Bridge(ctx context.Context, chanStream <-chan <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}

func (c *Channel[T]) Or(channels ...<-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

func (c *Channel[T]) OrDone(ctx context.Context, channel <-chan T) <-chan T {
	_ = "STUB: not implemented"
	return nil
}
