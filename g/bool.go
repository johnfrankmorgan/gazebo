package g

var _ Object = &Bool{}

type Bool struct {
	Partial
	h     ObjectHelper
	value bool
}

func NewBool(value bool) *Bool {
	object := &Bool{value: value}
	object.self = object
	return object
}

func (m *Bool) Value() interface{} {
	return m.value
}

func (m *Bool) Bool() bool {
	return m.value
}

func (m *Bool) CallMethod(name string, args *Args) Object {
	return m.h.CallMethod(m, name, args)
}

func (m *Bool) HasAttr(name string) bool {
	return m.h.HasAttr(m, name)
}

func (m *Bool) GetAttr(name string) Object {
	return m.h.GetAttr(m, name)
}

func (m *Bool) SetAttr(name string, value Object) {
	m.h.SetAttr(m, name, value)
}

func (m *Bool) DelAttr(name string) {
	m.h.DelAttr(m, name)
}

// GAZEBO BOOL OBJECT METHODS

func (m *Bool) G_str() *String {
	if m.value {
		return NewString("true")
	}

	return NewString("false")
}

func (m *Bool) G_num() *Number {
	if m.value {
		return NewNumber(1)
	}

	return NewNumber(0)
}

func (m *Bool) G_bool() *Bool {
	return NewBool(m.value)
}

func (m *Bool) G_not() *Bool {
	return NewBool(!m.value)
}
