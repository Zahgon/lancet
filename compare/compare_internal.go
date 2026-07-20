package compare

import (
	"math/big"
	"reflect"
)

func compareValue(operator string, left, right any) bool { _ = "STUB: not implemented"; return false }

func compareRefValue(operator string, leftObj, rightObj any, kind reflect.Kind) bool {
	_ = "STUB: not implemented"
	return false
}

func objectsAreEqualValues(expected, actual interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func objectsAreEqual(expected, actual interface{}) bool { _ = "STUB: not implemented"; return false }

func compareBasicValue(operator string, leftValue, rightValue any) bool {
	_ = "STUB: not implemented"
	return false
}

func compareBigInt(operator string, left, right *big.Int) bool {
	_ = "STUB: not implemented"
	return false
}

func compareFloats(operator string, left, right float64) bool {
	_ = "STUB: not implemented"
	return false
}

func compareStrings(operator string, left, right string) bool {
	_ = "STUB: not implemented"
	return false
}

func compareBools(operator string, left, right bool) bool { _ = "STUB: not implemented"; return false }
