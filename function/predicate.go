package function

func And[T any](predicates ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

func Nand[T any](predicates ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

func Negate[T any](predicate func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

func Or[T any](predicates ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

func Nor[T any](predicates ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }

func Xnor[T any](predicates ...func(T) bool) func(T) bool { _ = "STUB: not implemented"; return nil }
