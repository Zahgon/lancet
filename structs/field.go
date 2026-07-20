package structs

import (
	"reflect"
)

type Field struct {
	Struct
	field reflect.StructField
	tag   *Tag
}

func newField(v reflect.Value, f reflect.StructField, tagName string) *Field {
	_ = "STUB: not implemented"
	return nil
}

func (f *Field) Tag() *Tag { _ = "STUB: not implemented"; return nil }

func (f *Field) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (f *Field) IsEmbedded() bool { _ = "STUB: not implemented"; return false }

func (f *Field) IsExported() bool { _ = "STUB: not implemented"; return false }

func (f *Field) IsZero() bool { _ = "STUB: not implemented"; return false }

func (f *Field) IsNil() bool { _ = "STUB: not implemented"; return false }

func (f *Field) Name() string { _ = "STUB: not implemented"; return "" }

func (f *Field) Kind() reflect.Kind { _ = "STUB: not implemented"; return *new(reflect.Kind) }

func (f *Field) IsSlice() bool { _ = "STUB: not implemented"; return false }

func (f *Field) IsTargetType(targetType reflect.Kind) bool { _ = "STUB: not implemented"; return false }

func (f *Field) mapValue(value any) any { _ = "STUB: not implemented"; return *new(any) }
