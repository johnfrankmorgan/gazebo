package g

var _ Object = &String{}

type String struct {
	Base
	value string
}

func NewString(value string) *String {
	object := &String{value: value}

	object.SetType(TypeString)
	object.SetSelf(object)

	return object
}

func (m *String) Value() interface{} {
	return m.value
}

func (m *String) String() string {
	return m.value
}
