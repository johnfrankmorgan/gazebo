package objects

import (
	"gazebo/compiler"
	"strings"
)

type Func struct {
	Object

	runner    FuncRunner
	name      string
	arguments []string
	code      *compiler.Code
}

type FuncRunner interface {
	RunFunc(*Func, []*Object) *Object
}

func NewFunc(runner FuncRunner, name string, arguments []string, code *compiler.Code) *Func {
	return &Func{
		Object:    ObjectInit(Types.Func),
		runner:    runner,
		name:      name,
		arguments: arguments,
		code:      code,
	}
}

func (self *Func) Name() string {
	return self.name
}

func (self *Func) Arguments() []string {
	return self.arguments
}

func (self *Func) Code() *compiler.Code {
	return self.code
}

var FuncMethods = TypeMethods{
	Repr: func(self *Object) *String {
		assert(self.Type.Is(Types.Func), "todo")

		fn := (*Func)(self.Ptr())

		arguments := strings.Builder{}

		for i, argument := range fn.Arguments() {
			if i > 0 {
				arguments.WriteString(", ")
			}

			arguments.WriteString(argument)
		}

		return NewStringf("%s { %s(%s) }", fn.Type.Name, fn.Name(), arguments.String())
	},

	Call: func(self *Object, args ...*Object) *Object {
		assert(self.Type.Is(Types.Func), "todo")

		fn := (*Func)(self.Ptr())

		assert(len(args) == len(fn.Arguments()), "todo")

		return fn.runner.RunFunc(fn, args)
	},
}

var FuncAttributes = TypeAttributes{
	"name": TypeAttribute{
		Get: func(self *Object) *Object {
			assert(self.Type.Is(Types.Func), "todo")

			return NewString((*Func)(self.Ptr()).Name()).AsObject()
		},
	},
}
