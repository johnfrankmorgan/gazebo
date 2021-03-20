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

func (m *Bool) Bool() bool {
	return m.value
}
