package objects

import "unsafe"

type Object struct {
	Type *Type
}

func ObjectInit(typ *Type) Object {
	return Object{Type: typ}
}

func (obj *Object) Ptr() unsafe.Pointer {
	return unsafe.Pointer(obj)
}

func (obj *Object) AsObject() *Object {
	return obj
}

func (obj *Object) GoString() string {
	return obj.Type.Repr(obj).Value()
}

var ObjectMethods = TypeMethods{
	Repr: func(self *Object) *String {
		return NewStringf("%s { %p }", self.Type.Name, self)
	},

	String: func(self *Object) *String {
		return self.Type.Repr(self)
	},

	Bool: func(self *Object) *Bool {
		return Singletons.True
	},

	Equals: func(self, other *Object) *Bool {
		return Singletons.Bool(self == other)
	},
}
