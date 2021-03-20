package g

type Bool struct {
	Base
	value bool
}

func NewBool(value bool) *Bool {
	object := &Bool{value: value}

	object.SetType(TypeBool)
	object.SetSelf(object)

	return object
}

func (m *Bool) Value() interface{} {
	return m.value
}

func (m *Bool) ToBool() *Bool {
	return NewBool(m.value)
}

func (m *Bool) ToNumber() *Number {
	if m.value {
		return NewNumberFromInt(1)
	}

	return NewNumberFromInt(0)
}

func (m *Bool) ToString() *String {
	if m.value {
		return NewString("true")
	}

	return NewString("false")
}

func (m *Bool) Bool() bool {
	return m.value
}
