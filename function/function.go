package function

import (
	"reflect"
	"time"
)

func After(n int, fn any) func(args ...any) []reflect.Value { _ = "STUB: not implemented"; return nil }

func Before(n int, fn any) func(args ...any) []reflect.Value { _ = "STUB: not implemented"; return nil }

type CurryFn[T any] func(...T) T

func (cf CurryFn[T]) New(val T) func(...T) T { _ = "STUB: not implemented"; return nil }

func Compose[T any](fnList ...func(...T) T) func(...T) T { _ = "STUB: not implemented"; return nil }

func Delay(delay time.Duration, fn any, args ...any) { _ = "STUB: not implemented"; return }

func Debounced(fn func(), delay time.Duration) func() { _ = "STUB: not implemented"; return nil }

func Debounce(fn func(), delay time.Duration) (debouncedFn func(), cancelFn func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Throttle(fn func(), interval time.Duration) func() { _ = "STUB: not implemented"; return nil }

func Schedule(duration time.Duration, fn any, args ...any) chan bool {
	_ = "STUB: not implemented"
	return nil
}

func Pipeline[T any](funcs ...func(T) T) func(T) T { _ = "STUB: not implemented"; return nil }

func AcceptIf[T any](predicate func(T) bool, apply func(T) T) func(T) (T, bool) {
	_ = "STUB: not implemented"
	return nil
}

func unsafeInvokeFunc(fn any, args ...any) []reflect.Value { _ = "STUB: not implemented"; return nil }

func mustBeFunction(function any) { _ = "STUB: not implemented"; return }
