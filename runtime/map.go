package runtime

import "errors"

type Map struct {
	ht ht
}

var MapType = &Type{
	Name:   "Map",
	Parent: ObjectType,
	Protocols: TypeProtocols{
		Bool:   func(self Object) Bool { return self.(*Map).Bool() },
		String: func(self Object) String { return String(self.(*Map).String()) },
	},
	Ops: TypeOps{
		Contains: func(self, other Object) Bool { return self.(*Map).Contains(other) },
		GetIndex: func(self, index Object) Object { return self.(*Map).GetIndex(index) },
		SetIndex: func(self, index, value Object) { self.(*Map).Set(index, value) },
		GetAttribute: func(self Object, name String) (value Object) {
			defer func() {
				if err := recover(); err != nil && errors.Is(err.(error), ErrInvalidAttribute) {
					value = self.(*Map).GetIndex(name)
				}
			}()

			return ObjectType.Ops.GetAttribute(self, name)
		},
		SetAttribute: func(self Object, name String, value Object) {
			defer func() {
				if err := recover(); err != nil && errors.Is(err.(error), ErrInvalidAttribute) {
					self.(*Map).Set(name, value)
				}
			}()

			ObjectType.Ops.SetAttribute(self, name, value)
		},
	},
	Attributes: TypeAttributes{
		"len": Attribute{
			Get: func(self Object) Object { return self.(*Map).Len() },
		},
	},
}

func NewMap() *Map {
	return new(Map)
}

func (m *Map) Type() *Type {
	return MapType
}

func (m *Map) Bool() Bool {
	return m.Len() != 0
}

func (m *Map) String() String {
	panic("todo")
}

func (m *Map) Len() Int {
	return Int(m.ht.size)
}

func (m *Map) Contains(key Object) Bool {
	_, ok := m.Get(key)
	return ok
}

func (m *Map) Get(key Object) (Object, Bool) {
	return m.ht.get(key)
}

func (m *Map) GetIndex(index Object) Object {
	if value, ok := m.Get(index); ok {
		return value
	}

	panic(ErrUnimplemented)
}

func (m *Map) Set(key, value Object) {
	m.ht.set(key, value)
}
