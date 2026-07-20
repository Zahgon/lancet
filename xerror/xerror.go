package xerror

import (
	"fmt"
)

type XError struct {
	id      string
	message string
	stack   *stack
	cause   error
	values  map[string]any
}

func New(format string, args ...any) *XError { _ = "STUB: not implemented"; return nil }

func Wrap(cause error, message ...any) *XError { _ = "STUB: not implemented"; return nil }

func Unwrap(err error) *XError { _ = "STUB: not implemented"; return nil }

func newXError() *XError { _ = "STUB: not implemented"; return nil }

func (e *XError) copy(dest *XError) { _ = "STUB: not implemented"; return }

func (e *XError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *XError) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (e *XError) Wrap(cause error) *XError { _ = "STUB: not implemented"; return nil }

func (e *XError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *XError) With(key string, value any) *XError { _ = "STUB: not implemented"; return nil }

func (e *XError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e *XError) Id(id string) *XError { _ = "STUB: not implemented"; return nil }

func (e *XError) Values() map[string]any { _ = "STUB: not implemented"; return nil }

type errInfo struct {
	Message    string         `json:"message"`
	Id         string         `json:"id"`
	StackTrace []*Stack       `json:"stacktrace"`
	Cause      error          `json:"cause"`
	Values     map[string]any `json:"values"`
}

func (e *XError) Info() *errInfo { _ = "STUB: not implemented"; return nil }

func TryUnwrap[T any](val T, err error) T { _ = "STUB: not implemented"; return *new(T) }
