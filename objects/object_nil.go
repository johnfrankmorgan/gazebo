package objects

var _ Object = &Bool{}

type Nil struct{}

func NewNil() *Nil {
	return &Nil{}
}

/* Object methods */

func (m *Nil) GoVal() interface{} {
	return nil
}

func (m *Nil) Hash() interface{} {
	return nil
}

func (m *Nil) Type() Type {
	return TypeNil
}

func (m *Nil) ToString() *String {
	return NewString("nil")
}
