package gazebo

import "fmt"

type StringObject struct {
	Object

	value string
}

func NewStringObject(value string) *StringObject {
	return &StringObject{
		Object: ObjectInit(Types.String),
		value:  value,
	}
}

func NewStringObjectf(format string, args ...any) *StringObject {
	return NewStringObject(fmt.Sprintf(format, args...))
}

func (self *StringObject) Value() string {
	return self.value
}

var StringObjectMethods = TypeObjectMethods{
	Repr: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.String), "todo")

		return NewStringObjectf("%s { %q }", self.Type.Name, (*StringObject)(self.Ptr()).Value())
	},

	String: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.String), "todo")

		return (*StringObject)(self.Ptr())
	},

	Bool: func(self *Object) *BoolObject {
		assert(self.Type.Is(Types.String), "todo")

		if (*StringObject)(self.Ptr()).Value() == "" {
			return Bools.False
		}

		return Bools.True
	},

	Add: func(self *Object, other *Object) *Object {
		assert(self.Type.Is(Types.String), "todo")
		assert(other.Type.Is(Types.String), "todo")

		return (*StringObject)(self.Ptr()).Add((*StringObject)(other.Ptr())).AsObject()
	},
}

func (self *StringObject) Add(other *StringObject) *StringObject {
	return &StringObject{
		Object: ObjectInit(self.Type),
		value:  self.Value() + other.Value(),
	}
}
