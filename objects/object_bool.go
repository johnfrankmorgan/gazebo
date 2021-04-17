package objects

var _ Object = &Bool{}

var (
	True  *Bool
	False *Bool
)

func init() {
	True = &Bool{true}
	False = &Bool{false}
}

type Bool struct {
	value bool
}

func FromBool(value bool) *Bool {
	if value {
		return True
	}

	return False
}

func (m *Bool) Value() bool {
	return m.value
}

/* Object methods */

func (m *Bool) GoVal() interface{} {
	return m.Value()
}

func (m *Bool) Hash() interface{} {
	return m.value
}

func (m *Bool) Type() Type {
	return TypeBool
}

func (m *Bool) ToString() *String {
	if m.Value() {
		return NewString("true")
	}

	return NewString("false")
}
