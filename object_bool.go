package gazebo

type BoolObject struct {
	IntObject
}

var Bools = struct {
	True  *BoolObject
	False *BoolObject
}{
	False: &BoolObject{
		IntObject: IntObject{
			Object: ObjectInit(Types.Bool),
			value:  0,
		},
	},

	True: &BoolObject{
		IntObject: IntObject{
			Object: ObjectInit(Types.Bool),
			value:  1,
		},
	},
}

func (self *BoolObject) Value() bool {
	return self.IntObject.Value() != 0
}

var BoolObjectMethods = TypeObjectMethods{
	Repr: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Bool), "todo")

		return NewStringObjectf("%s { %v }", self.Type.Name, (*BoolObject)(self.Ptr()).Value())
	},

	String: func(self *Object) *StringObject {
		assert(self.Type.Is(Types.Bool), "todo")

		return NewStringObjectf("%v", (*BoolObject)(self.Ptr()).Value())
	},

	Bool: func(self *Object) *BoolObject {
		assert(self.Type.Is(Types.Bool), "todo")

		return (*BoolObject)(self.Ptr())
	},
}
