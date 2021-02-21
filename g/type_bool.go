package g

import "github.com/johnfrankmorgan/gazebo/protocols"

func initbool() {
	TypeBool = &Type{
		Name:   "Bool",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Bool: Method(func(self Object, _ Args) Object {
				return NewObjectBool(EnsureBool(self).Bool())
			}),

			protocols.Number: Method(func(self Object, _ Args) Object {
				if EnsureBool(self).Bool() {
					return NewObjectNumber(1)
				}

				return NewObjectNumber(0)
			}),
		},
	}
}
