package objects

type Bool struct {
	Integer
}

func (self *Bool) Value() bool {
	return self.Integer.Value() != 0
}

var BoolMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Bool), "todo")

		return NewStringf("%s { %v }", self.Type.Name, (*Bool)(self.Ptr()).Value())
	},

	String: func(self *Object) *String {
		assert(self.Type.Is(Types.Bool), "todo")

		return NewStringf("%v", (*Bool)(self.Ptr()).Value())
	},

	Bool: func(self *Object) *Bool {
		assert(self.Type.Is(Types.Bool), "todo")

		return (*Bool)(self.Ptr())
	},
}
