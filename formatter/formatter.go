package formatter

import (
	"io"

	"golang.org/x/exp/constraints"
)

func Comma[T constraints.Float | constraints.Integer | string](value T, prefixSymbol string) string {
	_ = "STUB: not implemented"
	return ""
}

func Pretty(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func PrettyToWriter(v any, out io.Writer) error { _ = "STUB: not implemented"; return nil }
