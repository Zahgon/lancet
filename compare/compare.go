package compare

import (
	"reflect"
	"time"

	"golang.org/x/exp/constraints"
)

const (
	equal          = "eq"
	lessThan       = "lt"
	greaterThan    = "gt"
	lessOrEqual    = "le"
	greaterOrEqual = "ge"
)

var (
	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

func Equal(left, right any) bool { _ = "STUB: not implemented"; return false }

func EqualValue(left, right any) bool { _ = "STUB: not implemented"; return false }

func LessThan(left, right any) bool { _ = "STUB: not implemented"; return false }

func GreaterThan(left, right any) bool { _ = "STUB: not implemented"; return false }

func LessOrEqual(left, right any) bool { _ = "STUB: not implemented"; return false }

func GreaterOrEqual(left, right any) bool { _ = "STUB: not implemented"; return false }

func InDelta[T constraints.Integer | constraints.Float](left, right T, delta float64) bool {
	_ = "STUB: not implemented"
	return false
}
