package vm

import (
	"encoding/gob"
	"fmt"
	"hash/maphash"
)

type TypeBase struct {
	LazyAttributes
	self Type
}

func (m *TypeBase) Self() Type {
	return m.self
}

func (m *TypeBase) SetSelf(self Type) {
	m.self = self
}

func (m *TypeBase) Unimplemented(method string) {
	panic(
		fmt.Errorf(
			"unimplemented method %q for type %q",
			method,
			m.Self().Name(),
		),
	)
}

func (m *TypeBase) Value() interface{} {
	m.Unimplemented("Value")
	return nil
}

func (m *TypeBase) Type() Type {
	return Types.Type
}

func (m *TypeBase) Methods() map[string]Object {
	return map[string]Object{
		"bool":    NewNativeMethodReturningBool(m.Self().ToBool),
		"isnil":   NewNativeMethodReturningBool(m.Self().IsNil),
		"str":     NewNativeMethodReturningString(m.Self().ToString),
		"num":     NewNativeMethodReturningNumber(m.Self().ToNumber),
		"eq":      NewNativeMethodReturningBool(m.Self().Eq),
		"neq":     NewNativeMethodReturningBool(m.Self().NEq),
		"gt":      NewNativeMethodReturningBool(m.Self().Gt),
		"gte":     NewNativeMethodReturningBool(m.Self().GtE),
		"lt":      NewNativeMethodReturningBool(m.Self().Lt),
		"lte":     NewNativeMethodReturningBool(m.Self().LtE),
		"add":     NewNativeMethod(m.Self().Add),
		"sub":     NewNativeMethod(m.Self().Sub),
		"mul":     NewNativeMethod(m.Self().Mul),
		"div":     NewNativeMethod(m.Self().Div),
		"hasattr": NewNativeMethodReturningBool(m.Self().HasAttr),
		"getattr": NewNativeMethod(m.Self().GetAttr),
		"setattr": NewNativeMethod(m.Self().SetAttr),
		"delattr": NewNativeMethod(m.Self().DelAttr),
		"hash":    NewNativeMethodReturningNumber(m.Self().Hash),
		"call":    NewNativeMethod(m.Self().Call),
	}
}

func (m *TypeBase) ToBool(Object, Args) *Bool {
	m.Unimplemented("ToBool")
	return nil
}

func (m *TypeBase) IsNil(Object, Args) *Bool {
	return NewBool(false)
}

func (m *TypeBase) ToString(Object, Args) *String {
	m.Unimplemented("ToString")
	return nil
}

func (m *TypeBase) ToNumber(Object, Args) *Number {
	m.Unimplemented("ToNumber")
	return nil
}

func (m *TypeBase) Eq(Object, Args) *Bool {
	m.Unimplemented("Eq")
	return nil
}

func (m *TypeBase) NEq(Object, Args) *Bool {
	m.Unimplemented("NEq")
	return nil
}

func (m *TypeBase) Gt(Object, Args) *Bool {
	m.Unimplemented("Gt")
	return nil
}

func (m *TypeBase) GtE(Object, Args) *Bool {
	m.Unimplemented("GtE")
	return nil
}

func (m *TypeBase) Lt(Object, Args) *Bool {
	m.Unimplemented("Lt")
	return nil
}

func (m *TypeBase) LtE(Object, Args) *Bool {
	m.Unimplemented("LtE")
	return nil
}

func (m *TypeBase) Add(Object, Args) Object {
	m.Unimplemented("Add")
	return nil
}

func (m *TypeBase) Sub(Object, Args) Object {
	m.Unimplemented("Sub")
	return nil
}

func (m *TypeBase) Mul(Object, Args) Object {
	m.Unimplemented("Mul")
	return nil
}

func (m *TypeBase) Div(Object, Args) Object {
	m.Unimplemented("Div")
	return nil
}

func (m *TypeBase) HasAttr(self Object, args Args) *Bool {
	var attr *String

	args.ExpectsExactly(1)
	args.Parse(&attr)

	return NewBool(self.Attrs().Has(attr))
}

func (m *TypeBase) GetAttr(self Object, args Args) Object {
	var attr *String

	args.ExpectsExactly(1)
	args.Parse(&attr)

	if self.Attrs().Has(attr) {
		return self.Attrs().Get(attr)
	}

	if method, ok := m.Methods()[attr.String()]; ok {
		return method.(*NativeMethod).Bind(self)
	}

	return NewNil()
}

func (m *TypeBase) SetAttr(self Object, args Args) Object {
	var (
		attr  *String
		value Object
	)

	args.ExpectsExactly(2)
	args.Parse(&attr, &value)

	self.Attrs().Set(attr, value)
	return NewNil()
}

func (m *TypeBase) DelAttr(self Object, args Args) Object {
	var attr *String

	args.ExpectsExactly(2)
	args.Parse(&attr)

	self.Attrs().Del(attr)
	return NewNil()
}

var _hash maphash.Hash

func (m *TypeBase) Hash(self Object, _ Args) *Number {
	defer _hash.Reset()

	if err := gob.NewEncoder(&_hash).Encode(self.Value()); err != nil {
		panic(err)
	}

	return NewNumber(float64(_hash.Sum64()))
}

func (m *TypeBase) Call(Object, Args) Object {
	m.Unimplemented("Call")
	return nil
}
