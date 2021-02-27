package g

var _ Object = &Bool{}

type Bool struct {
	Base
	value bool
}

func NewBool(value bool) *Bool {
	object := &Bool{value: value}
	object.SetSelf(object)
	return object
}

func (m *Bool) Value() interface{} {
	return m.value
}

func (m *Bool) Bool() bool {
	return m.value
}

// GAZEBO BOOL OBJECT PROTOCOLS

func (m *Bool) G_repr() *String {
	return m.G_str()
}

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
