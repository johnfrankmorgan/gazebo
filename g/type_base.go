package g

import (
	"fmt"
	"reflect"

	"github.com/johnfrankmorgan/gazebo/errors"
	"github.com/johnfrankmorgan/gazebo/protocols"
)

func initbase() {
	TypeBase = &Type{
		Name:   "Base",
		Parent: nil,
		Methods: Methods{
			protocols.Bool: Method(func(_ Object, _ Args) Object {
				return NewObjectBool(true)
			}),

			protocols.Not: Method(func(self Object, _ Args) Object {
				return NewObjectBool(!IsTruthy(self))
			}),

			protocols.String: Method(func(self Object, _ Args) Object {
				return NewObjectString(fmt.Sprintf("%v", self.Value()))
			}),

			protocols.Number: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(0)
			}),

			protocols.Inspect: Method(func(self Object, _ Args) Object {
				inspection := fmt.Sprintf(
					"<gtypes.%s>(%v)",
					self.Type().Name,
					self.Value(),
				)

				return NewObjectString(inspection)
			}),

			protocols.Equal: Method(func(self Object, args Args) Object {
				for _, arg := range args {
					if !reflect.DeepEqual(self.Value(), arg.Value()) {
						return NewObjectBool(false)
					}
				}

				return NewObjectBool(true)
			}),

			protocols.NotEqual: Method(func(self Object, args Args) Object {
				return NewObjectBool(!EnsureBool(self.Call(protocols.Equal, args)).Bool())
			}),

			protocols.HasAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()
				if self.Attributes().Has(name) {
					return NewObjectBool(true)
				}

				return NewObjectBool(self.Type().Implements(name))
			}),

			protocols.GetAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()

				if self.Attributes().Has(name) {
					return self.Attributes().Get(name)
				}

				if self.Type().Implements(name) {
					return NewObjectInternalFunc(func(args Args) Object {
						return self.Call(name, args)
					})
				}

				errors.ErrRuntime.Panic(
					"undefined attribute or method %s for type %s",
					name,
					self.Type().Name,
				)

				return nil
			}),

			protocols.SetAttr: Method(func(self Object, args Args) Object {
				args.Expects(2)

				name := EnsureString(args.Self()).String()

				self.Attributes().Set(name, args[1])

				return NewObjectNil()
			}),

			protocols.DelAttr: Method(func(self Object, args Args) Object {
				name := EnsureString(args.Self()).String()

				self.Attributes().Delete(name)

				return NewObjectNil()
			}),
		},
	}
}
