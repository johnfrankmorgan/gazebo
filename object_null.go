package gazebo

type NullObject struct {
	Object
}

var Null = struct {
	Object *NullObject
}{
	Object: &NullObject{
		Object: ObjectInit(Types.Null),
	},
}

var NullObjectMethods = TypeObjectMethods{
	Repr: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Null), "todo")

		return NewStringObjectf("%s { null }", self.Type.Name)
	},

	String: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Null), "todo")

		return NewStringObject("null")
	},

	Bool: func(self *Object) *BoolObject {
		assert(self.Type.Is(Types.Null), "todo")

		return Bools.False
	},
}
