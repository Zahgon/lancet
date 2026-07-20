package pointer

func Of[T any](v T) *T { _ = "STUB: not implemented"; return nil }

func Unwrap[T any](p *T) T { _ = "STUB: not implemented"; return *new(T) }

func UnwarpOr[T any](p *T, fallback T) T { _ = "STUB: not implemented"; return *new(T) }

func UnwarpOrDefault[T any](p *T) T { _ = "STUB: not implemented"; return *new(T) }

func UnwrapOr[T any](p *T, fallback ...T) T { _ = "STUB: not implemented"; return *new(T) }

func ExtractPointer(value any) any { _ = "STUB: not implemented"; return *new(any) }

func IsNil(i interface{}) bool { _ = "STUB: not implemented"; return false }
