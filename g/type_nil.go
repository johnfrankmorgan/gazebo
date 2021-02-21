package g

import "github.com/johnfrankmorgan/gazebo/protocols"

func initnil() {
	TypeNil = &Type{
		Name:   "Nil",
		Parent: TypeBase,
		Methods: Methods{
			protocols.Bool: Method(func(_ Object, _ Args) Object {
				return NewObjectBool(false)
			}),
		},
	}
}
