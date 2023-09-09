package objects

type Method struct {
	Object

	name string
	self *Object
	fn   func(*Object, ...*Object) *Object
}

func NewMethod(name string, self *Object, fn func(*Object, ...*Object) *Object) *Method {
	return &Method{
		Object: ObjectInit(Types.Method),
		name:   name,
		self:   self,
		fn:     fn,
	}
}

func MethodAttribute(name string, fn func(*Object, ...*Object) *Object) TypeAttribute {
	return TypeAttribute{
		Get: func(self *Object) *Object {
			return NewMethod(name, self, fn).AsObject()
		},
	}
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) Self() *Object {
	return self.self
}

func (self *Method) Func() func(*Object, ...*Object) *Object {
	return self.fn
}

var MethodMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Method), "todo")

		method := (*Method)(self.Ptr())

		return NewStringf("%s{ %s.%s@%p }", method.Type.Name, method.Self().Type.Name, method.Name(), method.Func())
	},

	Call: func(self *Object, args ...*Object) *Object {
		assert(self.Type.Is(Types.Method), "todo")

		method := (*Method)(self.Ptr())

		return method.Func()(method.self, args...)
	},
}

var MethodAttributes = TypeAttributes{
	"name": TypeAttribute{
		Get: func(self *Object) *Object {
			assert(self.Type.Is(Types.Method), "todo")

			return NewString((*Method)(self.Ptr()).Name()).AsObject()
		},
	},

	"self": TypeAttribute{
		Get: func(self *Object) *Object {
			assert(self.Type.Is(Types.Method), "todo")

			return (*Method)(self.Ptr()).Self()
		},
	},
}
