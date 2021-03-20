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
