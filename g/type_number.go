package g

import "github.com/johnfrankmorgan/gazebo/protocols"

func initnumber() {
	TypeNumber = &Type{
		Name:   "Number",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Bool: Method(func(self Object, _ Args) Object {
				return NewObjectBool(EnsureNumber(self).Float() != 0)
			}),

			protocols.Number: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(EnsureNumber(self).Float())
			}),

			protocols.Inverse: Method(func(self Object, _ Args) Object {
				return NewObjectNumber(-EnsureNumber(self).Float())
			}),

			protocols.Add: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float()

				for _, arg := range args {
					result += ToFloat(arg)
				}

				return NewObjectNumber(result)
			}),

			protocols.Sub: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float() - ToFloat(args.Self())
				return NewObjectNumber(result)
			}),

			protocols.Mul: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float() * ToFloat(args.Self())
				return NewObjectNumber(result)
			}),

			protocols.Div: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float() / ToFloat(args.Self())
				return NewObjectNumber(result)
			}),

			protocols.LessThan: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float() < ToFloat(args.Self())
				return NewObjectBool(result)
			}),

			protocols.GreaterThan: Method(func(self Object, args Args) Object {
				result := EnsureNumber(self).Float() > ToFloat(args.Self())
				return NewObjectBool(result)
			}),
		},
	}
}
