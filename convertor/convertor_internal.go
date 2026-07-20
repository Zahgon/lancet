package convertor

import "reflect"

type cloner struct {
	ptrs map[reflect.Type]map[uintptr]reflect.Value
}

func (c *cloner) clone(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (c *cloner) cloneArray(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (c *cloner) cloneMap(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func isNillable(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func (c *cloner) clonePtr(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func (c *cloner) cloneStruct(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
