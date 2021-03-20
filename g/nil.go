package g

var _ Object = &Nil{}

type Nil struct {
	Base
}

func NewNil() *Nil {
	object := &Nil{}

	object.SetType(TypeNil)
	object.SetSelf(object)

	return object
}

func (m *Nil) Value() interface{} {
	return nil
}

func (m *Nil) ToBool() *Bool {
	return NewBool(false)
}

func (m *Nil) ToNumber() *Number {
	return NewNumberFromInt(0)
}

func (m *Nil) ToString() *String {
	return NewString("nil")
}
