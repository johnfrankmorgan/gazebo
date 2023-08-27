package objects

import "golang.org/x/exp/constraints"

type Integer struct {
	Object

	value int64
}

func NewInteger[T constraints.Integer](value T) *Integer {
	return &Integer{
		Object: ObjectInit(Types.Integer),
		value:  int64(value),
	}
}

func (self *Integer) Value() int64 {
	return self.value
}

var IntegerMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Integer), "todo")

		return NewStringf("%s { %d }", self.Type.Name, (*Integer)(self.Ptr()).Value())
	},

	String: func(self *Object) *String {
		assert(self.Type.Is(Types.Integer), "todo")

		return NewStringf("%d", (*Integer)(self.Ptr()).Value())
	},

	Bool: func(self *Object) *Bool {
		assert(self.Type.Is(Types.Integer), "todo")

		return Singletons.Bool((*Integer)(self.Ptr()).Value() != 0)
	},

	Equals: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Equals((*Integer)(other.Ptr()))
	},

	Less: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Less((*Integer)(other.Ptr()))
	},

	Greater: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Greater((*Integer)(other.Ptr()))
	},

	Add: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Add((*Integer)(other.Ptr())).AsObject()
	},

	Subtract: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Subtract((*Integer)(other.Ptr())).AsObject()
	},

	Multiply: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Multiply((*Integer)(other.Ptr())).AsObject()
	},

	Divide: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Divide((*Integer)(other.Ptr())).AsObject()
	},

	Modulus: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.Integer), "todo")
		assert(other.Type.Is(Types.Integer), "todo")

		return (*Integer)(self.Ptr()).Modulus((*Integer)(other.Ptr())).AsObject()
	},
}

func (self *Integer) Equals(other *Integer) *Bool {
	return Singletons.Bool(self.Value() == other.Value())
}

func (self *Integer) Less(other *Integer) *Bool {
	return Singletons.Bool(self.Value() < other.Value())
}

func (self *Integer) Greater(other *Integer) *Bool {
	return Singletons.Bool(self.Value() > other.Value())
}

func (self *Integer) Add(other *Integer) *Integer {
	return &Integer{
		Object: ObjectInit(self.Type),
		value:  self.Value() + other.Value(),
	}
}

func (self *Integer) Subtract(other *Integer) *Integer {
	return &Integer{
		Object: ObjectInit(self.Type),
		value:  self.Value() - other.Value(),
	}
}

func (self *Integer) Multiply(other *Integer) *Integer {
	return &Integer{
		Object: ObjectInit(self.Type),
		value:  self.Value() * other.Value(),
	}
}

func (self *Integer) Divide(other *Integer) *Integer {
	return &Integer{
		Object: ObjectInit(self.Type),
		value:  self.Value() / other.Value(),
	}
}

func (self *Integer) Modulus(other *Integer) *Integer {
	return &Integer{
		Object: ObjectInit(self.Type),
		value:  self.Value() % other.Value(),
	}
}
