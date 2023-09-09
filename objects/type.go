package objects

import "reflect"

type Type struct {
	Object

	Parent *Type

	Name       string
	Methods    TypeMethods
	Attributes TypeAttributes
}

var Types = struct {
	Object  *Type
	Type    *Type
	Null    *Type
	Integer *Type
	Bool    *Type
	String  *Type
	Builtin *Type
	Func    *Type
	Method  *Type
}{
	Object:  &Type{Name: "Object"},
	Type:    &Type{Name: "Type"},
	Null:    &Type{Name: "Null"},
	Integer: &Type{Name: "Integer"},
	Bool:    &Type{Name: "Bool"},
	String:  &Type{Name: "String"},
	Builtin: &Type{Name: "Builtin"},
	Func:    &Type{Name: "Func"},
	Method:  &Type{Name: "Method"},
}

func init() {
	Types.Object.Type = Types.Type
	Types.Object.Methods = ObjectMethods
	Types.Object.Attributes = ObjectAttributes

	Types.Type.Type = Types.Type
	Types.Type.Parent = Types.Object
	Types.Type.Methods = TypeMethods{
		Repr: func(self *Object) *String {
			assert(self.Type.Is(Types.Type), "todo")

			return NewStringf("%s { %q }", self.Type.Name, (*Type)(self.Ptr()).Name)
		},
	}
	Types.Type.Attributes = TypeAttributes{
		"name": TypeAttribute{
			Get: func(self *Object) *Object {
				assert(self.Type.Is(Types.Type), "todo")

				return NewString((*Type)(self.Ptr()).Name).AsObject()
			},
		},
		"parent": TypeAttribute{
			Get: func(self *Object) *Object {
				assert(self.Type.Is(Types.Type), "todo")

				t := (*Type)(self.Ptr())

				if t.Parent == nil {
					return Singletons.Null.AsObject()
				}

				return t.Parent.AsObject()
			},
		},
	}

	Types.Null.Type = Types.Type
	Types.Null.Parent = Types.Object
	Types.Null.Methods = NullMethods

	Types.Integer.Type = Types.Type
	Types.Integer.Parent = Types.Object
	Types.Integer.Methods = IntegerMethods

	Types.Bool.Type = Types.Type
	Types.Bool.Parent = Types.Integer
	Types.Bool.Methods = BoolMethods

	Types.String.Type = Types.Type
	Types.String.Parent = Types.Object
	Types.String.Methods = StringMethods
	Types.String.Attributes = StringAttributes

	Types.Builtin.Type = Types.Type
	Types.Builtin.Parent = Types.Object
	Types.Builtin.Methods = BuiltinMethods
	Types.Builtin.Attributes = BuiltinAttributes

	Types.Func.Type = Types.Type
	Types.Func.Parent = Types.Object
	Types.Func.Methods = FuncMethods
	Types.Func.Attributes = FuncAttributes

	Types.Method.Type = Types.Type
	Types.Method.Parent = Types.Object
	Types.Method.Methods = MethodMethods
	Types.Method.Attributes = MethodAttributes

	rval := reflect.ValueOf(Types)
	for i := 0; i < rval.NumField(); i++ {
		f := rval.Field(i)

		Builtins[f.Interface().(*Type).Name] = f.Interface().(*Type).AsObject()
	}
}

type TypeMethods struct {
	// standard methods
	Bool   func(self *Object) *Bool
	Repr   func(self *Object) *String
	String func(self *Object) *String
	Clone  func(self *Object) *Object

	// unary
	Negate func(self *Object) *Object

	// comparisons
	Equals  func(self, other *Object) *Bool
	Less    func(self, other *Object) *Bool
	Greater func(self, other *Object) *Bool

	// numeric
	Add      func(self, other *Object) *Object
	Subtract func(self, other *Object) *Object
	Multiply func(self, other *Object) *Object
	Divide   func(self, other *Object) *Object
	Modulus  func(self, other *Object) *Object

	// attributes
	GetAttribute func(self *Object, name string) *Object

	// callables
	Call func(self *Object, args ...*Object) *Object
}

type TypeAttribute struct {
	Name string
	Get  func(self *Object) *Object
	Set  func(self, other *Object)
}

type TypeAttributes map[string]TypeAttribute

func (t *Type) Is(typ *Type) bool {
	for t := t; t != nil; t = t.Parent {
		if t == typ {
			return true
		}
	}

	return false
}

func (t *Type) Bool(self *Object) *Bool {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Bool != nil {
			return t.Methods.Bool(self)
		}
	}

	panic("todo")
}

func (t *Type) Repr(self *Object) *String {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Repr != nil {
			return t.Methods.Repr(self)
		}
	}

	panic("todo")
}

func (t *Type) String(self *Object) *String {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.String != nil {
			return t.Methods.String(self)
		}
	}

	panic("todo")
}

func (t *Type) Clone(self *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Clone != nil {
			return t.Methods.Clone(self)
		}
	}

	panic("todo")
}

func (t *Type) Equals(self, other *Object) *Bool {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Equals != nil {
			return t.Methods.Equals(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Less(self, other *Object) *Bool {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Less != nil {
			return t.Methods.Less(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Greater(self, other *Object) *Bool {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Greater != nil {
			return t.Methods.Greater(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Negate(self *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Negate != nil {
			return t.Methods.Negate(self)
		}
	}

	panic("todo")
}

func (t *Type) Add(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Add != nil {
			return t.Methods.Add(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Subtract(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Subtract != nil {
			return t.Methods.Subtract(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Multiply(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Multiply != nil {
			return t.Methods.Multiply(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Divide(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Divide != nil {
			return t.Methods.Divide(self, other)
		}
	}

	panic("todo")
}

func (t *Type) Modulus(self, other *Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Modulus != nil {
			return t.Methods.Modulus(self, other)
		}
	}

	panic("todo")
}

func (t *Type) GetAttribute(self *Object, name string) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.GetAttribute != nil {
			return t.Methods.GetAttribute(self, name)
		}
	}

	panic("todo")
}

func (t *Type) Call(self *Object, args ...*Object) *Object {
	for t := t; t != nil; t = t.Parent {
		if t.Methods.Call != nil {
			return t.Methods.Call(self, args...)
		}
	}

	panic("todo")
}
