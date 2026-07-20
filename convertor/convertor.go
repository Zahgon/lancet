package convertor

import (
	"math/big"
	"reflect"
)

func ToBool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func ToBytes(value any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ToChar(s string) []string { _ = "STUB: not implemented"; return nil }

func ToChannel[T any](array []T) <-chan T { _ = "STUB: not implemented"; return nil }

func ToString(value any) string { _ = "STUB: not implemented"; return "" }

func ToJson(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ToFloat(value any) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToInt(value any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func ToPointer[T any](value T) *T { _ = "STUB: not implemented"; return nil }

func ToPointers[T any](values []T) []*T { _ = "STUB: not implemented"; return nil }

func FromPointer[T any](ptr *T) T { _ = "STUB: not implemented"; return *new(T) }

func FromPointers[T any](pointers []*T) []T { _ = "STUB: not implemented"; return nil }

func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func StructToMap(value any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T {
	_ = "STUB: not implemented"
	return nil
}

func ColorHexToRGB(colorHex string) (red, green, blue int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func ColorRGBToHex(red, green, blue int) string { _ = "STUB: not implemented"; return "" }

func EncodeByte(data any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func DecodeByte(data []byte, target any) error { _ = "STUB: not implemented"; return nil }

func DeepClone[T any](src T) T { _ = "STUB: not implemented"; return *new(T) }

func CopyProperties[T, U any](dst T, src U) error { _ = "STUB: not implemented"; return nil }

func ToInterface(v reflect.Value) (value interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func Utf8ToGbk(bs []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GbkToUtf8(bs []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ToStdBase64(value any) string { _ = "STUB: not implemented"; return "" }

func ToUrlBase64(value any) string { _ = "STUB: not implemented"; return "" }

func ToRawStdBase64(value any) string { _ = "STUB: not implemented"; return "" }

func ToRawUrlBase64(value any) string { _ = "STUB: not implemented"; return "" }

func ToBigInt[T any](v T) (*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }
