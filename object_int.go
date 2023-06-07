package gazebo

import "golang.org/x/exp/constraints"

type IntObject struct {
	Object

	value int64
}

func NewIntObject[T constraints.Integer](value T) *IntObject {
	return &IntObject{
		Object: ObjectInit(Types.Int),
		value:  int64(value),
	}
}

func (self *IntObject) Value() int64 {
	return self.value
}

var IntObjectMethods = TypeObjectMethods{
	Repr: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Int), "todo")

		return NewStringObjectf("%s { %d }", self.Type.Name, (*IntObject)(self.Ptr()).Value())
	},

	String: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Int), "todo")

		return NewStringObjectf("%d", (*IntObject)(self.Ptr()).Value())
	},

	Bool: func(self *Object) *BoolObject {
		assert(self.Type.Is(Types.Int), "todo")

		if (*IntObject)(self.Ptr()).Value() != 0 {
			return Bools.True
		}

		return Bools.False
	},

	Add: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.Int), "todo")
		assert(other.Type.Is(Types.Int), "todo")

		return (*IntObject)(self.Ptr()).Add((*IntObject)(other.Ptr())).AsObject()
	},

	Sub: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.Int), "todo")
		assert(other.Type.Is(Types.Int), "todo")

		return (*IntObject)(self.Ptr()).Sub((*IntObject)(other.Ptr())).AsObject()
	},

	Mul: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.Int), "todo")
		assert(other.Type.Is(Types.Int), "todo")

		return (*IntObject)(self.Ptr()).Mul((*IntObject)(other.Ptr())).AsObject()
	},

	Div: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.Int), "todo")
		assert(other.Type.Is(Types.Int), "todo")

		return (*IntObject)(self.Ptr()).Div((*IntObject)(other.Ptr())).AsObject()
	},

	Mod: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.Int), "todo")
		assert(other.Type.Is(Types.Int), "todo")

		return (*IntObject)(self.Ptr()).Mod((*IntObject)(other.Ptr())).AsObject()
	},
}

func (self *IntObject) Add(other *IntObject) *IntObject {
	return &IntObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() + other.Value(),
	}
}

func (self *IntObject) Sub(other *IntObject) *IntObject {
	return &IntObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() - other.Value(),
	}
}

func (self *IntObject) Mul(other *IntObject) *IntObject {
	return &IntObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() * other.Value(),
	}
}

func (self *IntObject) Div(other *IntObject) *IntObject {
	return &IntObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() / other.Value(),
	}
}

func (self *IntObject) Mod(other *IntObject) *IntObject {
	return &IntObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() % other.Value(),
	}
}
