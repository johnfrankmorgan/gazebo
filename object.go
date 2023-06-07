package gazebo

import "unsafe"

type Object struct {
	Type *TypeObject
}

func ObjectInit(typ *TypeObject) Object {
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

var ObjectMethods = TypeObjectMethods{
	Repr: func(self *Object) *StringObject {
		return NewStringObjectf("%s { %p }", self.Type.Name, self)
	},

	String: func(self *Object) *StringObject {
		return self.Type.Repr(self)
	},

	Bool: func(self *Object) *BoolObject {
		return Bools.True
	},
}
