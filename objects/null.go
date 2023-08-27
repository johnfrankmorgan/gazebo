package objects

type Null struct {
	Object
}

var NullMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Null), "todo")

		return NewStringf("%s { null }", self.Type.Name)
	},

	String: func(self *Object) *String {
		assert(self.Type.Is(Types.Null), "todo")

		return NewString("null")
	},

	Bool: func(self *Object) *Bool {
		assert(self.Type.Is(Types.Null), "todo")

		return Singletons.True
	},
}
