package objects

import "fmt"

type String struct {
	Object

	value string
}

func NewString(value string) *String {
	return &String{
		Object: ObjectInit(Types.String),
		value:  value,
	}
}

func NewStringf(format string, args ...any) *String {
	return NewString(fmt.Sprintf(format, args...))
}

func (self *String) Value() string {
	return self.value
}

var StringMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.String), "todo")

		return NewStringf("%s { %q }", self.Type.Name, (*String)(self.Ptr()).Value())
	},

	String: func(self *Object) *String {
		assert(self.Type.Is(Types.String), "todo")

		return (*String)(self.Ptr())
	},

	Bool: func(self *Object) *Bool {
		assert(self.Type.Is(Types.String), "todo")

		return Singletons.Bool((*String)(self.Ptr()).Value() != "")
	},

	Equals: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.String), "todo")
		assert(other.Type.Is(Types.String), "todo")

		return (*String)(self.Ptr()).Equals((*String)(other.Ptr()))
	},

	Less: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.String), "todo")
		assert(other.Type.Is(Types.String), "todo")

		return (*String)(self.Ptr()).Less((*String)(other.Ptr()))
	},

	Greater: func(self, other *Object) *Bool {
		assert(self.Type.Is(Types.String), "todo")
		assert(other.Type.Is(Types.String), "todo")

		return (*String)(self.Ptr()).Greater((*String)(other.Ptr()))
	},

	Add: func(self, other *Object) *Object {
		assert(self.Type.Is(Types.String), "todo")
		assert(other.Type.Is(Types.String), "todo")

		return (*String)(self.Ptr()).Add((*String)(other.Ptr())).AsObject()
	},
}

func (self *String) Equals(other *String) *Bool {
	return Singletons.Bool(self.Value() == other.Value())
}

func (self *String) Less(other *String) *Bool {
	return Singletons.Bool(self.Value() < other.Value())
}

func (self *String) Greater(other *String) *Bool {
	return Singletons.Bool(self.Value() > other.Value())
}

func (self *String) Add(other *String) *String {
	return &String{
		Object: ObjectInit(self.Type),
		value:  self.Value() + other.Value(),
	}
}
