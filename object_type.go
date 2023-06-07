package gazebo

type TypeObject struct {
	Object

	Parent *TypeObject

	Name    string
	Methods TypeObjectMethods
}

var Types = struct {
	Object *TypeObject
	Type   *TypeObject
	Null   *TypeObject
	Int    *TypeObject
	Bool   *TypeObject
	String *TypeObject
}{
	Object: &TypeObject{Name: "Object"},
	Type:   &TypeObject{Name: "Type"},
	Null:   &TypeObject{Name: "Null"},
	Int:    &TypeObject{Name: "Int"},
	Bool:   &TypeObject{Name: "Bool"},
	String: &TypeObject{Name: "String"},
}

func init() {
	Types.Object.Type = Types.Type
	Types.Object.Methods = ObjectMethods

	Types.Type.Type = Types.Type
	Types.Type.Parent = Types.Object
	Types.Type.Methods = TypeObjectMethods{
		Repr: func(self *Object) *StringObject {
			assert(self.Type.Is(Types.Type), "todo")

			return NewStringObjectf("%s { %q }", self.Type.Name, (*TypeObject)(self.Ptr()).Name)
		},
	}

	Types.Null.Type = Types.Type
	Types.Null.Parent = Types.Object
	Types.Null.Methods = NullObjectMethods

	Types.Int.Type = Types.Type
	Types.Int.Parent = Types.Object
	Types.Int.Methods = IntObjectMethods

	Types.Bool.Type = Types.Type
	Types.Bool.Parent = Types.Int
	Types.Bool.Methods = BoolObjectMethods

	Types.String.Type = Types.Type
	Types.String.Parent = Types.Object
	Types.String.Methods = StringObjectMethods
}

type TypeObjectMethods struct {
	// standard methods
	Bool   func(self *Object) *BoolObject
	Repr   func(self *Object) *StringObject
	String func(self *Object) *StringObject

	// number methods
	Add func(self, other *Object) *Object
	Sub func(self, other *Object) *Object
	Mul func(self, other *Object) *Object
	Div func(self, other *Object) *Object
	Mod func(self, other *Object) *Object
}

func (t *TypeObject) Is(typ *TypeObject) bool {
	for t := t; t != nil; t = t.Parent {
		if t == typ {
			return true
		}
	}

	return false
}

func (t *TypeObject) Bool(self *Object) *BoolObject {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Bool != nil {
			return t.Methods.Bool(self)
		}
	}

	panic("todo")
}

func (t *TypeObject) Repr(self *Object) *StringObject {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Repr != nil {
			return t.Methods.Repr(self)
		}
	}

	panic("todo")
}

func (t *TypeObject) String(self *Object) *StringObject {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.String != nil {
			return t.Methods.String(self)
		}
	}

	panic("todo")
}

func (t *TypeObject) Add(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Add != nil {
			return t.Methods.Add(self, other)
		}
	}

	panic("todo")
}

func (t *TypeObject) Sub(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Sub != nil {
			return t.Methods.Sub(self, other)
		}
	}

	panic("todo")
}

func (t *TypeObject) Mul(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Mul != nil {
			return t.Methods.Mul(self, other)
		}
	}

	panic("todo")
}

func (t *TypeObject) Div(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Div != nil {
			return t.Methods.Div(self, other)
		}
	}

	panic("todo")
}

func (t *TypeObject) Mod(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Mod != nil {
			return t.Methods.Mod(self, other)
		}
	}

	panic("todo")
}
