package runtime

import "fmt"

type Func struct {
	name String
	fn   func(Tuple) Object
}

func (fn *Func) Type() *Type {
	return Types.Func
}

func (fn *Func) Name() String {
	return fn.name
}

func (fn *Func) Repr() String {
	return Stringf("%s(%s@%p)", fn.Type().Name, fn.Name(), fn)
}

func (fn *Func) Call(args Tuple) Object {
	return fn.fn(args)
}

var Builtins = []*Func{
	{
		"print", func(args Tuple) Object {
			for i, arg := range args {
				if i > 0 {
					fmt.Print(" ")
				}

				fmt.Print(Objects.String(arg))
			}

			fmt.Print("\n")

			return Nil
		},
	},

	{
		"repr", func(args Tuple) Object {
			if len(args) != 1 {
				panic("repr() expects exactly 1 argument")
			}

			return Objects.Repr(args[0])
		},
	},
}
