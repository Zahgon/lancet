package condition

func Bool[T any](value T) bool { _ = "STUB: not implemented"; return false }

func reflectValue(vp any) bool { _ = "STUB: not implemented"; return false }

func And[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Or[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Xor[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Nor[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Xnor[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Nand[T, U any](a T, b U) bool { _ = "STUB: not implemented"; return false }

func Ternary[T, U any](isTrue T, ifValue U, elseValue U) U {
	_ = "STUB: not implemented"
	return *new(U)
}

func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U {
	_ = "STUB: not implemented"
	return *new(U)
}
