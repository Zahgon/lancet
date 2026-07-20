package xerror

import (
	"fmt"
)

type Stack struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func (e *XError) Stacks() []*Stack { _ = "STUB: not implemented"; return nil }

func (e *XError) StackTrace() StackTrace { _ = "STUB: not implemented"; return *new(StackTrace) }

type frame uintptr
type stack []uintptr

type StackTrace []frame

func (f frame) pc() uintptr { _ = "STUB: not implemented"; return 0 }

func (f frame) file() string { _ = "STUB: not implemented"; return "" }

func (f frame) line() int { _ = "STUB: not implemented"; return 0 }

func (f frame) name() string { _ = "STUB: not implemented"; return "" }

func (f frame) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (f frame) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (st StackTrace) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (st StackTrace) formatSlice(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (s *stack) Format(st fmt.State, verb rune) { _ = "STUB: not implemented"; return }

func (s *stack) StackTrace() StackTrace { _ = "STUB: not implemented"; return *new(StackTrace) }

func callers() *stack { _ = "STUB: not implemented"; return nil }

func funcname(name string) string { _ = "STUB: not implemented"; return "" }
