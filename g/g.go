package g

import "fmt"

func NewObject(value interface{}) Object {
	switch value := value.(type) {
	case nil:
		return NewNil()

	case bool:
		return NewBool(value)

	case int:
		return NewNumberFromInt(value)

	case float64:
		return NewNumber(value)

	case string:
		return NewString(value)
	}

	panic(fmt.Errorf("can't coerce %T to Object", value))
}

type Object interface {
	Value() interface{}
	Type() Type
	Attrs() *Map
	HasAttr(string) bool
	GetAttr(string) Object
	SetAttr(string, Object)
	DelAttr(string)
	ToBool() *Bool
	ToNumber() *Number
	ToString() *String
}

type Type interface {
	Name() string
	Parent() Type
	Methods() Methods
	Object
}

type Base struct {
	_type  Type
	_self  Object
	_attrs *Map
}

func (m *Base) Type() Type {
	return m._type
}

func (m *Base) SetType(t Type) {
	m._type = t
}

func (m *Base) Self() Object {
	return m._self
}

func (m *Base) SetSelf(self Object) {
	m._self = self
}

func (m *Base) Attrs() *Map {
	if m._attrs == nil {
		m._attrs = NewMap()
	}

	return m._attrs
}

func (m *Base) HasAttr(name string) bool {
	if Resolve(m.Type(), name) != nil {
		return true
	}

	return m.Attrs().Has(NewString(name))
}

func (m *Base) GetAttr(name string) Object {
	if method := Resolve(m.Type(), name); method != nil {
		return NewBoundMethod(m.Self(), method)
	}

	return m.Attrs().Get(NewString(name))
}

func (m *Base) SetAttr(name string, value Object) {
	m.Attrs().Set(NewString(name), value)
}

func (m *Base) DelAttr(name string) {
	m.Attrs().Del(NewString(name))
}

func (m *Base) SetAttrs(attrs *Map) {
	m._attrs = attrs
}

func (m *Base) ToBool() *Bool {
	return NewBool(true)
}

func (m *Base) ToNumber() *Number {
	panic(fmt.Errorf("ToNumber not implemented for %q", m.Self().Type().Name()))
}

func (m *Base) ToString() *String {
	panic(fmt.Errorf("ToString not implemented for %q", m.Self().Type().Name()))
}
