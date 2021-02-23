package g

var _ Object = &String{}

type String struct {
	AttrsNoOp
	H     ObjectHelper
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (m *String) Value() interface{} {
	return m.value
}

func (m *String) GetAttr(name string) Object {
	return m.H.GetAttr(m, name)
}

// GAZEBO STRING OBJECT METHODS

func (m *String) G_str() Object {
	return &String{value: m.value}
}

func (m *String) G_eq(other *String) Object {
	if m.value == other.value {
		return NewString("true")
	}

	return NewString("false")
}
