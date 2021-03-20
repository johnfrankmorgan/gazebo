package g

type Object interface {
	Value() interface{}
	Type() Type
	Attrs() *Map
}

type Type interface {
	Name() string
	Parent() Type
	Methods() Methods
	Object
}

type Base struct {
	_type  Type
	_self  Object
	_attrs *Map
}

func (m *Base) Type() Type {
	return m._type
}

func (m *Base) SetType(t Type) {
	m._type = t
}

func (m *Base) SetSelf(self Object) {
	m._self = self
}

func (m *Base) Attrs() *Map {
	if m._attrs == nil {
		m._attrs = NewMap()
	}

	return m._attrs
}

func (m *Base) SetAttrs(attrs *Map) {
	m._attrs = attrs
}
