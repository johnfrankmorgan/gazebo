package g

import (
	"fmt"
	"math"

	"github.com/johnfrankmorgan/gazebo/protocols"
)

// Builtins returns all of the builtin objects in gazebo
func Builtins() map[string]Object {
	methodcall := func(name string) Object {
		return NewObjectInternalFunc(func(args Args) Object {
			self, args := args.SelfWithArgs()
			return self.Call(name, args)
		})
	}

	wrapmethods := func(builtins map[string]Object) map[string]Object {
		for _, method := range protocols.All() {
			builtins[method] = methodcall(method)
		}

		return builtins
	}

	return wrapmethods(map[string]Object{
		"nil": NewObjectNil(),

		"false": NewObjectBool(false),

		"true": NewObjectBool(true),

		"!": NewObjectInternalFunc(func(args Args) Object {
			return NewObjectBool(!IsTruthy(args.Self()))
		}),

		"%": NewObjectInternalFunc(func(args Args) Object {
			args.Expects(2)

			result := math.Mod(ToFloat(args[0]), ToFloat(args[1]))
			return NewObjectNumber(result)
		}),

		"call": NewObjectInternalFunc(func(args Args) Object {
			name, args := args.SelfWithArgs()
			self, args := args.SelfWithArgs()
			return self.Call(EnsureString(name).String(), args)
		}),

		"nil?": NewObjectInternalFunc(func(args Args) Object {
			return NewObjectBool(args.Self().Type() == TypeNil)
		}),

		"println": NewObjectInternalFunc(func(args Args) Object {
			fmt.Println(args.Values()...)
			return NewObjectNil()
		}),

		"printf": NewObjectInternalFunc(func(args Args) Object {
			format, args := args.SelfWithArgs()
			fmt.Printf(EnsureString(format).String(), args.Values()...)
			return NewObjectNil()
		}),
	})
}
