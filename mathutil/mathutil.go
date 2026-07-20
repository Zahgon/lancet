package mathutil

import (
	"golang.org/x/exp/constraints"
)

func Exponent(x, n int64) int64 { _ = "STUB: not implemented"; return 0 }

func Fibonacci(first, second, n int) int { _ = "STUB: not implemented"; return 0 }

func Factorial(n uint) uint { _ = "STUB: not implemented"; return 0 }

func Percent(val, total float64, n int) float64 { _ = "STUB: not implemented"; return 0 }

func RoundToString[T constraints.Float | constraints.Integer](x T, n int) string {
	_ = "STUB: not implemented"
	return ""
}

func RoundToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func TruncRound[T constraints.Float | constraints.Integer](x T, n int) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func FloorToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func FloorToString[T constraints.Float | constraints.Integer](x T, n int) string {
	_ = "STUB: not implemented"
	return ""
}

func CeilToFloat[T constraints.Float | constraints.Integer](x T, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func CeilToString[T constraints.Float | constraints.Integer](x T, n int) string {
	_ = "STUB: not implemented"
	return ""
}

func Max[T constraints.Ordered](items ...T) T { _ = "STUB: not implemented"; return *new(T) }

func MaxBy[T any](slice []T, comparator func(T, T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func Min[T constraints.Ordered](items ...T) T { _ = "STUB: not implemented"; return *new(T) }

func MinBy[T any](slice []T, comparator func(T, T) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func Sum[T constraints.Integer | constraints.Float](numbers ...T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func Average[T constraints.Integer | constraints.Float](numbers ...T) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Range[T constraints.Integer | constraints.Float](start T, count int) []T {
	_ = "STUB: not implemented"
	return nil
}

func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T {
	_ = "STUB: not implemented"
	return nil
}

func AngleToRadian(angle float64) float64 { _ = "STUB: not implemented"; return 0 }

func RadianToAngle(radian float64) float64 { _ = "STUB: not implemented"; return 0 }

func PointDistance(x1, y1, x2, y2 float64) float64 { _ = "STUB: not implemented"; return 0 }

func IsPrime(n int) bool { _ = "STUB: not implemented"; return false }

func GCD[T constraints.Integer](integers ...T) T { _ = "STUB: not implemented"; return *new(T) }

func gcd[T constraints.Integer](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

func LCM[T constraints.Integer](integers ...T) T { _ = "STUB: not implemented"; return *new(T) }

func lcm[T constraints.Integer](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

func Cos(radian float64, precision ...int) float64 { _ = "STUB: not implemented"; return 0 }

func Sin(radian float64, precision ...int) float64 { _ = "STUB: not implemented"; return 0 }

func Log(n, base float64) float64 { _ = "STUB: not implemented"; return 0 }

func Abs[T constraints.Integer | constraints.Float](x T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

func Div[T constraints.Float | constraints.Integer](x T, y T) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Variance[T constraints.Float | constraints.Integer](numbers []T) float64 {
	_ = "STUB: not implemented"
	return 0
}

func StdDev[T constraints.Float | constraints.Integer](numbers []T) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Permutation(n, k uint) uint { _ = "STUB: not implemented"; return 0 }

func Combination(n, k uint) uint { _ = "STUB: not implemented"; return 0 }
