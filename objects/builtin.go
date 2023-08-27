package objects

type Builtin struct {
	Object

	name  string
	value func(...*Object) *Object
}

func NewBuiltin(name string, value func(...*Object) *Object) *Builtin {
	return &Builtin{
		Object: ObjectInit(Types.Builtin),
		name:   name,
		value:  value,
	}
}

func (self *Builtin) Name() string {
	return self.name
}

func (self *Builtin) Value() func(...*Object) *Object {
	return self.value
}

var BuiltinMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Builtin), "todo")

		builtin := (*Builtin)(self.Ptr())

		return NewStringf("%s { %s@%p }", builtin.Type.Name, builtin.Name(), builtin.Value())
	},

	Call: func(self *Object, args ...*Object) *Object {
		assert(self.Type.Is(Types.Builtin), "todo")

		return (*Builtin)(self.Ptr()).Call(args...)
	},
}

func (self *Builtin) Call(args ...*Object) *Object {
	return self.Value()(args...)
}
