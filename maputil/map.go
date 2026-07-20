package maputil

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

func Keys[K comparable, V any](m map[K]V) []K { _ = "STUB: not implemented"; return nil }

func Values[K comparable, V any](m map[K]V) []V { _ = "STUB: not implemented"; return nil }

func KeysBy[K comparable, V any, T any](m map[K]V, mapper func(item K) T) []T {
	_ = "STUB: not implemented"
	return nil
}

func ValuesBy[K comparable, V any, T any](m map[K]V, mapper func(item V) T) []T {
	_ = "STUB: not implemented"
	return nil
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V { _ = "STUB: not implemented"; return nil }

func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	_ = "STUB: not implemented"
	return
}

func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func FilterByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func OmitBy[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func OmitByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func OmitByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func Intersect[K comparable, V any](maps ...map[K]V) map[K]V { _ = "STUB: not implemented"; return nil }

func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V { _ = "STUB: not implemented"; return nil }

func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool {
	_ = "STUB: not implemented"
	return false
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] { _ = "STUB: not implemented"; return nil }

func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	_ = "STUB: not implemented"
	return nil
}

func Transform[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	_ = "STUB: not implemented"
	return nil
}

func MapKeys[K comparable, V any, T comparable](m map[K]V, iteratee func(key K, value V) T) map[T]V {
	_ = "STUB: not implemented"
	return nil
}

func MapValues[K comparable, V any, T any](m map[K]V, iteratee func(key K, value V) T) map[K]T {
	_ = "STUB: not implemented"
	return nil
}

func HasKey[K comparable, V any](m map[K]V, key K) bool { _ = "STUB: not implemented"; return false }

func MapToStruct(m map[string]any, structObj any) error { _ = "STUB: not implemented"; return nil }

func setStructField(structObj any, fieldName string, fieldValue any) error {
	_ = "STUB: not implemented"
	return nil
}

func getFieldNameByJsonTag(structObj any, jsonTag string) string {
	_ = "STUB: not implemented"
	return ""
}

func ToSortedSlicesDefault[K constraints.Ordered, V any](m map[K]V) ([]K, []V) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToSortedSlicesWithComparator[K comparable, V any](m map[K]V, comparator func(a, b K) bool) ([]K, []V) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOrSet[K comparable, V any](m map[K]V, key K, value V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

func SortByKey[K constraints.Ordered, V any](m map[K]V, less func(a, b K) bool) (sortedKeysMap map[K]V) {
	_ = "STUB: not implemented"
	return nil
}

var mapHandlers = map[reflect.Kind]func(reflect.Value, reflect.Value) error{
	reflect.String:     convertNormal,
	reflect.Int:        convertNormal,
	reflect.Int16:      convertNormal,
	reflect.Int32:      convertNormal,
	reflect.Int64:      convertNormal,
	reflect.Uint:       convertNormal,
	reflect.Uint16:     convertNormal,
	reflect.Uint32:     convertNormal,
	reflect.Uint64:     convertNormal,
	reflect.Float32:    convertNormal,
	reflect.Float64:    convertNormal,
	reflect.Uint8:      convertNormal,
	reflect.Int8:       convertNormal,
	reflect.Struct:     convertNormal,
	reflect.Complex64:  convertNormal,
	reflect.Complex128: convertNormal,
}

var _ = func() struct{} {
	mapHandlers[reflect.Map] = convertMap
	mapHandlers[reflect.Array] = convertSlice
	mapHandlers[reflect.Slice] = convertSlice
	return struct{}{}
}()

func MapTo(src any, dst any) error { _ = "STUB: not implemented"; return nil }

func convertNormal(src reflect.Value, dst reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func convertSlice(src reflect.Value, dst reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func convertMap(src reflect.Value, dst reflect.Value) error { _ = "STUB: not implemented"; return nil }

func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	_ = "STUB: not implemented"
	return *new(V)
}

func FindValuesBy[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) []V {
	_ = "STUB: not implemented"
	return nil
}

func ToMarkdownTable(data []map[string]interface{}, headerMap map[string]string, columnOrder []string) string {
	_ = "STUB: not implemented"
	return ""
}

func formatValue(v interface{}) string { _ = "STUB: not implemented"; return "" }

func commaInt64(n int64) string { _ = "STUB: not implemented"; return "" }

func commaUint64(n uint64) string { _ = "STUB: not implemented"; return "" }

func addCommas(s string) string { _ = "STUB: not implemented"; return "" }
