package structs

import (
	"reflect"
)

var defaultTagName = "json"

type Struct struct {
	raw     any
	rtype   reflect.Type
	rvalue  reflect.Value
	TagName string
}

func New(value any, tagName ...string) *Struct { _ = "STUB: not implemented"; return nil }

func (s *Struct) ToMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Struct) Fields() []*Field { _ = "STUB: not implemented"; return nil }

func (s *Struct) Field(name string) (*Field, bool) { _ = "STUB: not implemented"; return nil, false }

func (s *Struct) IsStruct() bool { _ = "STUB: not implemented"; return false }

func ToMap(v any) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Struct) TypeName() string { _ = "STUB: not implemented"; return "" }
